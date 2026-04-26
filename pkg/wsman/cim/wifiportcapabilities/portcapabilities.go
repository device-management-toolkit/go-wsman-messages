/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package wifiportcapabilities implements CIM_WiFiPortCapabilities.
//
// CIM_WiFiPortCapabilities: Represents the capabilities of a CIM_WiFiPort.
package wifiportcapabilities

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type PortCapabilities struct {
	base.WSManService[Response]
}

// NewWiFiPortCapabilitiesWithClient returns a new instance of the PortCapabilities struct.
func NewWiFiPortCapabilitiesWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) PortCapabilities {
	return PortCapabilities{
		base.NewService[Response](wsmanMessageCreator, CIMWiFiPortCapabilities, client),
	}
}
