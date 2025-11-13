/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package ethernetport facilitates communication with Intel® AMT devices to configure all Intel® AMT network specific settings (IP, DHCP, VLAN).
//
// Intel® AMT devices support a single wired and a single wireless network adapter.  If an Intel® AMT device has multiple wired or wireless network adapters only one of each will be connected to AMT.
package ethernetport

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewEthernetPortSettingsWithClient instantiates a new Ethernet Port Settings service.
func NewEthernetPortSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTEthernetPortSettings, client),
	}
}

// Get retrieves the representation of the instance.
func (s Settings) Get(instanceID string) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: instanceID,
	}
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Get(&selector),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (s Settings) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// // Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (s Settings) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}

// Put will change properties of the selected instance.
func (s Settings) Put(instanceID string, ethernetPortSettings SettingsRequest) (response Response, err error) {
	ethernetPortSettings.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTEthernetPortSettings)
	selector := []message.Selector{{
		Name:  "InstanceID",
		Value: instanceID,
	}}
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Put(ethernetPortSettings, true, selector),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
