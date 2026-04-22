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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewEthernetPortSettingsWithClient instantiates a new Ethernet Port Settings service.
func NewEthernetPortSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base.NewService[Response](wsmanMessageCreator, AMTEthernetPortSettings, client),
	}
}

// Get retrieves the representation of the instance identified by the InstanceID
// selector. This shadows the generic parameterless Get because the public API
// has historically required an InstanceID argument here.
func (s Settings) Get(instanceID string) (response Response, err error) {
	return s.GetByInstanceID(instanceID)
}

// Put overrides the generic Put because each instance must be addressed by an
// InstanceID selector, which the generic Put does not provide.
func (s Settings) Put(instanceID string, ethernetPortSettings SettingsRequest) (response Response, err error) {
	ethernetPortSettings.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTEthernetPortSettings)
	selector := []message.Selector{{
		Name:  "InstanceID",
		Value: instanceID,
	}}
	response = Response{
		Message: &client.Message{
			XMLInput: s.Base.Put(ethernetPortSettings, true, selector),
		},
	}
	// send the message to AMT
	err = s.Base.Execute(response.Message)
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

// SetLinkPreference sets the link preference (ME or Host) and timeout on an ethernet port.
// This is an AMT method call that changes the link preference setting.
// linkPreference: 1 for ME, 2 for Host.
// timeout: timeout value in seconds.
// instanceID: the InstanceID of the AMT_EthernetPortSettings to modify.
func (s Settings) SetLinkPreference(linkPreference, timeout uint32, instanceID string) (response Response, err error) {
	selector := message.Selector{
		Name:  "InstanceID",
		Value: instanceID,
	}
	header := s.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTEthernetPortSettings, SetLinkPreference), AMTEthernetPortSettings, []message.Selector{selector}, "", "")

	request := SetLinkPreferenceRequest{
		H:              fmt.Sprintf("%s%s", message.AMTSchema, AMTEthernetPortSettings),
		LinkPreference: linkPreference,
		Timeout:        timeout,
	}

	body := s.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetLinkPreference), AMTEthernetPortSettings, &request)

	response = Response{
		Message: &client.Message{
			XMLInput: s.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	// send the message to AMT
	err = s.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
