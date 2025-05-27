/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package kerberos

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type SettingData struct {
	base message.Base
}

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
		XMLName                        xml.Name `xml:"Body"`
		GetResponse                    KerberosSettingDataResponse
		EnumerateResponse              common.EnumerateResponse
		PullResponse                   PullResponse
		GetCredentialCacheState_OUTPUT GetCredentialCacheState_OUTPUT
		SetCredentialCacheState_OUTPUT SetCredentialCacheState_OUTPUT
	}

	PullResponse struct {
		XMLName                  xml.Name                      `xml:"PullResponse"`
		KerberosSettingDataItems []KerberosSettingDataResponse `xml:"Items>AMT_KerberosSettingData"`
	}
	KerberosSettingDataResponse struct {
		XMLName                        xml.Name                         `xml:"AMT_KerberosSettingData"`
		ElementName                    string                           `xml:"ElementName,omitempty"`           // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID                     string                           `xml:"InstanceID,omitempty"`            // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm:	<OrgID>:<LocalID>	Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>.	<LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance.	For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		RealmName                      string                           `xml:"RealmName,omitempty"`             // Kerberos Realm name.
		ServicePrincipalName           []string                         `xml:"ServicePrincipalName,omitempty"`  // An array of strings, each of which names a distinct service principal.
		ServicePrincipalProtocol       []ServicePrincipalProtocol       `xml:"ServicePrincipalProtocol"`        // An array of 16-bit enumeration values, each of which corresponds to the string in the same position of ServicePrincipalName.
		KeyVersion                     int                              `xml:"KeyVersion,omitempty"`            // Key version number. User can update the value each time the master key is changed.
		EncryptionAlgorithm            EncryptionAlgorithm              `xml:"EncryptionAlgorithm,omitempty"`   // A 16-bit enumeration value that identifiers the encryption algorithm used in Kerberos authentication.
		MasterKey                      []int                            `xml:"MasterKey"`                       // A 128-bit binary key value. MasterKey cannot be used if the key generation method is used (using the Passphrase property)
		MaximumClockTolerance          int                              `xml:"MaximumClockTolerance,omitempty"` // The number of minutes by which the clocks of the Intel® AMT device and the client and KDC can be out of sync - typically 5 minutes.
		KrbEnabled                     bool                             `xml:"KrbEnabled"`                      // Indicates whether Kerberos authentication is enabled or disable.
		Passphrase                     string                           `xml:"Passphrase,omitempty"`            // Used when the key generation method is chosen (RFC 3961,3962). Salt and IterationCount must be supplied also.
		Salt                           string                           `xml:"Salt,omitempty"`                  // Used when the key generation method is chosen (RFC 3961,3962)
		IterationCount                 int                              `xml:"IterationCount,omitempty"`        // Can be used when the key generation method is chosen (RFC 3961,3962)
		SupportedEncryptionAlgorithms  []SupportedEncryptionAlgorithms  `xml:"SupportedEncryptionAlgorithms"`   // A 16-bit enumeration values that identifier the supported encryption algorithms used in Kerberos authentication.
		ConfiguredEncryptionAlgorithms []ConfiguredEncryptionAlgorithms `xml:"ConfiguredEncryptionAlgorithms"`  // A 16-bit enumeration values that identifier the configured encryption algorithms used in Kerberos authentication.
	}
	GetCredentialCacheState_OUTPUT struct {
		XMLName     xml.Name    `xml:"GetCredentialCacheState_OUTPUT"`
		Enabled     bool        `xml:"Enabled"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}

	SetCredentialCacheState_OUTPUT struct {
		XMLName     xml.Name    `xml:"SetCredentialCacheState_OUTPUT"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}
)

// INPUTS
// Request Types.
type (
	SetCredentialCacheStateInput struct {
		XMLName xml.Name `xml:"h:SetCredentialCacheState_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Enabled bool     `xml:"h:Enabled"`
	}
)

// An array of 16-bit enumeration values, each of which corresponds to the string in the same position of ServicePrincipalName. In Intel AMT Release 6.0 and later releases this field is not in use and has no impact
//
// ValueMap={0, 1, 2, 3}
//
// Values={HTTP Protocol definition, HTTPS Protocol definition, SOL&IDER protocol definition, SOL&IDER protocol definition (using SSL)}.
type ServicePrincipalProtocol int

// A 16-bit enumeration values that identifier the supported encryption algorithms used in Kerberos authentication. Note: While RC4-HMAC is supported, Intel recommends using AES256-CTS-HMAC-SHA1-96.
//
// ValueMap={0, 1, 2, ..}
//
// Values={RC4-HMAC, AES128-CTS-HMAC-SHA1-96, AES256-CTS-HMAC-SHA1-96, Reserved}.
type SupportedEncryptionAlgorithms int

// A 16-bit enumeration values that identifier the configured encryption algorithms used in Kerberos authentication. Note: While RC4-HMAC is supported, Intel recommends using AES256-CTS-HMAC-SHA1-96. Note: Intel AMT does not choose the encryption algorithm to configure based on the values specified by the user. Intel AMT attempts to enable RC4. If a Passphrase and Salt are provided, the AES suites are also configured.
//
// ValueMap={0, 1, 2, ..}
//
// Values={RC4-HMAC, AES128-CTS-HMAC-SHA1-96, AES256-CTS-HMAC-SHA1-96, Reserved}.
type ConfiguredEncryptionAlgorithms int

// A 16-bit enumeration value that identifiers the encryption algorithm used in Kerberos authentication.
//
// ValueMap={0}
//
// Values={RC4 encryption and HMAC authentication}.
type EncryptionAlgorithm int

// ReturnValue is a 16-bit enumeration value that indicates the success or failure of an operation.
type ReturnValue int
