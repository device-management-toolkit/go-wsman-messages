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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type SettingData struct {
	base.WSManService[Response]
}

// NewKerberosSettingDataWithClient instantiates a new kerberos SettingData.
func NewKerberosSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base.NewService[Response](wsmanMessageCreator, AMTKerberosSettingData, client),
	}
}

// GetCredentialCacheState gets the current state of the credential caching functionality.
func (settingData SettingData) GetCredentialCacheState() (response Response, err error) {
	header := settingData.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTKerberosSettingData, GetCredentialCacheState), AMTKerberosSettingData, nil, "", "")
	body := settingData.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetCredentialCacheState), AMTKerberosSettingData, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: settingData.Base.WSManMessageCreator.CreateXML(header, body),
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

// SetCredentialCacheState enables/disables the credential caching functionality
// TODO: Current gets SOAP schema violation from AMT.
func (settingData SettingData) SetCredentialCacheState(enabled bool) (response Response, err error) {
	credentialCasheState := SetCredentialCacheStateInput{
		H:       fmt.Sprintf("%s%s", message.AMTSchema, AMTKerberosSettingData),
		Enabled: enabled,
	}
	header := settingData.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMTKerberosSettingData, SetCredentialCacheState), AMTKerberosSettingData, nil, "", "")
	body := settingData.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetCredentialCacheState), AMTKerberosSettingData, credentialCasheState)

	response = Response{
		Message: &client.Message{
			XMLInput: settingData.Base.WSManMessageCreator.CreateXML(header, body),
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
