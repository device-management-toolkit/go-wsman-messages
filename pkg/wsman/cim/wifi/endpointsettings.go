/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package WiFi facilitates communication with Intel® AMT devices to access and configure WiFi Endpoint Settings and WiFi Port features of AMT.
//
// WiFiEndpointSettings:
// A class derived from SettingData that can be applied to an instance of CIM_WiFiEndpoint to enable it to associate to a particular Wi-Fi network.
//
// WiFiPort:
// A class derived from NetworkPort that provides the logical representation of wireless local area network communications hardware that conforms to the IEEE 802.11 series of specifications.
// It embodies properties at the lowest layers of a network stack, such as the antennas used for transmission and reception, the address permanently embedded into the hardware, and the operational bandwidth of the device.
package wifi

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type EndpointSettings struct {
	base.WSManService[Response]
}

// NewWiFiEndpointSettings returns a new instance of the WiFiEndpointSettings struct.
func NewWiFiEndpointSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) EndpointSettings {
	return EndpointSettings{
		base.NewService[Response](wsmanMessageCreator, CIMWiFiEndpointSettings, client),
	}
}

// Delete removes a the specified instance.
func (endpointSettings EndpointSettings) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "InstanceID", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: endpointSettings.Base.Delete(selector),
		},
	}

	err = endpointSettings.Base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
