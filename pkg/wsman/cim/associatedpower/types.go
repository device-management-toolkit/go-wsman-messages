/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package associatedpower

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
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
		XMLName                               xml.Name                               `xml:"PullResponse"`
		AssociatedPowerManagementServiceItems []CIM_AssociatedPowerManagementService `xml:"Items>CIM_AssociatedPowerManagementService"`
	}
	CIM_AssociatedPowerManagementService struct {
		AvailableRequestedPowerStates []models.AvailableRequestedPowerStates `xml:"AvailableRequestedPowerStates,omitempty"` // AvailableRequestedPowerStates indicates the possible values for the PowerState parameter of the method RequestPowerStateChange, used to initiate a power state change.
		PowerState                    models.PowerState                      `xml:"PowerState,omitempty"`                    // The current power state of the associated Managed System Element.
		OtherPowerState               string                                 `xml:"OtherPowerState,omitempty"`               // A string describing the additional power management state of the element, used when the PowerState is set to the value 1, "Other".
		RequestedPowerState           models.RequestedPowerState             `xml:"RequestedPowerState,omitempty"`           // The desired or the last requested power state of the associated Managed System Element.
		OtherRequestedPowerState      string                                 `xml:"OtherRequestedPowerState,omitempty"`      // A string describing the additional power management state of the element, used when the RequestedPowerState is set to the value 1, "Other".
		PowerOnTime                   string                                 `xml:"PowerOnTime,omitempty"`                   // The time when the element will be powered on again, used when the RequestedPowerState has the value 2, "On", 5, "Power Cycle (Off - Soft)" or 6, "Power Cycle (Off - Hard)".
		TransitioningToPowerState     models.TransitioningToPowerState       `xml:"TransitioningToPowerState,omitempty"`     // TransitioningToPowerState indicates the target power state to which the system is transitioning.
		ServiceProvided               ServiceProvided                        // The Service that is available.
		UserOfService                 UserOfService                          // The ManagedElement that can use the Service.
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
	SelectorSet struct {
		XMLName   xml.Name           `xml:"SelectorSet,omitempty"`
		Selectors []message.Selector `xml:"Selector"`
	}
	ReferenceParameters struct {
		XMLName     xml.Name `xml:"ReferenceParameters,omitempty"`
		ResourceURI string   `xml:"ResourceURI,omitempty"`
		SelectorSet SelectorSet
	}
)
