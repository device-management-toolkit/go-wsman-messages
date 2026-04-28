/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Endpoint struct {
	base.WSManService[Response]
}

// NewWiFiEndpointWithClient returns a new instance of the WiFiEndpoint struct.
func NewWiFiEndpointWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Endpoint {
	return Endpoint{
		base.NewService[Response](wsmanMessageCreator, CIMWiFiEndpoint, client),
	}
}
