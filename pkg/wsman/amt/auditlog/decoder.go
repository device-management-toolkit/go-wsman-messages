/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/common"
)

const (
	AMTAuditLog   string = "AMT_AuditLog"
	ReadRecords   string = "ReadRecords"
	ValueNotFound string = "Value not found in map"
)

const (
	OverwritePolicyUnknown                   OverwritePolicy = 0
	OverwritePolicyWrapsWhenFull             OverwritePolicy = 2
	OverwritePolicyNeverOverwrites           OverwritePolicy = 7
	OverwritePolicyPartialRestrictedRollover OverwritePolicy = 32768
)

const UnknownEventID = "Unknown Event ID"

var OverwritePolicyToString = map[OverwritePolicy]string{
	OverwritePolicyUnknown:                   "Unknown",
	OverwritePolicyWrapsWhenFull:             "WrapsWhenFull",
	OverwritePolicyNeverOverwrites:           "NeverOverwrites",
	OverwritePolicyPartialRestrictedRollover: "PartialRestrictedRollover",
}

// OverwritePolicyToString returns a string representation of a OverwritePolicy.
func (r OverwritePolicy) String() string {
	if value, exists := OverwritePolicyToString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

var EnabledStateToString = map[EnabledState]string{
	EnabledStateUnknown:           "Unknown",
	EnabledStateOther:             "Other",
	EnabledStateEnabled:           "Enabled",
	EnabledStateDisabled:          "Disabled",
	EnabledStateShuttingDown:      "ShuttingDown",
	EnabledStateNotApplicable:     "NotApplicable",
	EnabledStateEnabledButOffline: "EnabledButOffline",
	EnabledStateInTest:            "InTest",
	EnabledStateDeferred:          "Deferred",
	EnabledStateQuiesce:           "Quiesce",
	EnabledStateStarting:          "Starting",
}

// EnabledStateToString returns a string representation of a EnabledState.
func (r EnabledState) String() string {
	if value, exists := EnabledStateToString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)

var RequestedStateToString = map[RequestedState]string{
	RequestedStateUnknown:       "Unknown",
	RequestedStateEnabled:       "Enabled",
	RequestedStateDisabled:      "Disabled",
	RequestedStateShutDown:      "ShutDown",
	RequestedStateNoChange:      "NoChange",
	RequestedStateOffline:       "Offline",
	RequestedStateTest:          "Test",
	RequestedStateDeferred:      "Deferred",
	RequestedStateQuiesce:       "Quiesce",
	RequestedStateReboot:        "Reboot",
	RequestedStateReset:         "Reset",
	RequestedStateNotApplicable: "NotApplicable",
}

// RequestedStateToString returns a string representation of a RequestedState.
func (r RequestedState) String() string {
	if value, exists := RequestedStateToString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	StoragePolicyNoRollOver StoragePolicy = iota
	StoragePolicyRollOver
	StoragePolicyRestrictedRollOver
)

var StoragePolicyToString = map[StoragePolicy]string{
	StoragePolicyNoRollOver:         "NoRollOver",
	StoragePolicyRollOver:           "RollOver",
	StoragePolicyRestrictedRollOver: "RestrictedRollOver",
}

// StoragePolicyToString returns a string representation of a StoragePolicy.
func (r StoragePolicy) String() string {
	if value, exists := StoragePolicyToString[r]; exists {
		return value
	}

	return ValueNotFound
}

var provisioningMethodToString = map[int]string{
	2: "Remote Configuration",
	3: "Manual Provisioning via MEBx",
	5: "Host-Based Provisioning Admin Mode",
}

var initiatorTypeToString = map[int]string{
	0: "Unknown",
	1: "User",
	2: "Machine",
}

var optInPolicyToString = map[int]string{
	0:   "None",
	1:   "KVM",
	255: "All",
}

var operationStatusToString = map[int]string{
	0: "Remote operator entered a one-time password successfully",
	1: "Remote operator failed (3 times) to enter a one-time password correctly",
}

var provisioningHashTypeToString = map[int]string{
	1: "SHA1 160",
	2: "SHA 256",
	3: "SHA 384",
}

var ExtendedDataMap = map[int]string{
	0: "Invalid ME access",
	1: "Invalid MEBx access",
}

const (
	SecurityAdmin         = 16
	RemoteControl         = 17
	RedirectionManager    = 18
	FirmwareUpdateManager = 19
	SecurityAuditLog      = 20
	NetworkTime           = 21
	NetworkAdministration = 22
	StorageAdministration = 23
	EventManager          = 24
	SystemDefenseManager  = 25
	AgentPresenceManager  = 26
	WirelessConfiguration = 27
	EndpointAccessControl = 28
	KeyboardVideoMouse    = 29
	UserOptIn             = 30
	ScreenBlanking        = 32
	Watchdog              = 33
)

var AMTAppIDToString = map[int]string{
	16: "Security Admin Events",
	17: "Remote Control Events",
	18: "Redirection Manager Events",
	19: "Firmware Update Manager Events",
	20: "Security AuditLog Events",
	21: "Network Time Events",
	22: "Network Administration Events",
	23: "Storage Administration Events",
	24: "Event Manager Events",
	25: "System Defense Manager Events",
	26: "Agent Presence Manager Events",
	27: "Wireless Configuration Events",
	28: "Endpoint Access Control Events",
	29: "Keyboard Video Mouse Events",
	30: "User Opt-In Events",
	32: "Screen Blanking Events",
	33: "Watchdog Events",
}

var AMTAuditLogEventToString = map[int]string{
	1600: "AMT Provisioning Started",
	1601: "AMT Provisioning Completed",
	1602: "ACL Entry Added",
	1603: "ACL Entry Modified",
	1604: "ACL Entry Removed",
	1605: "ACL Access with Invalid Credentials",
	1606: "ACL Entry State Changed",
	1607: "TLS State Changed",
	1608: "TLS Server Certificate Set",
	1609: "TLS Server Certificate Removed",
	1610: "TLS Trusted Root Certificate Added",
	1611: "TLS Trusted Root Certificate Removed",
	1612: "TLS Pre-Shared Key Set",
	1613: "Kerberos Settings Modified",
	1614: "Kerberos Master Key or Passphrase Modified",
	1615: "Flash Wear out Counters Reset",
	1616: "Power Package Modified",
	1617: "Set Realm Authentication Mode",
	1618: "Upgrade Client to Admin Control Mode",
	1619: "AMT UnProvisioning Started",
	1700: "Performed Power Up",
	1701: "Performed Power Down",
	1702: "Performed Power Cycle",
	1703: "Performed Reset",
	1704: "Set Boot Options",
	1705: "Performed Graceful Power Down",
	1706: "Performed Graceful Power Reset",
	1707: "Preformed Standby",
	1708: "Performed Hibernate",
	1709: "Performed NMI",
	1800: "IDE-R Session Opened",
	1801: "IDE-R Session Closed",
	1802: "IDE-R Enabled",
	1803: "IDE-R Disabled",
	1804: "SoL Session Opened",
	1805: "SoL Session Closed",
	1806: "SoL Enabled",
	1807: "SoL Disabled",
	1808: "KVM Session Started",
	1809: "KVM Session Ended",
	1810: "KVM Enabled",
	1811: "KVM Disabled",
	1812: "VNC Password Failed 3 Times",
	1900: "Firmware Update Started",
	1901: "Firmware Update Failed",
	2000: "Security Audit Log Cleared",
	2001: "Security Audit Policy Modified",
	2002: "Security Audit Log Disabled",
	2003: "Security Audit Log Enabled",
	2004: "Security Audit Log Exported",
	2005: "Security Audit Log Recovered",
	2100: "AMT Time Set",
	2200: "TCP/IP Parameters Set",
	2201: "Host Name Set",
	2202: "Domain Name Set",
	2203: "VLAN Parameters Set",
	2204: "Link Policy Set",
	2205: "IPv6 Parameters Set",
	2300: "Global Storage Attributes Set",
	2301: "Storage EACL Modified",
	2302: "Storage FPACL Modified",
	2303: "Storage Write Operation",
	2400: "Alert Subscribed",
	2401: "Alert Unsubscribed",
	2402: "Event Log Cleared",
	2403: "Event Log Frozen",
	2500: "System Defense Filter Added",
	2501: "System Defense Filter Removed",
	2502: "System Defense Policy Added",
	2503: "System Defense Policy Removed",
	2504: "System Defense Default Policy Set",
	2505: "System Defense Heuristics Option Set",
	2506: "System Defense Heuristics State Cleared",
	2600: "Agent Watchdog Added",
	2601: "Agent Watchdog Removed",
	2602: "Agent Watchdog Action Set",
	2700: "Wireless Profile Added",
	2701: "Wireless Profile Removed",
	2702: "Wireless Profile Updated",
	2703: "Wireless Profile Modified",
	2704: "Wireless Link Preference Changed",
	2705: "Wireless Profile Share With UEFI Enabled Setting Changed",
	2800: "EAC Posture Signer Set",
	2801: "EAC Enabled",
	2802: "EAC Disabled",
	2803: "EAC Posture State Updated",
	2804: "EAC Set Options",
	2900: "KVM Opt-In Enabled",
	2901: "KVM Opt-In Disabled",
	2902: "KVM Password Changed",
	2903: "KVM Consent Succeeded",
	2904: "KVM Consent Failed",
	3000: "Opt-In Policy Change",
	3001: "Send Consent Code Event",
	3002: "Start Opt-In Blocked Event",
	3301: "Watchdog Reset Triggering Options Changed",
	3302: "Watchdog Action Pairing Changed",
}

var RealmNames = []string{
	"Redirection",
	"PT Administration",
	"Hardware Asset",
	"Remote Control",
	"Storage",
	"Event Manager",
	"Storage Admin",
	"Agent Presence Local",
	"Agent Presence Remote",
	"Circuit Breaker",
	"Network Time",
	"General Information",
	"Firmware Update",
	"EIT",
	"LocalUN",
	"Endpoint Access Control",
	"Endpoint Access Control Admin",
	"Event Log Reader",
	"Audit Log",
	"ACL Realm",
	"",
	"",
	"Local System",
	// Add more as needed
}

func convertToAuditLogResult(auditlogdata []string) []AuditLogRecord {
	records := []AuditLogRecord{}

	for _, eventRecord := range auditlogdata {
		ptr := 0

		decodedEventRecord, err := base64.StdEncoding.DecodeString(eventRecord)
		if err != nil {
			continue
		}

		decodedEventRecordStr := string(decodedEventRecord)
		auditLogRecord := AuditLogRecord{}

		auditLogRecord.AuditAppID = common.ReadShort(decodedEventRecordStr, 0)
		auditLogRecord.EventID = common.ReadShort(decodedEventRecordStr, 2)
		auditLogRecord.AuditApp = AMTAppIDToString[auditLogRecord.AuditAppID]
		auditLogRecord.Event = AMTAuditLogEventToString[(auditLogRecord.AuditAppID*100)+auditLogRecord.EventID]

		initiatorType, initiator, pointer := getInitiatorInfo(decodedEventRecordStr)
		auditLogRecord.InitiatorType = initiatorType
		auditLogRecord.Initiator = initiator
		ptr = pointer

		// Read timestamp
		timeStamp := common.ReadInt(decodedEventRecordStr, ptr)
		auditLogRecord.Time = time.Unix(int64(timeStamp), 0)
		ptr += 4

		// Read network access

		auditLogRecord.MCLocationType = []byte(decodedEventRecordStr[ptr : ptr+1])[0]
		ptr++

		netlen := []byte(decodedEventRecordStr[ptr : ptr+1])[0]
		ptr++

		auditLogRecord.NetAddress = strings.ReplaceAll(decodedEventRecordStr[ptr:ptr+int(netlen)], "0000:0000:0000:0000:0000:0000:0000:0001", "::1")

		// Read extended data
		ptr += int(netlen)

		exlen := []byte(decodedEventRecordStr[ptr : ptr+1])[0]
		ptr++

		auditLogRecord.Ex = decodedEventRecordStr[ptr : ptr+int(exlen)]
		auditLogRecord.ExStr = GetAuditLogExtendedDataString(auditLogRecord.AuditAppID, auditLogRecord.EventID, auditLogRecord.Ex)

		records = append([]AuditLogRecord{auditLogRecord}, records...)
	}

	return records
}

// Return human readable extended audit log data
// TODO: Just put some of them here, but many more still need to be added, helpful link here:
// https://software.intel.com/sites/manageability/AMT_Implementation_and_Reference_Guide/default.htm?turl=WordDocuments%2Fsecurityadminevents.htm
func GetAuditLogExtendedDataString(appId, eventId int, data string) string {
	var extendedDataString string

	switch appId {
	case SecurityAdmin:
		extendedDataString = parseSecurityAdminEvents(eventId, data)
	case RemoteControl:
		extendedDataString = parseRemoteControlEvents(eventId, data)
	case RedirectionManager:
		extendedDataString = parseRedirectionManagerEvents(eventId)
	case FirmwareUpdateManager:
		extendedDataString = parseFirmwareUpdateManagerEvents(eventId, data)
	case SecurityAuditLog:
		extendedDataString = parseSecurityAuditLog(eventId, data)
	case NetworkTime:
		extendedDataString = parseNetworkTimeEvents(eventId, data)
	case NetworkAdministration:
		extendedDataString = parseNetworkAdministrationEvents(eventId, data)
	case StorageAdministration:
		extendedDataString = parseStorageAdministrationEvents(eventId, data)
	case EventManager:
		extendedDataString = parseEventManagerEvents(eventId, data)
	case SystemDefenseManager:
		extendedDataString = parseSystemDefenseManagerEvents(eventId, data)
	case AgentPresenceManager:
		extendedDataString = parseAgentPresenceManagerEvents(eventId, data)
	case WirelessConfiguration:
		extendedDataString = parseWirelessConfigurationEvents(eventId, data)
	case EndpointAccessControl:
		extendedDataString = parseEndpointAccessControlEvents(eventId, data)
	case KeyboardVideoMouse:
		extendedDataString = parseKeyboardVideoMouseEvents(eventId)
	case UserOptIn:
		extendedDataString = parseUserOptInEvents(eventId, data)
	case Watchdog:
		extendedDataString = parseWatchdogEvents(eventId, data)
	default:
		extendedDataString = "Unknown Event Group ID"
	}

	return extendedDataString
}

func parseSecurityAdminEvents(eventId int, data string) string {
	const (
		ProvisioningStarted              = 0
		ProvisioningCompleted            = 1
		ACLEntryAdded                    = 2
		ACLEntryModified                 = 3
		ACLEntryRemoved                  = 4
		ACLAccessWithInvalidCredentials  = 5
		ACLEntryEnabled                  = 6
		TLSStateChanged                  = 7
		TLSServerCertificateSet          = 8
		TLSServerCertificateRemoved      = 9
		TLSTrustedRootCertificateAdded   = 10
		TLSTrustedRootCertificateRemoved = 11
		TLSPreSharedKeySet               = 12
		KerberosSettingsModified         = 13
		KerberosMasterKeyModified        = 14
		FlashWearOutCountersReset        = 15
		PowerPackageModified             = 16
		SetRealmAuthenticationMode       = 17
		UpgradeClientToAdmin             = 18
		AMTUnProvisioningStarted         = 19
	)

	var extendedDataString string

	byteData := []byte(data)

	switch eventId {
	case ProvisioningStarted:
		extendedDataString = "Intel AMT transitioned to setup mode."
	case ProvisioningCompleted:
		extendedDataString = "Intel AMT transitioned to operational mode."

		if len(byteData) > 0 {
			event := readProvisioningCompletedEventData(byteData)

			extendedDataString += provisioningCompletedToString(&event)
		}
	case ACLEntryAdded:
		extendedDataString = "User entry was added to the Intel AMT Device."

		if len(byteData) > 0 {
			entry := readACLData(ACLEntryAdded, byteData)

			extendedDataString += aclEntryToString(&entry)
		}
	case ACLEntryModified:
		extendedDataString = "User entry was updated in the Intel AMT device."

		if len(byteData) > 0 {
			entry := readACLData(ACLEntryModified, byteData)

			extendedDataString += aclEntryModifiedToString(&entry)
		}
	case ACLEntryRemoved:
		extendedDataString = "User entry was removed from the Intel AMT device."

		if len(byteData) > 0 {
			entry := readACLData(ACLEntryRemoved, byteData)

			extendedDataString += aclEntryToString(&entry)
		}
	case ACLAccessWithInvalidCredentials:
		if len(data) > 0 {
			extendedDataString = "User attempted to access " + []string{"Intel AMT", "MEBx"}[data[0]] + " with invalid credentials."
		}
	case ACLEntryEnabled:
		extendedDataString = "ACL entry state was changed."

		if len(byteData) > 0 {
			entry := readACLData(ACLEntryEnabled, byteData)

			extendedDataString += aclEntryEnabledToString(&entry)
		}
	case TLSStateChanged:
		extendedDataString = "TLS state changed."

		if len(byteData) > 0 {
			extendedDataString += "\nRemote: " +
				[]string{"No Auth", "Server Auth", "Mutual Auth"}[byteData[0]] +
				"\nLocal: " +
				[]string{"No Auth", "Server Auth", "Mutual Auth"}[byteData[1]]
		}
	case TLSServerCertificateSet:
		extendedDataString = "TLS server certificate was defined."

		if len(byteData) > 0 {
			extendedDataString += readCertificateSerialNumberToString(byteData)
		}
	case TLSServerCertificateRemoved:
		extendedDataString = "TLS server certificate was removed."

		if len(byteData) > 0 {
			extendedDataString += readCertificateSerialNumberToString(byteData)
		}
	case TLSTrustedRootCertificateAdded:
		extendedDataString = "TLS trusted root certificate was added."

		if len(byteData) > 0 {
			extendedDataString += readCertificateSerialNumberToString(byteData)
		}
	case TLSTrustedRootCertificateRemoved:
		extendedDataString = "TLS trusted root certificate was removed."

		if len(byteData) > 0 {
			extendedDataString += readCertificateSerialNumberToString(byteData)
		}
	case TLSPreSharedKeySet:
		extendedDataString = "TLS pre-shared key was defined."
	case KerberosSettingsModified:
		extendedDataString = "Kerberos settings were modified."

		if len(data) > 0 {
			extendedDataString += "\nTime tolerance: " + fmt.Sprint(byteData[0])
		}
	case KerberosMasterKeyModified:
		extendedDataString = "Kerberos master key or passphrase was modified."
	case FlashWearOutCountersReset:
		extendedDataString = "Flash wear out counter was reset."
	case PowerPackageModified:
		extendedDataString = "Active power package was set."

		if len(data) > 0 {
			extendedDataString += "\nPower policy: " + fmt.Sprint(byteData[0])
		}
	case SetRealmAuthenticationMode:
		extendedDataString = "Realm authentication mode changed."

		if len(data) > 0 {
			byteData := []byte(data)
			extendedDataString += "\n" + RealmNames[byteData[0]] + ", " + []string{"NoAuth", "Auth", "Disabled"}[byteData[4]]
		}
	case UpgradeClientToAdmin:
		extendedDataString = "The control mode of the Intel AMT was changed from Client control to Admin control."
	case AMTUnProvisioningStarted:
		extendedDataString = "Intel AMT UnProvisioned Started."

		if len(data) > 0 {
			byteData := []byte(data)
			extendedDataString += "\nInitiator: " + []string{"BIOS", "MEBx", "Local MEI", "Local WSMAN", "Remote WSMAN"}[byteData[0]]
		}
	default:
		extendedDataString = UnknownEventID
	}

	return extendedDataString
}

func parseRemoteControlEvents(eventId int, data string) string {
	byteData := []byte(data)

	var extendedDataString string

	const (
		PerformedPowerUp           = 0
		PerformedPowerDown         = 1
		PerformedPowerCycle        = 2
		PerformedReset             = 3
		SetBootOptions             = 4
		PerformedGracefulPowerDown = 5
		PerformedGracefulReset     = 6
		PerformedStandby           = 7
		PerformedHibernate         = 8
		PerformedNMI               = 9
	)

	switch eventId {
	case PerformedPowerUp:
		extendedDataString = "Remote power up initiated."

		if len(byteData) > 0 {
			rce := readBootOptionsData(byteData)
			extendedDataString += remoteControlEventToString(rce)
		}
	case PerformedPowerDown:
		extendedDataString = "Remote power down initiated."
	case PerformedPowerCycle:
		extendedDataString = "Remote power cycle initiated."

		if len(byteData) > 0 {
			rce := readBootOptionsData(byteData)
			extendedDataString += remoteControlEventToString(rce)
		}
	case PerformedReset:
		extendedDataString = "Remote reset initiated."

		if len(byteData) > 0 {
			rce := readBootOptionsData(byteData)
			extendedDataString += remoteControlEventToString(rce)
		}
	case SetBootOptions:
		extendedDataString = "Boot options were set."

		if len(byteData) > 0 {
			rce := readBootOptionsData(byteData)
			extendedDataString += remoteControlEventToString(rce)
		}
	case PerformedGracefulPowerDown:
		extendedDataString = "Remote graceful power down initiated."
	case PerformedGracefulReset:
		extendedDataString = "Remote graceful reset initiated."

		if len(byteData) > 0 {
			rce := readBootOptionsData(byteData)
			extendedDataString += remoteControlEventToString(rce)
		}
	case PerformedStandby:
		extendedDataString = "Remote standby initiated."
	case PerformedHibernate:
		extendedDataString = "Remote hibernate initiated."
	case PerformedNMI:
		extendedDataString = "Remote NMI initiated."
	default:
		extendedDataString = UnknownEventID
	}

	return extendedDataString
}

func parseRedirectionManagerEvents(eventId int) string {
	var extendedDataString string

	const (
		IDERSessionOpened = 0
		IDERSessionClosed = 1
		IDEREnabled       = 2
		IDERDisabled      = 3
		SoLSessionOpened  = 4
		SoLSessionClosed  = 5
		SoLEnabled        = 6
		SoLDisabled       = 7
		KVMSessionStarted = 8
		KVMSessionEnded   = 9
		KVMEnabled        = 10
		KVMDisabled       = 11
		VNCPasswordFailed = 12
	)

	switch eventId {
	case IDERSessionOpened:
		extendedDataString = "An application opened a Storage Redirection session."
	case IDERSessionClosed:
		extendedDataString = "An application or firmware closed a Storage Redirection session."
	case IDEREnabled:
		extendedDataString = "Storage Redirection was enabled."
	case IDERDisabled:
		extendedDataString = "Storage Redirection was disabled."
	case SoLSessionOpened:
		extendedDataString = "An application opened a Serial Over LAN session."
	case SoLSessionClosed:
		extendedDataString = "An application or firmware closed a Serial Over LAN session."
	case SoLEnabled:
		extendedDataString = "Serial Over LAN was enabled."
	case SoLDisabled:
		extendedDataString = "Serial Over LAN was disabled."
	case KVMSessionStarted:
		extendedDataString = "An application opened a Keyboard-Video-Mouse session."
	case KVMSessionEnded:
		extendedDataString = "An application or firmware closed a Keyboard-Video-Mouse session."
	case KVMEnabled:
		extendedDataString = "Keyboard-Video-Mouse was enabled."
	case KVMDisabled:
		extendedDataString = "Keyboard-Video-Mouse was disabled."
	case VNCPasswordFailed:
		extendedDataString = "Incorrect Remote Frame Buffer (RFB) password entered 3 times."
	}

	return extendedDataString
}

func parseFirmwareUpdateManagerEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		FirmwareUpdated = 0
		FirmwareFailed  = 1
	)

	if len(data) > 0 {
		byteData := []byte(data)
		buf := bytes.NewBuffer(byteData)

		switch eventId {
		case FirmwareUpdated:
			oldFWVersion := FWVersion{}
			newFWVersion := FWVersion{}

			readFWVersion(buf, &oldFWVersion)
			readFWVersion(buf, &newFWVersion)

			extendedDataString = "Firmware update was started.\nOld version: " +
				fmt.Sprint(oldFWVersion.Major) + "." +
				fmt.Sprint(oldFWVersion.Minor) + "." +
				fmt.Sprint(oldFWVersion.Hotfix) + "." +
				fmt.Sprint(oldFWVersion.Build) +
				"\nNew version: " +
				fmt.Sprint(newFWVersion.Major) + "." +
				fmt.Sprint(newFWVersion.Minor) + "." +
				fmt.Sprint(newFWVersion.Hotfix) + "." +
				fmt.Sprint(newFWVersion.Build)
		case FirmwareFailed:
			updateFailure := FWUpdateFailure{}

			binary.Read(buf, binary.LittleEndian, &updateFailure.Type)
			binary.Read(buf, binary.LittleEndian, &updateFailure.Reason)

			extendedDataString = "Firmware update failed.\nFailure Type: " +
				fmt.Sprint(updateFailure.Type) + "\nFailure Reason: " + fmt.Sprint(updateFailure.Reason)
		}
	}

	return extendedDataString
}

