/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbootreason

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUT.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name               `xml:"Body"`
		GetResponse       HostBootReasonResponse `xml:"IPS_HostBootReason"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
	}

	HostBootReasonResponse struct {
		XMLName         xml.Name `xml:"IPS_HostBootReason"`
		ElementName     string   `xml:"ElementName,omitempty"`
		InstanceID      string   `xml:"InstanceID,omitempty"`
		PreviousSxState int      `xml:"PreviousSxState,omitempty"`
		Reason          int      `xml:"Reason,omitempty"`
		ReasonDetails   string   `xml:"ReasonDetails,omitempty"`
	}

	PullResponse struct {
		XMLName             xml.Name                 `xml:"PullResponse"`
		HostBootReasonItems []HostBootReasonResponse `xml:"Items>IPS_HostBootReason"`
	}
)
