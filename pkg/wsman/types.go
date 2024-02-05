/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wsman

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips"
)

// Messages implements client.WSMan, amt.Messages, cim.Messages, and ips.Messages
type Messages struct {
	client client.WSMan
	AMT    amt.Messages
	CIM    cim.Messages
	IPS    ips.Messages
}
