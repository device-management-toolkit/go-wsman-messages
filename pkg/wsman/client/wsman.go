/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/amterror"
)

const (
	ContentType           = "application/soap+xml; charset=utf-8"
	NSWSMAN               = "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"
	NSWSMID               = "http://schemas.dmtf.org/wbem/wsman/identity/1/wsmanidentity.xsd"
	TLSPort               = "16993"
	NonTLSPort            = "16992"
	RedirectionTLSPort    = "16995"
	RedirectionNonTLSPort = "16994"
	WSManPath             = "/wsman"
	idleConnTimeout       = 90 * time.Second // keeps sockets warm as long as AMT holds them open
	// Keep at 2 to avoid overload during powered-off wake paths;
	// higher concurrency caused transient auth/transport failures.
	maxConcurrentPosts  = 2
	maxTransientRetries = 3
	retryDelay          = 500 * time.Millisecond // pause before each retry, lets the ME recover
	// Windows socket errors. Their text differs from the POSIX wording (and is
	// localized), so match on the numeric code. Declared here rather than using
	// syscall.WSAE*, which only exists on Windows builds.
	wsaeConnAborted syscall.Errno = 10053 // WSAECONNABORTED
	wsaeConnReset   syscall.Errno = 10054 // WSAECONNRESET
)

type Message struct {
	XMLInput  string
	XMLOutput string
}

// WSMan is an interface for the wsman.Client.
type WSMan interface {
	// HTTP Methods
	Post(msg string) (response []byte, err error)
	// TCP Methods
	Connect() error
	Send(data []byte) error
	Receive() ([]byte, error)
	CloseConnection() error
	IsAuthenticated() bool
	GetServerCertificate() (*tls.Certificate, error)
}

// Target is a thin wrapper around http.Target.
type Target struct {
	http.Client
	endpoint           string
	username           string
	password           string
	useDigest          bool
	logAMTMessages     bool
	challenge          *AuthChallenge
	challengeMu        sync.Mutex // serializes access to challenge digest state
	primeMu            sync.Mutex // single-flights the initial cold-wake auth handshake
	hostKey            string     // target:port, keys the shared per-device limiter
	conn               net.Conn
	bufferPool         sync.Pool
	UseTLS             bool
	InsecureSkipVerify bool
	PinnedCert         string
	tlsConfig          *tls.Config
}

// minTimeout is the floor applied to Parameters.Timeout. Declaring a Timeout
// field on Target would shadow the embedded http.Client.Timeout and leave
// requests with no deadline, so the promoted field is set directly.
const minTimeout = 30 * time.Second

// hostSems stores process-wide semaphores keyed by endpoint host:port.
// Each semaphore caps concurrent WSMAN Posts to that endpoint.
var (
	hostSemsMu sync.Mutex
	hostSems   = make(map[string]*hostLimiter)
)

// hostLimiter is a per-endpoint semaphore with a count of the Posts using it.
type hostLimiter struct {
	sem  chan struct{}
	refs int
}

// acquireHostSem returns the shared per-endpoint concurrency limiter.
// The same host:port gets the same semaphore across all Target instances.
func acquireHostSem(endpoint string) chan struct{} {
	hostSemsMu.Lock()
	defer hostSemsMu.Unlock()

	limiter, ok := hostSems[endpoint]
	if !ok {
		limiter = &hostLimiter{sem: make(chan struct{}, maxConcurrentPosts)}
		hostSems[endpoint] = limiter
	}

	limiter.refs++

	return limiter.sem
}

// releaseHostSem drops the entry once no Post is using the endpoint, so the map
// does not keep stale addresses (for example after a DHCP address change).
func releaseHostSem(endpoint string) {
	hostSemsMu.Lock()
	defer hostSemsMu.Unlock()

	limiter, ok := hostSems[endpoint]
	if !ok {
		return
	}

	limiter.refs--
	if limiter.refs <= 0 {
		delete(hostSems, endpoint)
	}
}

