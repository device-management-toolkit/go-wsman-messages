/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package mediaaccess facilitiates communication with Intel® AMT devices to represent the ability to access one or more media and use this media to store and retrieve data.
package mediaaccess

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewMediaAccessDevice returns a new instance of the MediaAccessDevice struct.
func NewMediaAccessDeviceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Device {
	return Device{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMMediaAccessDevice, client),
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors
// Get retrieves the representation of the instance

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (device Device) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: device.base.Enumerate(),
		},
	}

	err = device.base.Execute(response.Message)
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
func (device Device) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: device.base.Pull(enumerationContext),
		},
	}

	err = device.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
