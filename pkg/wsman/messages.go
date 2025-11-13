/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package wsman facilitates access to AMT, CIM, and IPS classes for communication with Intel® AMT devices.
package wsman

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips"
)

// NewMessages instantiates a new Messages class with client connection parameters.
func NewMessages(cp client.Parameters) Messages {
	var client1 *client.Target

	if cp.IsRedirection {
		client1 = client.NewWsmanTCP(cp)
	} else {
		client1 = client.NewWsman(cp)
	}

	m := Messages{
		Client: client1,
	}

	m.AMT = amt.NewMessages(client1)
	m.CIM = cim.NewMessages(client1)
	m.IPS = ips.NewMessages(client1)

	return m
}
