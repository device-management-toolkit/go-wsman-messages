package client

import (
	"crypto/tls"
	"net"
	"net/http"
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
}
