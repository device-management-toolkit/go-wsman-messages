/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ipv6portsettings

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
		XMLName           xml.Name         `xml:"Body"`
		GetResponse       IPv6PortSettings `xml:"IPS_IPv6PortSettings"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
	}

	IPv6PortSettings struct {
		XMLName              xml.Name `xml:"IPS_IPv6PortSettings"`
		CurrentDefaultRouter string   `xml:"CurrentDefaultRouter,omitempty"`
		CurrentPrimaryDNS    string   `xml:"CurrentPrimaryDNS,omitempty"`
		CurrentSecondaryDNS  string   `xml:"CurrentSecondaryDNS,omitempty"`
		DefaultRouter        string   `xml:"DefaultRouter,omitempty"`
		ElementName          string   `xml:"ElementName,omitempty"`
		IPv6Address          string   `xml:"IPv6Address,omitempty"`
		InstanceID           string   `xml:"InstanceID,omitempty"`
		InterfaceIDType      int      `xml:"InterfaceIDType"`
		ManualInterfaceID    string   `xml:"ManualInterfaceID,omitempty"`
		PrimaryDNS           string   `xml:"PrimaryDNS,omitempty"`
		SecondaryDNS         string   `xml:"SecondaryDNS,omitempty"`
	}

	PullResponse struct {
		XMLName               xml.Name           `xml:"PullResponse"`
		IPv6PortSettingsItems []IPv6PortSettings `xml:"Items>IPS_IPv6PortSettings"`
	}
)
