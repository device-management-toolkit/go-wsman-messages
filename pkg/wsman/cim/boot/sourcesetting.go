/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type SourceSetting struct {
	base.WSManService[Response]
}

// NewBootSourceSetting returns a new instance of the BootSourceSetting struct.
func NewBootSourceSettingWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SourceSetting {
	return SourceSetting{
		base.NewService[Response](wsmanMessageCreator, CIMBootSourceSetting, client),
	}
}
