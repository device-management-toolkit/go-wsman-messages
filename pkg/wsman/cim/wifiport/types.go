/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiport

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
		XMLName       xml.Name   `xml:"PullResponse"`
		WiFiPortItems []WiFiPort `xml:"Items>CIM_WiFiPort"`
	}

	WiFiPort struct {
		XMLName                 xml.Name       `xml:"CIM_WiFiPort"`
		LinkTechnology          LinkTechnology // An enumeration of the types of links.
		DeviceID                string         // An address or other identifying information to uniquely name the LogicalDevice.
		CreationClassName       string         // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		SystemName              string         // The scoping System's Name.
		SystemCreationClassName string         // The scoping System's CreationClassName.
		ElementName             string         // A user-friendly name for the object.
		HealthState             HealthState    // Indicates the current health of the element.
		EnabledState            EnabledState   // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState // RequestedState is an integer enumeration that indicates the last requested or desired state for the element.
		PortType                PortType       // PortType shall contain the specific 802.11 operating mode that is currently enabled on the Port.
		PermanentAddress        string         // IEEE 802 EUI-48 MAC address, formatted as twelve hexadecimal digits.
	}
)

type (
	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int
	// HealthState shall indicate the current health of the element.
	HealthState int
	// LinkTechnology shall contain an enumeration of the types of links.
	LinkTechnology int
	// PortType shall contain the specific 802.11 operating mode that is currently enabled on the Port.
	PortType int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int
	// ReturnValue is a 16-bit unsigned integer enumeration that specifies the completion status of the operation.
	ReturnValue int
)
