/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package kvm facilitates communication with Intel® AMT devices derived from Service Access Point, that describes an access point to start the KVM redirection. One access point represents access to a single KVM redirection stream.
package kvm

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewKVMRedirectionSAP returns a new instance of the KVMRedirectionSAP struct.
func NewKVMRedirectionSAPWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) RedirectionSAP {
	return RedirectionSAP{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMKVMRedirectionSAP, client),
	}
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (redirectionSAP RedirectionSAP) RequestStateChange(requestedState KVMRedirectionSAPRequestStateChangeInput) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: redirectionSAP.base.RequestStateChange(methods.RequestStateChange(CIMKVMRedirectionSAP), int(requestedState)),
		},
	}

	err = redirectionSAP.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Get retrieves the representation of the instance.
func (redirectionSAP RedirectionSAP) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: redirectionSAP.base.Get(nil),
		},
	}

	err = redirectionSAP.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (redirectionSAP RedirectionSAP) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: redirectionSAP.base.Enumerate(),
		},
	}

	err = redirectionSAP.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (redirectionSAP RedirectionSAP) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: redirectionSAP.base.Pull(enumerationContext),
		},
	}

	err = redirectionSAP.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
