/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package service facilitates communication with Intel® AMT devices to convey the semantics of a Service that is available for the use of a ManagedElement.
//
// An example of an available Service is that a Processor and an enclosure (a PhysicalElement) can use AlertOnLAN Services to signal an incomplete or erroneous boot.
//
// In reality, AlertOnLAN is simply a HostedService on a computer system that is generally available for use and is not a dependency of the processor or enclosure.
//
// To describe that the use of this service might be restricted or have limited availability or applicability, the CIM_ServiceAvailableToElement association would be instantiated between the Service and specific CIM_Processors and CIM_Chassis.
package service

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type AvailableToElement struct {
	base.WSManService[Response]
}

// NewServiceAvailableToElement returns a new instance of the ServiceAvailableToElement struct.
func NewServiceAvailableToElementWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) AvailableToElement {
	return AvailableToElement{
		base.NewService[Response](wsmanMessageCreator, CIMServiceAvailableToElement, client),
	}
}
