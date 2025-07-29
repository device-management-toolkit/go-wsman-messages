/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ips

import (
	"reflect"
	"testing"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/alarmclock"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/hostbasedsetup"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/http"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/ieee8021x"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/optin"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/power"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestNewMessages(t *testing.T) {
	mock := wsmantesting.MockClient{}
	m := NewMessages(&mock)

	if m.wsmanMessageCreator == nil {
		t.Error("wsmanMessageCreator is not initialized")
	}

	if reflect.DeepEqual(m.OptInService, optin.Service{}) {
		t.Error("AlarmClockService is not initialized")
	}

	if reflect.DeepEqual(m.HostBasedSetupService, hostbasedsetup.Service{}) {
		t.Error("AuditLog is not initialized")
	}

	if reflect.DeepEqual(m.AlarmClockOccurrence, alarmclock.Occurrence{}) {
		t.Error("AuthorizationService is not initialized")
	}

	if reflect.DeepEqual(m.IEEE8021xCredentialContext, ieee8021x.CredentialContext{}) {
		t.Error("BootCapabilities is not initialized")
	}

	if reflect.DeepEqual(m.IEEE8021xSettings, ieee8021x.Settings{}) {
		t.Error("BootSettingData is not initialized")
	}

	if reflect.DeepEqual(m.PowerManagementService, power.ManagementService{}) {
		t.Error("PowerManagementService is not initialized")
	}

	if reflect.DeepEqual(m.HTTPProxyService, http.ProxyService{}) {
		t.Error("HTTPProxyService is not initialized")
	}
}
