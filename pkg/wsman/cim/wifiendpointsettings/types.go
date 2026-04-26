/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiendpointsettings

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
		XMLName                   xml.Name `xml:"Body"`
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
		RequestStateChange_OUTPUT common.ReturnValue
	}

	PullResponse struct {
		XMLName               xml.Name                       `xml:"PullResponse"`
		EndpointSettingsItems []WiFiEndpointSettingsResponse `xml:"Items>CIM_WiFiEndpointSettings"`
	}

	WiFiEndpointSettingsResponse struct {
		XMLName              xml.Name             `xml:"CIM_WiFiEndpointSettings"`
		AuthenticationMethod AuthenticationMethod // AuthenticationMethod shall specify the 802.11 authentication method used when the settings are applied.
		BSSType              BSSType              // BSSType shall indicate the Basic Service Set (BSS) Type that shall be used when the settings are applied.
		ElementName          string               // The user-friendly name for this instance of SettingData.
		EncryptionMethod     EncryptionMethod     // EncryptionMethod shall specify the 802.11 encryption method used when the settings are applied.
		InstanceID           string               // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		Priority             int                  // Priority shall indicate the priority of the instance among all WiFiEndpointSettings instances.
		SSID                 string               // SSID shall indicate the Service Set Identifier (SSID) that shall be used when the settings are applied to a WiFiEndpoint.
	}
)

// INPUT
// Request Types.
type (
	WiFiEndpointSettings_INPUT struct {
		XMLName              xml.Name `xml:"CIM_WiFiEndpointSettings"`
		H                    string   `xml:"xmlns:q,attr"`
		AuthenticationMethod AuthenticationMethod
		BSSType              BSSType
		ElementName          string
		EncryptionMethod     EncryptionMethod
		InstanceID           string
		Priority             int
		SSID                 string
	}
	WiFiEndpointSettingsRequest struct {
		XMLName xml.Name `xml:"h:WiFiEndpointSettingsInput"`
		H       string   `xml:"xmlns:q,attr"`
		// SettingData
		ElementName          string               `xml:"q:ElementName,omitempty"`
		InstanceID           string               `xml:"q:InstanceID,omitempty"`
		AuthenticationMethod AuthenticationMethod `xml:"q:AuthenticationMethod,omitempty"`
		EncryptionMethod     EncryptionMethod     `xml:"q:EncryptionMethod,omitempty"`
		SSID                 string               `xml:"q:SSID,omitempty"` // Max Length 32
		Priority             int                  `xml:"q:Priority,omitempty"`
		PSKPassPhrase        string               `xml:"q:PSKPassPhrase,omitempty"` // Min Length 8 Max Length 63
		BSSType              BSSType              `xml:"q:BSSType,omitempty"`
		Keys                 []string             `xml:"q:Keys,omitempty"` // OctetString ArrayType=Indexed Max Length 256
		KeyIndex             int                  `xml:"q:KeyIndex,omitempty"`
		PSKValue             int                  `xml:"q:PSKValue,omitempty"` // OctetString
	}
)

type (
	// AuthenticationMethod shall specify the 802.11 authentication method used when the settings are applied.
	AuthenticationMethod int
	// BSSType shall indicate the Basic Service Set (BSS) Type that shall be used when the settings are applied.
	BSSType int
	// EncryptionMethod shall specify the 802.11 encryption method used when the settings are applied.
	EncryptionMethod int
)
