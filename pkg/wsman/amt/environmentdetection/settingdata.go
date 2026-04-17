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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewEnvironmentDetectionSettingDataWithClient instantiates a new Environment Detection Setting Data service.
func NewEnvironmentDetectionSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base.NewService[Response](wsmanMessageCreator, AMTEnvironmentDetectionSettingData, client),
	}
}

// Put overrides the generic Put because this instance must be addressed by a
// specific InstanceID selector ("Intel(r) AMT Environment Detection Settings"),
// which the generic Put does not provide.
func (sd SettingData) Put(environmentDetectionSettingData EnvironmentDetectionSettingDataRequest) (response Response, err error) {
	environmentDetectionSettingData.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTEnvironmentDetectionSettingData)
	selector := []message.Selector{{
		Name:  "InstanceID",
		Value: "Intel(r) AMT Environment Detection Settings",
	}}
	response = Response{
		Message: &client.Message{
			XMLInput: sd.Base.Put(environmentDetectionSettingData, true, selector),
		},
	}
	// send the message to AMT
	err = sd.Base.Execute(response.Message)
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
