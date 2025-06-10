/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package client

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var tlsconfig = &tls.Config{}

func TestNewWsTransport(t *testing.T) {
	trans := NewWsTransport("wss://localhost/mps/ws/relay/webrelay.ashx", 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "pass", 16992, false, false, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}
}

func TestNewWsTransportBuildUrl(t *testing.T) {
	baseurl := "wss://localhost/mps/ws/relay/webrelay.ashx"

	trans := NewWsTransport(baseurl, 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "pass", 16992, false, false, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}

	url := trans.buildURL()
	if url == "" {
		t.Error("Failed to build url")
	}

	// second path
	trans = NewWsTransport(baseurl, 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "pass", 16992, true, true, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}

	url = trans.buildURL()
	if url == "" {
		t.Error("Failed to build url")
	}
}

var upgrader = websocket.Upgrader{}

func relayTester(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.RequestURI, "simulate_fail") {
		_, err := w.Write([]byte("Hello"))
		if err != nil {
			return
		}

		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	if strings.Contains(r.RequestURI, "simulate_close") {
		// wait for few ms before close
		timer := time.NewTimer(time.Duration(100) * time.Millisecond)
		<-timer.C
		c.Close()

		return
	}

	simulateDelay := strings.Contains(r.RequestURI, "simulate_delay")

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}

		if simulateDelay {
			okHTTPResponse := "HTTP/1.1 200 OK\r\nServer: dummy\r\n\r\nasdf"

			err = c.WriteMessage(mt, []byte(okHTTPResponse))
			if err != nil {
				break
			}

			timer := time.NewTimer(time.Duration(1000) * time.Millisecond)
			<-timer.C

			err = c.WriteMessage(mt, []byte("asdf\r\n\r\n"))
			if err != nil {
				break
			}

			break
		}

		if strings.HasPrefix(string(message), "GET") {
			// It is a GET request
			okHTTPResponse := "HTTP/1.1 200 OK\r\nServer: dummy\r\n\r\n<html></html>\r\n\r\n"

			err = c.WriteMessage(mt, []byte(okHTTPResponse))
			if err != nil {
				break
			}
		} else if strings.HasPrefix(string(message), "POST") {
			// It is a POST request
			okHTTPResponse := "HTTP/1.1 200 OK\r\nServer: dummy\r\n\r\n<a:Envelope></a:Envelope>\r\n\r\n"

			err = c.WriteMessage(mt, []byte(okHTTPResponse))
			if err != nil {
				break
			}
		} else {
			err = c.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	}
}

func TestNewWsTransportRoundtripBadParam(t *testing.T) {
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(relayTester))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	baseurl := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	trans := NewWsTransport(baseurl, 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "", 16992, false, false, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}

	defer trans.disconnectWebsocket()

	req := httptest.NewRequest("GET", "http://localhost", http.NoBody)

	defer req.Body.Close()

	_, err := trans.RoundTrip(req)
	if err == nil {
		t.Error("Roundtripper should failed")
	}
}

func TestNewWsTransportRoundtripBadUrl(t *testing.T) {
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(relayTester))
	defer s.Close()

	// bad url
	baseurl := "ws://localhot/"

	// Connect to the server
	trans := NewWsTransport(baseurl, 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "pass", 16992, false, false, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}

	defer trans.disconnectWebsocket()

	req := httptest.NewRequest("GET", "http://localhost", http.NoBody)

	defer req.Body.Close()

	_, err := trans.RoundTrip(req)
	if err == nil {
		t.Error("Roundtripper should failed")
	}
}

func TestNewWsTransportRoundtripGet(t *testing.T) {
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(relayTester))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	baseurl := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	trans := NewWsTransport(baseurl, 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "pass", 16992, false, false, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}

	defer trans.disconnectWebsocket()

	req := httptest.NewRequest("GET", "http://localhost", http.NoBody)

	resp, err := trans.RoundTrip(req)
	if err != nil {
		t.Error("Roundtripper failed")
	}

	if resp.StatusCode != 200 {
		t.Error("Dummy Get request failed")
	}
}

func TestNewWsTransportRoundtripPost(t *testing.T) {
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(relayTester))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	baseurl := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	trans := NewWsTransport(baseurl, 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "pass", 16992, false, false, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}

	defer trans.disconnectWebsocket()

	req := httptest.NewRequest("POST", "http://localhost", http.NoBody)

	resp, err := trans.RoundTrip(req)
	if err != nil {
		t.Error("Roundtripper failed")
	}

	if resp.StatusCode != 200 {
		t.Error("Dummy POST request failed")
	}
}

func TestNewWsTransportFailedConnection(t *testing.T) {
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(relayTester))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	baseurl := "ws" + strings.TrimPrefix(s.URL, "http") + "/simulate_fail"

	// Connect to the server
	trans := NewWsTransport(baseurl, 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "pass", 16992, false, false, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}

	defer trans.disconnectWebsocket()

	req := httptest.NewRequest("POST", "http://localhost", http.NoBody)

	_, err := trans.RoundTrip(req)
	if err != nil {
		t.Error("Roundtripper should not fail")
	}
}

func TestNewWsTransportCloseConnection(t *testing.T) {
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(relayTester))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	baseurl := "ws" + strings.TrimPrefix(s.URL, "http") + "/simulate_close"

	// Connect to the server
	trans := NewWsTransport(baseurl, 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "pass", 16992, false, false, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}

	defer trans.disconnectWebsocket()

	req := httptest.NewRequest("POST", "http://localhost", http.NoBody)

	_, err := trans.RoundTrip(req)
	if err == nil {
		t.Error("Roundtripper should fail")
	}
}

func TestNewWsTransportDelay(t *testing.T) {
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(relayTester))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	baseurl := "ws" + strings.TrimPrefix(s.URL, "http") + "/simulate_delay"

	// Connect to the server
	trans := NewWsTransport(baseurl, 1, "9b3ee6a0-c1dc-5546-f7f3-54b2039edfb9", "user", "pass", 16992, false, false, "token", tlsconfig)
	if trans == nil {
		t.Error("NewWSTransporter constructor fails")
	}

	defer trans.disconnectWebsocket()

	req := httptest.NewRequest("POST", "http://localhost", http.NoBody)

	_, err := trans.RoundTrip(req)
	if err != nil {
		t.Error("Roundtripper should not fail")
	}
}
