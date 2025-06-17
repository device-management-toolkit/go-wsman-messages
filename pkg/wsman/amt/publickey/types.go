/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publickey

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Certificate struct {
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
		XMLName                               xml.Name                         `xml:"Body"`
		AddTrustedRootCertificate_OUTPUT      AddTrustedRootCertificate_OUTPUT `xml:"AddTrustedRootCertificate_OUTPUT,omitempty"`
		AddCertificate_OUTPUT                 AddCertificate_OUTPUT            `xml:"AddCertificate_OUTPUT,omitempty"`
		AddKey_OUTPUT                         AddKey_OUTPUT                    `xml:"AddKey_OUTPUT,omitempty"`
		GenerateKeyPair_OUTPUT                GenerateKeyPair_OUTPUT           `xml:"GenerateKeyPair_OUTPUT,omitempty"`
		GeneratePKCS10RequestEx_OUTPUT        GeneratePKCS10RequestEx_OUTPUT   `xml:"GeneratePKCS10RequestEx_OUTPUT,omitempty"`
		KeyManagementGetResponse              KeyManagementResponse            `xml:"AMT_PublicKeyManagementService,omitempty"`
		PublicKeyCertificateGetAndPutResponse PublicKeyCertificateResponse     `xml:"AMT_PublicKeyCertificate,omitempty"`
		EnumerateResponse                     common.EnumerateResponse
		PullResponse                          PullResponse
		RefinedPullResponse                   RefinedPullResponse
	}
	RefinedPullResponse struct {
		KeyManagementItems        []RefinedKeyManagementResponse        `json:"KeyManagementItems,omitempty"`
		PublicKeyCertificateItems []RefinedPublicKeyCertificateResponse `json:"PublicKeyCertificateItems,omitempty"`
	}
	PullResponse struct {
		XMLName                   xml.Name                       `xml:"PullResponse,omitempty"`
		KeyManagementItems        []KeyManagementResponse        `xml:"Items>AMT_PublicKeyManagementService,omitempty"`
		PublicKeyCertificateItems []PublicKeyCertificateResponse `xml:"Items>AMT_PublicKeyCertificate,omitempty"`
	}
	KeyManagementResponse struct {
		XMLName                 xml.Name            `xml:"AMT_PublicKeyManagementService,omitempty"`
		CreationClassName       string              `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		ElementName             string              `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledDefault          EnabledDefault      `xml:"EnabledDefault"`                    // An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element. By default, the element is "Enabled" (value=2).
		EnabledState            EnabledState        `xml:"EnabledState"`                      // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		Name                    string              `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		OperationalStatus       []OperationalStatus `xml:"OperationalStatus,omitempty"`       // Indicates the current statuses of the element.
		RequestedState          RequestedState      `xml:"RequestedState"`                    // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		SystemCreationClassName string              `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		SystemName              string              `xml:"SystemName,omitempty"`              // The Name of the scoping System.
	}
	RefinedKeyManagementResponse struct {
		CreationClassName       string
		ElementName             string
		EnabledDefault          EnabledDefault
		EnabledState            EnabledState
		Name                    string
		RequestedState          RequestedState
		SystemCreationClassName string
		SystemName              string
	}

	PublicKeyCertificateResponse struct {
		XMLName                xml.Name `xml:"AMT_PublicKeyCertificate,omitempty"`
		ElementName            string   `xml:"ElementName,omitempty"`     // A user-friendly name for the object . . .
		InstanceID             string   `xml:"InstanceID,omitempty"`      // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		X509Certificate        string   `xml:"X509Certificate,omitempty"` // uint8[4100] // The X.509 Certificate blob.
		TrustedRootCertificate bool     `xml:"TrustedRootCertficate"`     // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true. FYI Certificate is spelled wrong comimg from AMT.
		Issuer                 string   `xml:"Issuer,omitempty"`          // The Issuer field of this certificate.
		Subject                string   `xml:"Subject,omitempty"`         // The Subject field of this certificate.
		ReadOnlyCertificate    bool     `xml:"ReadOnlyCertificate"`       // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
	}

	RefinedPublicKeyCertificateResponse struct {
		ElementName            string   `json:"ElementName,omitempty"`     // A user-friendly name for the object . . .
		InstanceID             string   `json:"InstanceID,omitempty"`      // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		X509Certificate        string   `json:"X509Certificate,omitempty"` // uint8[4100] // The X.509 Certificate blob.
		TrustedRootCertificate bool     `json:"TrustedRootCertficate"`     // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true. FYI Certificate is spelled wrong comimg from AMT.
		Issuer                 string   `json:"Issuer,omitempty"`          // The Issuer field of this certificate.
		Subject                string   `json:"Subject,omitempty"`         // The Subject field of this certificate.
		ReadOnlyCertificate    bool     `json:"ReadOnlyCertificate"`       // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
		PublicKeyHandle        string   `json:"PublicKeyHandle,omitempty"`
		AssociatedProfiles     []string `json:"AssociatedProfiles,omitempty"`
	}
	AddTrustedRootCertificate_OUTPUT struct {
		XMLName            xml.Name                   `xml:"AddTrustedRootCertificate_OUTPUT"`
		CreatedCertificate CreatedCertificateResponse `xml:"CreatedCertificate,omitempty"`
		ReturnValue        ReturnValue                `xml:"ReturnValue,omitempty"`
	}
	AddCertificate_OUTPUT struct {
		XMLName            xml.Name                   `xml:"AddCertificate_OUTPUT"`
		CreatedCertificate CreatedCertificateResponse `xml:"CreatedCertificate,omitempty"`
		ReturnValue        ReturnValue                `xml:"ReturnValue,omitempty"`
	}
	AddKey_OUTPUT struct {
		XMLName     xml.Name           `xml:"AddKey_OUTPUT,omitempty"`
		CreatedKey  CreatedKeyResponse `xml:"CreatedKey,omitempty"`
		ReturnValue ReturnValue        `xml:"ReturnValue,omitempty"`
	}
	GenerateKeyPair_OUTPUT struct {
		XMLName     xml.Name        `xml:"GenerateKeyPair_OUTPUT,omitempty"`
		KeyPair     KeyPairResponse `xml:"KeyPair,omitempty"`
		ReturnValue ReturnValue     `xml:"ReturnValue,omitempty"`
	}
	GeneratePKCS10RequestEx_OUTPUT struct {
		XMLName                  xml.Name    `xml:"GeneratePKCS10RequestEx_OUTPUT,omitempty"`
		SignedCertificateRequest string      `xml:"SignedCertificateRequest,omitempty"`
		ReturnValue              ReturnValue `xml:"ReturnValue,omitempty"`
	}
	KeyPairResponse struct {
		XMLName             xml.Name                    `xml:"KeyPair,omitempty"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	CreatedKeyResponse struct {
		XMLName             xml.Name                    `xml:"CreatedKey,omitempty"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	CreatedCertificateResponse struct {
		XMLName             xml.Name                    `xml:"CreatedCertificate,omitempty"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	ReferenceParametersResponse struct {
		XMLName     xml.Name            `xml:"ReferenceParameters,omitempty"`
		ResourceURI string              `xml:"ResourceURI,omitempty"`
		SelectorSet SelectorSetResponse `xml:"SelectorSet,omitempty"`
	}
	SelectorSetResponse struct {
		XMLName   xml.Name           `xml:"SelectorSet,omitempty"`
		Selectors []SelectorResponse `xml:"Selector,omitempty"`
	}
	SelectorResponse struct {
		XMLName xml.Name `xml:"Selector,omitempty"`
		Name    string   `xml:"Name,attr"`
		Text    string   `xml:",chardata"`
	}

	// EnabledDefault is an integer enumeration that indicates an administrator's default or startup configuration for the Enabled State of an element.
	EnabledDefault int

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int

	// OperationalStatus is an integer enumeration that indicates the current statuses of the element.
	OperationalStatus int

	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int

	// ReturnValue is an integer enumeration that indicates the return status of the method.
	ReturnValue int
)

// INPUTS
// Request Types.
type (
	PublicKeyCertificateRequest struct {
		XMLName                xml.Name `xml:"h:AMT_PublicKeyCertificate"`
		H                      string   `xml:"xmlns:h,attr"`
		ElementName            string   `xml:"h:ElementName"`            // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		InstanceID             string   `xml:"h:InstanceID"`             // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		X509Certificate        string   `xml:"h:X509Certificate"`        // uint8[4100] // The X.509 Certificate blob.
		TrustedRootCertificate bool     `xml:"h:TrustedRootCertificate"` // For root certificate [that were added by AMT_PublicKeyManagementService.AddTrustedRootCertificate()]this property will be true.
		Issuer                 string   `xml:"h:Issuer"`                 // The Issuer field of this certificate.
		Subject                string   `xml:"h:Subject"`                // The Subject field of this certificate.
		ReadOnlyCertificate    bool     `xml:"h:ReadOnlyCertificate"`    // Indicates whether the certificate is an Intel AMT self-signed certificate. If True, the certificate cannot be deleted.
	}
	AddCertificate_INPUT struct {
		XMLName         xml.Name `xml:"h:AddCertificate_INPUT"`
		H               string   `xml:"xmlns:h,attr"`
		CertificateBlob string   `xml:"h:CertificateBlob"` // The use of ECC 192/224 is blocked starting from Intel® CSME 18.0.
	}
	AddTrustedRootCertificate_INPUT struct {
		XMLName         xml.Name `xml:"h:AddTrustedRootCertificate_INPUT"`
		H               string   `xml:"xmlns:h,attr"`
		CertificateBlob string   `xml:"h:CertificateBlob"` // The use of ECC 192/224 is blocked starting from Intel® CSME 18.0.
	}
	AddKey_INPUT struct {
		XMLName xml.Name `xml:"h:AddKey_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		KeyBlob string   `xml:"h:KeyBlob"` // The use of ECC 192/224 is blocked starting from Intel® CSME 18.0.
	}
	GenerateKeyPair_INPUT struct {
		XMLName      xml.Name     `xml:"h:GenerateKeyPair_INPUT"`
		H            string       `xml:"xmlns:h,attr"`
		KeyAlgorithm KeyAlgorithm `xml:"h:KeyAlgorithm"` // The algorithm of the generated key.
		KeyLength    KeyLength    `xml:"h:KeyLength"`    // The length of the generatd key in bits.
	}
	PKCS10Request struct {
		XMLName                      xml.Name         `xml:"h:GeneratePKCS10RequestEx_INPUT"`
		H                            string           `xml:"xmlns:h,attr"`
		KeyPair                      KeyPair          `xml:"h:KeyPair"`
		SigningAlgorithm             SigningAlgorithm `xml:"h:SigningAlgorithm"`             // The signing algorithm that the FW should use for signing the certificate request
		NullSignedCertificateRequest string           `xml:"h:NullSignedCertificateRequest"` // A binary representation of the null-signed PKCS#10 request.the request must include a valid PKCS10RequestInfo, that will be signed by AMT FW. The Public Key specified in the request must match the public key of the referenced KeyPair parameter.
	}

	KeyPair struct {
		XMLName             xml.Name                   `xml:"h:KeyPair,omitempty"`
		Address             string                     `xml:"a:Address,omitempty"`
		ReferenceParameters ReferenceParametersRequest `xml:"a:ReferenceParameters,omitempty"`
	}
	ReferenceParametersRequest struct {
		XMLName     xml.Name           `xml:"a:ReferenceParameters,omitempty"`
		ResourceURI string             `xml:"w:ResourceURI"`
		SelectorSet SelectorSetRequest `xml:"w:SelectorSet,omitempty"`
	}
	SelectorSetRequest struct {
		XMLName   xml.Name          `xml:"w:SelectorSet,omitempty"`
		Selectors []SelectorRequest `xml:"w:Selector"`
	}
	SelectorRequest struct {
		XMLName xml.Name `xml:"w:Selector,omitempty"`
		Name    string   `xml:"Name,attr"`
		Text    string   `xml:",chardata"`
	}

	// The signing algorithm that the FW should use for signing the certificate request
	//
	// Required
	//
	// ValueMap={0, 1, ..}
	//
	// Values={SHA1-RSA, SHA256-RSA, Reserved} Note: SHA1 is no longer available starting from Intel CSME 18.0.
	SigningAlgorithm int

	// The algorithm of the generated key.
	//
	// Required
	//
	// ValueMap={0, ..}
	//
	// Values={RSA, Reserved}.
	KeyAlgorithm int

	// The length of the generatd key in bits.
	KeyLength int
)
