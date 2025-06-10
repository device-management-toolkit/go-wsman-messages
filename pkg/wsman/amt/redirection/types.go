/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirection

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
		XMLName                   xml.Name                  `xml:"Body"`
		GetAndPutResponse         RedirectionResponse       `xml:"AMT_RedirectionService"`
		RequestStateChange_OUTPUT RequestStateChange_OUTPUT `xml:"RequestStateChange_OUTPUT"`
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
	}
	RedirectionResponse struct {
		XMLName                 xml.Name     `xml:"AMT_RedirectionService"`
		CreationClassName       string       `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string       `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledState            EnabledState `xml:"EnabledState"`                      // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		ListenerEnabled         bool         `xml:"ListenerEnabled"`                   // Describes the listener state of this service. The service is operational only if its EnabledState is set to Enabled (by the Intel® AMT Administrator) and the ListenerEnabled is set to TRUE (by a RedirectionRealm user)
		Name                    string       `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		SystemCreationClassName string       `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string       `xml:"SystemName,omitempty"`              // The Name of the scoping System.
		AccessLog               []string     `xml:"AccessLog,omitempty"`               // A list of string elements, describing recent IDE redirection operations. The format of each element is: 'Date (MM/DD/YYYY), Time (hh:mm:ss), IP:Port'
	}
	PullResponse struct {
		XMLName          xml.Name              `xml:"PullResponse"`
		RedirectionItems []RedirectionResponse `xml:"Items>AMT_RedirectionService"`
	}
	RequestStateChange_OUTPUT struct {
		XMLName xml.Name `xml:"RequestStateChange_OUTPUT"`
		// ValueMap={0, 1, 2, 3, 4, 5, 6, .., 4096, 4097, 4098, 4099, 4100..32767, 32768..65535}
		//
		// Values={Completed with No Error, Not Supported, Unknown or Unspecified Error, Cannot complete within Timeout Period, Failed, Invalid Parameter, In Use, DMTF Reserved, Method Parameters Checked - Job Started, Invalid State Transition, Use of Timeout Parameter Not Supported, Busy, Method Reserved, Vendor Specific}
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}

	// ReturnValue is an integer enumeration that indicates the completion status of the method. A value of 0 indicates success. A non-zero value indicates an error.
	ReturnValue int
)

// INPUT
// Request Types.
type (
	RedirectionRequest struct {
		XMLName                 xml.Name     `xml:"h:AMT_RedirectionService"`
		H                       string       `xml:"xmlns:h,attr"`
		CreationClassName       string       `xml:"h:CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string       `xml:"h:ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledState            EnabledState `xml:"h:EnabledState"`                      // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		ListenerEnabled         bool         `xml:"h:ListenerEnabled"`                   // Describes the listener state of this service. The service is operational only if its EnabledState is set to Enabled (by the Intel® AMT Administrator) and the ListenerEnabled is set to TRUE (by a RedirectionRealm user)
		Name                    string       `xml:"h:Name,omitempty"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		SystemCreationClassName string       `xml:"h:SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string       `xml:"h:SystemName,omitempty"`              // The Name of the scoping System.
	}
)

// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element. It can also indicate the transitions between these requested states. For example, shutting down (value=4) and starting (value=10) are transient states between enabled and disabled. The following text briefly summarizes the various enabled and disabled states:
//
// Enabled (2) indicates that the element is or could be executing commands, will process any queued commands, and queues new requests.
//
// Disabled (3) indicates that the element will not execute commands and will drop any new requests.
//
// Shutting Down (4) indicates that the element is in the process of going to a Disabled state.
//
// Not Applicable (5) indicates the element does not support being enabled or disabled.
//
// Enabled but Offline (6) indicates that the element might be completing commands, and will drop any new requests.
//
// Test (7) indicates that the element is in a test state.
//
// Deferred (8) indicates that the element might be completing commands, but will queue any new requests.
//
// Quiesce (9) indicates that the element is enabled but in a restricted mode. The behavior of the element is similar to the Enabled state, but it processes only a restricted set of commands. All other requests are queued.
//
// Starting (10) indicates that the element is in the process of going to an Enabled state. New requests are queued.
//
// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11..32767, 32768, 32769, 32770, 32771, 32772..65535}
//
// Values={Unknown, Other, Enabled, Disabled, Shutting Down, Not Applicable, Enabled but Offline, In Test, Deferred, Quiesce, Starting, DMTF Reserved, IDER and SOL are disabled, IDER is enabled and SOL is disabled, SOL is enabled and IDER is disabled, IDER and SOL are enabled, Vendor Reserved}.
type EnabledState int

// The state requested for the element. The only valid values that should be used when requesting a state change to redirection are: 32768, 32769, 32770, and 32771.
//
// ValueMap={32768, 32769, 32770, 32771}
//
// Values={disable IDER and SOL, enable IDER and disable SOL, enable SOL and disable IDER, enable IDER and SOL}.
type RequestedState int
