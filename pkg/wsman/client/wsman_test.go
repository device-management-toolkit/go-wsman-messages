/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

const (
	testMsg      = "<SampleRequest>Request</SampleRequest>"
	testResponse = "<SampleResponse>OK</SampleResponse>"
)

func TestNewClient(t *testing.T) {
	cp := Parameters{
		Target:            "example.com",
		Username:          "user",
		Password:          "password",
		UseDigest:         false,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}
	expectedTarget := "http://example.com:16992/wsman"

	client := NewWsman(cp)

	if client.endpoint != expectedTarget {
		t.Errorf("Expected endpoint to be %s, but got %s", cp.Target, client.endpoint)
	}

	if client.username != cp.Username {
		t.Errorf("Expected username to be %s, but got %s", cp.Username, client.username)
	}

	if client.password != cp.Password {
		t.Errorf("Expected password to be %s, but got %s", cp.Password, client.password)
	}

	if client.useDigest != cp.UseDigest {
		t.Errorf("Expected useDigest to be %v, but got %v", cp.UseDigest, client.useDigest)
	}
}

func TestNewClient_TLS(t *testing.T) {
	expectedTarget := "https://example.com:16993/wsman"
	cp := Parameters{
		Target:            "example.com",
		Username:          "user",
		Password:          "password",
		UseDigest:         false,
		UseTLS:            true,
		SelfSignedAllowed: true,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)

	if client.endpoint != expectedTarget {
		t.Errorf("Expected endpoint to be %s, but got %s", cp.Target, client.endpoint)
	}

	if client.username != cp.Username {
		t.Errorf("Expected username to be %s, but got %s", cp.Username, client.username)
	}

	if client.password != cp.Password {
		t.Errorf("Expected password to be %s, but got %s", cp.Password, client.password)
	}

	if client.useDigest != cp.UseDigest {
		t.Errorf("Expected useDigest to be %v, but got %v", cp.UseDigest, client.useDigest)
	}
}

func TestNewClient_WithTLSConfig(t *testing.T) {
	expectedTarget := "https://example2.com:16993/wsman"
	tlsConfig := &tls.Config{InsecureSkipVerify: true}

	cp := Parameters{
		Target:            "example2.com",
		Username:          "user",
		Password:          "password",
		UseDigest:         false,
		UseTLS:            true,
		SelfSignedAllowed: true,
		LogAMTMessages:    false,
		TlsConfig:         tlsConfig,
	}

	client := NewWsman(cp)

	if client.endpoint != expectedTarget {
		t.Errorf("Expected endpoint to be %s, but got %s", expectedTarget, client.endpoint)
	}

	if client.tlsConfig != tlsConfig {
		t.Errorf("Expected tlsConfig to be the same instance as the one passed, but got a different instance")
	}

	transport, ok := client.Transport.(*http.Transport)
	if !ok {
		t.Fatal("Expected client.Transport to be of type *http.Transport")
	}

	if transport.TLSClientConfig != tlsConfig {
		t.Errorf("Expected TLSClientConfig to be the provided tlsConfig, but got a different instance")
	}
}

// The timeout must land on the embedded http.Client, not on a field that
// shadows it -- a shadowed Timeout leaves requests with no deadline at all.
func TestNewClient_Timeout(t *testing.T) {
	tests := []struct {
		name     string
		timeout  time.Duration
		expected time.Duration
	}{
		{"unset falls back to the minimum", 0, minTimeout},
		{"below the minimum is raised", 3 * time.Second, minTimeout},
		{"above the minimum is honored", 45 * time.Second, 45 * time.Second},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewWsman(Parameters{Target: "example3.com", Timeout: tt.timeout})

			if client.Timeout != tt.expected {
				t.Errorf("Expected http.Client.Timeout to be %v, but got %v", tt.expected, client.Timeout)
			}
		})
	}
}

