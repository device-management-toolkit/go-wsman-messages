/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package hostipsettings facilitates communication with Intel(R) AMT devices for IPS_HostIPSettings data.
package hostipsettings

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Settings struct {
	base.WSManService[Response]
}

// NewHostIPSettingsWithClient creates a new IPS_HostIPSettings service.
func NewHostIPSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base.NewService[Response](wsmanMessageCreator, IPSHostIPSettings, client),
	}
}
