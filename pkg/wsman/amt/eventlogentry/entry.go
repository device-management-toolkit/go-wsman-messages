/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package eventlogentry facilitates communication with Intel AMT devices for event log entry data.
package eventlogentry

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Service struct {
	base.WSManService[Response]
}

// NewServiceWithClient instantiates a new Event Log Entry service.
func NewServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base.NewService[Response](wsmanMessageCreator, AMTEventLogEntry, client),
	}
}
