/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportcapabilities

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUT
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                         xml.Name `xml:"Body"`
		WiFiPortCapabilitiesGetResponse WiFiPortCapabilities
		EnumerateResponse               common.EnumerateResponse
		PullResponse                    PullResponse
	}

	PullResponse struct {
		XMLName                   xml.Name               `xml:"PullResponse"`
		WiFiPortCapabilitiesItems []WiFiPortCapabilities `xml:"Items>CIM_WiFiPortCapabilities"`
	}

	WiFiPortCapabilities struct {
		XMLName            xml.Name `xml:"CIM_WiFiPortCapabilities"`
		ElementName        string   `xml:"ElementName,omitempty"`        // A user-friendly name for the object.
		InstanceID         string   `xml:"InstanceID,omitempty"`         // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		SupportedPortTypes []int    `xml:"SupportedPortTypes,omitempty"` // The supported 802.11 operating modes.
	}
)