// isTransientTransportError reports whether err is a recoverable transport error.
func isTransientTransportError(err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, io.EOF) ||
		errors.Is(err, syscall.ECONNRESET) ||
		errors.Is(err, syscall.ECONNABORTED) ||
		errors.Is(err, syscall.EPIPE) {
		return true
	}

	var errno syscall.Errno
	if errors.As(err, &errno) && (errno == wsaeConnReset || errno == wsaeConnAborted) {
		return true
	}

	errText := strings.ToLower(err.Error())

	return strings.Contains(errText, "connection reset by peer") ||
		strings.Contains(errText, "bad record mac") ||
		strings.Contains(errText, "tls: internal error") ||
		strings.Contains(errText, "broken pipe")
}

// retryOnTransient retries req on transient transport errors until ctx expires.
func (t *Target) retryOnTransient(ctx context.Context, msgBody []byte, origReq *http.Request, origErr error) (*http.Response, error) {
	if !isTransientTransportError(origErr) {
		return nil, origErr
	}

	authenticated := t.useDigest && origReq.Header.Get("Authorization") != ""

	lastErr := origErr

	for attempt := 1; attempt <= maxTransientRetries; attempt++ {
		// Give the ME a moment to recover; retrying immediately usually fails again.
		time.Sleep(retryDelay)

		// Overall budget guard: the shared context caps the whole Post, so stop here
		// once it has expired instead of starting another attempt.
		if ctx.Err() != nil {
			return nil, lastErr
		}

		retryReq, reqErr := t.newPostRequest(ctx, msgBody)
		if reqErr != nil {
			return nil, reqErr
		}

		retryReq.Header = origReq.Header.Clone()

		if authenticated {
			auth, ok, authErr := t.digestAuthHeader()
			if authErr != nil {
				return nil, fmt.Errorf("failed digest auth: %w", authErr)
			}

			if ok {
				retryReq.Header.Set("Authorization", auth)
			}
		}

		res, err := t.Do(retryReq)
		if err == nil {
			return res, nil
		}

		lastErr = err

		logrus.WithError(err).Warnf("wsman transient retry %d/%d failed", attempt, maxTransientRetries)

		if !isTransientTransportError(err) {
			return nil, err
		}
	}

	return nil, lastErr
}

func NewWsman(cp Parameters) *Target {
	path := WSManPath
	port := NonTLSPort

	if cp.UseTLS {
		port = TLSPort
	}

	protocol := "http"
	if port == TLSPort {
		protocol = "https"
	}

	res := &Target{
		endpoint:           protocol + "://" + cp.Target + ":" + port + path,
		username:           cp.Username,
		password:           cp.Password,
		useDigest:          cp.UseDigest,
		logAMTMessages:     cp.LogAMTMessages,
		UseTLS:             cp.UseTLS,
		InsecureSkipVerify: cp.SelfSignedAllowed,
		conn:               cp.Connection,
		tlsConfig:          cp.TlsConfig,
		Client:             http.Client{Timeout: max(minTimeout, cp.Timeout)},
		hostKey:            cp.Target + ":" + port,
	}

	if cp.Transport == nil {
		// Use CIRATransport for CIRA APF tunnel connections
		if cp.IsCIRA && cp.CIRAManager != nil {
			res.Transport = NewCIRATransport(cp.CIRAManager, cp.LogAMTMessages)
		} else {
			// Standard HTTP transport setup
			var config *tls.Config
			if len(cp.PinnedCert) > 0 {
				// check if pinnedCert is not null and not empty
				config = &tls.Config{
					InsecureSkipVerify:    cp.SelfSignedAllowed,
					VerifyPeerCertificate: pinnedCertVerifier(cp.PinnedCert),
				}
			} else {
				if res.tlsConfig != nil {
					config = res.tlsConfig
				} else {
					config = &tls.Config{InsecureSkipVerify: cp.SelfSignedAllowed}

					if cp.AllowInsecureCipherSuites {
						defaultCipherSuites := tls.CipherSuites()
						config.CipherSuites = make([]uint16, 0, len(defaultCipherSuites)+3)

						for _, suite := range defaultCipherSuites {
							config.CipherSuites = append(config.CipherSuites, suite.ID)
						}
						// add the weak cipher suites
						config.CipherSuites = append(config.CipherSuites,
							tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
							tls.TLS_RSA_WITH_AES_128_CBC_SHA,
							tls.TLS_RSA_WITH_AES_256_CBC_SHA,
						)
					}
				}
			}
			// keep DisableKeepAlives false to reuse connections; true makes every request log in again.
			res.Transport = &http.Transport{
				MaxIdleConns:      maxConcurrentPosts,
				IdleConnTimeout:   idleConnTimeout,
				DisableKeepAlives: false,
				TLSClientConfig:   config,
			}
		}
	} else {
		res.Transport = cp.Transport
	}

	if res.useDigest {
		res.challenge = &AuthChallenge{Username: res.username, Password: res.password}
	}

	return res
}

