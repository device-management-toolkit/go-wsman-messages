/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package processor facilitates communication with Intel® AMT devices capabilities and management of the Processor LogicalDevice
package processor

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewProcessor returns a new instance of the Processor struct.
func NewProcessorWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMProcessor, client),
	}
}

// Get retrieves the representation of the instance.
func (processor Package) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: processor.base.Get(nil),
		},
	}

	err = processor.base.Execute(response.Message)
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
func (processor Package) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: processor.base.Enumerate(),
		},
	}

	err = processor.base.Execute(response.Message)
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
func (processor Package) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: processor.base.Pull(enumerationContext),
		},
	}

	err = processor.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
