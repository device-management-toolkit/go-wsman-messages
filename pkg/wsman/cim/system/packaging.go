/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package system facilitates communication with Intel® AMT devices in a way similar to the way that LogicalDevices are 'Realized' by PhysicalElements, Systems can be associated with specific packaging or PhysicalElements.
//
// This association explicitly defines the relationship between a System and its packaging.
package system

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewSystemPackaging returns a new instance of the SystemPackaging struct.
func NewSystemPackageWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMSystemPackaging, client),
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors
// Get retrieves the representation of the instance. No Route

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (packaging Package) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: packaging.base.Enumerate(),
		},
	}

	err = packaging.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (packaging Package) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: packaging.base.Pull(enumerationContext),
		},
	}

	err = packaging.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
