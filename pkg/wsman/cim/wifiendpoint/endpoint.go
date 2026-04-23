/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package wifiendpoint implements CIM_WiFiEndpoint.
//
// CIM_WiFiEndpoint: A CIM_WiFiEndpoint is a CIM_ProtocolEndpoint that has the capability to utilize IEEE 802.11-based communication technology.
package wifiendpoint

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Endpoint struct {
	base.WSManService[Response]
}

// NewWiFiEndpointWithClient returns a new instance of the WiFiEndpoint struct.
func NewWiFiEndpointWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Endpoint {
	return Endpoint{
		base.NewService[Response](wsmanMessageCreator, CIMWiFiEndpoint, client),
	}
}