func TestClient_Post(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(testResponse))
		if err != nil {
			t.Errorf("Unexpected error during write: %v", err)
		}
	}))
	defer ts.Close()

	cp := Parameters{
		Target:            ts.URL,
		Username:          "user",
		Password:          "password",
		UseDigest:         false,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)
	msg := testMsg

	client.endpoint = ts.URL

	response, err := client.Post(msg)
	if err != nil {
		t.Errorf("Unexpected error during POST: %v", err)
	}

	expectedResponse := testResponse
	if string(response) != expectedResponse {
		t.Errorf("Expected response to be %s, but got %s", expectedResponse, response)
	}
}

func newMockDigestAuthHandler(username, password string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Digest ") {
			// Check for the correct username and password in the Authorization header
			if strings.Contains(authHeader, `username="`+username+`"`) { // &&strings.Contains(authHeader, `uri="`+r.URL.RequestURI()
				handler.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		} else {
			// Simulate a server requesting digest authentication with required fields
			w.Header().Set("WWW-Authenticate", `Digest realm="example2.com", nonce="mock-nonce", qop="auth", opaque="opaque-data", algorithm=MD5`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

func TestClient_PostWithDigestAuth(t *testing.T) {
	// Use a simple digest auth implementation for testing purposes
	ts := httptest.NewServer(newMockDigestAuthHandler("user", "password", http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(testResponse))
		if err != nil {
			t.Errorf("Unexpected error during write: %v", err)
		}
	}))))

	defer ts.Close()

	cp := Parameters{
		Target:            ts.URL,
		Username:          "user",
		Password:          "password",
		UseDigest:         true,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)
	msg := testMsg

	client.endpoint = ts.URL

	response, err := client.Post(msg)
	if err != nil {
		t.Errorf("Unexpected error during POST with digest auth: %v", err)
	}

	expectedResponse := testResponse
	if string(response) != expectedResponse {
		t.Errorf("Expected response to be %s, but got %s", expectedResponse, response)
	}
}

func TestClient_PostWithDigestAuthUnauthorized(t *testing.T) {
	ts := httptest.NewServer(newMockDigestAuthHandler("user", "password", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusOK)
	})))

	defer ts.Close()

	cp := Parameters{
		Target:            ts.URL,
		Username:          "wronguser",
		Password:          "wrongpassword",
		UseDigest:         true,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)
	msg := testMsg

	client.endpoint = ts.URL

	_, err := client.Post(msg)
	if err == nil {
		t.Error("Expected error during POST with wrong digest auth credentials, but got nil")
	}
}

func TestClient_PostWithBasicAuth(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "user" || password != "password" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)

			return
		}

		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(testResponse))
		if err != nil {
			t.Errorf("Unexpected error during write: %v", err)
		}
	}))

	defer ts.Close()

	cp := Parameters{
		Target:            ts.URL,
		Username:          "user",
		Password:          "password",
		UseDigest:         false,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)
	msg := testMsg

	client.endpoint = ts.URL

	response, err := client.Post(msg)
	if err != nil {
		t.Errorf("Unexpected error during POST with basic auth: %v", err)
	}

	expectedResponse := testResponse

	if string(response) != expectedResponse {
		t.Errorf("Expected response to be %s, but got %s", expectedResponse, response)
	}
}

func TestClient_PostUnauthorized(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}))

	defer ts.Close()

	cp := Parameters{
		Target:            ts.URL,
		Username:          "wronguser",
		Password:          "wrongpassword",
		UseDigest:         false,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}
	client := NewWsman(cp)
	msg := testMsg

	client.endpoint = ts.URL

	_, err := client.Post(msg)
	if err == nil {
		t.Error("Expected error during POST with wrong credentials, but got nil")
	}
}

func TestClient_PostInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusInternalServerError)

		_, err := w.Write([]byte("Internal Server Error"))
		if err != nil {
			t.Errorf("Unexpected error during write: %v", err)
		}
	}))

	defer ts.Close()

	cp := Parameters{
		Target:            ts.URL,
		Username:          "user",
		Password:          "password",
		UseDigest:         false,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)
	msg := testMsg

	client.endpoint = ts.URL

	_, err := client.Post(msg)
	if err == nil {
		t.Error("Expected error during POST with invalid response, but got nil")
	}
}

