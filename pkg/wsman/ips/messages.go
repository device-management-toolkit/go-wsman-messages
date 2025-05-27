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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/ieee8021x"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/kvmredirection"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/optin"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/power"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/screensetting"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/secio"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

type Messages struct {
	wsmanMessageCreator        *message.WSManMessageCreator
	OptInService               optin.Service
	HostBasedSetupService      hostbasedsetup.Service
	AlarmClockOccurrence       alarmclock.Occurrence
	IEEE8021xCredentialContext ieee8021x.CredentialContext
	IEEE8021xSettings          ieee8021x.Settings
	PowerManagementService     power.ManagementService
	ScreenSettingData          screensetting.Data
	SecIOService               secio.Service
	KVMRedirectionSettingData  kvmredirection.SettingData
}

func NewMessages(client client.WSMan) Messages {
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	m := Messages{
		wsmanMessageCreator: wsmanMessageCreator,
	}
	m.OptInService = optin.NewOptInServiceWithClient(wsmanMessageCreator, client)
	m.HostBasedSetupService = hostbasedsetup.NewHostBasedSetupServiceWithClient(wsmanMessageCreator, client)
	m.AlarmClockOccurrence = alarmclock.NewAlarmClockOccurrenceWithClient(wsmanMessageCreator, client)
	m.IEEE8021xCredentialContext = ieee8021x.NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator, client)
	m.IEEE8021xSettings = ieee8021x.NewIEEE8021xSettingsWithClient(wsmanMessageCreator, client)
	m.PowerManagementService = power.NewPowerManagementServiceWithClient(wsmanMessageCreator, client)
	m.ScreenSettingData = screensetting.NewScreenSettingDataWithClient(wsmanMessageCreator, client)
	m.SecIOService = secio.NewSecIOServiceWithClient(wsmanMessageCreator, client)
	m.KVMRedirectionSettingData = kvmredirection.NewKVMRedirectionSettingDataWithClient(wsmanMessageCreator, client)

	return m
}
