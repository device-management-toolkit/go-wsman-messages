/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

func TestExtractWSMANSummary(t *testing.T) {
	t.Parallel()

	request := `<s:Envelope><s:Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService</w:ResourceURI></s:Header></s:Envelope>`

	service, action := extractWSMANSummary(request)

	if service != "AMT_SetupAndConfigurationService" {
		t.Fatalf("expected service AMT_SetupAndConfigurationService, got %s", service)
	}

	if action != "Enumerate" {
		t.Fatalf("expected action Enumerate, got %s", action)
	}
}

func TestExtractWSMANSummaryResponse(t *testing.T) {
	t.Parallel()

	response := `<a:Envelope><a:Header><b:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</b:Action><c:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService</c:ResourceURI></a:Header></a:Envelope>`

	service, action := extractWSMANSummary(response)

	if service != "AMT_SetupAndConfigurationService" {
		t.Fatalf("expected service AMT_SetupAndConfigurationService, got %s", service)
	}

	if action != "GetResponse" {
		t.Fatalf("expected action GetResponse, got %s", action)
	}
}

func TestExtractWSMANSummaryNoData(t *testing.T) {
	t.Parallel()

	service, action := extractWSMANSummary("not xml")

	if service != "" || action != "" {
		t.Fatalf("expected empty service and action, got service=%s action=%s", service, action)
	}
}
