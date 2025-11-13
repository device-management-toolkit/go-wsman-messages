/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package mediaaccess facilitiates communication with Intel® AMT devices to represent the ability to access one or more media and use this media to store and retrieve data.
package mediaaccess

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Device struct {
	base.WSManService[Response]
}

// NewMediaAccessDevice returns a new instance of the MediaAccessDevice struct.
func NewMediaAccessDeviceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Device {
	return Device{
		base.NewService[Response](wsmanMessageCreator, CIMMediaAccessDevice, client),
	}
}
