/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package bios facilitiates communication with Intel® AMT devices to get information about the device bios element
package bios

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewBIOSElementWithClient instantiates a new Element
func NewBIOSElementWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Element {
	return Element{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIM_BIOSElement, client),
	}
}

// Get retrieves the representation of the instance
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

// Enumerates the instances of this class
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

// Pulls instances of this class, following an Enumerate operation
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
