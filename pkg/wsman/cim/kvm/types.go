/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kvm

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
		XMLName                   xml.Name                  `xml:"Body"`
		GetResponse               KVMRedirectionSAP         `xml:"CIM_KVMRedirectionSAP"`
		RequestStateChange_OUTPUT RequestStateChange_OUTPUT `xml:"RequestStateChange_OUTPUT"`
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse `xml:"PullResponse"`
	}

	KVMRedirectionSAP struct {
		CreationClassName       string         `xml:"CreationClassName"`        // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string         `xml:"ElementName"`              // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		Name                    string         `xml:"Name"`                     // The Name property uniquely identifies the ServiceAccessPoint and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		SystemCreationClassName string         `xml:"SystemCreationClassName"`  // The CreationClassName of the scoping System.
		SystemName              string         `xml:"SystemName"`               // The Name of the scoping System.
		EnabledState            EnabledState   `xml:"EnabledState,omitempty"`   // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState `xml:"RequestedState,omitempty"` // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		KVMProtocol             KVMProtocol    `xml:"KVMProtocol,omitempty"`    // An enumeration specifying the type of the KVM stream supported on this SAP.
	}

	Time struct {
		DateTime string `xml:"Datetime"`
	}

	PullResponse struct {
		XMLName xml.Name            `xml:"PullResponse"`
		Items   []KVMRedirectionSAP `xml:"Items>CIM_KVMRedirectionSAP"`
	}

	RequestStateChange_OUTPUT struct {
		XMLName xml.Name `xml:"RequestStateChange_OUTPUT"`
		// ValueMap={0, 1, 2, 3, 4, 5, 6, .., 4096, 4097, 4098, 4099, 4100..32767, 32768..65535}
		//
		// Values={Completed with No Error, Not Supported, Unknown or Unspecified Error, Cannot complete within Timeout Period, Failed, Invalid Parameter, In Use, DMTF Reserved, Method Parameters Checked - Job Started, Invalid State Transition, Use of Timeout Parameter Not Supported, Busy, Method Reserved, Vendor Specific}
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}

	// KVMProtocol is an enumeration specifying the type of the KVM stream supported on this SAP.
	KVMProtocol int
	// ReturnValue is an enumeration specifying the return value of the operation.
	ReturnValue int
	// EnabledState is an enumeration that indicates the enabled and disabled states of an element.
	EnabledState int
	// RequestedState is an enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int
)

// Request Types.
type (
	KVMRedirectionSAPRequestStateChangeInput int
)
