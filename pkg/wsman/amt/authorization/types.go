/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package authorization

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUTS
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                    xml.Name `xml:"Body"`
		GetResponse                AuthorizationOccurrence
		EnumerateResponse          common.EnumerateResponse
		PullResponse               PullResponse
		SetAdminResponse           SetAdminAclEntryEx_OUTPUT        `xml:"SetAdminAclEntryEx_OUTPUT"`
		AddUserResponse            AddUserAclEntryEx_OUTPUT         `xml:"AddUserAclEntryEx_OUTPUT"`
		RemoveUserResponse         RemoveUserAclEntry_OUTPUT        `xml:"RemoveUserAclEntry_OUTPUT"`
		GetUserResponse            GetUserAclEntryEx_OUTPUT         `xml:"GetUserAclEntryEx_OUTPUT"`
		GetACLEnabledStateResponse GetAclEnabledState_OUTPUT        `xml:"GetAclEnabledState_OUTPUT"`
		GetAdminResponse           GetAdminAclEntry_OUTPUT          `xml:"GetAdminAclEntry_OUTPUT"`
		GetAdminStatusResponse     GetAdminAclEntryStatus_OUTPUT    `xml:"GetAdminAclEntryStatus_OUTPUT"`
		GetAdminNetStatusResponse  GetAdminNetAclEntryStatus_OUTPUT `xml:"GetAdminNetAclEntryStatus_OUTPUT"`
		SetACLEnabledStateResponse SetAclEnabledState_OUTPUT        `xml:"SetAclEnabledState_OUTPUT"`
		EnumerateUserResponse      EnumerateUserAclEntries_OUTPUT   `xml:"EnumerateUserAclEntries_OUTPUT"`
	}
	SetAdminAclEntryEx_OUTPUT struct {
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
	AddUserAclEntry_OUTPUT struct {
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
	AddUserAclEntryEx_OUTPUT struct {
		ReturnValue ReturnValue `xml:"ReturnValue"`
		Handle      int         `xml:"Handle"`
	}
	RemoveUserAclEntry_OUTPUT struct {
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
	GetAclEnabledState_OUTPUT struct {
		Enabled     bool        `xml:"Enabled"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
	GetAdminAclEntry_OUTPUT struct {
		Username    string      `xml:"Username"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
	GetAdminAclEntryStatus_OUTPUT struct {
		IsDefault   bool        `xml:"IsDefault"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
	GetAdminNetAclEntryStatus_OUTPUT struct {
		IsDefault   bool        `xml:"IsDefault"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
	SetAclEnabledState_OUTPUT struct {
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
	GetUserAclEntryEx_OUTPUT struct {
		DigestUsername   string           `xml:"DigestUsername,omitempty"`
		KerberosUserSid  string           `xml:"KerberosUserSid,omitempty"`
		AccessPermission AccessPermission `xml:"AccessPermission"`
		Realms           []RealmValues    `xml:"Realms"`
		ReturnValue      ReturnValue      `xml:"ReturnValue"`
	}
	EnumerateUserAclEntries_OUTPUT struct {
		TotalCount   int         `xml:"TotalCount"`
		HandlesCount int         `xml:"HandlesCount"`
		Handles      []int       `xml:"Handles"`
		ReturnValue  ReturnValue `xml:"ReturnValue"`
	}
	AuthorizationOccurrence struct {
		XMLName                 xml.Name       `xml:"AMT_AuthorizationService"`
		AllowHttpQopAuthOnly    int            `xml:"AllowHttpQopAuthOnly"`    // Indicates whether using the http "quality of protection" (qop) directive with value auth is allowed
		CreationClassName       string         `xml:"CreationClassName"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified. In Intel AMT Release 6.0 and later releases value is 'AMT_AuthorizationService'
		ElementName             string         `xml:"ElementName"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information.  Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledState            EnabledState   `xml:"EnabledState"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		Name                    string         `xml:"Name"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.  In Intel AMT Release 6.0 and later releases value is 'Intel® AMT Authorization Service'
		RequestedState          RequestedState `xml:"RequestedState"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		SystemCreationClassName string         `xml:"SystemCreationClassName"` // The CreationClassName of the scoping System. In Intel AMT Release 6.0 and later releases value is 'CIM_ComputerSystem'
		SystemName              string         `xml:"SystemName"`              // The Name of the scoping System.  In Intel AMT Release 6.0 and later releases value is 'Intel® AMT'
	}
	PullResponse struct {
		XMLName                      xml.Name                  `xml:"PullResponse"`
		AuthorizationOccurrenceItems []AuthorizationOccurrence `xml:"Items>AMT_AuthorizationService"`
	}

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int

	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int

	// ReturnValue is an integer enumeration that indicates the success or failure of an operation.
	ReturnValue int
)

// ValueMap={0, 1, 2}
//
// Values={LocalAccessPermission, NetworkAccessPermission, AnyAccessPermission}.
type AccessPermission int

// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, ..}
//
// Values={InvalidRealm, ReservedRealm0, RedirectionRealm, PTAdministrationRealm, HardwareAssetRealm, RemoteControlRealm, StorageRealm, EventManagerRealm, StorageAdminRealm, AgentPresenceLocalRealm, AgentPresenceRemoteRealm, CircuitBreakerRealm, NetworkTimeRealm, GeneralInfoRealm, FirmwareUpdateRealm, EITRealm, LocalUN, EndpointAccessControlRealm, EndpointAccessControlAdminRealm, EventLogReaderRealm, AuditLogRealm, ACLRealm, ReservedRealm1, ReservedRealm2, LocalSystemRealm, Reserved}.
type RealmValues int

// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, ..}
//
// Values={InvalidRealm, ReservedRealm0, RedirectionRealm, PTAdministrationRealm, HardwareAssetRealm, RemoteControlRealm, StorageRealm, EventManagerRealm, StorageAdminRealm, AgentPresenceLocalRealm, AgentPresenceRemoteRealm, CircuitBreakerRealm, NetworkTimeRealm, GeneralInfoRealm, FirmwareUpdateRealm, EITRealm, LocalUN, EndpointAccessControlRealm, EndpointAccessControlAdminRealm, EventLogReaderRealm, AuditLogRealm, ACLRealm, ReservedRealm1, ReservedRealm2, LocalSystemRealm, Reserved}.
type Realms int

// INPUTS
// Request Types.
type (
	EnumerateUserAclEntries_INPUT struct {
		XMLName    xml.Name `xml:"h:EnumerateUserAclEntries_INPUT"`
		H          string   `xml:"xmlns:h,attr"`
		StartIndex int      `xml:"h:StartIndex"` // Indicates the first ACL entry to retrieve. To enumerate the entire list, an application sends this message with StartIndex set to 1.
	}

	GetAclEnabledState_INPUT struct {
		XMLName xml.Name `xml:"h:GetAclEnabledState_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"` // Specifies the ACL entry to fetch.
	}
	GetUserAclEntryEx_INPUT struct {
		XMLName xml.Name `xml:"h:GetUserAclEntryEx_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"` // Specifies the ACL entry to fetch.
	}
	RemoveUserAclEntry_INPUT struct {
		XMLName xml.Name `xml:"h:RemoveUserAclEntry_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"` // Specifies the ACL entry to be removed.
	}

	SetAclEnabledState_INPUT struct {
		XMLName xml.Name `xml:"h:SetAclEnabledState_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Handle  int      `xml:"h:Handle"`  // Specifies the ACL entry to update
		Enabled bool     `xml:"h:Enabled"` // Specifies the state of the ACL entry
	}

	SetAdminAclEntryEx_INPUT struct {
		XMLName        xml.Name `xml:"h:SetAdminAclEntryEx_INPUT"`
		H              string   `xml:"xmlns:h,attr"`
		Username       string   `xml:"h:Username"`       // Username for access control. Contains 7-bit ASCII characters. String length is limited to 16 characters. Username cannot be an empty string.
		DigestPassword string   `xml:"h:DigestPassword"` // An MD5 Hash of these parameters concatenated together (Username + ":" + DigestRealm + ":" + Password). The DigestRealm is a field in AMT_GeneralSettings
	}
	AddUserAclEntryEx_INPUT struct {
		XMLName          xml.Name         `xml:"h:AddUserAclEntryEx_INPUT"`
		H                string           `xml:"xmlns:h,attr"`
		DigestUsername   string           `xml:"h:DigestUsername"`            // Username for access control. Contains 7-bit ASCII characters. String length is limited to 16 characters. Username cannot be an empty string.
		DigestPassword   string           `xml:"h:DigestPassword"`            // An MD5 Hash of these parameters concatenated together (Username + ":" + DigestRealm + ":" + Password). The DigestRealm is a field in AMT_GeneralSettings
		KerberosUserSid  string           `xml:"h:KerberosUserSid,omitempty"` // Descriptor for user (SID) which is authenticated using the Kerberos Authentication. Byte array, specifying the Security Identifier (SID) according to the Kerberos specification. Current requirements imply that SID should be not smaller than 1 byte length and no longer than 28 bytes. SID length should also be a multiplicand of 4.
		AccessPermission AccessPermission `xml:"h:AccessPermission"`          // Indicates whether the User is allowed to access Intel® AMT from the Network or Local Interfaces. Note: this definition is restricted by the Default Interface Access Permissions of each Realm.
		Realms           []RealmValues    `xml:"h:Realms,omitempty"`          // Array of interface names the ACL entry is allowed to access.
	}
	UpdateUserAclEntryEx_INPUT struct {
		XMLName          xml.Name         `xml:"h:UpdateUserAclEntryEx_INPUT"`
		H                string           `xml:"xmlns:h,attr"`
		Handle           int              `xml:"h:Handle"`                    // Contains a creation handle.
		DigestUsername   string           `xml:"h:DigestUsername"`            // Username for access control. Contains 7-bit ASCII characters. String length is limited to 16 characters. Username cannot be an empty string.
		DigestPassword   string           `xml:"h:DigestPassword"`            // An MD5 Hash of these parameters concatenated together (Username + ":" + DigestRealm + ":" + Password). The DigestRealm is a field in AMT_GeneralSettings
		KerberosUserSid  string           `xml:"h:KerberosUserSid,omitempty"` // Descriptor for user (SID) which is authenticated using the Kerberos Authentication. Byte array, specifying the Security Identifier (SID) according to the Kerberos specification. Current requirements imply that SID should be not smaller than 1 byte length and no longer than 28 bytes. SID length should also be a multiplicand of 4.
		AccessPermission AccessPermission `xml:"h:AccessPermission"`          // Indicates whether the User is allowed to access Intel® AMT from the Network or Local Interfaces. Note: this definition is restricted by the Default Interface Access Permissions of each Realm.
		Realms           []RealmValues    `xml:"h:Realms,omitempty"`          // Array of interface names the ACL entry is allowed to access.
	}
)
