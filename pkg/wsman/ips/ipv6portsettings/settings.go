/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package ipv6portsettings facilitates communication with Intel(R) AMT devices for IPS_IPv6PortSettings data.
package ipv6portsettings

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Settings struct {
	base.WSManService[Response]
}

// NewIPv6PortSettingsWithClient creates a new IPS_IPv6PortSettings service.
func NewIPv6PortSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base.NewService[Response](wsmanMessageCreator, IPSIPv6PortSettings, client),
	}
}
