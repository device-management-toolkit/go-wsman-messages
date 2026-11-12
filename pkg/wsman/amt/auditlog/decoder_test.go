/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverwritePolicy_String(t *testing.T) {
	tests := []struct {
		state    OverwritePolicy
		expected string
	}{
		{OverwritePolicyUnknown, "Unknown"},
		{OverwritePolicyWrapsWhenFull, "WrapsWhenFull"},
		{OverwritePolicyNeverOverwrites, "NeverOverwrites"},
		{OverwritePolicyPartialRestrictedRollover, "PartialRestrictedRollover"},
		{OverwritePolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestStoragePolicy_String(t *testing.T) {
	tests := []struct {
		state    StoragePolicy
		expected string
	}{
		{StoragePolicyNoRollOver, "NoRollOver"},
		{StoragePolicyRollOver, "RollOver"},
		{StoragePolicyRestrictedRollOver, "RestrictedRollOver"},
		{StoragePolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestEnabledState_String(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
		{EnabledStateUnknown, "Unknown"},
		{EnabledStateOther, "Other"},
		{EnabledStateEnabled, "Enabled"},
		{EnabledStateDisabled, "Disabled"},
		{EnabledStateShuttingDown, "ShuttingDown"},
		{EnabledStateNotApplicable, "NotApplicable"},
		{EnabledStateEnabledButOffline, "EnabledButOffline"},
		{EnabledStateInTest, "InTest"},
		{EnabledStateDeferred, "Deferred"},
		{EnabledStateQuiesce, "Quiesce"},
		{EnabledStateStarting, "Starting"},
		{EnabledState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestRequestedState_String(t *testing.T) {
	tests := []struct {
		state    RequestedState
		expected string
	}{
		{RequestedStateUnknown, "Unknown"},
		{RequestedStateEnabled, "Enabled"},
		{RequestedStateDisabled, "Disabled"},
		{RequestedStateShutDown, "ShutDown"},
		{RequestedStateNoChange, "NoChange"},
		{RequestedStateOffline, "Offline"},
		{RequestedStateTest, "Test"},
		{RequestedStateDeferred, "Deferred"},
		{RequestedStateQuiesce, "Quiesce"},
		{RequestedStateReboot, "Reboot"},
		{RequestedStateReset, "Reset"},
		{RequestedStateNotApplicable, "NotApplicable"},
		{RequestedState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestGetAuditLogExtendedDataString(t *testing.T) {
	tests := []struct {
		name     string
		appId    int
		eventId  int
		data     string
		expected string
	}{
		{"Unknown Event Group ID", 999, 0, "", "Unknown Event Group ID"},
		{"Security Admin - Provisioning Started", SecurityAdmin, 0, "", "Intel AMT transitioned to setup mode."},
		{"Security Admin - Provisioning Completed Manual", SecurityAdmin, 1, "\u0003", "Intel AMT transitioned to operational mode.\nProvisioning Method: Manual Provisioning via MEBx"},
		{"Security Admin - Provisioning Completed PKI", SecurityAdmin, 1, "\x05\x02\xcb<˷`1\xe5\xe0\x13\x8f\x8dӚ#\xf9\xdeG\xff\xc3^C\xc1\x14L\xea'\xd4jZ\xb1\xcb_\x02\f\x8e\xe0\xc9\rj\x89\x15\x88\x04\x06\x1e\xe2A\xf9\xaf\x03:\xf1\xe6\xa7\x11\xa9\xa0\xbb(d\xb1\x1d\t\xfa\xe5\x00\x12Intel.vprodemo.com", "Intel AMT transitioned to operational mode.\nProvisioning Method: Host-Based Provisioning Admin Mode\nHash Type: SHA 256\nTrusted Root Cert Hash: cb3ccbb76031e5e0138f8dd39a23f9de47ffc35e43c1144cea27d46a5ab1cb5f\nNumber of Certificates: 2\nCert Serial Numbers (first 3): [0c8ee0c90d6a89158804061ee241f9af 033af1e6a711a9a0bb2864b11d09fae5]\nProvisioning Server FQDN: Intel.vprodemo.com"},
		{"Security Admin - ACL Added without data", SecurityAdmin, 2, "", "User entry was added to the Intel AMT Device."},
		{"Security Admin - ACL Added with data", SecurityAdmin, 2, "\u0000\u0004test", "User entry was added to the Intel AMT Device.\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Modified without data", SecurityAdmin, 3, "", "User entry was updated in the Intel AMT device."},
		{"Security Admin - ACL Modified with data", SecurityAdmin, 3, "\u0000\u0000\u0004test", "User entry was updated in the Intel AMT device.\nParameter(s) Modified: None\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Modified with data", SecurityAdmin, 3, "\u0001\u0000\u0004test", "User entry was updated in the Intel AMT device.\nParameter(s) Modified: Username\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Modified with data", SecurityAdmin, 3, "\u0002\u0000\u0004test", "User entry was updated in the Intel AMT device.\nParameter(s) Modified: Password\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Modified with data", SecurityAdmin, 3, "\u0004\u0000\u0004test", "User entry was updated in the Intel AMT device.\nParameter(s) Modified: Local realms\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Modified with data", SecurityAdmin, 3, "\u0008\u0000\u0004test", "User entry was updated in the Intel AMT device.\nParameter(s) Modified: Remote realms\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Modified with data", SecurityAdmin, 3, "\u0010\u0000\u0004test", "User entry was updated in the Intel AMT device.\nParameter(s) Modified: Kerberos domain\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Modified with data", SecurityAdmin, 3, "\u0020\u0000\u0004test", "User entry was updated in the Intel AMT device.\nParameter(s) Modified: SID\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Modified with data", SecurityAdmin, 3, "\u0021\u0000\u0004test", "User entry was updated in the Intel AMT device.\nParameter(s) Modified: Username, SID\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Modified with data", SecurityAdmin, 3, "\u0031\u0000\u0004test", "User entry was updated in the Intel AMT device.\nParameter(s) Modified: Username, Kerberos domain, SID\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - ACL Removed without data", SecurityAdmin, 4, "", "User entry was removed from the Intel AMT device."},
		{"Security Admin - ACL Removed with data", SecurityAdmin, 4, "\x00\x04test", "User entry was removed from the Intel AMT device.\nInitiator Type: Unknown\nUsername: test"},
		{"Security Admin - Invalid credentials Intel AMT", SecurityAdmin, 5, "\x00", "User attempted to access Intel AMT with invalid credentials."},
		{"Security Admin - Invalid credentials MEBx", SecurityAdmin, 5, "\x01", "User attempted to access MEBx with invalid credentials."},
		{"Security Admin - ACL State change without data ", SecurityAdmin, 6, "", "ACL entry state was changed."},
		{"Security Admin - ACL State change with data ", SecurityAdmin, 6, "\x01\x01\x04test", "ACL entry state was changed.\nEntry State: Enabled\nInitiator Type: User\nUsername: test"},
		{"Security Admin - TLS State change without data", SecurityAdmin, 7, "", "TLS state changed."},
		{"Security Admin - TLS State change with data", SecurityAdmin, 7, "\x01\x00", "TLS state changed.\nRemote: Server Auth\nLocal: No Auth"},
		{"Security Admin - TLS Server Certificate Set without data", SecurityAdmin, 8, "", "TLS server certificate was defined."},
		{"Security Admin - TLS Server Certificate Set with data", SecurityAdmin, 8, "M\xf5\xa0`\xe1\xe1>p\xc0S_e\xf23\b%\xa2\x831\x93", "TLS server certificate was defined.\nCertificate serial number: 4df5a060e1e13e70c0535f65f2330825a2833193"},
		{"Security Admin - TLS Server Certificate Removed without data", SecurityAdmin, 9, "", "TLS server certificate was removed."},
		{"Security Admin - TLS Server Certificate Removed with data", SecurityAdmin, 9, "M\xf5\xa0`\xe1\xe1>p\xc0S_e\xf23\b%\xa2\x831\x93", "TLS server certificate was removed.\nCertificate serial number: 4df5a060e1e13e70c0535f65f2330825a2833193"},
		{"Security Admin - TLS Trusted Root Certificate Added without data", SecurityAdmin, 10, "", "TLS trusted root certificate was added."},
		{"Security Admin - TLS Trusted Root Certificate Added with data", SecurityAdmin, 10, "M\xf5\xa0`\xe1\xe1>p\xc0S_e\xf23\b%\xa2\x831\x93", "TLS trusted root certificate was added.\nCertificate serial number: 4df5a060e1e13e70c0535f65f2330825a2833193"},
		{"Security Admin - TLS Trusted Root Certificate Removed without data", SecurityAdmin, 11, "", "TLS trusted root certificate was removed."},
		{"Security Admin - TLS Trusted Root Certificate Removed with data", SecurityAdmin, 11, "M\xf5\xa0`\xe1\xe1>p\xc0S_e\xf23\b%\xa2\x831\x93", "TLS trusted root certificate was removed.\nCertificate serial number: 4df5a060e1e13e70c0535f65f2330825a2833193"},
		{"Security Admin - TLS Pre-Shared Key Set", SecurityAdmin, 12, "", "TLS pre-shared key was defined."},
		{"Security Admin - Kerberos Settings Modified without data", SecurityAdmin, 13, "", "Kerberos settings were modified."},
		{"Security Admin - Kerberos Settings Modified with data", SecurityAdmin, 13, "\x01", "Kerberos settings were modified.\nTime tolerance: 1"},
		{"Security Admin - Kerberos Master Key Modified", SecurityAdmin, 14, "", "Kerberos master key or passphrase was modified."},
		{"Security Admin - Flash Wear Out Counter Reset", SecurityAdmin, 15, "", "Flash wear out counter was reset."},
		{"Security Admin - Power Package Modified without data", SecurityAdmin, 16, "", "Active power package was set."},
		{"Security Admin - Power Package Modified with data", SecurityAdmin, 16, "\x01", "Active power package was set.\nPower policy: 1"},
		{"Security Admin - Set Realm Authentication Mode without data", SecurityAdmin, 17, "", "Realm authentication mode changed."},
		{"Security Admin - Set Realm Authentication Mode with data", SecurityAdmin, 17, "\x01\x00\x00\x00\x01", "Realm authentication mode changed.\nPT Administration, Auth"},
		{"Security Admin - Upgrade Client to Admin", SecurityAdmin, 18, "", "The control mode of the Intel AMT was changed from Client control to Admin control."},
		{"Security Admin - AMT UnProvisioning Started - BIOS", SecurityAdmin, 19, "\x00", "Intel AMT UnProvisioned Started.\nInitiator: BIOS"},
		{"Security Admin - AMT UnProvisioning Started - MEBx", SecurityAdmin, 19, "\x01", "Intel AMT UnProvisioned Started.\nInitiator: MEBx"},
		{"Security Admin - AMT UnProvisioning Started - Local MEI", SecurityAdmin, 19, "\x02", "Intel AMT UnProvisioned Started.\nInitiator: Local MEI"},
		{"Security Admin - AMT UnProvisioning Started - Local WSMAN", SecurityAdmin, 19, "\x03", "Intel AMT UnProvisioned Started.\nInitiator: Local WSMAN"},
		{"Security Admin - AMT UnProvisioning Started - Remote WSMAN", SecurityAdmin, 19, "\u0004", "Intel AMT UnProvisioned Started.\nInitiator: Remote WSMAN"},
		{"Security Admin - Unknown Event", SecurityAdmin, 20, "", "Unknown Event ID"},
		// {"Remote Control - Performed Power-Up", RemoteControl, 0, "\x00\x00\x00\x00\x00\x00\x00", "Remote power up initiated.\nBoot Media: None\n Boot Media Override: Disabled\n BIOS Pause: Disabled\n BIOS Pause Key: None"},
		{"Remote Control - Performed Power-Down", RemoteControl, 1, "", "Remote power down initiated."},
		// {"Remote Control - Performed Power-Cycle", RemoteControl, 2, "\x00\x00\x00\x00\x00\x00\x00", "Remote power cycle initiated.\nBoot Media: None\n Boot Media Override: Disabled\n BIOS Pause: Disabled\n BIOS Pause Key: None"},
		// {"Remote Control - Performed Reset", RemoteControl, 3, "\x00\x00\x00\x00\x00\x00\x00", "Remote reset initiated.\nBoot Media: None\n Boot Media Override: Disabled\n BIOS Pause: Disabled\n BIOS Pause Key: None"},
		// {"Remote Control - Set Boot Options", RemoteControl, 4, "\x05\x00\x00\x00\x00\x00\x00", "Boot options were set.\nBoot Media: PXE\n Boot Media Override: Disabled\n BIOS Pause: Disabled\n BIOS Pause Key: None"},
		{"Remote Control - Performed Graceful Power Down", RemoteControl, 5, "", "Remote graceful power down initiated."},
		// {"Remote Control - Performed Graceful Power Reset", RemoteControl, 6, "\x00\x00\x00\x00\x00\x00\x00", "Remote reset initiated.\nBoot Media: None\n Boot Media Override: Disabled\n BIOS Pause: Disabled\n BIOS Pause Key: None"},
		{"Remote Control - Performed Standby", RemoteControl, 7, "", "Remote standby initiated."},
		{"Remote Control - Performed Hibernate", RemoteControl, 8, "", "Remote hibernate initiated."},
		{"Remote Control - Performed NMI", RemoteControl, 9, "", "Remote NMI initiated."},
		{"Remote Control - Unknown Event", RemoteControl, 10, "", "Unknown Event ID"},
		{"Redirection Manager - IDER Session Opened", RedirectionManager, 0, "", "An application opened a Storage Redirection session."},
		{"Redirection Manager - IDER Session Closed", RedirectionManager, 1, "", "An application or firmware closed a Storage Redirection session."},
		{"Redirection Manager - IDER Enabled", RedirectionManager, 2, "", "Storage Redirection was enabled."},
		{"Redirection Manager - IDER Disabled", RedirectionManager, 3, "", "Storage Redirection was disabled."},
		{"Redirection Manager - SoL Session Opened", RedirectionManager, 4, "", "An application opened a Serial Over LAN session."},
		{"Redirection Manager - SoL Session Closed", RedirectionManager, 5, "", "An application or firmware closed a Serial Over LAN session."},
		{"Redirection Manager - SoL Enabled", RedirectionManager, 6, "", "Serial Over LAN was enabled."},
		{"Redirection Manager - SoL Disabled", RedirectionManager, 7, "", "Serial Over LAN was disabled."},
		{"Redirection Manager - KVM Session Started", RedirectionManager, 8, "", "An application opened a Keyboard-Video-Mouse session."},
		{"Redirection Manager - KVM Session Ended", RedirectionManager, 9, "", "An application or firmware closed a Keyboard-Video-Mouse session."},
		{"Redirection Manager - KVM Enabled", RedirectionManager, 10, "", "Keyboard-Video-Mouse was enabled."},
		{"Redirection Manager - KVM Disabled", RedirectionManager, 11, "", "Keyboard-Video-Mouse was disabled."},
		{"Redirection Manager - VNC Password Failed", RedirectionManager, 12, "", "Incorrect Remote Frame Buffer (RFB) password entered 3 times."},
		{"Firmware Update Manager - Updated", FirmwareUpdateManager, 0, "\x10\x00\x01\x00\x19\x00\x7F\x09\x10\x00\x01\x00\x1A\x00\xE8\x03", "Firmware update was started.\nOld version: 16.1.25.2431\nNew version: 16.1.26.1000"},
		{"Firmware Update Manager - Update Failed", FirmwareUpdateManager, 1, "\x01\x02", "Firmware update failed.\nFailure Type: 1\nFailure Reason: 2"},
		{"Security Audit Log - Log cleared", SecurityAuditLog, 0, "", "Audit log was cleared."},
		{"Security Audit Log - Policy event was enabled or disabled", SecurityAuditLog, 1, "", "Audit policy event was enabled or disabled."},
		{"Security Audit Log - Disable monitor feature", SecurityAuditLog, 2, "", "Access monitor feature was disabled."},
		{"Security Audit Log - Enable monitor feature", SecurityAuditLog, 3, "", "Access monitor feature was enabled."},
		{"Security Audit Log - Log exported", SecurityAuditLog, 4, "", "Audit log signature and log-related information was exported."},
		{"Security Audit Log - Log recovery - Unknown", SecurityAuditLog, 5, "\x00", "Internal check of audit log resulted in a recovery action.\nReason: Unknown"},
		{"Security Audit Log - Log recovery - Migration failure", SecurityAuditLog, 5, "\x01", "Internal check of audit log resulted in a recovery action.\nReason: Migration failure"},
		{"Security Audit Log - Log recovery - Initialization failure", SecurityAuditLog, 5, "\x02", "Internal check of audit log resulted in a recovery action.\nReason: Initialization failure"},
		{"Network Time", NetworkTime, 0, "", "Command received to set Intel AMT local time."},
		{"Network Administration - TCP/IP Parameters Set", NetworkAdministration, 0, "\x00\x00\x00\x00\x01\x02\x00\xA8\xC0\x00\xFF\xFF\xFF\x01\x00\xA8\xC0\x04\x04\x04\x02\x02\x02\x02\x01", "TCP/IP parameters were set.\nDHCP Enabled: Enabled\nStatic IP: 192.168.0.2\nSubnet Mask: 255.255.255.0\nGateway: 192.168.0.1"},
		{"Network Administration - Host Name Set", NetworkAdministration, 1, "\x04test", "Host name was set to test"},
		{"Network Administration - Domain Name Set", NetworkAdministration, 2, "\x04test", "Domain name was set to test"},
		{"Network Administration - VLAN Parameters Set", NetworkAdministration, 3, "\u0000\u0000\u0000\u0000\u0000\u0000", "VLAN tag was set to disabled"},
		{"Network Administration - VLAN Parameters Set", NetworkAdministration, 3, "\u0000\u0000\u0000\u0000\u0000\u000F", "VLAN tag was set to 3840"},
		{"Network Administration - Link Policy Set", NetworkAdministration, 4, "\u0000\u0000\u0000\u0000\u0001\u0000\u0000\u0000", "Link policy was set to 1"},
		{"Network Administration - Link Policy Set", NetworkAdministration, 4, "\u0000\u0000\u0000\u0000\u0001\u0000\u0000\u0000", "Link policy was set to 1"},
		{"Network Administration - IPv6 Parameters Set", NetworkAdministration, 5, "\x00\x00\x00\x00\x01\x01\x20\x01\x0d\xb8\x85\xa3\x00\x00\x00\x00\x8a\x2e\x03\x70\x73\x34\x20\x01\x0d\xb8\x85\xa3\x00\x00\x00\x00\x8a\x2e\x03\x70\x73\x34\x20\x01\x0d\xb8\x85\xa3\x00\x00\x00\x00\x8a\x2e\x03\x70\x73\x34\x20\x01\x0d\xb8\x85\xa3\x00\x00\x00\x00\x8a\x2e\x03\x70\x73\x34", "IPv6 parameters were set.\nIPv6: Enabled\nInterface Gen Type: Intel ID\nIPv6 Address: 2001:db8:85a3::8a2e:370:7334\nIPv6 Gateway: 2001:db8:85a3::8a2e:370:7334\nIPv6 Primary DNS: 2001:db8:85a3::8a2e:370:7334\nIPv6 Secondary DNS: 2001:db8:85a3::8a2e:370:7334"},
		{"Network Administration - IPv6 Parameters Set", NetworkAdministration, 5, "\x00\x00\x00\x00\x01\x02\x02\x1a\x2b\xff\xfe\x3c\x4d\x5e\x20\x01\x0d\xb8\x85\xa3\x00\x00\x00\x00\x8a\x2e\x03\x70\x73\x34\x20\x01\x0d\xb8\x85\xa3\x00\x00\x00\x00\x8a\x2e\x03\x70\x73\x34\x20\x01\x0d\xb8\x85\xa3\x00\x00\x00\x00\x8a\x2e\x03\x70\x73\x34\x20\x01\x0d\xb8\x85\xa3\x00\x00\x00\x00\x8a\x2e\x03\x70\x73\x34", "IPv6 parameters were set.\nIPv6: Enabled\nInterface Gen Type: Manual ID\nInterface ID: 021a:2bff:fe3c:4d5e\nIPv6 Address: 2001:db8:85a3::8a2e:370:7334\nIPv6 Gateway: 2001:db8:85a3::8a2e:370:7334\nIPv6 Primary DNS: 2001:db8:85a3::8a2e:370:7334\nIPv6 Secondary DNS: 2001:db8:85a3::8a2e:370:7334"},
		{"Storage Administration - Global Storage Attributes Set", StorageAdministration, 0, "\x00\x00\x01\x00\x00\x00\x01\x00", "Global storage attributes were set.\nMax Partner Storage: 65536\nMax Non-Partner Total Allocation Size: 65536"},
		{"Storage Administration - Storage EACL Modified", StorageAdministration, 1, "", "Storage EACL entry was added or removed."},
		{"Storage Administration - Storage FPACL Modified", StorageAdministration, 2, "", "Storage FPACL entry was added, removed, or updated."},
		{"Storage Administration - Storage Write Operation", StorageAdministration, 3, "", "Application wrote a block to storage."},
		{"Event Manager - Alert Subscribed", EventManager, 0, "\x01\x01\x00\xc0\xa8\x00\x01", "An alert subscription was created successfully.\nPolicy ID: 1\nSubscription Alert Type: SNMP\nIP Address Type: IPv4\nAlert Target IP Address: 192.168.0.1"},
		{"Event Manager - Alert Unsubscribed", EventManager, 1, "\x01\x01\x00\xc0\xa8\x00\x01", "An existing alert subscription was cancelled.\nPolicy ID: 1\nSubscription Alert Type: SNMP\nIP Address Type: IPv4\nAlert Target IP Address: 192.168.0.1"},
		{"Event Manager - Alert Unsubscribed", EventManager, 1, "\x01\x01\x01\x20\x01\x0d\xb8\x85\xa3\x00\x00\x00\x00\x8a\x2e\x03\x70\x73\x34", "An existing alert subscription was cancelled.\nPolicy ID: 1\nSubscription Alert Type: SNMP\nIP Address Type: IPv6\nAlert Target IP Address: 2001:db8:85a3::8a2e:370:7334"},
		{"Event Manager - Event Log Cleared", EventManager, 2, "", "Event log was cleared of existing records."},
		{"Event Manager - Event Log Frozen", EventManager, 3, "\x01", "Event log was frozen"},
		{"Event Manager - Event Log Frozen", EventManager, 3, "\x00", "Event log was unfrozen"},
		{"System Defense Manager - SD Filter Added", SystemDefenseManager, 0, "", "Attempt made to add a system defense filter."},
		{"System Defense Manager - SD Filter Removed", SystemDefenseManager, 1, "\x01\x00\x00\x00", "System defense filter was removed successfully.\nFilter Handle: 1"},
		{"System Defense Manager - SD Policy Added", SystemDefenseManager, 2, "", "Attempt made to add a system defense policy."},
		{"System Defense Manager - SD Policy Removed", SystemDefenseManager, 3, "\x01\x00\x00\x00", "System defense policy was removed successfully.\nPolicy Handle: 1"},
		{"System Defense Manager - SD Default Policy Set", SystemDefenseManager, 4, "\x01\x00\x00\x00\x01\x00\x00\x00", "System defense policy selected.\nHardware Interface: 1\nPolicy Handle: 1"},
		{"System Defense Manager - SD Heuristics Option Set", SystemDefenseManager, 5, "\x01\x00\x00\x00\x01\xd9\x00\x01\x00\x00", "System defense heuristics settings were set successfully.\nInterface Handle: 1\nBlock All: 1\nBlock Offensive Port: 217\nPolicy Handle: 256"},
		{"System Defense Manager - SD Heuristics State Cleared", SystemDefenseManager, 6, "\x01\x00\x00\x00", "System defense heuristics settings were removed successfully.\nInterface Handle: 1"},
		{"Agent Presence Manager - Agent Watchdog Added", AgentPresenceManager, 0, "\x3b\xd4\x6c\x70\x98\x85\x45\xb6\x9f\xe2\x62\x2f\x14\x97\x38\x7e\xa0\x00\x0a\x00", "An application entry was created to be monitored.\nAgent ID: 3bd46c70-9885-45b6-9fe2-622f1497387e\nAgent Heartbeat Time: 160\nAgent Startup Time: 10"},
		{"Agent Presence Manager - Agent Watchdog Removed", AgentPresenceManager, 1, "\x3b\xd4\x6c\x70\x98\x85\x45\xb6\x9f\xe2\x62\x2f\x14\x97\x38\x7e", "An application entry was removed.\nAgent ID: 3bd46c70-9885-45b6-9fe2-622f1497387e"},
		{"Agent Presence Manager - Agent Watchdog Action Set", AgentPresenceManager, 2, "\x3b\xd4\x6c\x70\x98\x85\x45\xb6\x9f\xe2\x62\x2f\x14\x97\x38\x7e", "Actions were set, added, or removed for an application watchdog entry.\nAgent ID: 3bd46c70-9885-45b6-9fe2-622f1497387e"},
		{"Wireless Configuration - Wireless Profile Added", WirelessConfiguration, 0, "test\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x04test", "A new profile was added.\nSSID: test\nProfile Priority: 1\nProfile Name Length: 4\nProfile Name: test"},
		{"Wireless Configuration - Wireless Profile Removed", WirelessConfiguration, 1, "\x04test", "An existing profile was removed.\nProfile Name Length: 4\nProfile Name: test"},
		{"Wireless Configuration - Wireless Profile Updated", WirelessConfiguration, 2, "test\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x04test", "An existing profile was updated.\nSSID: test\nProfile Priority: 1\nProfile Name Length: 4\nProfile Name: test"},
		{"Wireless Configuration - Wireless Profile Modified", WirelessConfiguration, 3, "\u0000\u0000\u0000\u0000", "An existing profile sync was modified.\nProfile sync is disabled"},
		{"Wireless Configuration - Wireless Profile Modified", WirelessConfiguration, 3, "\u0001\u0000\u0000\u0000", "An existing profile sync was modified.\nProfile sync user"},
		{"Wireless Configuration - Wireless Profile Modified", WirelessConfiguration, 3, "\u0002\u0000\u0000\u0000", "An existing profile sync was modified.\nProfile sync admin"},
		{"Wireless Configuration - Wireless Profile Modified", WirelessConfiguration, 3, "\u0003\u0000\u0000\u0000", "An existing profile sync was modified.\nProfile sync is unrestricted"},
		{"Wireless Configuration - Wireless Link Preference Changed", WirelessConfiguration, 4, "\u0003\u0000\u0000\u0000\u0001\u0000\u0000\u0000", "An existing profile link preference was changed.\nTimeout: 3\nLink Preference: ME"},
		{"Wireless Configuration - Wireless Link Preference Changed", WirelessConfiguration, 4, "\u0003\u0000\u0000\u0000\u0002\u0000\u0000\u0000", "An existing profile link preference was changed.\nTimeout: 3\nLink Preference: Host"},
		{"Wireless Configuration - Wireless Profile Share with UEFI Enabled Setting Changed", WirelessConfiguration, 5, "\u0001", "Wireless profile share with UEFI was set to Enabled."},
		{"Wireless Configuration - Wireless Profile Share with UEFI Enabled Setting Changed", WirelessConfiguration, 5, "\u0000", "Wireless profile share with UEFI was set to Disabled."},
		{"Endpoint Access Control - EAC Posture Signer Set", EndpointAccessControl, 0, "", "A certificate handle for signing EAC postures was either set or removed."},
		{"Endpoint Access Control - EAC Enabled", EndpointAccessControl, 1, "", "EAC was set to enabled by WS-MAN interface."},
		{"Endpoint Access Control - EAC Disabled", EndpointAccessControl, 2, "", "EAC was set to disabled by WS-MAN interface."},
		{"Endpoint Access Control - EAC Posture State Updated", EndpointAccessControl, 3, "", "Controllable fields of EAC posture were reset manually by WS-MAN interface."},
		{"Endpoint Access Control - EAC Set Options", EndpointAccessControl, 4, "\x96\x1f\u0000\u0000", "EAC options were changed.\nEAC Vendors: 8086"},
		{"Keyboard Video Mouse - KVM Opt-In Enabled", KeyboardVideoMouse, 0, "", "User consent for a KVM session is now required."},
		{"Keyboard Video Mouse - KVM Opt-In Disabled", KeyboardVideoMouse, 1, "", "User consent for a KVM session is no longer required."},
		{"Keyboard Video Mouse - KVM Password Changed", KeyboardVideoMouse, 2, "", "RFB password for KVM session has changed."},
		{"Keyboard Video Mouse - KVM Consent Succeeded", KeyboardVideoMouse, 3, "", "Remote operator entered a one-time password successfully"},
		{"Keyboard Video Mouse - KVM Consent Failed", KeyboardVideoMouse, 4, "", "Remote operator failed (3 times) to enter a one-time password correctly"},
		{"User Opt-In - Policy Change", UserOptIn, 0, "\x01\x00", "A user has modified the opt-in policy settings.\nPrevious Opt-In Policy: KVM\nCurrent Opt-In Policy: None"},
		{"User Opt-In - Policy Change", UserOptIn, 0, "\x00\x01", "A user has modified the opt-in policy settings.\nPrevious Opt-In Policy: None\nCurrent Opt-In Policy: KVM"},
		{"User Opt-In - Policy Change", UserOptIn, 0, "\x00\xff", "A user has modified the opt-in policy settings.\nPrevious Opt-In Policy: None\nCurrent Opt-In Policy: All"},
		{"User Opt-In - Send Consent Code Event", UserOptIn, 1, "\x00", "The remote operator sent a consent code.\nRemote operator entered a one-time password successfully"},
		{"User Opt-In - Send Consent Code Event", UserOptIn, 1, "\x01", "The remote operator sent a consent code.\nRemote operator failed (3 times) to enter a one-time password correctly"},
		{"User Opt-In - Start Opt-In Blocked Event", UserOptIn, 2, "", "The remote operator attempted to send a start opt-in request, but the request was blocked (denial-of-service attack prevention)."},
		{"Watchdog - Reset-Triggering Options Changed", Watchdog, 0, "", "A user has modified the watchdog action settings."},
		{"Watchdog - Action Pairing Changed", Watchdog, 1, "\x00", "A user has modified a watchdog to add, remove, or alter the watchdog action connected to it.\nRemote operator entered a one-time password successfully"},
		{"Watchdog - Action Pairing Changed", Watchdog, 1, "\x01", "A user has modified a watchdog to add, remove, or alter the watchdog action connected to it.\nRemote operator failed (3 times) to enter a one-time password correctly"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetAuditLogExtendedDataString(tt.appId, tt.eventId, tt.data)
			if result != tt.expected {
				t.Errorf("GetAuditLogExtendedDataString(%d, %d, %q) = %v; want %v", tt.appId, tt.eventId, tt.data, result, tt.expected)
			}
		})
	}
}

func TestConvertToAuditLogResultSkipsTruncatedRecords(t *testing.T) {
	valid := "ABMAAAI/9M1uAgAQAAwAAAAoBZkADAAAACQFlA=="
	raw, _ := base64.StdEncoding.DecodeString(valid)

	// a Kerberos initiator whose SID length points past the end of the record
	kerberos := base64.StdEncoding.EncodeToString([]byte{0x00, 0x10, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x01})

	records := []string{kerberos}
	for i := 1; i < len(raw); i++ {
		records = append(records, base64.StdEncoding.EncodeToString(raw[:i]))
	}

	records = append(records, valid)

	assert.NotPanics(t, func() {
		result := convertToAuditLogResult(records)
		assert.Len(t, result, 1)
	})
}
