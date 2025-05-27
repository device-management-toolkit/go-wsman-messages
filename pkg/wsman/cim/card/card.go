/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package card facilitates communication with Intel® AMT devices to represent a type of physical container that can be plugged into another Card or HostingBoard, or is itself a HostingBoard/Motherboard in a Chassis
//
// The CIM_Card class includes any package capable of carrying signals and providing a mounting point for PhysicalComponents, such as Chips, or other PhysicalPackages, such as other Cards.
package card

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewCard returns a new instance of the Card struct.
func NewCardWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMCard, client),
	}
}

// Get retrieves the representation of the instance.
func (card Package) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: card.base.Get(nil),
		},
	}

	err = card.base.Execute(response.Message)
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
func (card Package) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: card.base.Enumerate(),
		},
	}

	err = card.base.Execute(response.Message)
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
func (card Package) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: card.base.Pull(enumerationContext),
		},
	}

	err = card.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