func (t *Target) IsAuthenticated() bool {
	t.challengeMu.Lock()
	defer t.challengeMu.Unlock()

	return t.challenge != nil && t.challenge.Realm != ""
}

func (t *Target) GetServerCertificate() (*tls.Certificate, error) {
	httpTransport, ok := t.Transport.(*http.Transport)
	if !ok {
		return nil, errors.New("transport does not support TLSClientConfig")
	}

	if httpTransport.TLSClientConfig == nil {
		return nil, errors.New("TLSClientConfig is nil")
	}

	// Clone the transport's TLS config so the capture callback below is scoped
	// to this one throwaway dial.
	tlsConfig := httpTransport.TLSClientConfig.Clone()

	// Create a custom DialTLS to capture the server certificate
	capturedCert := &tls.Certificate{}
	tlsConfig.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
		if len(rawCerts) > 0 {
			cert, err := x509.ParseCertificate(rawCerts[0])
			if err != nil {
				return err
			}

			*capturedCert = tls.Certificate{
				Certificate: [][]byte{cert.Raw},
			}
		}

		return nil
	}

	// undo what we did in the constructor to get the endpoint (host and port)
	nohttps := strings.Replace(t.endpoint, "https://", "", 1)
	nohttps = strings.Replace(nohttps, "/wsman", "", 1)

	conn, err := tls.Dial("tcp", nohttps, tlsConfig)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	if len(capturedCert.Certificate) == 0 {
		return nil, errors.New("no server certificate captured")
	}

	return capturedCert, nil
}

// digestAuthHeader builds a Digest Authorization header under lock.
func (t *Target) digestAuthHeader() (auth string, ok bool, err error) {
	t.challengeMu.Lock()
	defer t.challengeMu.Unlock()

	if t.challenge.Realm == "" {
		return "", false, nil
	}

	auth, err = t.challenge.authorize("POST", "/wsman")
	if err != nil {
		return "", false, err
	}

	return auth, true, nil
}

// primeChallenge parses a 401 challenge and resets nc counter when nonce changes.
func (t *Target) primeChallenge(header string) error {
	t.challengeMu.Lock()
	defer t.challengeMu.Unlock()

	prevNonce := t.challenge.Nonce

	if err := t.challenge.parseChallenge(header); err != nil {
		return err
	}

	if t.challenge.Nonce != prevNonce {
		t.challenge.NonceCount = 0
	}

	return nil
}

func (t *Target) newPostRequest(ctx context.Context, msgBody []byte) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", t.endpoint, bytes.NewReader(msgBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", ContentType)
	// GetBody enables auto-retry on stale keep-alive connections.
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(msgBody)), nil
	}

	return req, nil
}

