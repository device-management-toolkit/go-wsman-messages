/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package powermanagementcapabilities facilitates communication with Intel® AMT devices to access CIM_PowerManagementCapabilities data.
package powermanagementcapabilities

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Capabilities struct {
	base.WSManService[Response]
}

// NewPowerManagementCapabilitiesWithClient instantiates a new Capabilities service.
func NewPowerManagementCapabilitiesWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Capabilities {
	return Capabilities{
		base.NewService[Response](wsmanMessageCreator, CIMPowerManagementCapabilities, client),
	}
}
