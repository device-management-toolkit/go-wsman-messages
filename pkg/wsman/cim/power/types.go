/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package power

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type ManagementService struct {
	base   message.Base
	client client.WSMan
}

type PowerState int

// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                         xml.Name            `xml:"Body"`
		RequestPowerStateChangeResponse PowerActionResponse `xml:"RequestPowerStateChange_OUTPUT"`
		GetResponse                     PowerManagementService
		EnumerateResponse               common.EnumerateResponse
		PullResponse                    PullResponse
	}

	PullResponse struct {
		XMLName                     xml.Name                 `xml:"PullResponse"`
		PowerManagementServiceItems []PowerManagementService `xml:"Items>CIM_PowerManagementService"`
	}

	PowerManagementService struct {
		XMLName                 xml.Name       `xml:"CIM_PowerManagementService"`
		CreationClassName       string         `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance.
		ElementName             string         `xml:"ElementName,omitempty"`             // A user-friendly name for the object.
		EnabledState            EnabledState   `xml:"EnabledState,omitempty"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		Name                    string         `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed.
		RequestedState          RequestedState `xml:"RequestedState,omitempty"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		SystemCreationClassName string         `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string         `xml:"SystemName,omitempty"`              // The Name of the scoping System.
	}

	PowerActionResponse struct {
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int
	// ReturnValue is an integer enumeration that indicates the success or failure of an operation.
	ReturnValue int
)