func parseSecurityAuditLog(eventId int, data string) string {
	var extendedDataString string

	const (
		SecurityAuditLogCleared     = 0
		SecurityAuditPolicyModified = 1
		SecurityAuditLogDisabled    = 2
		SecurityAuditLogEnabled     = 3
		SecurityAuditLogExported    = 4
		SecurityAuditLogRecovered   = 5
	)

	switch eventId {
	case SecurityAuditLogCleared:
		extendedDataString = "Audit log was cleared."
	case SecurityAuditPolicyModified:
		extendedDataString = "Audit policy event was enabled or disabled."
	case SecurityAuditLogDisabled:
		extendedDataString = "Access monitor feature was disabled."
	case SecurityAuditLogEnabled:
		extendedDataString = "Access monitor feature was enabled."
	case SecurityAuditLogExported:
		extendedDataString = "Audit log signature and log-related information was exported."
	case SecurityAuditLogRecovered:
		if len(data) > 0 {
			extendedDataString = "Internal check of audit log resulted in a recovery action.\nReason: " +
				[]string{"Unknown", "Migration failure", "Initialization failure"}[data[0]]
		}
	}

	return extendedDataString
}

func parseNetworkTimeEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		IntelMETimeSet = 0
	)

	switch eventId {
	case IntelMETimeSet:
		extendedDataString = "Command received to set Intel AMT local time."

		if len(data) > 0 {
			extendedDataString += "\nTime: " + time.Unix(int64(common.ReadInt(data, 0)), 0).String()
		}
	}

	return extendedDataString
}

