/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package bios facilitiates communication with Intel® AMT devices to get information about the device bios element
package bios

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Element struct {
	base.WSManService[Response]
}

// NewBIOSElementWithClient instantiates a new Element.
func NewBIOSElementWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Element {
	return Element{
		base.NewService[Response](wsmanMessageCreator, CIMBIOSElement, client),
	}
}