// Post overrides http.Client's Post method.
func (t *Target) Post(msg string) (response []byte, err error) {
	// Cap the whole Post, queue wait and retries included. Targets built outside
	// NewWsman carry no timeout, so fall back to the same floor it applies.
	budget := t.Timeout
	if budget <= 0 {
		budget = minTimeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), budget)
	defer cancel()

	// Bound concurrent Posts to match AMT firmware connection capacity.
	// hostKey is empty for redirection targets from NewWsmanTCP, which do not Post.
	if t.hostKey != "" {
		sem := acquireHostSem(t.hostKey)

		select {
		case sem <- struct{}{}:
			defer func() {
				<-sem
				releaseHostSem(t.hostKey)
			}()
		case <-ctx.Done():
			// Give up instead of queueing forever; the caller has likely gone away.
			releaseHostSem(t.hostKey)

			return nil, fmt.Errorf("timed out waiting for a connection slot to %s: %w", t.hostKey, ctx.Err())
		}
	}

	// Single-flight the first authentication handshake when digest is enabled.
	if t.useDigest && !t.IsAuthenticated() {
		t.primeMu.Lock()
		if t.IsAuthenticated() {
			t.primeMu.Unlock()
		} else {
			defer t.primeMu.Unlock()
		}
	}

	msgBody := []byte(msg)

	req, err := t.newPostRequest(ctx, msgBody)
	if err != nil {
		return nil, err
	}

	if t.username != "" && t.password != "" {
		if t.useDigest {
			auth, ok, authErr := t.digestAuthHeader()
			if authErr != nil {
				return nil, fmt.Errorf("failed digest auth: %w", authErr)
			}

			if ok {
				req.Header.Set("Authorization", auth)
			}
		} else {
			req.SetBasicAuth(t.username, t.password)
		}
	}

	if t.logAMTMessages {
		logrus.Trace(msg)
	}

	res, err := t.Do(req)
	if err != nil {
		res, err = t.retryOnTransient(ctx, msgBody, req, err)
		if err != nil {
			return nil, err
		}
	}

	if t.useDigest && res.StatusCode == 401 {
		_, _ = io.Copy(io.Discard, res.Body)
		_ = res.Body.Close()

		if parseErr := t.primeChallenge(res.Header.Get("WWW-Authenticate")); parseErr != nil {
			return nil, parseErr
		}

		auth, _, authErr := t.digestAuthHeader()
		if authErr != nil {
			return nil, fmt.Errorf("failed digest auth: %w", authErr)
		}

		req, err = t.newPostRequest(ctx, msgBody)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", auth)

		res, err = t.Do(req)
		if err != nil {
			res, err = t.retryOnTransient(ctx, msgBody, req, err)
			if err != nil {
				return nil, err
			}
		}
	}

	defer res.Body.Close()

	response, err = io.ReadAll(res.Body)

	if t.logAMTMessages {
		logrus.Trace(string(response))
	}

	if err != nil && err.Error() != io.EOF.Error() {
		return nil, err
	}

	if res.StatusCode == 400 {
		amterr := amterror.DecodeAMTErrorString(string(response))

		return nil, amterr
	}

	if res.StatusCode >= 401 {
		errPostResponse := errors.New("wsman.Client post received")

		return nil, fmt.Errorf("%w: %v\n%v", errPostResponse, res.Status, string(response))
	}

	return response, nil
}

// ProxyURL sets proxy address for the underlying Transport if supported.
func (t *Target) ProxyURL(proxyStr string) (err error) {
	// check if c.Transport is *http.Transport, otherwise currently it is not supported
	_, ok := t.Transport.(*http.Transport)
	if !ok {
		return errors.New("transport does not support proxy")
	}

	// check if proxy parsing failed or check if scheme is not nil
	proxyURL, err := url.Parse(proxyStr)
	if err != nil || (proxyURL != nil && proxyURL.Scheme == "") {
		return errors.New("unknown URL Scheme")
	}

	t.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)

	return nil
}
