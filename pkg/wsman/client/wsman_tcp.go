package client

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	// defaultBufferPoolSize controls the per-Receive temporary buffer size used to read from the socket.
	// Larger values reduce syscalls and fragmentation for KVM streaming payloads.
	defaultBufferPoolSize = 64 * 1024

	// tcpSocketBufferSize sets OS-level socket read/write buffer hints for throughput.
	tcpSocketBufferSize = 256 * 1024

	// defaultKeepAlive configures TCP keepalive probe interval on the dialer.
	defaultKeepAlive = 30 * time.Second

	// defaultReadTimeout sets the maximum time to wait for read operations.
	defaultReadTimeout = 30 * time.Second

	// defaultWriteTimeout sets the maximum time to wait for write operations.
	defaultWriteTimeout = 30 * time.Second
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
		bufferPool: sync.Pool{
			New: func() interface{} {
				// Larger buffer to reduce read syscalls and frame fragmentation for KVM streams
				return make([]byte, defaultBufferPoolSize)
			},
		},
	}
}

// Connect establishes a TCP connection to the endpoint specified in the Target struct.
func (t *Target) Connect() error {
	// Use a Dialer so we can enable TCP keep-alives and TCP_NODELAY for lower latency.
	d := &net.Dialer{KeepAlive: defaultKeepAlive}

	if t.UseTLS {
		// Build TLS config with optional pinning
		var config *tls.Config
		if len(t.PinnedCert) > 0 {
			config = &tls.Config{
				InsecureSkipVerify: t.InsecureSkipVerify,
				VerifyPeerCertificate: func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
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

		// Establish plain TCP first to set socket options
		plainConn, err := d.Dial("tcp", t.endpoint)
		if err != nil {
			return fmt.Errorf("failed to connect to %s: %w", t.endpoint, err)
		}

		if tcp, ok := plainConn.(*net.TCPConn); ok {
			// Best-effort; ignore error to avoid failing connection setup
			_ = tcp.SetNoDelay(true)
			_ = tcp.SetReadBuffer(tcpSocketBufferSize)
			_ = tcp.SetWriteBuffer(tcpSocketBufferSize)
		}

		tlsConn := tls.Client(plainConn, config)
		if err := tlsConn.Handshake(); err != nil {
			_ = plainConn.Close()

			return fmt.Errorf("TLS handshake failed with %s: %w", t.endpoint, err)
		}

		t.conn = tlsConn

		return nil
	}

	// Non-TLS path
	c, err := d.Dial("tcp", t.endpoint)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", t.endpoint, err)
	}

	if tcp, ok := c.(*net.TCPConn); ok {
		_ = tcp.SetNoDelay(true)
		_ = tcp.SetReadBuffer(tcpSocketBufferSize)
		_ = tcp.SetWriteBuffer(tcpSocketBufferSize)
	}

	t.conn = c

	return nil
}

// Send sends data to the connected TCP endpoint in the Target struct.
func (t *Target) Send(data []byte) error {
	if t.conn == nil {
		return fmt.Errorf("no active connection")
	}

	// Optimize: Only set deadline if enough time has passed or it's not set
	if err := t.setWriteDeadlineIfNeeded(); err != nil {
		return fmt.Errorf("failed to set write deadline: %w", err)
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

	// Optimize: Only set deadline if enough time has passed or it's not set
	if err := t.setReadDeadlineIfNeeded(); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}

	tmp := t.bufferPool.Get().([]byte)
	defer t.bufferPool.Put(tmp) //nolint:staticcheck // changing the argument to be pointer-like to avoid allocations caused issues.

	n, err := t.conn.Read(tmp)
	if err != nil {
		return nil, err
	}

	return append([]byte(nil), tmp[:n]...), nil
}

// setReadDeadlineIfNeeded sets read deadline only when needed to reduce syscall overhead
func (t *Target) setReadDeadlineIfNeeded() error {
	t.deadlineMutex.Lock()
	defer t.deadlineMutex.Unlock()
	
	now := time.Now()
	// Only set deadline if it's been more than half the timeout period since last set
	// or if it's never been set (zero value)
	const deadlineRefreshInterval = defaultReadTimeout / 2
	
	if t.lastReadDeadlineSet.IsZero() || now.Sub(t.lastReadDeadlineSet) > deadlineRefreshInterval {
		newDeadline := now.Add(defaultReadTimeout)
		if err := t.conn.SetReadDeadline(newDeadline); err != nil {
			return err
		}
		t.lastReadDeadlineSet = now
	}
	
	return nil
}

// setWriteDeadlineIfNeeded sets write deadline only when needed to reduce syscall overhead  
func (t *Target) setWriteDeadlineIfNeeded() error {
	t.deadlineMutex.Lock()
	defer t.deadlineMutex.Unlock()
	
	now := time.Now()
	// Only set deadline if it's been more than half the timeout period since last set
	// or if it's never been set (zero value)
	const deadlineRefreshInterval = defaultWriteTimeout / 2
	
	if t.lastWriteDeadlineSet.IsZero() || now.Sub(t.lastWriteDeadlineSet) > deadlineRefreshInterval {
		newDeadline := now.Add(defaultWriteTimeout)
		if err := t.conn.SetWriteDeadline(newDeadline); err != nil {
			return err
		}
		t.lastWriteDeadlineSet = now
	}
	
	return nil
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
	
	// Reset deadline tracking when connection is closed
	t.deadlineMutex.Lock()
	t.lastReadDeadlineSet = time.Time{}
	t.lastWriteDeadlineSet = time.Time{}
	t.deadlineMutex.Unlock()

	return nil
}