func parseNetworkAdministrationEvents(eventId int, data string) string {
	var extendedDataString string

	InterfaceIDGenType := map[int]string{
		0: "Random ID",
		1: "Intel ID",
		2: "Manual ID",
		3: "Invalid ID",
	}

	const (
		TCPIPParametersSet = 0
		HostNameSet        = 1
		DomainNameSet      = 2
		VLANParametersSet  = 3
		LinkPolicySet      = 4
		IPv6ParametersSet  = 5
	)

	switch eventId {
	case TCPIPParametersSet:
		extendedDataString = "TCP/IP parameters were set."

		if len(data) > 0 {
			event := readNetworkAdministrationEventData(TCPIPParametersSet, []byte(data))
			extendedDataString += "\nDHCP Enabled: " + []string{"Disabled", "Enabled"}[event.DHCPEnabled] +
				"\nStatic IP: " + convertUINT32ToIPv4(event.IPV4Address).String() +
				"\nSubnet Mask: " + convertUINT32ToIPv4(event.SubnetMask).String() +
				"\nGateway: " + convertUINT32ToIPv4(event.Gateway).String()
		}
	case HostNameSet:
		extendedDataString = "Host name was set"

		if len(data) > 0 {
			event := readNetworkAdministrationEventData(HostNameSet, []byte(data))
			extendedDataString += " to " + event.HostName
		}
	case DomainNameSet:
		extendedDataString = "Domain name was set"

		if len(data) > 0 {
			event := readNetworkAdministrationEventData(DomainNameSet, []byte(data))
			extendedDataString += " to " + event.DomainName
		}
	case VLANParametersSet:
		extendedDataString = "VLAN tag was set"

		if len(data) > 0 {
			event := readNetworkAdministrationEventData(VLANParametersSet, []byte(data))

			if event.VLANTag == 0 {
				extendedDataString += " to disabled"
			} else {
				extendedDataString += " to " + fmt.Sprint(event.VLANTag)
			}
		}
	case LinkPolicySet:
		extendedDataString = "Link policy was set"

		if len(data) > 0 {
			event := readNetworkAdministrationEventData(LinkPolicySet, []byte(data))
			extendedDataString += " to " + fmt.Sprint(event.LinkPolicy)
		}
	case IPv6ParametersSet:
		extendedDataString = "IPv6 parameters were set."

		if len(data) > 0 {
			event := readNetworkAdministrationEventData(IPv6ParametersSet, []byte(data))
			extendedDataString += "\nIPv6: " + []string{"Disabled", "Enabled"}[event.IPV6Enabled] +
				"\nInterface Gen Type: " + InterfaceIDGenType[int(event.InterfaceIDGenType)]

			if event.InterfaceIDGenType == 2 {
				interfaceId := fmt.Sprintf("%02x%02x:%02x%02x:%02x%02x:%02x%02x",
					event.InterfaceID[0], event.InterfaceID[1], event.InterfaceID[2], event.InterfaceID[3],
					event.InterfaceID[4], event.InterfaceID[5], event.InterfaceID[6], event.InterfaceID[7])
				extendedDataString += "\nInterface ID: " + interfaceId
			}

			extendedDataString += "\nIPv6 Address: " + net.IP(event.IPV6Address).String() +
				"\nIPv6 Gateway: " + net.IP(event.IPV6Gateway).String() +
				"\nIPv6 Primary DNS: " + net.IP(event.IPV6PrimaryDNS).String() +
				"\nIPv6 Secondary DNS: " + net.IP(event.IPV6SecondaryDNS).String()
		}
	}

	return extendedDataString
}

func parseStorageAdministrationEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		GlobalStorageAttributesSet = 0
		StorageEACLModified        = 1
		StorageFPACLModified       = 2
		StorageWriteOperation      = 3
	)

	switch eventId {
	case GlobalStorageAttributesSet:
		extendedDataString = "Global storage attributes were set."

		if len(data) > 0 {
			byteData := []byte(data)
			buf := bytes.NewBuffer(byteData)

			event := StorageAdministrationEvent{}

			binary.Read(buf, binary.LittleEndian, &event.MaxPartnerStorage)
			binary.Read(buf, binary.LittleEndian, &event.MaxNonPartnerTotalAllocationSize)

			extendedDataString += "\nMax Partner Storage: " + fmt.Sprint(event.MaxPartnerStorage) +
				"\nMax Non-Partner Total Allocation Size: " + fmt.Sprint(event.MaxNonPartnerTotalAllocationSize)
		}
	case StorageEACLModified:
		extendedDataString = "Storage EACL entry was added or removed."
	case StorageFPACLModified:
		extendedDataString = "Storage FPACL entry was added, removed, or updated."
	case StorageWriteOperation:
		extendedDataString = "Application wrote a block to storage."
	}

	return extendedDataString
}

func parseEventManagerEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		AlertSubscribed   = 0
		AlertUnsubscribed = 1
		EventLogCleared   = 2
		EventLogFrozen    = 3
	)

	switch eventId {
	case AlertSubscribed, AlertUnsubscribed:
		if eventId == AlertSubscribed {
			extendedDataString = "An alert subscription was created successfully."
		} else {
			extendedDataString = "An existing alert subscription was cancelled."
		}

		if len(data) > 0 {
			event := readEventManagerEventData(AlertSubscribed, []byte(data))
			extendedDataString += eventManagerEventDataToString(event)
		}
	case EventLogCleared:
		extendedDataString = "Event log was cleared of existing records."
	case EventLogFrozen:
		if len(data) > 0 {
			event := readEventManagerEventData(EventLogFrozen, []byte(data))
			extendedDataString = "Event log was " + []string{"unfrozen", "frozen"}[event.Freeze]
		}
	}

	return extendedDataString
}

func parseSystemDefenseManagerEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		SDFilterAdded            = 0
		SDFilterRemoved          = 1
		SDPolicyAdded            = 2
		SDPolicyRemoved          = 3
		SDDefaultPolicySet       = 4
		SDHeuristicsOptionSet    = 5
		SDHeuristicsStateCleared = 6
	)

	switch eventId {
	case SDFilterAdded:
		extendedDataString = "Attempt made to add a system defense filter."
	case SDFilterRemoved:
		extendedDataString = "System defense filter was removed successfully."

		if len(data) > 0 {
			event := readSystemDefenseManagerEventData(SDFilterRemoved, []byte(data))
			extendedDataString += "\nFilter Handle: " + fmt.Sprint(event.FilterHandle)
		}
	case SDPolicyAdded:
		extendedDataString = "Attempt made to add a system defense policy."
	case SDPolicyRemoved:
		extendedDataString = "System defense policy was removed successfully."

		if len(data) > 0 {
			event := readSystemDefenseManagerEventData(SDPolicyRemoved, []byte(data))
			extendedDataString += "\nPolicy Handle: " + fmt.Sprint(event.PolicyHandle)
		}
	case SDDefaultPolicySet:
		extendedDataString = "System defense policy selected."

		if len(data) > 0 {
			event := readSystemDefenseManagerEventData(SDDefaultPolicySet, []byte(data))
			extendedDataString += "\nHardware Interface: " + fmt.Sprint(event.HardwareInterface) +
				"\nPolicy Handle: " + fmt.Sprint(event.PolicyHandle)
		}
	case SDHeuristicsOptionSet:
		extendedDataString = "System defense heuristics settings were set successfully."

		if len(data) > 0 {
			event := readSystemDefenseManagerEventData(SDHeuristicsOptionSet, []byte(data))
			extendedDataString += "\nInterface Handle: " + fmt.Sprint(event.InterfaceHandle) +
				"\nBlock All: " + fmt.Sprint(event.BlockAll) +
				"\nBlock Offensive Port: " + fmt.Sprint(event.BlockOffensivePort) +
				"\nPolicy Handle: " + fmt.Sprint(event.PolicyHandle)
		}
	case SDHeuristicsStateCleared:
		extendedDataString = "System defense heuristics settings were removed successfully."

		if len(data) > 0 {
			event := readSystemDefenseManagerEventData(SDHeuristicsStateCleared, []byte(data))
			extendedDataString += "\nInterface Handle: " + fmt.Sprint(event.InterfaceHandle)
		}
	}

	return extendedDataString
}

func parseAgentPresenceManagerEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		AgentWatchdogAdded     = 0
		AgentWatchdogRemoved   = 1
		AgentWatchdogActionSet = 2
	)

	switch eventId {
	case AgentWatchdogAdded:
		extendedDataString = "An application entry was created to be monitored."

		if len(data) > 0 {
			event := readAgentPresenceManagerEventData(AgentWatchdogAdded, []byte(data))
			extendedDataString += "\nAgent ID: " + fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
				event.AgentID[0:4],
				event.AgentID[4:6],
				event.AgentID[6:8],
				event.AgentID[8:10],
				event.AgentID[10:]) +
				"\nAgent Heartbeat Time: " + fmt.Sprint(event.AgentHeartBeatTime) +
				"\nAgent Startup Time: " + fmt.Sprint(event.AgentStartupTime)
		}
	case AgentWatchdogRemoved:
		extendedDataString = "An application entry was removed."

		if len(data) > 0 {
			event := readAgentPresenceManagerEventData(AgentWatchdogRemoved, []byte(data))
			extendedDataString += "\nAgent ID: " + fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
				event.AgentID[0:4],
				event.AgentID[4:6],
				event.AgentID[6:8],
				event.AgentID[8:10],
				event.AgentID[10:])
		}
	case AgentWatchdogActionSet:
		extendedDataString = "Actions were set, added, or removed for an application watchdog entry."

		if len(data) > 0 {
			event := readAgentPresenceManagerEventData(AgentWatchdogActionSet, []byte(data))
			extendedDataString += "\nAgent ID: " + fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
				event.AgentID[0:4],
				event.AgentID[4:6],
				event.AgentID[6:8],
				event.AgentID[8:10],
				event.AgentID[10:])
		}
	}

	return extendedDataString
}

func parseWirelessConfigurationEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		WirelessProfileAdded                              = 0
		WirelessProfileRemoved                            = 1
		WirelessProfileUpdated                            = 2
		WirelessProfileModified                           = 3
		WirelessLinkPreferenceChanged                     = 4
		WirelessProfileShareWithUEFIEnabledSettingChanged = 5
	)

	switch eventId {
	case WirelessProfileAdded:
		extendedDataString = "A new profile was added."

		if len(data) > 0 {
			event := readWirelessConfigurationEventData(WirelessProfileAdded, []byte(data))
			extendedDataString += "\nSSID: " + strings.Trim(fmt.Sprint(string(event.SSID)), "\x00") +
				"\nProfile Priority: " + fmt.Sprint(event.ProfilePriority) +
				"\nProfile Name Length: " + fmt.Sprint(event.ProfileNameLength) +
				"\nProfile Name: " + fmt.Sprint(string(event.ProfileName))
		}
	case WirelessProfileRemoved:
		extendedDataString = "An existing profile was removed."

		if len(data) > 0 {
			event := readWirelessConfigurationEventData(WirelessProfileRemoved, []byte(data))
			extendedDataString += "\nProfile Name Length: " + fmt.Sprint(event.ProfileNameLength) +
				"\nProfile Name: " + fmt.Sprint(string(event.ProfileName))
		}
	case WirelessProfileUpdated:
		extendedDataString = "An existing profile was updated."

		if len(data) > 0 {
			event := readWirelessConfigurationEventData(WirelessProfileUpdated, []byte(data))
			extendedDataString += "\nSSID: " + strings.Trim(fmt.Sprint(string(event.SSID)), "\x00") +
				"\nProfile Priority: " + fmt.Sprint(event.ProfilePriority) +
				"\nProfile Name Length: " + fmt.Sprint(event.ProfileNameLength) +
				"\nProfile Name: " + fmt.Sprint(string(event.ProfileName))
		}
	case WirelessProfileModified:
		extendedDataString = "An existing profile sync was modified."

		if len(data) > 0 {
			event := readWirelessConfigurationEventData(WirelessProfileModified, []byte(data))
			extendedDataString += "\nProfile sync " + []string{"is disabled", "user", "admin", "is unrestricted"}[event.ProfileSync]
		}
	case WirelessLinkPreferenceChanged:
		extendedDataString = "An existing profile link preference was changed."

		if len(data) > 0 {
			event := readWirelessConfigurationEventData(WirelessLinkPreferenceChanged, []byte(data))
			extendedDataString += "\nTimeout: " + fmt.Sprint(event.Timeout)
			extendedDataString += "\nLink Preference: " + []string{"none", "ME", "Host"}[int(event.LinkPreference)]
		}
	case WirelessProfileShareWithUEFIEnabledSettingChanged:
		extendedDataString = fmt.Sprintf("Wireless profile share with UEFI was set to %s.", []string{"Disabled", "Enabled"}[data[0]])
	}

	return extendedDataString
}

func parseEndpointAccessControlEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		EACPostureSignerSet   = 0
		EACEnabled            = 1
		EACDisabled           = 2
		EASPostureStateUpdate = 3
		EACSetOptions         = 4
	)

	switch eventId {
	case EACPostureSignerSet:
		extendedDataString = "A certificate handle for signing EAC postures was either set or removed."
	case EACEnabled:
		extendedDataString = "EAC was set to enabled by WS-MAN interface."
	case EACDisabled:
		extendedDataString = "EAC was set to disabled by WS-MAN interface."
	case EASPostureStateUpdate:
		extendedDataString = "Controllable fields of EAC posture were reset manually by WS-MAN interface."
	case EACSetOptions:
		extendedDataString = "EAC options were changed."

		if len(data) > 0 {
			byteData := []byte(data)
			buf := bytes.NewBuffer(byteData)

			var eacVendors uint32
			_ = binary.Read(buf, binary.LittleEndian, &eacVendors)
			extendedDataString += "\nEAC Vendors: " + fmt.Sprint(eacVendors)
		}
	}

	return extendedDataString
}

func parseKeyboardVideoMouseEvents(eventId int) string {
	var extendedDataString string

	const (
		KVMOptInEnabled     = 0
		KVMOptInDisabled    = 1
		KVMPasswordChanged  = 2
		KVMConsentSucceeded = 3
		KVMConsentFailed    = 4
	)

	switch eventId {
	case KVMOptInEnabled:
		extendedDataString = "User consent for a KVM session is now required."
	case KVMOptInDisabled:
		extendedDataString = "User consent for a KVM session is no longer required."
	case KVMPasswordChanged:
		extendedDataString = "RFB password for KVM session has changed."
	case KVMConsentSucceeded:
		extendedDataString = operationStatusToString[0]
	case KVMConsentFailed:
		extendedDataString = operationStatusToString[1]
	}

	return extendedDataString
}

func parseUserOptInEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		OptInPolicyChange      = 0
		SendConsentCodeEvent   = 1
		StartOptInBlockedEvent = 2
	)

	switch eventId {
	case OptInPolicyChange:
		extendedDataString = "A user has modified the opt-in policy settings."

		if len(data) > 0 {
			event := readUserOptInEventData(OptInPolicyChange, []byte(data))
			extendedDataString += "\nPrevious Opt-In Policy: " + optInPolicyToString[int(event.PreviousOptInPolicy)] +
				"\nCurrent Opt-In Policy: " + optInPolicyToString[int(event.CurrentOptInPolicy)]
		}
	case SendConsentCodeEvent:
		extendedDataString = "The remote operator sent a consent code."

		if len(data) > 0 {
			event := readUserOptInEventData(SendConsentCodeEvent, []byte(data))
			extendedDataString += "\n" + operationStatusToString[int(event.OperationStatus)]
		}
	case StartOptInBlockedEvent:
		extendedDataString = "The remote operator attempted to send a start opt-in request, but the request was blocked (denial-of-service attack prevention)."
	}

	return extendedDataString
}

func parseWatchdogEvents(eventId int, data string) string {
	var extendedDataString string

	const (
		WatchdogResetTriggeringOptionsChanged = 0
		WatchdogActionPairingChanged          = 1
	)

	switch eventId {
	case WatchdogResetTriggeringOptionsChanged:
		extendedDataString = "A user has modified the watchdog action settings."
	case WatchdogActionPairingChanged:
		extendedDataString = "A user has modified a watchdog to add, remove, or alter the watchdog action connected to it."

		if len(data) > 0 {
			byteData := []byte(data)

			var opStatus uint8

			buf := bytes.NewBuffer(byteData)

			_ = binary.Read(buf, binary.LittleEndian, &opStatus)

			extendedDataString += "\n" + operationStatusToString[int(opStatus)]
		}
	}

	return extendedDataString
}

const (
	HTTPDigest     byte = 0
	Kerberos       byte = 1
	Local          byte = 2
	KvmDefaultPort byte = 3
)

func readProvisioningCompletedEventData(data []byte) ProvisioningParameters {
	buf := bytes.NewBuffer(data)
	event := ProvisioningParameters{}

	// Read ProvisioningMethod
	_ = binary.Read(buf, binary.LittleEndian, &event.ProvisioningMethod)

	// Read HashType
	_ = binary.Read(buf, binary.LittleEndian, &event.HashType)

	// Read TrustedRootCertHash based on HashType
	switch event.HashType {
	case 1: // SHA1_160
		event.TrustedRootCertHash = make([]byte, 20)
	case 2: // SHA_256
		event.TrustedRootCertHash = make([]byte, 32)
	case 3: // SHA_384
		event.TrustedRootCertHash = make([]byte, 48)
	}

	buf.Read(event.TrustedRootCertHash)

	// Read NumberOfCertificates
	_ = binary.Read(buf, binary.LittleEndian, &event.NumberOfCertificates)

	// Read CertSerialNumbers
	for i := 0; i < int(event.NumberOfCertificates); i++ {
		serialNumber := make([]byte, 16)
		buf.Read(serialNumber)
		event.CertSerialNumbers = append(event.CertSerialNumbers, hex.EncodeToString(serialNumber))
	}

	// Read AdditionalCaSerialNumbers
	_ = binary.Read(buf, binary.LittleEndian, &event.AdditionalCaSerialNumbers)

	// Read ProvServFQDNLength
	_ = binary.Read(buf, binary.LittleEndian, &event.ProvServFQDNLength)

	// Read ProvServFQDN
	fqdn := make([]byte, event.ProvServFQDNLength)
	buf.Read(fqdn)

	event.ProvServFQDN = string(fqdn)

	return event
}

func provisioningCompletedToString(provisioningCompleted *ProvisioningParameters) string {
	s := fmt.Sprintf("\nProvisioning Method: %s", provisioningMethodToString[int(provisioningCompleted.ProvisioningMethod)])

	if provisioningCompleted.HashType != 0 {
		s += fmt.Sprintf("\nHash Type: %s", provisioningHashTypeToString[int(provisioningCompleted.HashType)])
	}

	if len(provisioningCompleted.TrustedRootCertHash) > 0 {
		s += fmt.Sprintf("\nTrusted Root Cert Hash: %s", hex.EncodeToString(provisioningCompleted.TrustedRootCertHash))
	}

	if provisioningCompleted.NumberOfCertificates > 0 {
		s += fmt.Sprintf("\nNumber of Certificates: %d", provisioningCompleted.NumberOfCertificates)
		s += fmt.Sprintf("\nCert Serial Numbers (first 3): %v", provisioningCompleted.CertSerialNumbers)

		if provisioningCompleted.AdditionalCaSerialNumbers > 3 {
			s += fmt.Sprintf("\nThere are %d additional certificates", provisioningCompleted.AdditionalCaSerialNumbers)
		}
	}

	if provisioningCompleted.ProvServFQDNLength > 0 {
		s += fmt.Sprintf("\nProvisioning Server FQDN: %s", provisioningCompleted.ProvServFQDN)
	}

	return s
}

func aclEntryToString(entry *ACLEntry) string {
	s := fmt.Sprintf("\nInitiator Type: %s", initiatorTypeToString[int(entry.InitiatorType)])

	if entry.UsernameLength == 0 {
		s += fmt.Sprintf("\nSID: %d", entry.SID)

		if entry.DomainLength > 0 {
			s += fmt.Sprintf("\nDomain: %s", entry.Domain)
		}
	} else {
		s += fmt.Sprintf("\nUsername: %s", entry.Username)
	}

	return s
}

