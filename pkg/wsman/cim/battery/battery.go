/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package battery models the DMTF CIM_Battery class as exposed by
// Intel(R) AMT firmware over WS-Management. CIM_Battery describes the
// capabilities and management of a battery — typically a laptop battery,
// but the schema also covers other internal or external batteries such as
// a UPS.
//
// Compatible with Intel AMT releases 6.2, 7.0, 8.0, 8.1, 9.0, 9.5, 10.0,
// and 11.0. CIM schema version 2.26.0 (DMTF UML path
// CIM::Device::CoolingAndPower).
//
// Schema inheritance:
//
//	CIM_ManagedElement
//	  CIM_ManagedSystemElement
//	    CIM_LogicalElement
//	      CIM_EnabledLogicalElement
//	        CIM_LogicalDevice
//	          CIM_Battery
//
// Supported operations:
//
//   - Get       — retrieve the battery instance. Permitted realm:
//     ADMIN_SECURITY_HARDWARE_ASSET_REALM.
//   - Enumerate — open an enumeration of CIM_Battery instances. Open to
//     all authenticated users; results are filtered to the
//     instances the caller is permitted to see.
//   - Pull      — fetch instances against the open enumeration context.
//   - Release   — release the enumeration context (handled implicitly by
//     the underlying WS-Management transport).
//
// See the Intel AMT SDK for the firmware-side reference:
// https://www.intel.com/content/www/us/en/developer/tools/active-management-technology-sdk/overview.html
package battery

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// Package is the CIM_Battery WS-Management service. It exposes Get,
// Enumerate, and Pull via the generic base.WSManService[Response].
type Package struct {
	base.WSManService[Response]
}

// NewBatteryWithClient constructs a battery service bound to the given
// WS-Management message creator and transport client. The returned
// Package is safe to share across goroutines provided the underlying
// client is.
func NewBatteryWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base.NewService[Response](wsmanMessageCreator, CIMBattery, client),
	}
}
