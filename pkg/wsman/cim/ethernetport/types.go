/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

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
		GetResponse       EthernetPort
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	EthernetPort struct {
		XMLName                 xml.Name       `xml:"CIM_EthernetPort"`
		DeviceID                string         `xml:"DeviceID"`                   // An address or other identifying information to uniquely name the LogicalDevice.
		CreationClassName       string         `xml:"CreationClassName"`          // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		SystemName              string         `xml:"SystemName"`                 // The scoping System's Name.
		SystemCreationClassName string         `xml:"SystemCreationClassName"`    // The scoping System's CreationClassName.
		ElementName             string         `xml:"ElementName"`                // A user-friendly name for the object.
		EnabledState            EnabledState   `xml:"EnabledState,omitempty"`     // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState `xml:"RequestedState,omitempty"`   // RequestedState is an integer enumeration that indicates the last requested or desired state for the element.
		LinkTechnology          LinkTechnology `xml:"LinkTechnology,omitempty"`   // An enumeration of the types of links.
		PermanentAddress        string         `xml:"PermanentAddress,omitempty"` // The IEEE 802 EUI-48 MAC address.
		Speed                   uint64         `xml:"Speed,omitempty"`            // The current bandwidth of the Port in Bits per Second.
		PortType                PortType       `xml:"PortType,omitempty"`         // The specific mode that is currently enabled on the Port.
		MaxDataSize             uint32         `xml:"MaxDataSize,omitempty"`      // The maximum size of the INFO (non-MAC) field that will be received or transmitted.
		Capabilities            []int          `xml:"Capabilities,omitempty"`     // Capabilities of the EthernetPort.
	}

	PullResponse struct {
		XMLName           xml.Name       `xml:"PullResponse"`
		EthernetPortItems []EthernetPort `xml:"Items>CIM_EthernetPort"`
	}

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element.
	RequestedState int
	// LinkTechnology is an enumeration of the types of links.
	LinkTechnology int
	// PortType is the specific mode that is currently enabled on the Port.
	PortType int
)
