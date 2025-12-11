/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/apf"
)

const (
	maxCIRAChannels     = 6
	ciraTimeout         = 60 * time.Second
	ciraChannelOpenPort = 16992
)

// CIRAChannelManager is an interface that must be implemented by the console
// to manage APF channels through the CIRA tunnel.
type CIRAChannelManager interface {
	// RegisterAPFChannel creates and registers a new APF channel.
	RegisterAPFChannel() CIRAChannel
	// UnregisterAPFChannel removes an APF channel.
	UnregisterAPFChannel(senderChannel uint32)
	// GetConnection returns the underlying network connection for writes.
	GetConnection() net.Conn
}

// CIRAChannel represents an APF channel that can send/receive data.
type CIRAChannel interface {
	// GetSenderChannel returns our channel ID.
	GetSenderChannel() uint32
	// GetRecipientChannel returns the device's channel ID (set after open confirmation).
	GetRecipientChannel() uint32
	// SetRecipientChannel sets the device's channel ID.
	SetRecipientChannel(channel uint32)
	// GetTXWindow returns the current transmit window.
	GetTXWindow() uint32
	// SetTXWindow sets the transmit window.
	SetTXWindow(window uint32)
	// AddTXWindow adds to the transmit window.
	AddTXWindow(bytes uint32)
	// WaitForOpen waits for the channel open result.
	WaitForOpen(timeout time.Duration) error
	// ReceiveData receives data from the channel with timeout.
	ReceiveData(timeout time.Duration) ([]byte, error)
	// ReceiveWindowAdjust receives a window adjust notification with timeout.
	ReceiveWindowAdjust(timeout time.Duration) (uint32, error)
	// IsClosed returns true if the channel is closed.
	IsClosed() bool
}

// CIRATransport implements http.RoundTripper for CIRA APF tunnels.
// It routes HTTP requests through an APF (Application Protocol Forwarder)
// connection to an AMT device.
type CIRATransport struct {
	manager     CIRAChannelManager
	channelSem  chan struct{}
	timeout     time.Duration
	logMessages bool
}

// NewCIRATransport creates a new CIRA transport that routes HTTP requests
// through an APF tunnel connection.
func NewCIRATransport(manager CIRAChannelManager, logMessages bool) *CIRATransport {
	return &CIRATransport{
		manager:     manager,
		channelSem:  make(chan struct{}, maxCIRAChannels),
		timeout:     ciraTimeout,
		logMessages: logMessages,
	}
}

// RoundTrip executes a single HTTP request through the CIRA APF tunnel.
func (c *CIRATransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Acquire semaphore slot (blocks if 6 channels in use)
	select {
	case c.channelSem <- struct{}{}:
		defer func() { <-c.channelSem }()
	case <-time.After(c.timeout):
		return nil, errors.New("timeout waiting for available CIRA channel slot")
	}

	// Register and open APF channel
	channel := c.manager.RegisterAPFChannel()
	defer c.manager.UnregisterAPFChannel(channel.GetSenderChannel())

	// Send APF_CHANNEL_OPEN
	if err := c.sendChannelOpen(channel); err != nil {
		return nil, fmt.Errorf("failed to send channel open: %w", err)
	}

	// Wait for channel open confirmation (routed by tunnel)
	if err := channel.WaitForOpen(c.timeout); err != nil {
		return nil, fmt.Errorf("failed to open APF channel: %w", err)
	}

	// Build HTTP request manually to match AMT's expected format
	// (similar to MPS Node.js implementation)
	reqBytes, err := c.buildHTTPRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to build HTTP request: %w", err)
	}

	if c.logMessages {
		logrus.Tracef("CIRA TX HTTP Request:\n%s", string(reqBytes))
	}

	// Send via APF_CHANNEL_DATA
	if err := c.sendData(channel, reqBytes); err != nil {
		return nil, fmt.Errorf("failed to send request via APF: %w", err)
	}

	// Read response from channel (data routed by tunnel)
	respBytes, err := c.readResponse(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to read APF response: %w", err)
	}

	if c.logMessages {
		logrus.Tracef("CIRA RX HTTP Response:\n%s", string(respBytes))
	}

	// Send channel close
	c.sendChannelClose(channel)

	// Parse HTTP response
	resp, err := http.ReadResponse(bufio.NewReader(bytes.NewReader(respBytes)), req)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTTP response: %w", err)
	}

	return resp, nil
}

// buildHTTPRequest constructs an HTTP request in the format AMT expects.
// This matches the format used by the MPS Node.js implementation.
func (c *CIRATransport) buildHTTPRequest(req *http.Request) ([]byte, error) {
	var buf bytes.Buffer

	// Read body if present
	var body []byte

	if req.Body != nil {
		var err error

		body, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		// Reset body for potential re-reads
		req.Body = io.NopCloser(bytes.NewReader(body))
	}

	// Request line: POST /wsman HTTP/1.1
	path := req.URL.Path
	if path == "" {
		path = "/wsman"
	}

	buf.WriteString(fmt.Sprintf("%s %s HTTP/1.1\r\n", req.Method, path))

	// Write headers - Authorization first if present (important for digest auth)
	if auth := req.Header.Get("Authorization"); auth != "" {
		buf.WriteString(fmt.Sprintf("Authorization: %s\r\n", auth))
	}

	// Host header
	host := req.Host
	if host == "" {
		host = req.URL.Host
	}

	buf.WriteString(fmt.Sprintf("Host: %s\r\n", host))

	// Content-Type header
	if ct := req.Header.Get("Content-Type"); ct != "" {
		buf.WriteString(fmt.Sprintf("Content-Type: %s\r\n", ct))
	}

	// Content-Length header
	buf.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(body)))

	// End of headers
	buf.WriteString("\r\n")

	// Body
	buf.Write(body)

	return buf.Bytes(), nil
}

