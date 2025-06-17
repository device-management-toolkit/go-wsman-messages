/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package kvm facilitates communication with Intel® AMT devices derived from Service Access Point, that describes an access point to start the KVM redirection. One access point represents access to a single KVM redirection stream.
package kvm

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type RedirectionSAP struct {
	base.WSManService[Response]
}

// NewKVMRedirectionSAP returns a new instance of the KVMRedirectionSAP struct.
func NewKVMRedirectionSAPWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) RedirectionSAP {
	return RedirectionSAP{
		base.NewService[Response](wsmanMessageCreator, CIMKVMRedirectionSAP, client),
	}
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (redirectionSAP RedirectionSAP) RequestStateChange(requestedState KVMRedirectionSAPRequestStateChangeInput) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: redirectionSAP.Base.RequestStateChange(methods.RequestStateChange(CIMKVMRedirectionSAP), int(requestedState)),
		},
	}

	err = redirectionSAP.Base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
