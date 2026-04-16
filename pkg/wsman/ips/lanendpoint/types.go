/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package lanendpoint

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name    `xml:"Body"`
		GetResponse       LANEndpoint `xml:"IPS_LANEndpoint"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
	}

	LANEndpoint struct {
		XMLName                 xml.Name `xml:"IPS_LANEndpoint"`
		CreationClassName       string   `xml:"CreationClassName,omitempty"`
		EnabledState            int      `xml:"EnabledState"`
		LANType                 int      `xml:"LANType"`
		MACAddress              string   `xml:"MACAddress,omitempty"`
		Name                    string   `xml:"Name,omitempty"`
		ProtocolType            int      `xml:"ProtocolType"`
		RequestedState          int      `xml:"RequestedState"`
		SystemCreationClassName string   `xml:"SystemCreationClassName,omitempty"`
		SystemName              string   `xml:"SystemName,omitempty"`
	}

	PullResponse struct {
		XMLName          xml.Name      `xml:"PullResponse"`
		LANEndpointItems []LANEndpoint `xml:"Items>IPS_LANEndpoint"`
	}
)
