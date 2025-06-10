/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package chip facilitates communication with Intel® AMT devices to represent any type of integrated circuit hardware, including ASICs, processors, memory chips, etc.
package chip

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Package struct {
	base.WSManService[Response]
}

// NewChip returns a new instance of the Chip struct.
func NewChipWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base.NewService[Response](wsmanMessageCreator, CIMChip, client),
	}
}
