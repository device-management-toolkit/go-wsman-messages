/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package processor facilitates communication with Intel® AMT devices capabilities and management of the Processor LogicalDevice
package processor

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Package struct {
	base.WSManService[Response]
}

// NewProcessor returns a new instance of the Processor struct.
func NewProcessorWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Package {
	return Package{
		base.NewService[Response](wsmanMessageCreator, CIMProcessor, client),
	}
}
