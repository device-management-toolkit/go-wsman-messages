/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package boot facilitates communication with Intel® AMT devices to access the boot capabilities and boot setting data.
//
// Capabilities reports what boot options that the Intel® AMT device supports.
//
// SettingData provides configuration-related and operational parameters for the boot service in the Intel® AMT device.  In order to activate these settings use [pkg/github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/power] RequestPowerStateChange().  Notice that you can't set certain values while others are enabled (for example: You can't set UseIDER or UseSOL if a CIM_BootSourceSetting is chosen).
package boot

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Capabilities struct {
	base.WSManService[Response]
}

// NewBootCapabilitiesWithClient instantiates a new Boot Capabilities service.
func NewBootCapabilitiesWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Capabilities {
	return Capabilities{
		base.NewService[Response](wsmanMessageCreator, AMTBootCapabilities, client),
	}
}
