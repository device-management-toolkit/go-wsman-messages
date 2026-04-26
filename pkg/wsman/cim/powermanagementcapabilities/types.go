/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package powermanagementcapabilities

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
		GetResponse       PowerManagementCapabilities
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	PowerManagementCapabilities struct {
		XMLName              xml.Name `xml:"CIM_PowerManagementCapabilities"`
		InstanceID           string   `xml:"InstanceID"`                     // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		ElementName          string   `xml:"ElementName"`                    // A user-friendly name for the object.
		PowerCapabilities    []int    `xml:"PowerCapabilities,omitempty"`    // An enumeration indicating the specific power-related capabilities of a managed element.
		PowerStatesSupported []int    `xml:"PowerStatesSupported,omitempty"` // An enumeration indicating the power states supported by a managed element.
	}

	PullResponse struct {
		XMLName                          xml.Name                      `xml:"PullResponse"`
		PowerManagementCapabilitiesItems []PowerManagementCapabilities `xml:"Items>CIM_PowerManagementCapabilities"`
	}
)