// sendChannelOpen sends APF_CHANNEL_OPEN to the device.
func (c *CIRATransport) sendChannelOpen(channel CIRAChannel) error {
	openMsg := apf.ChannelOpen(int(channel.GetSenderChannel()))

	conn := c.manager.GetConnection()
	if err := conn.SetWriteDeadline(time.Now().Add(c.timeout)); err != nil {
		return err
	}

	_, err := conn.Write(openMsg.Bytes())

	return err
}

// sendData sends data via APF_CHANNEL_DATA, respecting flow control.
func (c *CIRATransport) sendData(channel CIRAChannel, data []byte) error {
	conn := c.manager.GetConnection()
	offset := 0

	for offset < len(data) {
		// Wait for transmit window if needed
		for channel.GetTXWindow() == 0 {
			bytesToAdd, err := channel.ReceiveWindowAdjust(c.timeout)
			if err != nil {
				return fmt.Errorf("timeout waiting for window adjust: %w", err)
			}

			channel.AddTXWindow(bytesToAdd)
		}

		// Calculate chunk size (respect window and reasonable packet size)
		chunkSize := len(data) - offset
		if uint32(chunkSize) > channel.GetTXWindow() {
			chunkSize = int(channel.GetTXWindow())
		}

		if chunkSize > apf.LME_RX_WINDOW_SIZE {
			chunkSize = apf.LME_RX_WINDOW_SIZE
		}

		chunk := data[offset : offset+chunkSize]

		// Build and send APF_CHANNEL_DATA
		dataMsg := apf.BuildChannelDataBytes(channel.GetRecipientChannel(), chunk)

		if err := conn.SetWriteDeadline(time.Now().Add(c.timeout)); err != nil {
			return err
		}

		if _, err := conn.Write(dataMsg); err != nil {
			return fmt.Errorf("failed to send CHANNEL_DATA: %w", err)
		}

		channel.AddTXWindow(^uint32(chunkSize - 1)) // Subtract chunkSize
		offset += chunkSize
	}

	return nil
}

// readResponse reads the HTTP response from the channel.
func (c *CIRATransport) readResponse(channel CIRAChannel) ([]byte, error) {
	var response bytes.Buffer

	conn := c.manager.GetConnection()
	bytesReceived := uint32(0)

	deadline := time.Now().Add(c.timeout)

	for time.Now().Before(deadline) {
		if channel.IsClosed() {
			if response.Len() > 0 {
				return response.Bytes(), nil
			}

			return nil, errors.New("channel closed by remote")
		}

		// Try to receive data from channel (non-blocking with short timeout)
		data, err := channel.ReceiveData(100 * time.Millisecond)
		if err != nil {
			// Timeout is expected, continue waiting
			continue
		}

		response.Write(data)
		bytesReceived += uint32(len(data))

		// Send WINDOW_ADJUST to allow more data
		if bytesReceived >= apf.LME_RX_WINDOW_SIZE/2 {
			adjustMsg := apf.BuildChannelWindowAdjustBytes(channel.GetRecipientChannel(), bytesReceived)

			_ = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
			_, _ = conn.Write(adjustMsg)

			bytesReceived = 0
		}

		// Check if we have a complete HTTP response
		if isCompleteHTTPResponse(response.Bytes()) {
			return response.Bytes(), nil
		}
	}

	if response.Len() > 0 {
		return response.Bytes(), nil
	}

	return nil, errors.New("timeout waiting for response")
}

// isCompleteHTTPResponse checks if we have received a complete HTTP response.
func isCompleteHTTPResponse(data []byte) bool {
	// Look for end of SOAP envelope which indicates complete response
	if bytes.Contains(data, []byte("</a:Envelope>")) {
		return true
	}

	// Also check for HTML responses (error pages)
	if bytes.Contains(data, []byte("</html>")) || bytes.Contains(data, []byte("</HTML>")) {
		return true
	}

	// Check for chunked transfer encoding completion
	if bytes.Contains(data, []byte("\r\n0\r\n\r\n")) {
		return true
	}

	// For non-chunked responses, try to parse Content-Length
	headerEnd := bytes.Index(data, []byte("\r\n\r\n"))
	if headerEnd == -1 {
		return false
	}

	headers := string(data[:headerEnd])

	// Look for Content-Length header
	clStart := bytes.Index([]byte(headers), []byte("Content-Length:"))
	if clStart == -1 {
		clStart = bytes.Index([]byte(headers), []byte("content-length:"))
	}

	if clStart != -1 {
		clEnd := bytes.Index([]byte(headers[clStart:]), []byte("\r\n"))
		if clEnd != -1 {
			clValue := headers[clStart+15 : clStart+clEnd]

			var contentLength int

			if _, err := fmt.Sscanf(clValue, " %d", &contentLength); err == nil {
				bodyStart := headerEnd + 4
				bodyLen := len(data) - bodyStart

				return bodyLen >= contentLength
			}
		}
	}

	return false
}

// sendChannelClose sends APF_CHANNEL_CLOSE to the device.
func (c *CIRATransport) sendChannelClose(channel CIRAChannel) {
	if channel.GetRecipientChannel() == 0 {
		return
	}

	closeMsg := apf.BuildChannelCloseBytes(channel.GetRecipientChannel())

	conn := c.manager.GetConnection()
	_ = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	_, _ = conn.Write(closeMsg)

	logrus.Debugf("CIRA channel closed: recipient=%d", channel.GetRecipientChannel())
}
