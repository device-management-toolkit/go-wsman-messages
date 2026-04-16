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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/hostbootreason"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/hostipsettings"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/http"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/ieee8021x"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/ipv6portsettings"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/lanendpoint"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/optin"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/power"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/provisioningrecordlog"
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

	if reflect.DeepEqual(m.HostBootReason, hostbootreason.Service{}) {
		t.Error("HostBootReason is not initialized")
	}

	if reflect.DeepEqual(m.HostIPSettings, hostipsettings.Settings{}) {
		t.Error("HostIPSettings is not initialized")
	}

	if reflect.DeepEqual(m.IPv6PortSettings, ipv6portsettings.Settings{}) {
		t.Error("IPv6PortSettings is not initialized")
	}

	if reflect.DeepEqual(m.LANEndpoint, lanendpoint.Endpoint{}) {
		t.Error("LANEndpoint is not initialized")
	}

	if reflect.DeepEqual(m.ProvisioningRecordLog, provisioningrecordlog.Log{}) {
		t.Error("ProvisioningRecordLog is not initialized")
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
