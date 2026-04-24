/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package provisioningrecordlog facilitates communication with Intel(R) AMT devices for IPS_ProvisioningRecordLog data.
package provisioningrecordlog

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Log struct {
	base.WSManService[Response]
}

// NewProvisioningRecordLogWithClient creates a new IPS_ProvisioningRecordLog service.
func NewProvisioningRecordLogWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Log {
	return Log{
		base.NewService[Response](wsmanMessageCreator, IPSProvisioningRecordLog, client),
	}
}
