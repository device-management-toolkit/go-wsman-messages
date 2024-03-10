/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wsman

import (
	"reflect"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/amt"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips"
)

func TestNewMessages(t *testing.T) {
	clientParams := client.Parameters{
		Target:            "test",
		Username:          "username",
		Password:          "testPassword",
		UseDigest:         true,
		UseTLS:            true,
		SelfSignedAllowed: true,
		LogAMTMessages:    true,
	}
	m := NewMessages(clientParams)

	if m.client == nil {
		t.Error("client is not initialized")
	}
	if reflect.DeepEqual(m.AMT, amt.Messages{}) {
		t.Error("AMT is not initialized")
	}
	if reflect.DeepEqual(m.CIM, cim.Messages{}) {
		t.Error("CIM is not initialized")
	}
	if reflect.DeepEqual(m.IPS, ips.Messages{}) {
		t.Error("IPS is not initialized")
	}
}
