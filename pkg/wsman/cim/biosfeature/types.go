/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package biosfeature

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
		GetResponse       BIOSFeature
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	BIOSFeature struct {
		XMLName                xml.Name `xml:"CIM_BIOSFeature"`
		InstanceID             string   `xml:"InstanceID"`             // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		Name                   string   `xml:"Name"`                   // The label by which the object is known.
		CapabilityDescriptions []string `xml:"CapabilityDescriptions"` // An array of free-form strings providing more detailed explanations for any of the features indicated in the Capabilities array.
		ElementName            string   `xml:"ElementName"`            // A user-friendly name for the object.
	}

	PullResponse struct {
		XMLName          xml.Name      `xml:"PullResponse"`
		BIOSFeatureItems []BIOSFeature `xml:"Items>CIM_BIOSFeature"`
	}
)

// PutRequest is used to modify a BIOSFeature instance.
type PutRequest struct {
	FeatureName string
}
