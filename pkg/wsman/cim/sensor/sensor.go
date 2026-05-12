/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package sensor facilitates communication with Intel® AMT devices to enumerate CIM_Sensor instances.
//
// A CIM_Sensor represents an entity capable of measuring or reporting the
// characteristics of some physical property — for example, the temperature or
// voltage characteristics of a managed computer system. On Intel® AMT, sensor
// instances are populated from the ASF_ALRT table in the ACPI "ASF!" namespace:
// each row of the ASF_ALERTDATA structure becomes one CIM_Sensor instance, and
// the EventSensorType field on that row drives the SensorType property
// (01h->02h Temperature, 02h->03h Voltage, 03h->04h Current).
//
// CIM_Sensor extends CIM_LogicalDevice and belongs to the Hardware Asset and
// Event Manager feature groups. It is supported on Intel® AMT releases 3.0
// through 11.0.
//
// Inheritance:
//
//	CIM_ManagedElement
//	  CIM_ManagedSystemElement
//	    CIM_LogicalElement
//	      CIM_EnabledLogicalElement
//	        CIM_LogicalDevice
//	          CIM_Sensor
//
// The Get, Pull, and Enumerate methods inherited from base.WSManService are
// reachable by any caller in the ADMIN_SECURITY_HARDWARE_ASSET_REALM or
// ADMIN_SECURITY_EVENT_MANAGER_REALM. Enumerate / Pull return only those
// instances the caller has permission to read.
package sensor

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// Package wraps base.WSManService[Response] to expose Get / Enumerate / Pull
// against the CIM_Sensor resource URI. It is named "Package" to stay parallel
// with the sibling CIM LogicalDevice services (chip.Package, processor.Package,
// physical.Memory, etc.) — these are all instance-bearing collections rather
// than singleton services.
type Package struct {
	base.WSManService[Response]
}

// NewSensorWithClient returns a Package bound to the supplied WSMAN message
// creator and transport client. The returned value's Get / Enumerate / Pull
// methods build CIM_Sensor-scoped envelopes and execute them through the
// client; callers do not need to construct headers manually.
func NewSensorWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base.NewService[Response](wsmanMessageCreator, CIMSensor, client),
	}
}
