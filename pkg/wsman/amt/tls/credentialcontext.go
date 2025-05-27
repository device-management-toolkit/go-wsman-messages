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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewTLSCredentialContextWithClient instantiates a new CredentialContext.
func NewTLSCredentialContextWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) CredentialContext {
	return CredentialContext{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTTLSCredentialContext, client),
	}
}

// Get retrieves the representation of the instance.
func (credentialContext CredentialContext) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Get(nil),
		},
	}
	// send the message to AMT
	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (credentialContext CredentialContext) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (credentialContext CredentialContext) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Delete removes a the specified instance.
func (credentialContext CredentialContext) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "Name", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Delete(selector),
		},
	}
	// send the message to AMT
	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Creates a new instance of this class.
func (credentialContext CredentialContext) Create(certHandle string) (response Response, err error) {
	header := credentialContext.base.WSManMessageCreator.CreateHeader(message.BaseActionsCreate, AMTTLSCredentialContext, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AMT_TLSCredentialContext xmlns:h="%sAMT_TLSCredentialContext"><h:ElementInContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementInContext><h:ElementProvidingContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_TLSProtocolEndpointCollection</w:ResourceURI><w:SelectorSet><w:Selector Name="ElementName">TLSProtocolEndpointInstances Collection</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementProvidingContext></h:AMT_TLSCredentialContext></Body>`, credentialContext.base.WSManMessageCreator.ResourceURIBase, credentialContext.base.WSManMessageCreator.ResourceURIBase, certHandle, credentialContext.base.WSManMessageCreator.ResourceURIBase)
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Put will update the certificate when TLS is enabled.
func (credentialContext CredentialContext) Put(certHandle string) (response Response, err error) {
	header := credentialContext.base.WSManMessageCreator.CreateHeader(message.BaseActionsPut, AMTTLSCredentialContext, nil, "", "")
	body := fmt.Sprintf(`<Body><h:AMT_TLSCredentialContext xmlns:h="%sAMT_TLSCredentialContext"><h:ElementInContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementInContext><h:ElementProvidingContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_TLSProtocolEndpointCollection</w:ResourceURI><w:SelectorSet><w:Selector Name="ElementName">TLSProtocolEndpointInstances Collection</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementProvidingContext></h:AMT_TLSCredentialContext></Body>`, credentialContext.base.WSManMessageCreator.ResourceURIBase, credentialContext.base.WSManMessageCreator.ResourceURIBase, certHandle, credentialContext.base.WSManMessageCreator.ResourceURIBase)
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = credentialContext.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
