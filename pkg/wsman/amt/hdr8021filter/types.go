/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hdr8021filter

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
		GetResponse       Hdr8021Filter
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	PullResponse struct {
		XMLName            xml.Name        `xml:"PullResponse"`
		Hdr8021FilterItems []Hdr8021Filter `xml:"Items>AMT_Hdr8021Filter"`
	}

	// Hdr8021Filter represents a single AMT 802.1 filter entry.
	Hdr8021Filter struct {
		XMLName                 xml.Name `xml:"AMT_Hdr8021Filter"`
		CreationClassName       string   `xml:"CreationClassName"`
		ElementName             string   `xml:"ElementName"`
		FilterDirection         int      `xml:"FilterDirection"`
		FilterEnabled           bool     `xml:"FilterEnabled"`
		Name                    string   `xml:"Name"`
		SystemCreationClassName string   `xml:"SystemCreationClassName"`
		SystemName              string   `xml:"SystemName"`
		VLANPriority            int      `xml:"VLANPriority"`
		VLANID                  int      `xml:"VLANID"`
	}
)
