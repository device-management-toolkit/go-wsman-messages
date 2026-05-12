/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package fan facilitates communication with Intel® AMT devices to represent
// the capabilities and management of a Fan CoolingDevice (CIM_Fan).
//
// CIM_Fan is part of the Hardware Asset feature and is compatible with Intel
// AMT releases 3.0, 3.2, 4.0, 5.0, 5.1, 6.0, 6.1, 6.2, 7.0, 8.0, 8.1, 9.0, 9.5,
// 10.0, 11.0 and later.
//
// Inheritance:
//
//	CIM_ManagedElement
//	  └── CIM_ManagedSystemElement
//	        └── CIM_LogicalElement
//	              └── CIM_EnabledLogicalElement
//	                    └── CIM_LogicalDevice
//	                          └── CIM_CoolingDevice
//	                                └── CIM_Fan
//
// AMT firmware translates the SMBIOS type 27 "Status" field bits 7:5 to the
// CIM_Fan OperationalStatus / HealthState values described on those fields'
// godoc.
//
// The service exposes the standard Get / Enumerate / Pull / Release surface
// inherited from base.WSManService. CIM_Fan also defines a SetSpeed method on
// the wire, but Intel AMT 5.1+ always returns 1 (NotSupported) for it, so this
// package intentionally does not expose a SetSpeed helper.
package fan

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// Device is the WSMAN service binding for CIM_Fan. Construct one via
// NewFanWithClient (or reach for the pre-wired field on cim.Messages).
//
// Permitted realms:
//   - Get / Enumerate / Pull: ADMIN_SECURITY_HARDWARE_ASSET_REALM,
//     ADMIN_SECURITY_GENERAL_INFO_REALM
type Device struct {
	base.WSManService[Response]
}

// NewFanWithClient returns a new instance of the Fan Device bound to the given
// WSMAN client and message creator.
func NewFanWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Device {
	return Device{
		base.NewService[Response](wsmanMessageCreator, CIMFan, client),
	}
}
