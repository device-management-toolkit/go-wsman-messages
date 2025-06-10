/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
// Package 8021xCredentialContext facilitates communication with Intel® AMT devices to create an association between an instance of IPS_IEEE8021xSettings and an instance of AMT_PublicKeyCertificate that it uses.
package ieee8021x

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type CredentialContext struct {
	base.WSManService[Response]
}

// NewIEEE8021xCredentialContext returns a new instance of the IPS_8021xCredentialContext struct.
func NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) CredentialContext {
	return CredentialContext{
		base.NewService[Response](wsmanMessageCreator, IPS8021xCredentialContext, client),
	}
}
