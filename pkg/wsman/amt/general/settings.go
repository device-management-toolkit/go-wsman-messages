/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package general facilitates communication with Intel® AMT to read and configure the device's Intel® AMT general settings.
package general

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewGeneralSettingsWithClient instantiates a new General Settings service
func NewGeneralSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_GeneralSettings, client),
	}
}

// Get retrieves the representation of the instance
func (GeneralSettings Settings) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: GeneralSettings.base.Get(nil),
		},
	}
	// send the message to AMT
	err = GeneralSettings.base.Execute(response.Message)
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

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (GeneralSettings Settings) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: GeneralSettings.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = GeneralSettings.base.Execute(response.Message)
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

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (GeneralSettings Settings) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: GeneralSettings.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = GeneralSettings.base.Execute(response.Message)
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
func (GeneralSettings Settings) Put(generalSettings GeneralSettingsResponse) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: GeneralSettings.base.Put(generalSettings, false, nil),
		},
	}
	// send the message to AMT
	err = GeneralSettings.base.Execute(response.Message)
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