func aclEntryModifiedToString(entry *ACLEntry) string {
	parameterModifiedStr := ""

	if entry.ParameterModified&0x01 != 0 {
		parameterModifiedStr += "Username, "
	}

	if entry.ParameterModified&0x02 != 0 {
		parameterModifiedStr += "Password, "
	}

	if entry.ParameterModified&0x04 != 0 {
		parameterModifiedStr += "Local realms, "
	}

	if entry.ParameterModified&0x08 != 0 {
		parameterModifiedStr += "Remote realms, "
	}

	if entry.ParameterModified&0x10 != 0 {
		parameterModifiedStr += "Kerberos domain, "
	}

	if entry.ParameterModified&0x20 != 0 {
		parameterModifiedStr += "SID, "
	}

	if len(parameterModifiedStr) > 0 {
		parameterModifiedStr = parameterModifiedStr[:len(parameterModifiedStr)-2] // Remove the trailing comma and space
	} else {
		parameterModifiedStr = "None"
	}

	s := fmt.Sprintf("\nParameter(s) Modified: %s\nInitiator Type: %s", parameterModifiedStr, initiatorTypeToString[int(entry.InitiatorType)])

	if entry.UsernameLength == 0 {
		s += fmt.Sprintf("\nSID: %d", entry.SID)
		if entry.DomainLength > 0 {
			s += fmt.Sprintf("\nDomain: %s", entry.Domain)
		}
	} else {
		s += fmt.Sprintf("\nUsername: %s", entry.Username)
	}

	return s
}

func aclEntryEnabledToString(entry *ACLEntry) string {
	s := fmt.Sprintf("\nEntry State: %s\nInitiator Type: %s", []string{"Disabled", "Enabled"}[int(entry.EntryState)], initiatorTypeToString[int(entry.InitiatorType)])
	if entry.UsernameLength == 0 {
		s += fmt.Sprintf("\nSID: %d", entry.SID)
		if entry.DomainLength > 0 {
			s += fmt.Sprintf("\nDomain: %s", entry.Domain)
		}
	} else {
		s += fmt.Sprintf("\nUsername: %s", entry.Username)
	}

	return s
}

func readCertificateSerialNumberToString(data []byte) string {
	hexString := hex.EncodeToString(data)

	return "\nCertificate serial number: " + hexString
}

func getCommonACLProperties(buf *bytes.Buffer, entry *ACLEntry) {
	// Read Initiator Type
	_ = binary.Read(buf, binary.LittleEndian, &entry.InitiatorType)

	// Read Username Length
	_ = binary.Read(buf, binary.LittleEndian, &entry.UsernameLength)

	if entry.UsernameLength == 0 {
		_ = binary.Read(buf, binary.LittleEndian, &entry.SID)

		_ = binary.Read(buf, binary.LittleEndian, &entry.DomainLength)

		domain := make([]byte, entry.DomainLength)

		buf.Read(domain)

		entry.Domain = string(domain)
	} else {
		username := make([]byte, entry.UsernameLength)

		buf.Read(username)

		entry.Username = string(username)
	}
}

func getACLParameters(buf *bytes.Buffer, entry *ACLEntry) {
	// Read Parameter Modified
	_ = binary.Read(buf, binary.LittleEndian, &entry.ParameterModified)
}

func readACLData(id int, data []byte) ACLEntry {
	buf := *bytes.NewBuffer(data)
	entry := ACLEntry{}

	switch id {
	case 2, 4:
		getCommonACLProperties(&buf, &entry)
	case 3:
		getACLParameters(&buf, &entry)
		getCommonACLProperties(&buf, &entry)
	case 6:
		_ = binary.Read(&buf, binary.LittleEndian, &entry.EntryState)
		getCommonACLProperties(&buf, &entry)
	}

	return entry
}

func readBootOptionsData(data []byte) RemoteControlEvent {
	buf := bytes.NewBuffer(data)
	rce := RemoteControlEvent{}

	_ = binary.Read(buf, binary.LittleEndian, &rce.SpecialCommand)
	_ = binary.Read(buf, binary.LittleEndian, &rce.SpecialCommandParameterHighByte)
	_ = binary.Read(buf, binary.LittleEndian, &rce.SpecialCommandParameterLowByte)
	_ = binary.Read(buf, binary.LittleEndian, &rce.BootOptionsMaskByte1)
	_ = binary.Read(buf, binary.LittleEndian, &rce.BootOptionsMaskByte2)
	_ = binary.Read(buf, binary.LittleEndian, &rce.OEMParameterByte1)
	_ = binary.Read(buf, binary.LittleEndian, &rce.OEMParameterByte2)

	return rce
}

var remoteControlSpecialCommandToString = map[int]string{
	0:   "None",
	1:   "PXE",
	2:   "HDD",
	3:   "HDD - Safe Mode",
	4:   "Diagnostic Partition",
	5:   "CD or DVD Boot",
	193: "Intel Command",
	194: "Intel Command - HTTPS Boot",
	195: "Intel Command - PBA Boot",
}

var remoteControlSpecialCommandParameterHighByteToString = map[int]string{
	0: "SRoU is connected to CD/DVD device",
	1: "SRoU is connected to floppy device",
}

var remoteControlSpecialCommandParameterLowByte193ToString = map[int]string{
	1:   "SRoU is to be used on next boot.",
	2:   "BIOS will enforce secure boot over SRoU.",
	4:   "BIOS is to be re-flashed on next boot.",
	8:   "Boot into BIOS SETUP screen.",
	16:  "BIOS Pause on the next boot.",
	32:  "BIOS is to participate in a KVM session.",
	64:  "BIOS is to start a Remote Secure Erase (RSE) session.",
	128: "BIOS is to start a Remote Platform Erase (RPE) session.",
}

var remoteControlSpecialCommandParameterLowByte194or195ToString = map[int]string{
	2:  "BIOS will enforce secure boot on next boot.",
	32: "BIOS is to participate in a KVM session.",
}

var remoteControlSpecialCommandParameterLowByteOthersToString = map[int]string{
	32: "BIOS is to participate in a KVM session.",
}

func remoteControlEventToString(rce RemoteControlEvent) string {
	// s := "\nBoot Media: " + []string{"None", "PXE", "HDD", "HDD - Safe Mode", "Diagnostic Partition", "CD or DVD Boot", "USB", "PXE", "BIOS Setup"}[rce.BootOptions] +
	// 	"\n " + "Boot Media Override: " + []string{"Disabled", "Enabled"}[rce.SpecialCommand] +
	// 	"\n " + "BIOS Pause: " + []string{"Disabled", "Enabled"}[rce.OEMParameters] +
	// 	"\n " + "BIOS Pause Key: " + []string{"None", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F11", "F12"}[rce.SpecialCommandParameter]
	s := "\nSpecial Command: " + remoteControlSpecialCommandToString[int(rce.SpecialCommand)]

	if rce.SpecialCommand == 193 {
		s += "\nSpecial Command Parameter High Byte: " + remoteControlSpecialCommandParameterHighByteToString[int(rce.SpecialCommandParameterHighByte)]
		s += "\nSpecial Command Parameter Low Byte: " + remoteControlSpecialCommandParameterLowByte193ToString[int(rce.SpecialCommandParameterLowByte)]
	} else if rce.SpecialCommand > 193 {
		s += "\nSpecial Command Parameter Low Byte: " + remoteControlSpecialCommandParameterLowByte194or195ToString[int(rce.SpecialCommandParameterLowByte)]
	} else {
		s += "\nSpecial Command Parameter Low Byte: " + remoteControlSpecialCommandParameterLowByteOthersToString[int(rce.SpecialCommandParameterLowByte)]
	}

	return s
}

func readFWVersion(buf *bytes.Buffer, version *FWVersion) {
	_ = binary.Read(buf, binary.LittleEndian, &version.Major)
	_ = binary.Read(buf, binary.LittleEndian, &version.Minor)
	_ = binary.Read(buf, binary.LittleEndian, &version.Hotfix)
	_ = binary.Read(buf, binary.LittleEndian, &version.Build)
}

