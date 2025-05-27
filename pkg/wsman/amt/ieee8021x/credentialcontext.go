/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package ieee8021x facilitiates communication with Intel® AMT devices to access the ieee8021x credential context and profile settings
//
// CredentialContext gets the association between an instance of AMT_8021XProfile and an instance of AMT_PublicKeyCertificate that it uses.
//
// Profile represents a 802.1X profile in the Intel® AMT system.
package ieee8021x

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type CredentialContext struct {
	base message.Base
}

// NewIEEE8021xCredentialContextWithClient instantiates a new CredentialContext service.
func NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) CredentialContext {
	return CredentialContext{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTIEEE8021xCredentialContext, client),
	}
}

// TODO: Handle GET input
// Get retrieves the representation of the instance

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
