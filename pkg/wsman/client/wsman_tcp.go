package client

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"net"
	"sync"
)

func NewWsmanTCP(cp Parameters) *Target {
	port := RedirectionNonTLSPort
	if cp.UseTLS {
		port = RedirectionTLSPort
	}

	return &Target{
		endpoint:           cp.Target + ":" + port,
		username:           cp.Username,
		password:           cp.Password,
		useDigest:          cp.UseDigest,
		logAMTMessages:     cp.LogAMTMessages,
		challenge:          &AuthChallenge{},
		UseTLS:             cp.UseTLS,
		InsecureSkipVerify: cp.SelfSignedAllowed,
		PinnedCert:         cp.PinnedCert,
		conn:               cp.Connection,
		bufferPool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 4096) // Adjust size according to your needs.
			},
		},
	}
}

// Connect establishes a TCP connection to the endpoint specified in the Target struct.
func (t *Target) Connect() error {
	var err error
	// already connected and connection has been provided
	if t.conn != nil {
		return nil
	}

	if t.UseTLS {
		// check if pinnedCert is not null and not empty
		var config *tls.Config
		if len(t.PinnedCert) > 0 {
			config = &tls.Config{
				InsecureSkipVerify: t.InsecureSkipVerify,
				VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
					for _, rawCert := range rawCerts {
						cert, err := x509.ParseCertificate(rawCert)
						if err != nil {
							return err
						}

						// Compare the current certificate with the pinned certificate
						sha256Fingerprint := sha256.Sum256(cert.Raw)
						if hex.EncodeToString(sha256Fingerprint[:]) == t.PinnedCert {
							return nil // Success: The certificate matches the pinned certificate
						}
					}

					return fmt.Errorf("certificate pinning failed")
				},
			}
		} else {
			config = &tls.Config{InsecureSkipVerify: t.InsecureSkipVerify}
		}

		t.conn, err = tls.Dial("tcp", t.endpoint, config)
	} else {
		t.conn, err = net.Dial("tcp", t.endpoint)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", t.endpoint, err)
	}

	return nil
}

// Send sends data to the connected TCP endpoint in the Target struct.
func (t *Target) Send(data []byte) error {
	if t.conn == nil {
		return fmt.Errorf("no active connection")
	}

	_, err := t.conn.Write(data)
	if err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	return nil
}

// Receive reads data from the connected TCP endpoint in the Target struct.
func (t *Target) Receive() ([]byte, error) {
	if t.conn == nil {
		return nil, fmt.Errorf("no active connection")
	}

	tmp := t.bufferPool.Get().([]byte)
	defer t.bufferPool.Put(tmp)

	n, err := t.conn.Read(tmp)
	if err != nil {
		return nil, err
	}

	return append([]byte(nil), tmp[:n]...), nil
}

// CloseConnection cleanly closes the TCP connection.
func (t *Target) CloseConnection() error {
	if t.conn == nil {
		return fmt.Errorf("no active connection to close")
	}

	err := t.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	t.conn = nil

	return nil
}
