/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package kerberos facilitiates communication with Intel® AMT devices to access the configuration-related and operational parameters for the kerberos service in the Intel® AMT.
package kerberos

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/methods"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewKerberosSettingDataWithClient instantiates a new kerberos SettingData.
func NewKerberosSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTKerberosSettingData, client),
	}
}

// Get retrieves the representation of the instance.
func (settingData SettingData) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Get(nil),
		},
	}
	// send the message to AMT
	err = settingData.base.Execute(response.Message)
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
func (settingData SettingData) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = settingData.base.Execute(response.Message)
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
func (settingData SettingData) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = settingData.base.Execute(response.Message)
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

// GetCredentialCacheState gets the current state of the credential caching functionality.
func (settingData SettingData) GetCredentialCacheState() (response Response, err error) {
	header := settingData.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTKerberosSettingData, GetCredentialCacheState), AMTKerberosSettingData, nil, "", "")
	body := settingData.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetCredentialCacheState), AMTKerberosSettingData, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = settingData.base.Execute(response.Message)
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

// SetCredentialCacheState enables/disables the credential caching functionality
// TODO: Current gets SOAP schema violation from AMT.
func (settingData SettingData) SetCredentialCacheState(enabled bool) (response Response, err error) {
	credentialCasheState := SetCredentialCacheStateInput{
		H:       fmt.Sprintf("%s%s", message.AMTSchema, AMTKerberosSettingData),
		Enabled: enabled,
	}
	header := settingData.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTKerberosSettingData, SetCredentialCacheState), AMTKerberosSettingData, nil, "", "")
	body := settingData.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetCredentialCacheState), AMTKerberosSettingData, credentialCasheState)

	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = settingData.base.Execute(response.Message)
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
