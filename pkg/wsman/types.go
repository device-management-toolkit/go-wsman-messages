/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wsman

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips"
)

// Messages implements client.WSMan, amt.Messages, cim.Messages, and ips.Messages.
type Messages struct {
	Client client.WSMan
	AMT    amt.Messages
	CIM    cim.Messages
	IPS    ips.Messages
}
