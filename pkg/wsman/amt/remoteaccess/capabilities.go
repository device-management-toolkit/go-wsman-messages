/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// Capabilities provides access to AMT_RemoteAccessCapabilities.
type Capabilities struct {
	base.WSManService[Response]
}

// NewCapabilitiesWithClient instantiates a new capabilities service.
func NewCapabilitiesWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Capabilities {
	return Capabilities{
		base.NewService[Response](wsmanMessageCreator, AMTRemoteAccessCapabilities, client),
	}
}
