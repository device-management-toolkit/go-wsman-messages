/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
// Package 8021xCredentialContext facilitates communication with Intel® AMT devices to create an association between an instance of IPS_IEEE8021xSettings and an instance of AMT_PublicKeyCertificate that it uses.
package ieee8021x

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewIEEE8021xCredentialContext returns a new instance of the IPS_8021xCredentialContext struct.
func NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) CredentialContext {
	return CredentialContext{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPS8021xCredentialContext, client),
	}
}

// Get retrieves the representation of the instance.
func (credentialContext CredentialContext) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Get(nil),
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

	return response, nil
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (credentialContext CredentialContext) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Enumerate(),
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

	return response, nil
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (credentialContext CredentialContext) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: credentialContext.base.Pull(enumerationContext),
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

	return response, nil
}
