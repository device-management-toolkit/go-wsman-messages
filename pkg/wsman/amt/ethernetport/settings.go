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

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewEthernetPortSettingsWithClient instantiates a new Ethernet Port Settings service
func NewEthernetPortSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_EthernetPortSettings, client),
	}
}

// Get retrieves the representation of the instance
func (s Settings) Get(instanceId int) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: fmt.Sprintf("Intel(r) AMT Ethernet Port Settings %d", instanceId),
	}
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Get(&selector),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (s Settings) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// // Pulls instances of this class, following an Enumerate operation
func (s Settings) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Put will change properties of the selected instance
func (s Settings) Put(ethernetPortSettings SettingsRequest, instanceId int) (response Response, err error) {
	ethernetPortSettings.H = fmt.Sprintf("%s%s", message.AMTSchema, AMT_EthernetPortSettings)
	selector := message.Selector{
		Name:  "InstanceID",
		Value: fmt.Sprintf("Intel(r) AMT Ethernet Port Settings %d", instanceId),
	}
	response = Response{
		Message: &client.Message{
			XMLInput: s.base.Put(ethernetPortSettings, true, &selector),
		},
	}
	// send the message to AMT
	err = s.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
