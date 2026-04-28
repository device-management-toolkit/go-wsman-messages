/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostipsettings

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
		XMLName           xml.Name       `xml:"Body"`
		GetResponse       HostIPSettings `xml:"IPS_HostIPSettings"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
	}

	HostIPSettings struct {
		XMLName     xml.Name `xml:"IPS_HostIPSettings"`
		DHCPEnabled bool     `xml:"DHCPEnabled"`
		ElementName string   `xml:"ElementName,omitempty"`
		InstanceID  string   `xml:"InstanceID,omitempty"`
	}

	PullResponse struct {
		XMLName             xml.Name         `xml:"PullResponse"`
		HostIPSettingsItems []HostIPSettings `xml:"Items>IPS_HostIPSettings"`
	}
)
