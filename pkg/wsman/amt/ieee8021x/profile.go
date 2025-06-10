/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Profile struct {
	base.WSManService[Response]
}

// NewIEEE8021xProfileWithClient instantiates a new Profile service.
func NewIEEE8021xProfileWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Profile {
	return Profile{
		base.NewService[Response](wsmanMessageCreator, AMTIEEE8021xProfile, client),
	}
}
