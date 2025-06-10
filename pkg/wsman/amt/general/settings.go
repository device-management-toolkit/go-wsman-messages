/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package general facilitates communication with Intel® AMT to read and configure the device's Intel® AMT general settings.
package general

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Settings struct {
	base.WSManService[Response]
}

// NewGeneralSettingsWithClient instantiates a new General Settings service.
func NewGeneralSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Settings {
	return Settings{
		base.NewService[Response](wsmanMessageCreator, AMTGeneralSettings, client),
	}
}
