/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ips

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/alarmclock"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/hostbasedsetup"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/hostbootreason"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/hostipsettings"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/http"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/ieee8021x"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/ipv6portsettings"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/kvmredirection"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/lanendpoint"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/optin"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/power"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/provisioningrecordlog"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/screensetting"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/secio"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

type Messages struct {
	wsmanMessageCreator         *message.WSManMessageCreator
	OptInService                optin.Service
	HostBasedSetupService       hostbasedsetup.Service
	HostBootReason              hostbootreason.Service
	HostIPSettings              hostipsettings.Settings
	IPv6PortSettings            ipv6portsettings.Settings
	LANEndpoint                 lanendpoint.Endpoint
	ProvisioningRecordLog       provisioningrecordlog.Log
	AlarmClockOccurrence        alarmclock.Occurrence
	IEEE8021xCredentialContext  ieee8021x.CredentialContext
	IEEE8021xSettings           ieee8021x.Settings
	PowerManagementService      power.ManagementService
	ScreenSettingData           screensetting.Data
	SecIOService                secio.Service
	KVMRedirectionSettingData   kvmredirection.SettingData
	HTTPProxyService            http.ProxyService
	HTTPProxyAccessPointService http.ProxyAccessPointService
}

func NewMessages(client client.WSMan) Messages {
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	m := Messages{
		wsmanMessageCreator: wsmanMessageCreator,
	}
	m.OptInService = optin.NewOptInServiceWithClient(wsmanMessageCreator, client)
	m.HostBasedSetupService = hostbasedsetup.NewHostBasedSetupServiceWithClient(wsmanMessageCreator, client)
	m.HostBootReason = hostbootreason.NewHostBootReasonWithClient(wsmanMessageCreator, client)
	m.HostIPSettings = hostipsettings.NewHostIPSettingsWithClient(wsmanMessageCreator, client)
	m.IPv6PortSettings = ipv6portsettings.NewIPv6PortSettingsWithClient(wsmanMessageCreator, client)
	m.LANEndpoint = lanendpoint.NewLANEndpointWithClient(wsmanMessageCreator, client)
	m.ProvisioningRecordLog = provisioningrecordlog.NewProvisioningRecordLogWithClient(wsmanMessageCreator, client)
	m.AlarmClockOccurrence = alarmclock.NewAlarmClockOccurrenceWithClient(wsmanMessageCreator, client)
	m.IEEE8021xCredentialContext = ieee8021x.NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator, client)
	m.IEEE8021xSettings = ieee8021x.NewIEEE8021xSettingsWithClient(wsmanMessageCreator, client)
	m.PowerManagementService = power.NewPowerManagementServiceWithClient(wsmanMessageCreator, client)
	m.ScreenSettingData = screensetting.NewScreenSettingDataWithClient(wsmanMessageCreator, client)
	m.SecIOService = secio.NewSecIOServiceWithClient(wsmanMessageCreator, client)
	m.KVMRedirectionSettingData = kvmredirection.NewKVMRedirectionSettingDataWithClient(wsmanMessageCreator, client)
	m.HTTPProxyService = http.NewHTTPProxyServiceWithClient(wsmanMessageCreator, client)
	m.HTTPProxyAccessPointService = http.NewHTTPProxyAccessPointServiceWithClient(wsmanMessageCreator, client)

	return m
}
