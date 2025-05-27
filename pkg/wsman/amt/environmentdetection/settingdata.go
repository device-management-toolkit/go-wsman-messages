/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package environmentdetection facilitates communication with Intel® AMT device configuration-related and operational parameters for the Environment Detection service in Intel® AMT.
package environmentdetection

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewEnvironmentDetectionSettingDataWithClient instantiates a new Environment Detection Setting Data service.
func NewEnvironmentDetectionSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTEnvironmentDetectionSettingData, client),
	}
}

// Get retrieves the representation of the instance.
func (sd SettingData) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: sd.base.Get(nil),
		},
	}
	// send the message to AMT
	err = sd.base.Execute(response.Message)
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

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (sd SettingData) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: sd.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = sd.base.Execute(response.Message)
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
func (sd SettingData) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: sd.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = sd.base.Execute(response.Message)
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

// Put will change properties of the selected instance.
func (sd SettingData) Put(environmentDetectionSettingData EnvironmentDetectionSettingDataRequest) (response Response, err error) {
	environmentDetectionSettingData.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTEnvironmentDetectionSettingData)
	selector := []message.Selector{{
		Name:  "InstanceID",
		Value: "Intel(r) AMT Environment Detection Settings",
	}}
	response = Response{
		Message: &client.Message{
			XMLInput: sd.base.Put(environmentDetectionSettingData, true, selector),
		},
	}
	// send the message to AMT
	err = sd.base.Execute(response.Message)
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
