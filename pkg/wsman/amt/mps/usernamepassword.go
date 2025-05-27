/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package mps facilitiates communication with Intel® AMT devices to configure the username and password used to access an MPS.
package mps

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewMPSUsernamePasswordWithClient instantiates a new UsernamePassword.
func NewMPSUsernamePasswordWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) UsernamePassword {
	return UsernamePassword{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTMPSUsernamePassword, client),
	}
}

// Get retrieves the representation of the instance.
func (usernamePassword UsernamePassword) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: usernamePassword.base.Get(nil),
		},
	}
	// send the message to AMT
	err = usernamePassword.base.Execute(response.Message)
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
func (usernamePassword UsernamePassword) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: usernamePassword.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = usernamePassword.base.Execute(response.Message)
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
func (usernamePassword UsernamePassword) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: usernamePassword.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = usernamePassword.base.Execute(response.Message)
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

// Put will change properties of the selected instance.
func (usernamePassword UsernamePassword) Put(mpsUsernamePassword MPSUsernamePasswordRequest) (response Response, err error) {
	mpsUsernamePassword.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTMPSUsernamePassword)
	response = Response{
		Message: &client.Message{
			XMLInput: usernamePassword.base.Put(mpsUsernamePassword, false, nil),
		},
	}
	// send the message to AMT
	err = usernamePassword.base.Execute(response.Message)
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
