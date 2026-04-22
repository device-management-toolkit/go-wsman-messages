/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package publickey facilitiates communication with Intel® AMT devices to access and configure Public Key Certificates and Public Key Management Service classes for AMT
//
// Certificate:
// This class represents a X.509 Certificate in the Intel® AMT CertStore. Instances of this class can be created using the AMT_PublicKeyManagementService.AddCertificate and AMT_PublicKeyManagementService.AddTrustedRootCertificate methods. A certificate cannot be deleted while it is being used by any service (TLS/EAC).
//
// Management Service:
// This service contains the information necessary to represent and manage the functionality provided by the Intel® AMT CertStore.
package publickey

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewPublicKeyCertificateWithClient instantiates a new Certificate.
func NewPublicKeyCertificateWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Certificate {
	return Certificate{
		base.NewService[Response](wsmanMessageCreator, AMTPublicKeyCertificate, client),
	}
}

// Get retrieves the representation of the instance identified by the InstanceID
// selector. Shadows the generic parameterless Get to preserve the public API.
func (certificate Certificate) Get(instanceID string) (response Response, err error) {
	return certificate.GetByInstanceID(instanceID)
}

// Pull overrides the generic Pull to post-process the response into the
// RefinedPullResponse shape used by callers.
func (certificate Certificate) Pull(enumerationContext string) (response Response, err error) {
	var refinedOutput []RefinedPublicKeyCertificateResponse

	response = Response{
		Message: &client.Message{
			XMLInput: certificate.Base.Pull(enumerationContext),
		},
	}

	// send the message to AMT
	err = certificate.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	for _, item := range response.Body.PullResponse.PublicKeyCertificateItems {
		output := RefinedPublicKeyCertificateResponse{
			InstanceID:             item.InstanceID,
			X509Certificate:        item.X509Certificate,
			ElementName:            item.ElementName,
			TrustedRootCertificate: item.TrustedRootCertificate,
			Issuer:                 item.Issuer,
			Subject:                item.Subject,
			ReadOnlyCertificate:    item.ReadOnlyCertificate,
		}

		refinedOutput = append(refinedOutput, output)
	}

	response.Body.RefinedPullResponse.PublicKeyCertificateItems = refinedOutput

	return response, err
}

// Put overrides the generic Put because each certificate must be addressed by
// its InstanceID selector, which the generic Put does not provide.
func (certificate Certificate) Put(instanceID, cert string) (response Response, err error) {
	selector := []message.Selector{{
		Name:  "InstanceID",
		Value: instanceID,
	}}
	publicKeyCertificate := PublicKeyCertificateRequest{}
	publicKeyCertificate.X509Certificate = cert
	publicKeyCertificate.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicKeyCertificate)
	response = Response{
		Message: &client.Message{
			XMLInput: certificate.Base.Put(publicKeyCertificate, true, selector),
		},
	}
	// send the message to AMT
	err = certificate.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// Delete removes the specified instance.
func (certificate Certificate) Delete(instanceID string) (response Response, err error) {
	selector := message.Selector{Name: "InstanceID", Value: instanceID}
	response = Response{
		Message: &client.Message{
			XMLInput: certificate.Base.Delete(selector),
		},
	}
	// send the message to AMT
	err = certificate.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
