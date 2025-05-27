/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

// The state requested for the element. The valid input values for userinitiatedconnection request state change are: 32768, 32769, 32770, 32771.
//
// ValueMap={32768, 32769, 32770, 32771}
//
// Values={All Interfaces disabled, BIOS Interface enabled, OS Interface enabled, BIOS and OS Interfaces enabled}.
type RequestedState int

// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
type EnabledState int

// ReturnValue is a 32-bit unsigned integer indicating the success or failure of the operation.
type ReturnValue int

// OUTPUTS
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                   xml.Name                  `xml:"Body"`
		RequestStateChange_OUTPUT RequestStateChange_OUTPUT `xml:"RequestStateChange_OUTPUT"`
		GetResponse               UserResponse
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
	}
	UserResponse struct {
		XMLName                 xml.Name     `xml:"AMT_UserInitiatedConnectionService"`
		CreationClassName       string       `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string       `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledState            EnabledState `xml:"EnabledState"`                      // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		Name                    string       `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		SystemCreationClassName string       `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string       `xml:"SystemName,omitempty"`              // The Name of the scoping System.
	}
	PullResponse struct {
		XMLName   xml.Name       `xml:"PullResponse"`
		UserItems []UserResponse `xml:"Items>AMT_UserInitiatedConnectionService"`
	}

	RequestStateChange_OUTPUT struct {
		XMLName     xml.Name    `xml:"RequestStateChange_OUTPUT"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
)
