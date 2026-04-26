/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiendpointcapabilities

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
		XMLName                             xml.Name `xml:"Body"`
		WiFiEndpointCapabilitiesGetResponse WiFiEndpointCapabilities
		EnumerateResponse                   common.EnumerateResponse
		PullResponse                        PullResponse
	}

	PullResponse struct {
		XMLName                   xml.Name                   `xml:"PullResponse"`
		EndpointCapabilitiesItems []WiFiEndpointCapabilities `xml:"Items>CIM_WiFiEndpointCapabilities"`
	}

	WiFiEndpointCapabilities struct {
		XMLName                        xml.Name               `xml:"CIM_WiFiEndpointCapabilities"`
		ElementName                    string                 `xml:"ElementName,omitempty"`                    // A user-friendly name for the object.
		InstanceID                     string                 `xml:"InstanceID,omitempty"`                     // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		SupportedAuthenticationMethods []AuthenticationMethod `xml:"SupportedAuthenticationMethods,omitempty"` // Supported 802.11 authentication methods.
		SupportedEncryptionMethods     []EncryptionMethod     `xml:"SupportedEncryptionMethods,omitempty"`     // Supported 802.11 encryption methods.
	}
)

type (
	// AuthenticationMethod shall specify the 802.11 authentication method used when the settings are applied.
	AuthenticationMethod int
	// EncryptionMethod shall specify the 802.11 encryption method used when the settings are applied.
	EncryptionMethod int
)
