/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package physical facilitates communications with Intel® AMT devices to get the PhysicalMemory as a subclass of CIM_Chip, representing low level memory devices - SIMMS, DIMMs, raw memory chips, etc.
package physical

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewPhysicalMemory returns a new instance of the PhysicalMemory struct.
func NewPhysicalMemoryWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Memory {
	return Memory{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMPhysicalMemory, client),
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors
// Get retrieves the representation of the instance

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (memory Memory) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: memory.base.Enumerate(),
		},
	}

	err = memory.base.Execute(response.Message)
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
func (memory Memory) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: memory.base.Pull(enumerationContext),
		},
	}

	err = memory.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
