/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package chip facilitates communication with Intel® AMT devices to represent any type of integrated circuit hardware, including ASICs, processors, memory chips, etc.
package chip

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewChip returns a new instance of the Chip struct.
func NewChipWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMChip, client),
	}
}

// Get retrieves the representation of the instance.
func (chip Package) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: chip.base.Get(nil),
		},
	}

	err = chip.base.Execute(response.Message)
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
func (chip Package) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: chip.base.Enumerate(),
		},
	}

	err = chip.base.Execute(response.Message)
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
func (chip Package) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: chip.base.Pull(enumerationContext),
		},
	}

	err = chip.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
