/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package associatedpower facilitates communication with Intel® AMT devices to access the association between a Managed System Element and its power management service.
//
// CIM_AssociatedPowerManagementService provides details about the power state of the system,
// including the AvailableRequestedPowerStates property which contains the set of states the system can transition to.
package associatedpower

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type ManagementService struct {
	base.WSManService[Response]
}

// NewAssociatedPowerManagementServiceWithClient returns a new instance of the ManagementService struct.
func NewAssociatedPowerManagementServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ManagementService {
	return ManagementService{
		base.NewService[Response](wsmanMessageCreator, CIMAssociatedPowerManagementService, client),
	}
}
