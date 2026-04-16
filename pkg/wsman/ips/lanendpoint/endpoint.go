/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package lanendpoint facilitates communication with Intel(R) AMT devices for IPS_LANEndpoint data.
package lanendpoint

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Endpoint struct {
	base.WSManService[Response]
}

// NewLANEndpointWithClient creates a new IPS_LANEndpoint service.
func NewLANEndpointWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Endpoint {
	return Endpoint{
		base.NewService[Response](wsmanMessageCreator, IPSLANEndpoint, client),
	}
}
