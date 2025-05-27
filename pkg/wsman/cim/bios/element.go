/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package bios facilitiates communication with Intel® AMT devices to get information about the device bios element
package bios

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewBIOSElementWithClient instantiates a new Element.
func NewBIOSElementWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Element {
	return Element{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMBIOSElement, client),
	}
}

// Get retrieves the representation of the instance.
func (element Element) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: element.base.Get(nil),
		},
	}

	err = element.base.Execute(response.Message)
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
func (element Element) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: element.base.Enumerate(),
		},
	}

	err = element.base.Execute(response.Message)
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
func (element Element) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: element.base.Pull(enumerationContext),
		},
	}

	err = element.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
