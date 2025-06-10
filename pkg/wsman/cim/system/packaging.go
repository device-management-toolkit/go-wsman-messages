/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package system facilitates communication with Intel® AMT devices in a way similar to the way that LogicalDevices are 'Realized' by PhysicalElements, Systems can be associated with specific packaging or PhysicalElements.
//
// This association explicitly defines the relationship between a System and its packaging.
package system

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Package struct {
	base.WSManService[Response]
}

// NewSystemPackaging returns a new instance of the SystemPackaging struct.
func NewSystemPackageWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base.NewService[Response](wsmanMessageCreator, CIMSystemPackaging, client),
	}
}
