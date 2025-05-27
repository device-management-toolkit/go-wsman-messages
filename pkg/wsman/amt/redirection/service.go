/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package redirection facilitiates communication with Intel® AMT devices to configure the IDER and SOL redirection functionalities
package redirection

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewRedirectionServiceWithClient instantiates a new Service.
func NewRedirectionServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTRedirectionService, client),
	}
}

// Get retrieves the representation of the instance.
func (service Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Get(nil),
		},
	}

	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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

// Put changes properties of the selected instance.
// The following properties must be included in any representation of AMT_RedirectionService:
//
// - Name(cannot be modified)
//
// - CreationClassName(cannot be modified)
//
// - SystemName (cannot be modified)
//
// - SystemCreationClassName (cannot be modified)
//
// - ListenerEnabled.
func (service Service) Put(redirectionService RedirectionRequest) (response Response, err error) {
	redirectionService.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTRedirectionService)
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Put(redirectionService, false, nil),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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

// RequestStateChange requests that AMT change the state of the element to the value specified in the RequestedState parameter.
// When the requested state change takes place, the EnabledState and RequestedState of the element will be the same.
// Invoking the RequestStateChange method multiple times could result in earlier requests being overwritten or lost.
// If 0 is returned, then the task completed successfully and the use of ConcreteJob was not required.
// If 4096 (0x1000) is returned, then the task will take some time to complete, ConcreteJob will be created, and its reference returned in the output parameter Job.
// Any other return code indicates an error condition.
func (service Service) RequestStateChange(requestedState RequestedState) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.RequestStateChange(methods.GenerateAction(AMTRedirectionService, RequestStateChange), int(requestedState)),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	if response.Body.RequestStateChange_OUTPUT.ReturnValue != 0 {
		err = errors.New("RequestStateChange failed with return code " + response.Body.RequestStateChange_OUTPUT.ReturnValue.String())
	}

	return
}
