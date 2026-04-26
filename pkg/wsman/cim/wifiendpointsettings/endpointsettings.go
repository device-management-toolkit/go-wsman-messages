/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package wifiendpointsettings implements CIM_WiFiEndpointSettings.
//
// CIM_WiFiEndpointSettings: A class derived from SettingData that can be applied to an instance of
// CIM_WiFiEndpoint to enable it to associate to a particular Wi-Fi network.
package wifiendpointsettings

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type EndpointSettings struct {
	base.WSManService[Response]
}

// NewWiFiEndpointSettingsWithClient returns a new instance of the EndpointSettings struct.
func NewWiFiEndpointSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) EndpointSettings {
	return EndpointSettings{
		base.NewService[Response](wsmanMessageCreator, CIMWiFiEndpointSettings, client),
	}
}

// Delete removes the specified instance.
func (endpointSettings EndpointSettings) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "InstanceID", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: endpointSettings.Base.Delete(selector),
		},
	}

	err = endpointSettings.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
