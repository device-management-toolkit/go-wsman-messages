/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package cryptographiccapabilities

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUTS
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName           xml.Name `xml:"Body"`
		GetResponse       CryptographicCapabilities
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	PullResponse struct {
		XMLName                        xml.Name                    `xml:"PullResponse"`
		CryptographicCapabilitiesItems []CryptographicCapabilities `xml:"Items>AMT_CryptographicCapabilities"`
	}

	// CryptographicCapabilities represents AMT cryptographic capabilities information.
	CryptographicCapabilities struct {
		XMLName              xml.Name `xml:"AMT_CryptographicCapabilities"`
		ElementName          string   `xml:"ElementName"`
		HardwareAcceleration int      `xml:"HardwareAcceleration"`
		InstanceID           string   `xml:"InstanceID"`
	}
)
