/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package wifiendpointcapabilities implements CIM_WiFiEndpointCapabilities.
//
// CIM_WiFiEndpointCapabilities: Represents the capabilities of a CIM_WiFiEndpoint.
package wifiendpointcapabilities

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type EndpointCapabilities struct {
	base.WSManService[Response]
}

// NewWiFiEndpointCapabilitiesWithClient returns a new instance of the EndpointCapabilities struct.
func NewWiFiEndpointCapabilitiesWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) EndpointCapabilities {
	return EndpointCapabilities{
		base.NewService[Response](wsmanMessageCreator, CIMWiFiEndpointCapabilities, client),
	}
}
