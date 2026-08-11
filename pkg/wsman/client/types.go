/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// Parameters struct defines the connection settings for wsman client.
type Parameters struct {
	Target                    string
	Username                  string
	Password                  string
	UseDigest                 bool
	UseTLS                    bool
	SelfSignedAllowed         bool
	LogAMTMessages            bool
	Transport                 http.RoundTripper
	IsRedirection             bool
	PinnedCert                string
	Connection                net.Conn
	TlsConfig                 *tls.Config
	AllowInsecureCipherSuites bool
	IsCIRA                    bool               // Flag to indicate CIRA APF tunnel connection
	CIRAManager               CIRAChannelManager // Manager for CIRA channel operations
	// Timeout bounds the whole HTTP request (connect, write, read) on the
	// Target returned by NewWsman. It is a real time.Duration, so write
	// 30*time.Second, not a bare 30 -- an untyped constant is nanoseconds and
	// will be raised to the floor below. Values under 10s, including the zero
	// value, are raised to 10s. Note that the CIRA APF transport applies its
	// own 60s budget internally, which this timeout can cut short.
	Timeout time.Duration
}