func readNetworkAdministrationEventData(id int, data []byte) NetworkAdministrationEvent {
	buf := bytes.NewBuffer(data)
	event := NetworkAdministrationEvent{}

	switch id {
	case 0:
		_ = binary.Read(buf, binary.LittleEndian, &event.InterfaceHandle)
		_ = binary.Read(buf, binary.LittleEndian, &event.DHCPEnabled)
		_ = binary.Read(buf, binary.LittleEndian, &event.IPV4Address)
		_ = binary.Read(buf, binary.LittleEndian, &event.SubnetMask)
		_ = binary.Read(buf, binary.LittleEndian, &event.Gateway)
		_ = binary.Read(buf, binary.LittleEndian, &event.PrimaryDNS)
		_ = binary.Read(buf, binary.LittleEndian, &event.SecondaryDNS)
	case 1:
		_ = binary.Read(buf, binary.LittleEndian, &event.HostNameLength)

		hostname := make([]byte, event.HostNameLength)

		buf.Read(hostname)

		event.HostName = string(hostname)
	case 2:
		_ = binary.Read(buf, binary.LittleEndian, &event.DomainNameLength)

		domainName := make([]byte, event.DomainNameLength)

		buf.Read(domainName)

		event.DomainName = string(domainName)
	case 3:
		_ = binary.Read(buf, binary.LittleEndian, &event.InterfaceHandle)
		_ = binary.Read(buf, binary.LittleEndian, &event.VLANTag)
	case 4:
		_ = binary.Read(buf, binary.LittleEndian, &event.InterfaceHandle)
		_ = binary.Read(buf, binary.LittleEndian, &event.LinkPolicy)
	case 5:
		_ = binary.Read(buf, binary.LittleEndian, &event.InterfaceHandle)
		_ = binary.Read(buf, binary.LittleEndian, &event.IPV6Enabled)
		_ = binary.Read(buf, binary.LittleEndian, &event.InterfaceIDGenType)

		if event.InterfaceIDGenType == 2 {
			interfaceId := make([]byte, 8)
			buf.Read(interfaceId)
			event.InterfaceID = interfaceId
		}

		ipv6Address := make([]byte, 16)
		buf.Read(ipv6Address)
		event.IPV6Address = ipv6Address

		ipv6Gateway := make([]byte, 16)
		buf.Read(ipv6Gateway)
		event.IPV6Gateway = ipv6Gateway

		ipv6PrimaryDNS := make([]byte, 16)
		buf.Read(ipv6PrimaryDNS)
		event.IPV6PrimaryDNS = ipv6PrimaryDNS

		ipv6SecondaryDNS := make([]byte, 16)
		buf.Read(ipv6SecondaryDNS)
		event.IPV6SecondaryDNS = ipv6SecondaryDNS
	}

	return event
}

func readEventManagerEventData(id int, data []byte) EventManagerEvent {
	buf := bytes.NewBuffer(data)
	event := EventManagerEvent{}

	switch id {
	case 0, 1:
		_ = binary.Read(buf, binary.LittleEndian, &event.PolicyID)
		_ = binary.Read(buf, binary.LittleEndian, &event.SubscriptionAlertType)
		_ = binary.Read(buf, binary.LittleEndian, &event.IPAddrType)

		if event.IPAddrType == 0 {
			ipAddress := make([]byte, 4)
			buf.Read(ipAddress)
			event.AlertTargetIPAddress = ipAddress
		}

		if event.IPAddrType == 1 {
			ipAddress := make([]byte, 16)
			buf.Read(ipAddress)
			event.AlertTargetIPAddress = ipAddress
		}
	case 3:
		_ = binary.Read(buf, binary.LittleEndian, &event.Freeze)
	}

	return event
}

func eventManagerEventDataToString(event EventManagerEvent) string {
	s := "\nPolicy ID: " + fmt.Sprint(event.PolicyID)

	if event.SubscriptionAlertType == 1 {
		s += "\nSubscription Alert Type: SNMP"
	}

	s += "\nIP Address Type: " + []string{"IPv4", "IPv6"}[event.IPAddrType] +
		"\nAlert Target IP Address: " + net.IP(event.AlertTargetIPAddress).String()

	return s
}

func readSystemDefenseManagerEventData(id int, data []byte) SystemDefenseManagerEvent {
	buf := bytes.NewBuffer(data)
	event := SystemDefenseManagerEvent{}

	switch id {
	case 1:
		_ = binary.Read(buf, binary.LittleEndian, &event.FilterHandle)
	case 3:
		_ = binary.Read(buf, binary.LittleEndian, &event.PolicyHandle)
	case 4:
		_ = binary.Read(buf, binary.LittleEndian, &event.HardwareInterface)
		_ = binary.Read(buf, binary.LittleEndian, &event.PolicyHandle)
	case 5:
		_ = binary.Read(buf, binary.LittleEndian, &event.InterfaceHandle)
		_ = binary.Read(buf, binary.LittleEndian, &event.BlockAll)
		_ = binary.Read(buf, binary.LittleEndian, &event.BlockOffensivePort)
		_ = binary.Read(buf, binary.LittleEndian, &event.PolicyHandle)
	case 6:
		_ = binary.Read(buf, binary.LittleEndian, &event.InterfaceHandle)
	}

	return event
}

func readAgentPresenceManagerEventData(id int, data []byte) AgentPresenceManagerEvent {
	buf := bytes.NewBuffer(data)
	event := AgentPresenceManagerEvent{}

	agentId := make([]byte, 16)
	buf.Read(agentId)
	event.AgentID = agentId

	if id == 0 {
		_ = binary.Read(buf, binary.LittleEndian, &event.AgentHeartBeatTime)
		_ = binary.Read(buf, binary.LittleEndian, &event.AgentStartupTime)
	}

	return event
}

func readWirelessConfigurationEventData(id int, data []byte) WirelessConfigurationEvent {
	buf := bytes.NewBuffer(data)
	event := WirelessConfigurationEvent{}

	switch id {
	case 0, 2:
		ssid := make([]byte, 32)
		buf.Read(ssid)
		event.SSID = ssid
		_ = binary.Read(buf, binary.LittleEndian, &event.ProfilePriority)
		_ = binary.Read(buf, binary.LittleEndian, &event.ProfileNameLength)
		profileName := make([]byte, event.ProfileNameLength)
		buf.Read(profileName)
		event.ProfileName = profileName
	case 1:
		_ = binary.Read(buf, binary.LittleEndian, &event.ProfileNameLength)
		profileName := make([]byte, event.ProfileNameLength)
		buf.Read(profileName)
		event.ProfileName = profileName
	case 3:
		_ = binary.Read(buf, binary.LittleEndian, &event.ProfileSync)
	case 4:
		_ = binary.Read(buf, binary.LittleEndian, &event.Timeout)
		_ = binary.Read(buf, binary.LittleEndian, &event.LinkPreference)
	}

	return event
}

func readUserOptInEventData(id int, data []byte) UserOptInEvent {
	buf := bytes.NewBuffer(data)
	event := UserOptInEvent{}

	switch id {
	case 0:
		_ = binary.Read(buf, binary.LittleEndian, &event.PreviousOptInPolicy)
		_ = binary.Read(buf, binary.LittleEndian, &event.CurrentOptInPolicy)
	case 1:
		_ = binary.Read(buf, binary.LittleEndian, &event.OperationStatus)
	}

	return event
}

func convertUINT32ToIPv4(intIP uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, intIP)

	return ip
}

// [initiatorType: number, initiator: string, ptr: number] need to type it.
func getInitiatorInfo(decodedEventRecord string) (initatorType byte, initiator string, ptr int) {
	var userlen uint8

	initiatorType := []byte(decodedEventRecord[4:5])[0]

	switch initiatorType {
	case HTTPDigest:
		userlen = []byte(decodedEventRecord[5:6])[0]
		initiator = decodedEventRecord[6 : 6+userlen]
		ptr = 6 + int(userlen)
	case Kerberos:
		userlen = []byte(decodedEventRecord[9:10])[0]
		initiator = common.GetSidString(decodedEventRecord[10 : 10+userlen])
		ptr = 10 + int(userlen)
	case Local:
		initiator = "Local"
		ptr = 5
	case KvmDefaultPort:
		initiator = "KVM Default Port"
		ptr = 5
	default:
		initiator = ""
		ptr = 0
	}

	return initiatorType, initiator, ptr
}
