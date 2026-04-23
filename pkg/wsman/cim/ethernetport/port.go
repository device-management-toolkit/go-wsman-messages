/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package ethernetport facilitates communication with Intel® AMT devices to access CIM_EthernetPort data.
package ethernetport

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Port struct {
	base.WSManService[Response]
}

// NewEthernetPortWithClient instantiates a new EthernetPort service.
func NewEthernetPortWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Port {
	return Port{
		base.NewService[Response](wsmanMessageCreator, CIMEthernetPort, client),
	}
}
