/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package service

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
		XMLName                          xml.Name `xml:"Body"`
		PullResponse                     PullResponse
		EnumerateResponse                common.EnumerateResponse
		AssociatedPowerManagementService CIM_AssociatedPowerManagementService `xml:"CIM_AssociatedPowerManagementService"`
	}
	PullResponse struct {
		XMLName                          xml.Name                               `xml:"PullResponse"`
		AssociatedPowerManagementService []CIM_AssociatedPowerManagementService `xml:"Items>CIM_AssociatedPowerManagementService"`
	}
	CIM_AssociatedPowerManagementService struct {
		AvailableRequestedPowerStates []AvailableRequestedPowerStates `xml:"AvailableRequestedPowerStates,omitempty"` // AvailableRequestedPowerStates indicates the possible values for the PowerState parameter of the method RequestPowerStateChange, used to initiate a power state change.The values listed shall be a subset of the values contained in the RequestedPowerStatesSupported property of the CIM_PowerManagementCapabilities where the values selected are a function of the current power state of the system. This property shall be non-null if an implementation supports the advertisement of the set of possible values as a function of the current state. This property shall be null if an implementation does not support the advertisement of the set of possible values as a function of the current state.
		PowerState                    PowerState                      `xml:"PowerState,omitempty"`                    // The current power state of the associated Managed System Element.
		ServiceProvided               ServiceProvided                 // The Service that is available.
		UserOfService                 UserOfService                   // The ManagedElement that can use the Service.
	}
	ServiceProvided struct {
		XMLName             xml.Name `xml:"ServiceProvided,omitempty"`
		Address             string   `xml:"Address"`
		ReferenceParameters ReferenceParameters
	}
	UserOfService struct {
		XMLName             xml.Name `xml:"UserOfService,omitempty"`
		Address             string   `xml:"Address"`
		ReferenceParameters ReferenceParameters
	}
	ReferenceParameters struct {
		XMLName     xml.Name `xml:"ReferenceParameters,omitempty"`
		ResourceURI string   `xml:"ResourceURI,omitempty"`
		SelectorSet message.SelectorSet
	}

	// AvailableRequestedPowerStates indicates the possible values for the PowerState parameter of the method RequestPowerStateChange, used to initiate a power state change.
	AvailableRequestedPowerStates int
	// PowerState indicates the current power state of the associated Managed System Element.
	PowerState int
)
