/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package computer facilitates communication with Intel® AMT devices in a way similar to the way that LogicalDevices are 'Realized' by PhysicalElements, however ComputerSystem may be realized realized in one or more PhysicalPackages.
//
// The ComputerSystemPackage association explicitly defines this relationship.
package computer

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type SystemPackage struct {
	base.WSManService[Response]
}

// NewComputerSystemPackage returns a new instance of the ComputerSystemPackage struct.
func NewComputerSystemPackageWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SystemPackage {
	return SystemPackage{
		base.NewService[Response](wsmanMessageCreator, CIMComputerSystemPackage, client),
	}
}
