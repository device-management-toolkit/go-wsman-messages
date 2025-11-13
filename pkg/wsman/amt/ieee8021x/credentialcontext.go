/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package ieee8021x facilitiates communication with Intel® AMT devices to access the ieee8021x credential context and profile settings
//
// CredentialContext gets the association between an instance of AMT_8021XProfile and an instance of AMT_PublicKeyCertificate that it uses.
//
// Profile represents a 802.1X profile in the Intel® AMT system.
package ieee8021x

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type CredentialContext struct {
	base.WSManService[Response]
}

// NewIEEE8021xCredentialContextWithClient instantiates a new CredentialContext service.
func NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) CredentialContext {
	return CredentialContext{
		base.NewService[Response](wsmanMessageCreator, AMTIEEE8021xCredentialContext, client),
	}
}