func TestClient_PostWithDigestBlankRealm(t *testing.T) {
	ts := httptest.NewServer(newMockDigestAuthHandler("user", "password", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Digest ") {
			// Simulate internal server error
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			// Simulate a server requesting digest authentication with required fields
			w.Header().Set("WWW-Authenticate", `Digest realm="example.com", nonce="mock-nonce", qop="auth", opaque="opaque-data", algorithm=MD5`)
			w.WriteHeader(http.StatusUnauthorized)
		}
	})))
	defer ts.Close()

	cp := Parameters{
		Target:            ts.URL,
		Username:          "user",
		Password:          "password",
		UseDigest:         true,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)
	client.challenge.Realm = ""
	msg := testMsg

	client.endpoint = ts.URL

	_, err := client.Post(msg)
	if err == nil {
		t.Error("Expected error during POST with wrong digest auth credentials, but got nil")
	}

	if !strings.Contains(err.Error(), "500 Internal Server Error") {
		t.Error("Wsman client should not send digest on initial challenges")
	}
}

func TestClient_ProxyUrlTransport(t *testing.T) {
	cp := Parameters{
		Target:            "example.com",
		Username:          "user",
		Password:          "password",
		UseDigest:         true,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)

	err := client.ProxyURL("http://localhost:3128")
	if err != nil {
		t.Error("Failed to set proxy on proper Transport")
	}
}

func TestClient_InvalidProxyUrlGoodTransport(t *testing.T) {
	cp := Parameters{
		Target:            "example.com",
		Username:          "user",
		Password:          "password",
		UseDigest:         true,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)

	err := client.ProxyURL("localhost")
	if err == nil {
		t.Error("Failed to detect invalid proxy url")
	}
}

// inline struct for mock roundtripper.
type rt struct{}

func (*rt) RoundTrip(r *http.Request) (*http.Response, error) {
	return nil, nil
}

func TestClient_SimpleRountripper(t *testing.T) {
	cp := Parameters{
		Target:            "example.com",
		Username:          "user",
		Password:          "password",
		UseDigest:         true,
		UseTLS:            false,
		SelfSignedAllowed: false,
		LogAMTMessages:    false,
	}

	mockrt := rt{}
	client := NewWsman(cp)
	client.Transport = &mockrt

	err := client.ProxyURL("http://localhost:3128")
	if err == nil {
		t.Error("Failed to detect proper transport")
	}
}

func TestClient_GetServerCertificate(t *testing.T) {
	// Setting up a mock server to simulate a TLS handshake and provide a certificate
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	cp := Parameters{
		Target:            ts.URL,
		Username:          "user",
		Password:          "password",
		UseDigest:         false,
		UseTLS:            true,
		SelfSignedAllowed: true,
		LogAMTMessages:    false,
	}

	client := NewWsman(cp)
	client.endpoint = ts.URL

	cert, err := client.GetServerCertificate()
	if err != nil {
		t.Errorf("Unexpected error during GetServerCertificate: %v", err)
	}

	// Check that a certificate was indeed captured
	if cert == nil || len(cert.Certificate) == 0 {
		t.Error("Expected a server certificate, but none was captured")
	}
}

// TestClient_GetServerCertificate_LeavesTransportConfigUnchanged - certificate capture is scoped to a cloned config and does not mutate the Target's shared TLS config.
func TestClient_GetServerCertificate_LeavesTransportConfigUnchanged(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	cp := Parameters{
		Target:            ts.URL,
		Username:          "user",
		Password:          "password",
		UseTLS:            true,
		SelfSignedAllowed: true,
	}

	client := NewWsman(cp)
	client.endpoint = ts.URL

	httpTransport, ok := client.Transport.(*http.Transport)
	if !ok {
		t.Fatal("expected *http.Transport")
	}

	// Baseline: the freshly-built transport has no per-certificate callback.
	if httpTransport.TLSClientConfig.VerifyPeerCertificate != nil {
		t.Fatal("precondition failed: transport already has a VerifyPeerCertificate")
	}

	if _, err := client.GetServerCertificate(); err != nil {
		t.Fatalf("Unexpected error during GetServerCertificate: %v", err)
	}

	// The capture callback must have been scoped to the throwaway dial only;
	// the shared transport config must be unchanged.
	if httpTransport.TLSClientConfig.VerifyPeerCertificate != nil {
		t.Error("GetServerCertificate modified the shared transport TLS config; " +
			"VerifyPeerCertificate should remain unset")
	}
}

