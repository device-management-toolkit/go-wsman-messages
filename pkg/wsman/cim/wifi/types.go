/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

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
		WiFiPortGetResponse       WiFiPort
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
		RequestStateChange_OUTPUT common.ReturnValue
	}

	PullResponse struct {
		XMLName               xml.Name                       `xml:"PullResponse"`
		EndpointSettingsItems []WiFiEndpointSettingsResponse `xml:"Items>CIM_WiFiEndpointSettings"`
		WiFiPortItems         []WiFiPort                     `xml:"Items>CIM_WiFiPort"`
	}

	WiFiEndpointSettingsResponse struct {
		XMLName              xml.Name             `xml:"CIM_WiFiEndpointSettings"`
		AuthenticationMethod AuthenticationMethod // AuthenticationMethod shall specify the 802.11 authentication method used when the settings are applied.
		BSSType              BSSType              // BSSType shall indicate the Basic Service Set (BSS) Type that shall be used when the settings are applied.
		ElementName          string               // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		EncryptionMethod     EncryptionMethod     // EncryptionMethod shall specify the 802.11 encryption method used when the settings are applied.
		InstanceID           string               // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		Priority             int                  // Priority shall indicate the priority of the instance among all WiFiEndpointSettings instances.
		SSID                 string               // SSID shall indicate the Service Set Identifier (SSID) that shall be used when the settings are applied to a WiFiEndpoint. An SSID identifies a wireless network.
	}

	WiFiPort struct {
		XMLName                 xml.Name       `xml:"CIM_WiFiPort"`
		LinkTechnology          LinkTechnology // An enumeration of the types of links. When set to 1 ("Other"), the related property OtherLinkTechnology contains a string description of the type of link.
		DeviceID                string         // An address or other identifying information to uniquely name the LogicalDevice.
		CreationClassName       string         // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		SystemName              string         // The scoping System's Name.
		SystemCreationClassName string         // The scoping System's CreationClassName.
		ElementName             string         // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information.
		HealthState             HealthState    // Indicates the current health of the element.
		EnabledState            EnabledState   // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		PortType                PortType       // PortType shall contain the specific 802.11 operating mode that is currently enabled on the Port.
		PermanentAddress        string         // IEEE 802 EUI-48 MAC address, formatted as twelve hexadecimal digits (for example, "010203040506"), with each pair representing one of the six octets of the MAC address in "canonical" bit order.
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
	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int
	// EncryptionMethod shall specify the 802.11 encryption method used when the settings are applied.
	EncryptionMethod int
	// HealthState shall indicate the current health of the element.
	HealthState int
	// LinkTechnology shall contain an enumeration of the types of links. When set to 1 ("Other"), the related property OtherLinkTechnology shall contain a string description of the type of link.
	LinkTechnology int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int
	// ReturnValue is a 16-bit unsigned integer enumeration that specifies the completion status of the operation.
	ReturnValue int
	// PortType shall contain the specific 802.11 operating mode that is currently enabled on the Port.
	PortType int
)
