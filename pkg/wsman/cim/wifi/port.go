/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"
	"errors"
	"strconv"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewWiFiPort returns a new instance of the WiFiPort struct.
func NewWiFiPortWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Port {
	return Port{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMWiFiPort, client),
	}
}

// RequestStateChange requests that the state of the element be changed to the value specified in the RequestedState parameter . . .
func (port Port) RequestStateChange(requestedState int) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: port.base.RequestStateChange(methods.GenerateAction(CIMWiFiPort, "RequestStateChange"), requestedState),
		},
	}

	err = port.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	if response.Body.RequestStateChange_OUTPUT.ReturnValue != 0 {
		err = errors.New("RequestStateChange failed with return code " + strconv.Itoa(response.Body.RequestStateChange_OUTPUT.ReturnValue))
	}

	return response, err
}

// Get retrieves the representation of the instance.
func (port Port) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: port.base.Get(nil),
		},
	}

	err = port.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (port Port) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: port.base.Enumerate(),
		},
	}

	err = port.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (port Port) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: port.base.Pull(enumerationContext),
		},
	}

	err = port.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