const testHostKey = "10.0.0.1:16992"

var errNotTransient = errors.New("wsman: not a transport problem")

// wsaError builds the error shape a Windows socket read failure arrives in:
// net.OpError wrapping os.SyscallError wrapping the raw WSA code.
func wsaError(code syscall.Errno) error {
	return &net.OpError{
		Op:  "read",
		Err: &os.SyscallError{Syscall: "wsarecv", Err: code},
	}
}

func TestIsTransientTransportError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{"nil", nil, false},
		{"unrelated", errNotTransient, false},
		{"io.EOF", io.EOF, true},
		{"ECONNRESET", syscall.ECONNRESET, true},
		{"ECONNABORTED", syscall.ECONNABORTED, true},
		{"EPIPE", syscall.EPIPE, true},
		{"wrapped EOF", &net.OpError{Op: "read", Err: io.EOF}, true},
		{"WSAECONNRESET", wsaError(wsaeConnReset), true},
		{"WSAECONNABORTED", wsaError(wsaeConnAborted), true},
		{"other errno", wsaError(syscall.Errno(10061)), false},
		{"reset text", errors.New("read tcp: connection reset by peer"), true},
		{"bad record mac text", errors.New("local error: tls: bad record MAC"), true},
		{"tls internal text", errors.New("remote error: tls: internal error"), true},
		{"broken pipe text", errors.New("write tcp: broken pipe"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTransientTransportError(tt.err); got != tt.expected {
				t.Errorf("isTransientTransportError(%v) = %v, want %v", tt.err, got, tt.expected)
			}
		})
	}
}

func TestAcquireHostSemIsSharedAndReleased(t *testing.T) {
	first := acquireHostSem(testHostKey)
	second := acquireHostSem(testHostKey)

	if first != second {
		t.Error("expected both callers to share one semaphore for the same host")
	}

	if cap(first) != maxConcurrentPosts {
		t.Errorf("expected semaphore cap %d, got %d", maxConcurrentPosts, cap(first))
	}

	hostSemsMu.Lock()
	refs := hostSems[testHostKey].refs
	hostSemsMu.Unlock()

	if refs != 2 {
		t.Errorf("expected 2 refs, got %d", refs)
	}

	releaseHostSem(testHostKey)
	releaseHostSem(testHostKey)

	hostSemsMu.Lock()
	_, stillThere := hostSems[testHostKey]
	hostSemsMu.Unlock()

	if stillThere {
		t.Error("expected the entry to be dropped once no Post is using it")
	}
}

func TestReleaseHostSemUnknownKeyIsSafe(t *testing.T) {
	releaseHostSem("never-acquired:16992")
}

func TestPostConcurrencyIsCapped(t *testing.T) {
	var (
		inFlight int32
		peak     int32
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := atomic.AddInt32(&inFlight, 1)

		for {
			old := atomic.LoadInt32(&peak)
			if now <= old || atomic.CompareAndSwapInt32(&peak, old, now) {
				break
			}
		}

		time.Sleep(20 * time.Millisecond)
		atomic.AddInt32(&inFlight, -1)

		_, _ = w.Write([]byte(testResponse))
	}))

	defer ts.Close()

	client := NewWsman(Parameters{Target: ts.URL})
	client.endpoint = ts.URL

	var wg sync.WaitGroup

	for range 8 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if _, err := client.Post(testMsg); err != nil {
				t.Errorf("unexpected error during POST: %v", err)
			}
		}()
	}

	wg.Wait()

	if got := atomic.LoadInt32(&peak); got > maxConcurrentPosts {
		t.Errorf("expected at most %d concurrent posts, saw %d", maxConcurrentPosts, got)
	}

	hostSemsMu.Lock()
	remaining := len(hostSems)
	hostSemsMu.Unlock()

	if remaining != 0 {
		t.Errorf("expected hostSems to be empty after all posts finished, got %d entries", remaining)
	}
}

