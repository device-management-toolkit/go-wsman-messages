/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiendpoint

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
		XMLName                 xml.Name `xml:"Body"`
		WiFiEndpointGetResponse WiFiEndpoint
		EnumerateResponse       common.EnumerateResponse
		PullResponse            PullResponse
	}

	PullResponse struct {
		XMLName       xml.Name       `xml:"PullResponse"`
		EndpointItems []WiFiEndpoint `xml:"Items>CIM_WiFiEndpoint"`
	}

	WiFiEndpoint struct {
		XMLName                 xml.Name       `xml:"CIM_WiFiEndpoint"`
		CreationClassName       string         `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		DeviceID                string         `xml:"DeviceID,omitempty"`                // An address or other identifying information to uniquely name the LogicalDevice.
		ElementName             string         `xml:"ElementName,omitempty"`             // A user-friendly name for the object.
		EnabledState            EnabledState   `xml:"EnabledState,omitempty"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		LANID                   string         `xml:"LANID,omitempty"`                   // A label or identifier for the LAN Segment to which the Endpoint is connected.
		MACAddress              string         `xml:"MACAddress,omitempty"`              // The principal unicast address used in communication with the WiFiEndpoint.
		Name                    string         `xml:"Name,omitempty"`                    // A string that identifies this ProtocolEndpoint with either a port or an interface on a device.
		NameFormat              string         `xml:"NameFormat,omitempty"`              // NameFormat contains the naming heuristic that is chosen to ensure that the value of the Name property is unique.
		ProtocolIFType          int            `xml:"ProtocolIFType,omitempty"`          // ProtocolIFType's enumeration is limited to Wi-Fi and reserved values for this subclass of ProtocolEndpoint.
		RequestedState          RequestedState `xml:"RequestedState,omitempty"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element.
		SystemCreationClassName string         `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string         `xml:"SystemName,omitempty"`              // The Name of the scoping System.
	}
)

type (
	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int
)
