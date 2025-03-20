/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/amterror"
)

const (
	ContentType           = "application/soap+xml; charset=utf-8"
	NSWSMAN               = "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"
	NSWSMID               = "http://schemas.dmtf.org/wbem/wsman/identity/1/wsmanidentity.xsd"
	TLSPort               = "16993"
	NonTLSPort            = "16992"
	RedirectionTLSPort    = "16995"
	RedirectionNonTLSPort = "16994"
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
	conn               net.Conn
	bufferPool         sync.Pool
	UseTLS             bool
	InsecureSkipVerify bool
	PinnedCert         string
	tlsConfig          *tls.Config
}

const timeout = 10 * time.Second

func NewWsman(cp Parameters) *Target {
	path := "/wsman"
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
		tlsConfig:          cp.TlsConfig,
	}

	res.Timeout = timeout

	if cp.Transport == nil {
		// check if pinnedCert is not null and not empty
		var config *tls.Config
		if len(cp.PinnedCert) > 0 {
			config = &tls.Config{
				InsecureSkipVerify: cp.SelfSignedAllowed,
				VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
					for _, rawCert := range rawCerts {
						cert, err := x509.ParseCertificate(rawCert)
						if err != nil {
							return err
						}

						// Compare the current certificate with the pinned certificate
						sha256Fingerprint := sha256.Sum256(cert.Raw)
						if hex.EncodeToString(sha256Fingerprint[:]) == cp.PinnedCert {
							return nil // Success: The certificate matches the pinned certificate
						}
					}

					return fmt.Errorf("certificate pinning failed")
				},
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

		res.Transport = &http.Transport{
			MaxIdleConns:      10,
			IdleConnTimeout:   30 * time.Second,
			DisableKeepAlives: true,
			TLSClientConfig:   config,
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
	return t.challenge != nil && t.challenge.Realm != ""
}

func (t *Target) GetServerCertificate() (*tls.Certificate, error) {
	httpTransport, ok := t.Transport.(*http.Transport)
	if !ok {
		return nil, errors.New("transport does not support TLSClientConfig")
	}

	tlsConfig := httpTransport.TLSClientConfig
	if tlsConfig == nil {
		return nil, errors.New("TLSClientConfig is nil")
	}

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

// Post overrides http.Client's Post method.
func (t *Target) Post(msg string) (response []byte, err error) {
	msgBody := []byte(msg)

	var auth string

	bodyReader := bytes.NewReader(msgBody)

	req, err := http.NewRequest("POST", t.endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	if t.username != "" && t.password != "" {
		if t.useDigest {
			auth, err = t.challenge.authorize("POST", "/wsman")
			if err != nil {
				return nil, fmt.Errorf("failed digest auth %w", err)
			}

			if t.challenge.Realm != "" {
				req.Header.Set("Authorization", auth)
			}
		} else {
			req.SetBasicAuth(t.username, t.password)
		}
	}

	req.Header.Add("content-type", ContentType)

	if t.logAMTMessages {
		logrus.Trace(msg)
	}

	res, err := t.Do(req)
	if err != nil {
		return nil, err
	}

	if t.useDigest && res.StatusCode == 401 {
		if err := t.challenge.parseChallenge(res.Header.Get("WWW-Authenticate")); err != nil {
			return nil, err
		}

		auth, err = t.challenge.authorize("POST", "/wsman")
		if err != nil {
			return nil, fmt.Errorf("failed digest auth %w", err)
		}

		bodyReader = bytes.NewReader(msgBody)

		req, err = http.NewRequest("POST", t.endpoint, bodyReader)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", auth)
		req.Header.Add("content-type", ContentType)

		res, err = t.Do(req)
		if err != nil && err.Error() != io.EOF.Error() {
			return nil, err
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
