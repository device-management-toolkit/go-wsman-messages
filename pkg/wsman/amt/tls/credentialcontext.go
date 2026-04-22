/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package tls facilitiates communication with Intel® AMT devices to access and configure TLS Credential Context, TLS Protocol Endpoint Collection, and TLS Setting Data features of AMT
//
// Credential Context:
// This class represents the credential of the TLSProtocolEndpointCollection, by connecting a certficate to the service.
// The connected certificate must be a leaf certificate, and must have a matching private key.
// You can't enable the TLS service without a credential.
// When TLS is enabled the certificate can be changed using the Put method.
//
// Protocol Endpoint Collection:
// This class connects the 2 instances of AMT_TLSProtocolEndpoint and can be used in order to enable/disable TLS in the system.
//
// Setting Data:
// This class represents configuration-related and operational parameters for the TLS service in the Intel® AMT.
package tls

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewTLSCredentialContextWithClient instantiates a new CredentialContext.
func NewTLSCredentialContextWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) CredentialContext {
	return CredentialContext{
		base.NewService[Response](wsmanMessageCreator, AMTTLSCredentialContext, client),
	}
}

// Delete removes the specified instance.
func (credentialContext CredentialContext) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "Name", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.Base.Delete(selector),
		},
	}
	// send the message to AMT
	err = credentialContext.Base.Execute(response.Message)
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

// Creates a new instance of this class.
func (credentialContext CredentialContext) Create(certHandle string) (response Response, err error) {
	header := credentialContext.Base.WSManMessageCreator.CreateHeader(message.BaseActionsCreate, AMTTLSCredentialContext, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AMT_TLSCredentialContext xmlns:h="%sAMT_TLSCredentialContext"><h:ElementInContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementInContext><h:ElementProvidingContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_TLSProtocolEndpointCollection</w:ResourceURI><w:SelectorSet><w:Selector Name="ElementName">TLSProtocolEndpointInstances Collection</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementProvidingContext></h:AMT_TLSCredentialContext></Body>`, credentialContext.Base.WSManMessageCreator.ResourceURIBase, credentialContext.Base.WSManMessageCreator.ResourceURIBase, certHandle, credentialContext.Base.WSManMessageCreator.ResourceURIBase)
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = credentialContext.Base.Execute(response.Message)
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

// Put overrides the generic Put because this method takes a cert handle string
// (not a struct) and must craft a body containing two EPRs identifying the
// PublicKeyCertificate and TLSProtocolEndpointCollection instances.
func (credentialContext CredentialContext) Put(certHandle string) (response Response, err error) {
	header := credentialContext.Base.WSManMessageCreator.CreateHeader(message.BaseActionsPut, AMTTLSCredentialContext, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AMT_TLSCredentialContext xmlns:h="%sAMT_TLSCredentialContext"><h:ElementInContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementInContext><h:ElementProvidingContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_TLSProtocolEndpointCollection</w:ResourceURI><w:SelectorSet><w:Selector Name="ElementName">TLSProtocolEndpointInstances Collection</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementProvidingContext></h:AMT_TLSCredentialContext></Body>`, credentialContext.Base.WSManMessageCreator.ResourceURIBase, credentialContext.Base.WSManMessageCreator.ResourceURIBase, certHandle, credentialContext.Base.WSManMessageCreator.ResourceURIBase)
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = credentialContext.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
