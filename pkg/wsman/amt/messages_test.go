/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package amt

import (
	"reflect"
	"testing"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/alarmclock"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/auditlog"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/authorization"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/boot"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/environmentdetection"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/ethernetport"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/general"
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
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/timesynchronization"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/tls"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/userinitiatedconnection"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/wifiportconfiguration"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestNewMessages(t *testing.T) {
	mock := wsmantesting.MockClient{}
	m := NewMessages(&mock)

	if m.wsmanMessageCreator == nil {
		t.Error("wsmanMessageCreator is not initialized")
	}

	if reflect.DeepEqual(m.AlarmClockService, alarmclock.Service{}) {
		t.Error("AlarmClockService is not initialized")
	}

	if reflect.DeepEqual(m.AuditLog, auditlog.Service{}) {
		t.Error("AuditLog is not initialized")
	}

	if reflect.DeepEqual(m.AuthorizationService, authorization.Service{}) {
		t.Error("AuthorizationService is not initialized")
	}

	if reflect.DeepEqual(m.BootCapabilities, boot.Capabilities{}) {
		t.Error("BootCapabilities is not initialized")
	}

	if reflect.DeepEqual(m.BootSettingData, boot.SettingData{}) {
		t.Error("BootSettingData is not initialized")
	}

	if reflect.DeepEqual(m.EnvironmentDetectionSettingData, environmentdetection.SettingData{}) {
		t.Error("EnvironmentDetectionSettingData is not initialized")
	}

	if reflect.DeepEqual(m.EthernetPortSettings, ethernetport.Settings{}) {
		t.Error("EthernetPortSettings is not initialized")
	}

	if reflect.DeepEqual(m.GeneralSettings, general.Settings{}) {
		t.Error("GeneralSettings is not initialized")
	}

	if reflect.DeepEqual(m.IEEE8021xCredentialContext, ieee8021x.CredentialContext{}) {
		t.Error("IEEE8021xCredentialContext is not initialized")
	}

	if reflect.DeepEqual(m.IEEE8021xProfile, ieee8021x.Profile{}) {
		t.Error("IEEE8021xProfile is not initialized")
	}

	if reflect.DeepEqual(m.KerberosSettingData, kerberos.SettingData{}) {
		t.Error("KerberosSettingData is not initialized")
	}

	if reflect.DeepEqual(m.ManagementPresenceRemoteSAP, managementpresence.RemoteSAP{}) {
		t.Error("ManagementPresenceRemoteSAP is not initialized")
	}

	if reflect.DeepEqual(m.MessageLog, messagelog.Service{}) {
		t.Error("MessageLog is not initialized")
	}

	if reflect.DeepEqual(m.MPSUsernamePassword, mps.UsernamePassword{}) {
		t.Error("MPSUsernamePassword is not initialized")
	}

	if reflect.DeepEqual(m.PublicKeyCertificate, publickey.Certificate{}) {
		t.Error("PublicKeyCertificate is not initialized")
	}

	if reflect.DeepEqual(m.PublicKeyManagementService, publickey.ManagementService{}) {
		t.Error("PublicKeyManagementService is not initialized")
	}

	if reflect.DeepEqual(m.PublicPrivateKeyPair, publicprivate.KeyPair{}) {
		t.Error("PublicPrivateKeyPair is not initialized")
	}

	if reflect.DeepEqual(m.RedirectionService, redirection.Service{}) {
		t.Error("RedirectionService is not initialized")
	}

	if reflect.DeepEqual(m.RemoteAccessPolicyAppliesToMPS, remoteaccess.PolicyAppliesToMPS{}) {
		t.Error("RemoteAccessPolicyAppliesToMPS is not initialized")
	}

	if reflect.DeepEqual(m.RemoteAccessPolicyRule, remoteaccess.PolicyRule{}) {
		t.Error("RemoteAccessPolicyRule is not initialized")
	}

	if reflect.DeepEqual(m.RemoteAccessService, remoteaccess.Service{}) {
		t.Error("RemoteAccessService is not initialized")
	}

	if reflect.DeepEqual(m.SetupAndConfigurationService, setupandconfiguration.Service{}) {
		t.Error("SetupAndConfigurationService is not initialized")
	}

	if reflect.DeepEqual(m.TimeSynchronizationService, timesynchronization.Service{}) {
		t.Error("TimeSynchronizationService is not initialized")
	}

	if reflect.DeepEqual(m.TLSCredentialContext, tls.CredentialContext{}) {
		t.Error("TLSCredentialContext is not initialized")
	}

	if reflect.DeepEqual(m.TLSProtocolEndpointCollection, tls.ProtocolEndpointCollection{}) {
		t.Error("TLSProtocolEndpointCollection is not initialized")
	}

	if reflect.DeepEqual(m.TLSSettingData, tls.SettingData{}) {
		t.Error("TLSSettingData is not initialized")
	}

	if reflect.DeepEqual(m.UserInitiatedConnectionService, userinitiatedconnection.Service{}) {
		t.Error("UserInitiatedConnectionService is not initialized")
	}

	if reflect.DeepEqual(m.WiFiPortConfigurationService, wifiportconfiguration.Service{}) {
		t.Error("WiFiPortConfigurationService is not initialized")
	}
}
