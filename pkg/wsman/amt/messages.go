/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package amt implements AMT classes to support communicating with Intel® AMT Devices.
package amt

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/alarmclock"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/asset"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/auditlog"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/authorization"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/boot"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/cryptographiccapabilities"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/environmentdetection"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/ethernetport"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/eventlogentry"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/general"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/hdr8021filter"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/ieee8021x"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/kerberos"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/managementpresence"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/messagelog"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/mps"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/publickey"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/publicprivate"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/redirection"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/remoteaccess"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/setupandconfiguration"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/systempowerscheme"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/timesynchronization"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/tls"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/userinitiatedconnection"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/wifiportconfiguration"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// Messages contains the supported AMT classes.
type Messages struct {
	wsmanMessageCreator             *message.WSManMessageCreator
	AlarmClockService               alarmclock.Service
	AssetTable                      asset.Table
	AssetTableService               asset.Service
	AuditLog                        auditlog.Service
	AuthorizationService            authorization.Service
	BootCapabilities                boot.Capabilities
	CryptographicCapabilities       cryptographiccapabilities.Service
	BootSettingData                 boot.SettingData
	EventLogEntry                   eventlogentry.Service
	EnvironmentDetectionSettingData environmentdetection.SettingData
	EthernetPortSettings            ethernetport.Settings
	GeneralSettings                 general.Settings
	Hdr8021Filter                   hdr8021filter.Service
	IEEE8021xCredentialContext      ieee8021x.CredentialContext
	IEEE8021xProfile                ieee8021x.Profile
	KerberosSettingData             kerberos.SettingData
	ManagementPresenceRemoteSAP     managementpresence.RemoteSAP
	MessageLog                      messagelog.Service
	MPSUsernamePassword             mps.UsernamePassword
	PublicKeyCertificate            publickey.Certificate
	PublicKeyManagementService      publickey.ManagementService
	PublicPrivateKeyPair            publicprivate.KeyPair
	RedirectionService              redirection.Service
	RemoteAccessCapabilities        remoteaccess.Capabilities
	RemoteAccessPolicyAppliesToMPS  remoteaccess.PolicyAppliesToMPS
	RemoteAccessPolicyRule          remoteaccess.PolicyRule
	RemoteAccessService             remoteaccess.Service
	SetupAndConfigurationService    setupandconfiguration.Service
	SystemPowerScheme               systempowerscheme.Service
	TimeSynchronizationService      timesynchronization.Service
	TLSCredentialContext            tls.CredentialContext
	TLSProtocolEndpointCollection   tls.ProtocolEndpointCollection
	TLSSettingData                  tls.SettingData
	UserInitiatedConnectionService  userinitiatedconnection.Service
	WiFiPortConfigurationService    wifiportconfiguration.Service
}

// NewMessages instantiates a new instance of amt Messages.
func NewMessages(client client.WSMan) Messages {
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	m := Messages{
		wsmanMessageCreator: wsmanMessageCreator,
	}
	m.AlarmClockService = alarmclock.NewServiceWithClient(wsmanMessageCreator, client)
	m.AssetTable = asset.NewTableWithClient(wsmanMessageCreator, client)
	m.AssetTableService = asset.NewServiceWithClient(wsmanMessageCreator, client)
	m.AuditLog = auditlog.NewAuditLogWithClient(wsmanMessageCreator, client)
	m.AuthorizationService = authorization.NewServiceWithClient(wsmanMessageCreator, client)
	m.BootCapabilities = boot.NewBootCapabilitiesWithClient(wsmanMessageCreator, client)
	m.CryptographicCapabilities = cryptographiccapabilities.NewServiceWithClient(wsmanMessageCreator, client)
	m.BootSettingData = boot.NewBootSettingDataWithClient(wsmanMessageCreator, client)
	m.EventLogEntry = eventlogentry.NewServiceWithClient(wsmanMessageCreator, client)
	m.EnvironmentDetectionSettingData = environmentdetection.NewEnvironmentDetectionSettingDataWithClient(wsmanMessageCreator, client)
	m.EthernetPortSettings = ethernetport.NewEthernetPortSettingsWithClient(wsmanMessageCreator, client)
	m.GeneralSettings = general.NewGeneralSettingsWithClient(wsmanMessageCreator, client)
	m.Hdr8021Filter = hdr8021filter.NewServiceWithClient(wsmanMessageCreator, client)
	m.IEEE8021xCredentialContext = ieee8021x.NewIEEE8021xCredentialContextWithClient(wsmanMessageCreator, client)
	m.IEEE8021xProfile = ieee8021x.NewIEEE8021xProfileWithClient(wsmanMessageCreator, client)
	m.KerberosSettingData = kerberos.NewKerberosSettingDataWithClient(wsmanMessageCreator, client)
	m.ManagementPresenceRemoteSAP = managementpresence.NewManagementPresenceRemoteSAPWithClient(wsmanMessageCreator, client)
	m.MessageLog = messagelog.NewMessageLogWithClient(wsmanMessageCreator, client)
	m.MPSUsernamePassword = mps.NewMPSUsernamePasswordWithClient(wsmanMessageCreator, client)
	m.PublicKeyCertificate = publickey.NewPublicKeyCertificateWithClient(wsmanMessageCreator, client)
	m.PublicKeyManagementService = publickey.NewPublicKeyManagementServiceWithClient(wsmanMessageCreator, client)
	m.PublicPrivateKeyPair = publicprivate.NewPublicPrivateKeyPairWithClient(wsmanMessageCreator, client)
	m.RedirectionService = redirection.NewRedirectionServiceWithClient(wsmanMessageCreator, client)
	m.RemoteAccessCapabilities = remoteaccess.NewCapabilitiesWithClient(wsmanMessageCreator, client)
	m.RemoteAccessPolicyAppliesToMPS = remoteaccess.NewRemoteAccessPolicyAppliesToMPSWithClient(wsmanMessageCreator, client)
	m.RemoteAccessPolicyRule = remoteaccess.NewPolicyRuleWithClient(wsmanMessageCreator, client)
	m.RemoteAccessService = remoteaccess.NewRemoteAccessServiceWithClient(wsmanMessageCreator, client)
	m.SetupAndConfigurationService = setupandconfiguration.NewSetupAndConfigurationServiceWithClient(wsmanMessageCreator, client)
	m.SystemPowerScheme = systempowerscheme.NewServiceWithClient(wsmanMessageCreator, client)
	m.TimeSynchronizationService = timesynchronization.NewTimeSynchronizationServiceWithClient(wsmanMessageCreator, client)
	m.TLSCredentialContext = tls.NewTLSCredentialContextWithClient(wsmanMessageCreator, client)
	m.TLSProtocolEndpointCollection = tls.NewTLSProtocolEndpointCollectionWithClient(wsmanMessageCreator, client)
	m.TLSSettingData = tls.NewTLSSettingDataWithClient(wsmanMessageCreator, client)
	m.UserInitiatedConnectionService = userinitiatedconnection.NewUserInitiatedConnectionServiceWithClient(wsmanMessageCreator, client)
	m.WiFiPortConfigurationService = wifiportconfiguration.NewWiFiPortConfigurationServiceWithClient(wsmanMessageCreator, client)

	return m
}
