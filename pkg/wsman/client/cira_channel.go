/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"errors"
	"net"
	"sync"
	"time"
)

// APFChannel represents an active APF channel for CIRA connections.
// It implements the CIRAChannel interface.
type APFChannel struct {
	SenderChannel    uint32
	RecipientChannel uint32
	TXWindow         uint32
	DataChan         chan []byte // Incoming data from device
	OpenChan         chan error  // Signals channel open result
	WindowChan       chan uint32 // Window adjust notifications
	Closed           bool
	mu               sync.Mutex
}

// NewAPFChannel creates a new APF channel with the given sender channel ID.
func NewAPFChannel(senderChannel uint32) *APFChannel {
	return &APFChannel{
		SenderChannel: senderChannel,
		DataChan:      make(chan []byte, 100),
		OpenChan:      make(chan error, 1),
		WindowChan:    make(chan uint32, 10),
	}
}

// GetSenderChannel returns our channel ID.
func (c *APFChannel) GetSenderChannel() uint32 {
	return c.SenderChannel
}

// GetRecipientChannel returns the device's channel ID.
func (c *APFChannel) GetRecipientChannel() uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.RecipientChannel
}

// SetRecipientChannel sets the device's channel ID.
func (c *APFChannel) SetRecipientChannel(channel uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.RecipientChannel = channel
}

// GetTXWindow returns the current transmit window.
func (c *APFChannel) GetTXWindow() uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.TXWindow
}

// SetTXWindow sets the transmit window.
func (c *APFChannel) SetTXWindow(window uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.TXWindow = window
}

// AddTXWindow adds to the transmit window.
func (c *APFChannel) AddTXWindow(bytes uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.TXWindow += bytes
}

// SubtractTXWindow subtracts from the transmit window.
func (c *APFChannel) SubtractTXWindow(bytes uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if bytes > c.TXWindow {
		c.TXWindow = 0
	} else {
		c.TXWindow -= bytes
	}
}

// WaitForOpen waits for the channel open result with timeout.
func (c *APFChannel) WaitForOpen(timeout time.Duration) error {
	select {
	case err := <-c.OpenChan:
		return err
	case <-time.After(timeout):
		return errors.New("timeout waiting for channel open")
	}
}

// ReceiveData receives data from the channel with timeout.
func (c *APFChannel) ReceiveData(timeout time.Duration) ([]byte, error) {
	select {
	case data, ok := <-c.DataChan:
		if !ok {
			return nil, errors.New("channel closed")
		}

		return data, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout waiting for data")
	}
}

// ReceiveWindowAdjust receives a window adjust notification with timeout.
func (c *APFChannel) ReceiveWindowAdjust(timeout time.Duration) (uint32, error) {
	select {
	case bytes := <-c.WindowChan:
		return bytes, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout waiting for window adjust")
	}
}

// IsClosed returns true if the channel is closed.
func (c *APFChannel) IsClosed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.Closed
}

// SignalOpen signals that the channel has been opened (or failed to open).
func (c *APFChannel) SignalOpen(err error) {
	select {
	case c.OpenChan <- err:
	default:
	}
}

// SendData sends data to the channel's data buffer.
func (c *APFChannel) SendData(data []byte) {
	c.mu.Lock()
	closed := c.Closed
	c.mu.Unlock()

	if !closed {
		select {
		case c.DataChan <- data:
		default:
		}
	}
}

// SendWindowAdjust sends a window adjust notification.
func (c *APFChannel) SendWindowAdjust(bytes uint32) {
	select {
	case c.WindowChan <- bytes:
	default:
	}
}

// Close marks the channel as closed and closes the data channel.
func (c *APFChannel) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.Closed {
		c.Closed = true
		close(c.DataChan)
	}
}

// APFChannelStore manages APF channels for a CIRA connection.
// It implements the CIRAChannelManager interface.
type APFChannelStore struct {
	channels      map[uint32]*APFChannel
	mu            sync.Mutex
	nextChannelID uint32
	conn          net.Conn
}

// NewAPFChannelStore creates a new APF channel store.
func NewAPFChannelStore(conn net.Conn) *APFChannelStore {
	return &APFChannelStore{
		channels: make(map[uint32]*APFChannel),
		conn:     conn,
	}
}

// RegisterAPFChannel creates and registers a new APF channel.
// Implements CIRAChannelManager interface.
func (s *APFChannelStore) RegisterAPFChannel() CIRAChannel {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextChannelID++
	channelID := s.nextChannelID

	channel := NewAPFChannel(channelID)
	s.channels[channelID] = channel

	return channel
}

// UnregisterAPFChannel removes an APF channel.
// Implements CIRAChannelManager interface.
func (s *APFChannelStore) UnregisterAPFChannel(senderChannel uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if ch, ok := s.channels[senderChannel]; ok {
		ch.Close()
		delete(s.channels, senderChannel)
	}
}

// GetConnection returns the underlying network connection.
// Implements CIRAChannelManager interface.
func (s *APFChannelStore) GetConnection() net.Conn {
	return s.conn
}

// GetChannel retrieves an APF channel by sender channel ID.
func (s *APFChannelStore) GetChannel(senderChannel uint32) *APFChannel {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.channels[senderChannel]
}

// SetConnection updates the underlying network connection.
func (s *APFChannelStore) SetConnection(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.conn = conn
}

// CloseAll closes all channels in the store.
func (s *APFChannelStore) CloseAll() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ch := range s.channels {
		ch.Close()
	}

	s.channels = make(map[uint32]*APFChannel)
}
