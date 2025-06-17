/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package mps facilitiates communication with Intel® AMT devices to configure the username and password used to access an MPS.
package mps

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type UsernamePassword struct {
	base.WSManService[Response]
}

// NewMPSUsernamePasswordWithClient instantiates a new UsernamePassword.
func NewMPSUsernamePasswordWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) UsernamePassword {
	return UsernamePassword{
		base.NewService[Response](wsmanMessageCreator, AMTMPSUsernamePassword, client),
	}
}
