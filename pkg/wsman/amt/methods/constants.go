/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package methods

type Methods string

const (
	Get                       Methods = "Get"
	Pull                      Methods = "Pull"
	Enumerate                 Methods = "Enumerate"
	Put                       Methods = "Put"
	Delete                    Methods = "Delete"
	ReadRecords               Methods = "ReadRecords"
	AddTrustedRootCertificate Methods = "AddTrustedRootCertificate"
	AddCertificate            Methods = "AddCertificate"
	AddMps                    Methods = "AddMpServer"
	AddRemoteAccessPolicyRule Methods = "AddRemoteAccessPolicyRule"
	Create                    Methods = "Create"
	RequestStateChange        Methods = "RequestStateChange"
	SetBootConfigRole         Methods = "SetBootConfigRole"
	GetRecords                Methods = "GetRecords"
	PositionToFirstRecord     Methods = "PositionToFirstRecord"
	CommitChanges             Methods = "commitChanges"
	Unprovision               Methods = "Unprovision"
	SetMEBxPassword           Methods = "SetMEBxPassword"
	SetAdminAclEntryEx        Methods = "SetAdminAclEntryEx"
	GetLowAccuracyTimeSynch   Methods = "GetLowAccuracyTimeSynch"
	SetHighAccuracyTimeSynch  Methods = "SetHighAccuracyTimeSynch"
	GenerateKeyPair           Methods = "GenerateKeyPair"
	AddWiFiSettings           Methods = "AddWiFiSettings"
	AddAlarm                  Methods = "AddAlarm"
	GeneratePKCS10RequestEx   Methods = "GeneratePKCS10RequestEx"
	GetUuid                   Methods = "GetUuid"
	AddAdminAclEntryEx        Methods = "AddUserAclEntryEx"
	EnumerateUserAclEntries   Methods = "EnumerateUserAclEntries"
	GetAclEnabledState        Methods = "GetAclEnabledState"
	GetAdminAclEntry          Methods = "GetAdminAclEntry"
	GetAdminAclEntryStatus    Methods = "GetAdminAclEntryStatus"
	GetAdminNetAclEntryStatus Methods = "GetAdminNetAclEntryStatus"
	GetUserAclEntryEx         Methods = "GetUserAclEntryEx"
	RemoveUserAclEntry        Methods = "RemoveUserAclEntry"
	SetAclEnabledState        Methods = "SetAclEnabledState"
	UpdateUserAclEntryEx      Methods = "UpdateUserAclEntryEx"
	AddUserAclEntryEx         Methods = "AddUserAclEntryEx"
)
