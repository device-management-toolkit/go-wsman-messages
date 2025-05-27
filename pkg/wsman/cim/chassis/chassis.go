/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package chassis facilitates communication with Intel® AMT devices to represent the PhysicalElements that enclose other Elements and provide definable functionality, such as a desktop, processing node, UPS, disk or tape storage, or a combination of these.
package chassis

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewChassis returns a new instance of the Chassis struct.
func NewChassisWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base: message.NewBaseWithClient(wsmanMessageCreator, CIMChassis, client),
	}
}

// Get retrieves the representation of the instance.
func (chassis Package) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: chassis.base.Get(nil),
		},
	}

	err = chassis.base.Execute(response.Message)
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
func (chassis Package) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: chassis.base.Enumerate(),
		},
	}

	err = chassis.base.Execute(response.Message)
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
func (chassis Package) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: chassis.base.Pull(enumerationContext),
		},
	}

	err = chassis.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
