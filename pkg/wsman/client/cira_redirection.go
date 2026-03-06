/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/apf"
)

const (
	ciraRedirectionPort    = 16994
	ciraRedirectionTimeout = 60 * time.Second
	windowAdjustThreshold  = apf.LME_RX_WINDOW_SIZE / 2
)

// CIRARedirectionTarget implements the WSMan interface for raw binary streaming
// over a long-lived APF channel through a CIRA tunnel.
type CIRARedirectionTarget struct {
	manager       CIRAChannelManager
	channel       CIRAChannel
	timeout       time.Duration
	bytesReceived uint32
}

// NewCIRARedirectionTarget creates a new CIRA redirection target.
func NewCIRARedirectionTarget(manager CIRAChannelManager) *CIRARedirectionTarget {
	return &CIRARedirectionTarget{
		manager: manager,
		timeout: ciraRedirectionTimeout,
	}
}

// Connect opens an APF channel to port 16994 on the device through the CIRA tunnel.
func (c *CIRARedirectionTarget) Connect() error {
	channel := c.manager.RegisterAPFChannel()

	logrus.Infof("CIRA Redirection: opening channel %d to port %d", channel.GetSenderChannel(), ciraRedirectionPort)

	// Send APF_CHANNEL_OPEN for redirection port
	openMsg := apf.ChannelOpenPort(int(channel.GetSenderChannel()), ciraRedirectionPort)

	if err := c.manager.WriteToConnection(openMsg.Bytes()); err != nil {
		c.manager.UnregisterAPFChannel(channel.GetSenderChannel())

		return fmt.Errorf("failed to send channel open: %w", err)
	}

	// Wait for channel open confirmation
	if err := channel.WaitForOpen(c.timeout); err != nil {
		c.manager.UnregisterAPFChannel(channel.GetSenderChannel())

		return fmt.Errorf("failed to open APF channel for redirection: %w", err)
	}

	c.channel = channel

	logrus.Infof("CIRA Redirection: channel %d opened, recipient=%d, txWindow=%d",
		channel.GetSenderChannel(), channel.GetRecipientChannel(), channel.GetTXWindow())

	return nil
}

// Send sends raw bytes via APF_CHANNEL_DATA with TX window flow control.
func (c *CIRARedirectionTarget) Send(data []byte) error {
	if c.channel == nil {
		return errors.New("no active CIRA redirection channel")
	}

	offset := 0

	for offset < len(data) {
		// Wait for transmit window if needed
		for c.channel.GetTXWindow() == 0 {
			bytesToAdd, err := c.channel.ReceiveWindowAdjust(c.timeout)
			if err != nil {
				return fmt.Errorf("timeout waiting for window adjust: %w", err)
			}

			c.channel.AddTXWindow(bytesToAdd)
		}

		// Calculate chunk size
		chunkSize := len(data) - offset
		if uint32(chunkSize) > c.channel.GetTXWindow() {
			chunkSize = int(c.channel.GetTXWindow())
		}

		if chunkSize > apf.LME_RX_WINDOW_SIZE {
			chunkSize = apf.LME_RX_WINDOW_SIZE
		}

		chunk := data[offset : offset+chunkSize]

		dataMsg := apf.BuildChannelDataBytes(c.channel.GetRecipientChannel(), chunk)

		if err := c.manager.WriteToConnection(dataMsg); err != nil {
			return fmt.Errorf("failed to send redirection data: %w", err)
		}

		c.channel.SubtractTXWindow(uint32(chunkSize))
		offset += chunkSize
	}

	return nil
}

// Receive reads raw bytes from the APF channel as they arrive.
func (c *CIRARedirectionTarget) Receive() ([]byte, error) {
	if c.channel == nil {
		return nil, errors.New("no active CIRA redirection channel")
	}

	data, err := c.channel.ReceiveData(c.timeout)
	if err != nil {
		return nil, err
	}

	c.bytesReceived += uint32(len(data))

	// Send WINDOW_ADJUST when threshold reached
	if c.bytesReceived >= windowAdjustThreshold {
		adjustMsg := apf.BuildChannelWindowAdjustBytes(c.channel.GetRecipientChannel(), c.bytesReceived)

		if err := c.manager.WriteToConnection(adjustMsg); err != nil {
			logrus.Warnf("CIRA Redirection: failed to send window adjust: %v", err)
		}

		c.bytesReceived = 0
	}

	return data, nil
}

// CloseConnection sends APF_CHANNEL_CLOSE and unregisters the channel.
func (c *CIRARedirectionTarget) CloseConnection() error {
	if c.channel == nil {
		return nil
	}

	if c.channel.GetRecipientChannel() != 0 {
		closeMsg := apf.BuildChannelCloseBytes(c.channel.GetRecipientChannel())

		if err := c.manager.WriteToConnection(closeMsg); err != nil {
			logrus.Warnf("CIRA Redirection: failed to send channel close: %v", err)
		}
	}

	c.manager.UnregisterAPFChannel(c.channel.GetSenderChannel())
	c.channel = nil

	return nil
}

// Post is not used for redirection.
func (c *CIRARedirectionTarget) Post(_ string) ([]byte, error) {
	return nil, errors.New("Post not supported for CIRA redirection")
}

// IsAuthenticated returns false — redirection handles auth at the protocol level.
func (c *CIRARedirectionTarget) IsAuthenticated() bool {
	return false
}

// GetServerCertificate is not supported for CIRA redirection.
func (c *CIRARedirectionTarget) GetServerCertificate() (*tls.Certificate, error) {
	return nil, errors.New("GetServerCertificate not supported for CIRA redirection")
}
