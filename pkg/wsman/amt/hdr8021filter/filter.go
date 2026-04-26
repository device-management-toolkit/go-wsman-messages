/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package hdr8021filter facilitates communication with Intel AMT devices for 802.1 filter data.
package hdr8021filter

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Service struct {
	base.WSManService[Response]
}

// NewServiceWithClient instantiates a new Hdr8021Filter service.
func NewServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base.NewService[Response](wsmanMessageCreator, AMTHdr8021Filter, client),
	}
}
