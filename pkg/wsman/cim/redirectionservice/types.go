/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package redirectionservice

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
		GetResponse       RedirectionService
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	RedirectionService struct {
		XMLName                 xml.Name       `xml:"CIM_RedirectionService"`
		CreationClassName       string         `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		ElementName             string         `xml:"ElementName,omitempty"`             // A user-friendly name for the object.
		EnabledState            EnabledState   `xml:"EnabledState,omitempty"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		Name                    string         `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service.
		RequestedState          RequestedState `xml:"RequestedState,omitempty"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element.
		SystemCreationClassName string         `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string         `xml:"SystemName,omitempty"`              // The Name of the scoping System.
		AccessLog               []string       `xml:"AccessLog,omitempty"`               // An array of free-form strings that provide information about the access.
		MaxCurrentEnabledSAPs   int            `xml:"MaxCurrentEnabledSAPs,omitempty"`   // The maximum number of currently enabled SAPs supported by this service.
		RedirectionServiceType  []int          `xml:"RedirectionServiceType,omitempty"`  // An enumeration indicating the type(s) of redirection supported by this service.
		SharingMode             int            `xml:"SharingMode,omitempty"`             // An enumeration that indicates how the redirection service may be shared.
	}

	PullResponse struct {
		XMLName                 xml.Name             `xml:"PullResponse"`
		RedirectionServiceItems []RedirectionService `xml:"Items>CIM_RedirectionService"`
	}

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element.
	RequestedState int
)
