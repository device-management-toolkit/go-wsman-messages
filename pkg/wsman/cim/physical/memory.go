/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package physical facilitates communications with Intel® AMT devices to get the PhysicalMemory as a subclass of CIM_Chip, representing low level memory devices - SIMMS, DIMMs, raw memory chips, etc.
package physical

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Memory struct {
	base.WSManService[Response]
}

// NewPhysicalMemory returns a new instance of the PhysicalMemory struct.
func NewPhysicalMemoryWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Memory {
	return Memory{
		base.NewService[Response](wsmanMessageCreator, CIMPhysicalMemory, client),
	}
}