func TestRetryOnTransientRecovers(t *testing.T) {
	var calls int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)

		_, _ = w.Write([]byte(testResponse))
	}))

	defer ts.Close()

	client := NewWsman(Parameters{Target: ts.URL})
	client.endpoint = ts.URL

	ctx, cancel := context.WithTimeout(context.Background(), minTimeout)
	defer cancel()

	req, err := client.newPostRequest(ctx, []byte(testMsg))
	if err != nil {
		t.Fatalf("unexpected error building request: %v", err)
	}

	res, err := client.retryOnTransient(ctx, []byte(testMsg), req, io.EOF)
	if err != nil {
		t.Fatalf("expected the retry to recover, got %v", err)
	}

	defer res.Body.Close()

	if atomic.LoadInt32(&calls) != 1 {
		t.Errorf("expected 1 retry attempt, got %d", atomic.LoadInt32(&calls))
	}
}

func TestRetryOnTransientGivesUpAfterMaxAttempts(t *testing.T) {
	var calls int32

	// Hijack and close without answering, so every attempt sees a dropped connection.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)

		conn, _, err := w.(http.Hijacker).Hijack()
		if err == nil {
			_ = conn.Close()
		}
	}))

	defer ts.Close()

	client := NewWsman(Parameters{Target: ts.URL})
	client.endpoint = ts.URL

	ctx, cancel := context.WithTimeout(context.Background(), minTimeout)
	defer cancel()

	req, err := client.newPostRequest(ctx, []byte(testMsg))
	if err != nil {
		t.Fatalf("unexpected error building request: %v", err)
	}

	if _, err = client.retryOnTransient(ctx, []byte(testMsg), req, io.EOF); err == nil {
		t.Fatal("expected an error once every attempt failed")
	}

	if atomic.LoadInt32(&calls) != maxTransientRetries {
		t.Errorf("expected %d attempts, got %d", maxTransientRetries, atomic.LoadInt32(&calls))
	}
}

func TestRetryOnTransientSkipsOtherErrors(t *testing.T) {
	client := NewWsman(Parameters{Target: "10.0.0.1"})

	ctx, cancel := context.WithTimeout(context.Background(), minTimeout)
	defer cancel()

	req, err := client.newPostRequest(ctx, []byte(testMsg))
	if err != nil {
		t.Fatalf("unexpected error building request: %v", err)
	}

	_, err = client.retryOnTransient(ctx, []byte(testMsg), req, errNotTransient)
	if !errors.Is(err, errNotTransient) {
		t.Errorf("expected the original error back, got %v", err)
	}
}

// The deadline must cover the whole Post, so an expired context stops the retry
// loop instead of starting another attempt.
func TestRetryOnTransientStopsWhenContextExpired(t *testing.T) {
	var calls int32

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)

		_, _ = w.Write([]byte(testResponse))
	}))

	defer ts.Close()

	client := NewWsman(Parameters{Target: ts.URL})
	client.endpoint = ts.URL

	ctx, cancel := context.WithCancel(context.Background())

	req, err := client.newPostRequest(ctx, []byte(testMsg))
	if err != nil {
		t.Fatalf("unexpected error building request: %v", err)
	}

	cancel()

	_, err = client.retryOnTransient(ctx, []byte(testMsg), req, io.EOF)
	if !errors.Is(err, io.EOF) {
		t.Errorf("expected the last error back, got %v", err)
	}

	if atomic.LoadInt32(&calls) != 0 {
		t.Errorf("expected no attempt after the deadline passed, got %d", atomic.LoadInt32(&calls))
	}
}
