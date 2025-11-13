/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publicprivate

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type KeyPair struct {
	base message.Base
}

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
		XMLName             xml.Name `xml:"Body"`
		GetResponse         PublicPrivateKeyPair
		EnumerateResponse   common.EnumerateResponse
		PullResponse        PullResponse
		RefinedPullResponse RefinedPullResponse
	}

	PullResponse struct {
		XMLName                   xml.Name               `xml:"PullResponse"`
		PublicPrivateKeyPairItems []PublicPrivateKeyPair `xml:"Items>AMT_PublicPrivateKeyPair"`
	}
	RefinedPullResponse struct {
		PublicPrivateKeyPairItems []RefinedPublicPrivateKeyPair
	}

	PublicPrivateKeyPair struct {
		XMLName     xml.Name `xml:"AMT_PublicPrivateKeyPair"`
		ElementName string   // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		InstanceID  string   // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		DERKey      string   // RSA Key encoded as DES PKCS#1. The Exponent (E) is 65537 (0x010001).When this structure is used as an output parameter (GET or PULL method),only the public section of the key is exported. uint8[1210]
	}

	RefinedPublicPrivateKeyPair struct {
		ElementName       string // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		InstanceID        string // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		DERKey            string // RSA Key encoded as DES PKCS#1. The Exponent (E) is 65537 (0x010001).When this structure is used as an output parameter (GET or PULL method),only the public section of the key is exported. uint8[1210]
		CertificateHandle string `json:"CertificateHandle,omitempty"`
	}
	PublicPrivateSelector message.Selector
)
