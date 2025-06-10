/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

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
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
	}

	PullResponse struct {
		XMLName            xml.Name            `xml:"PullResponse"`
		MediaAccessDevices []MediaAccessDevice `xml:"Items>CIM_MediaAccessDevice"`
	}

	MediaAccessDevice struct {
		Capabilities            []Capabilities      `xml:"Capabilities,omitempty"`  // ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} Values={Unknown, Other, Sequential Access, Random Access, Supports Writing, Encryption, Compression, Supports Removeable Media, Manual Cleaning, Automatic Cleaning, SMART Notification, Supports Dual Sided Media, Predismount Eject Not Required} ArrayType=Indexed
		CreationClassName       string              `xml:"CreationClassName"`       // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		DeviceID                string              `xml:"DeviceID"`                // An address or other identifying information to uniquely name the LogicalDevice.
		ElementName             string              `xml:"ElementName"`             // This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information.
		EnabledDefault          EnabledDefault      `xml:"EnabledDefault"`          // An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element.
		EnabledState            EnabledState        `xml:"EnabledState"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states.
		MaxMediaSize            uint64              `xml:"MaxMediaSize,omitempty"`  // Maximum size, in KBytes, of media supported by this Device.
		OperationalStatus       []OperationalStatus `xml:"OperationalStatus"`       // Indicates the current statuses of the element.
		RequestedState          RequestedState      `xml:"RequestedState"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		Security                Security            `xml:"Security,omitempty"`      // ValueMap={1, 2, 3, 4, 5, 6, 7} Values={Other, Unknown, None, Read Only, Locked Out, Boot Bypass, Boot Bypass and Read Only}
		SystemCreationClassName string              `xml:"SystemCreationClassName"` // The scoping System's CreationClassName.
		SystemName              string              `xml:"SystemName"`              // The scoping System's Name.
	}

	// Capabilities is an integer enumeration that indicates the various capabilities of the media access device.
	Capabilities int
	// EnabledDefault is an integer enumeration that indicates the administrator's default or startup configuration for the Enabled State of an element.
	EnabledDefault int
	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states.
	EnabledState int
	// OperationalStatus is an integer enumeration that indicates the current statuses of the element.
	OperationalStatus int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int
	// Security is an integer enumeration that indicates the security supported by the media access device.
	Security int
)
