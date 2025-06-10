/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type ProtocolEndpointCollection struct {
	base.WSManService[Response]
}

// NewTLSProtocolEndpointCollectionWithClient instantiates a new ProtocolEndpointCollection.
func NewTLSProtocolEndpointCollectionWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) ProtocolEndpointCollection {
	return ProtocolEndpointCollection{
		base.NewService[Response](wsmanMessageCreator, AMTTLSProtocolEndpointCollection, client),
	}
}
