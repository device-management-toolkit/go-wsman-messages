/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewTLSSettingDataWithClient instantiates a new SettingData.
func NewTLSSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base.NewService[Response](wsmanMessageCreator, AMTTLSSettingData, client),
	}
}

// Get retrieves the representation of the instance identified by the InstanceID
// selector. Shadows the generic parameterless Get to preserve the public API.
func (settingData SettingData) Get(instanceID string) (response Response, err error) {
	return settingData.GetByInstanceID(instanceID)
}

// Put changes properties of the selected instance.
// The following properties must be included in any representation of SettingDataRequest:
//
// - ElementName(cannot be modified)
//
// - InstanceID (cannot be modified)
//
// - Enabled.
//
// This method will not modify the flash ("Enabled" property) until setupandconfiguration.CommitChanges() is issued and performed successfully.
// Overrides the generic Put because each TLS setting must be addressed by an
// InstanceID selector, which the generic Put does not provide.
func (settingData SettingData) Put(instanceID string, tlsSettingData SettingDataRequest) (response Response, err error) {
	tlsSettingData.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTTLSSettingData)
	selector := []message.Selector{{
		Name:  "InstanceID",
		Value: instanceID,
	}}
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.Base.Put(tlsSettingData, true, selector),
		},
	}
	// send the message to AMT
	err = settingData.Base.Execute(response.Message)
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
