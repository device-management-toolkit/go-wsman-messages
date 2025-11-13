/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ieee8021x

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
		XMLName                      xml.Name `xml:"Body"`
		ProfileGetAndPutResponse     ProfileResponse
		CredentialContextGetResponse CredentialContextResponse
		EnumerateResponse            common.EnumerateResponse
		PullResponse                 PullResponse
	}

	PullResponse struct {
		XMLName                xml.Name                    `xml:"PullResponse"`
		ProfileItems           []ProfileResponse           `xml:"Items>AMT_8021XProfile"`
		CredentialContextItems []CredentialContextResponse `xml:"Items>AMT_8021xCredentialContext"`
	}

	ProfileResponse struct {
		XMLName                         xml.Name                        `xml:"AMT_8021XProfile"`
		ElementName                     string                          `xml:"ElementName,omitempty"`                     // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID                      string                          `xml:"InstanceID,omitempty"`                      // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>. <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance. For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		Enabled                         bool                            `xml:"Enabled,omitempty"`                         // Indicates whether the 802.1x profile is enabled.
		ActiveInS0                      bool                            `xml:"ActiveInS0,omitempty"`                      // Indicates the activity setting of the 802.1X module in H0 state when the LAN driver is active. The default value for this property is 'true'. If the LAN driver is down, this property is not relevant. Functionality: when FALSE, AMT is not accessible (over 802.1x enabled port) in case the host is in S0 but fails to authenticate to the server. When TRUE, AMT handles the authentication in this case (but the host still can't be accessed until it authenticates successfully). If 802.1X is not configured, this API may still succeed as the setting may be stored for future use. The default factory setting is TRUE.
		AuthenticationProtocol          AuthenticationProtocol          `xml:"AuthenticationProtocol,omitempty"`          // Identifies the authentication protocol used to authenticate the access requestor to the AAA server.
		RoamingIdentity                 string                          `xml:"RoamingIdentity,omitempty"`                 // A string presented to the authentication server in 802.1x protocol exchange. The AAA server determines the format of this string. Formats supported by AAA servers include: username@domain.
		ServerCertificateName           string                          `xml:"ServerCertificateName,omitempty"`           // The name compared against the subject name field in the certificate provided by the AAA server. This name is either the full name of the AAA server, in which case ServerCertificateNameComparison is set to "FullName", or it is the domain suffix of the AAA server, in which case ServerCertificateNameComparison is set to "DomainSuffix"
		ServerCertificateNameComparison ServerCertificateNameComparison `xml:"ServerCertificateNameComparison,omitempty"` // Determines the comparison algorithm used between the ServerCertificateName value and the subject name field of the certificate presented by the AAA server. This field is mandatory if ServerCertificateName is defined.
		Username                        string                          `xml:"Username,omitempty"`                        // Within the domain specified by Domain, Identifies the user that is requesting access to the network. MaxLen=128
		Password                        string                          `xml:"Password,omitempty"`                        // The password associated with the user identified by Username and Domain. MaxLen=32
		Domain                          string                          `xml:"Domain,omitempty"`                          // The domain within which Username is unique. The Domain string shouldn't contain the suffix, so the user name (Domain\user) will be correct. If the Domain string contains a suffix (e.g. Domain = intel.com), the user trying to authenticate will be of the form intel.com\user (instead of intel\user) and thus authentication will fail. MaxLen=128
		ProtectedAccessCredential       []int                           `xml:"ProtectedAccessCredential,omitempty"`       // A credential used by the supplicant and AAA server to establish a mutually authenticated encrypted tunnel for confidential user authentication. This field is relevant for EAP-FAST only. It is not required if the server is configured for "PAC provisioning".
		PACPassword                     string                          `xml:"PACPassword,omitempty"`                     // Optional password to extract the PAC (Protected Access Credential)information from the PAC data. MaxLen=256
		ClientCertificate               string                          `xml:"ClientCertificate,omitempty"`               // The client certificate that should be used by the profile. The client certificate should be specified in the Put request while configuring the profile. This will delete the existing instance of AMT_8021xCredentialContext that represents the client certificate, and create a new instance if a client certificate EPR is provided. This property will never be returned in Get response.
		ServerCertificateIssue          string                          `xml:"ServerCertificateIssue,omitempty"`          // The trusted root CA that should be used while verifying the server certificate. The root certificate should be specified in the Put request while configuring the profile. This will delete the existing instance of AMT_8021xCredentialContext that represents the trusted root certificate, and create a new instance if a trusted root certificate EPR is provided. This property will never be returned in Get response. This field is optional. If not defined, AMT looks for a matching root certidicate in its repository.
		PxeTimeout                      int                             `xml:"PxeTimeout,omitempty"`                      // Timeout in seconds, in which the Intel® AMT will hold an authenticated 802.1X session. During the defined period, Intel® AMT manages the 802.1X negotiation while a PXE boot takes place. After the timeout, control of the negotiation passes to the host. The maximum value is 86400 seconds (one day). A value of 0 disables the feature. If you do not set a profile, the value of PxeTimeout is 0. If you set a profile without specifying a value for PxeTimeout, the firmware sets it to 120.
	}
	SelectorResponse struct {
		XMLName xml.Name `xml:"Selector,omitempty"`
		Name    string   `xml:"Name,attr"`
		Text    string   `xml:",chardata"`
	}
	SelectorSetResponse struct {
		XMLName   xml.Name           `xml:"SelectorSet,omitempty"`
		Selectors []SelectorResponse `xml:"Selector,omitempty"`
	}
	ReferenceParametersResponse struct {
		XMLName     xml.Name            `xml:"ReferenceParameters,omitempty"`
		ResourceURI string              `xml:"ResourceURI,omitempty"`
		SelectorSet SelectorSetResponse `xml:"SelectorSet,omitempty"`
	}
	ElementInContextResponse struct {
		XMLName             xml.Name                    `xml:"ElementInContext"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	ElementProvidingContextResponse struct {
		XMLName             xml.Name                    `xml:"ElementProvidingContext"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	CredentialContextResponse struct {
		XMLName                 xml.Name                        `xml:"AMT_8021xCredentialContext"`
		ElementInContext        ElementInContextResponse        `xml:"ElementInContext"`
		ElementProvidingContext ElementProvidingContextResponse `xml:"ElementProvidingContext"`
	}
)

type ProfileRequest struct {
	XMLName                         xml.Name                        `xml:"h:AMT_8021XProfile"`
	H                               string                          `xml:"xmlns:h,attr"`
	ElementName                     string                          `xml:"h:ElementName"`                         // Required. The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
	InstanceID                      string                          `xml:"h:InstanceID,omitempty"`                // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>. <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance. For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
	Enabled                         bool                            `xml:"h:Enabled"`                             // Required. Indicates whether the 802.1x profile is enabled.
	ActiveInS0                      bool                            `xml:"h:ActiveInS0"`                          // Indicates the activity setting of the 802.1X module in H0 state when the LAN driver is active. The default value for this property is 'true'. If the LAN driver is down, this property is not relevant. Functionality: when FALSE, AMT is not accessible (over 802.1x enabled port) in case the host is in S0 but fails to authenticate to the server. When TRUE, AMT handles the authentication in this case (but the host still can't be accessed until it authenticates successfully). If 802.1X is not configured, this API may still succeed as the setting may be stored for future use. The default factory setting is TRUE.
	AuthenticationProtocol          AuthenticationProtocol          `xml:"h:AuthenticationProtocol"`              // Identifies the authentication protocol used to authenticate the access requestor to the AAA server.
	RoamingIdentity                 string                          `xml:"h:RoamingIdentity,omitempty"`           // A string presented to the authentication server in 802.1x protocol exchange. The AAA server determines the format of this string. Formats supported by AAA servers include: username@domain.
	ServerCertificateName           string                          `xml:"h:ServerCertificateName,omitempty"`     // The name compared against the subject name field in the certificate provided by the AAA server. This name is either the full name of the AAA server, in which case ServerCertificateNameComparison is set to "FullName", or it is the domain suffix of the AAA server, in which case ServerCertificateNameComparison is set to "DomainSuffix"
	ServerCertificateNameComparison ServerCertificateNameComparison `xml:"h:ServerCertificateNameComparison"`     // Determines the comparison algorithm used between the ServerCertificateName value and the subject name field of the certificate presented by the AAA server. This field is mandatory if ServerCertificateName is defined.
	Username                        string                          `xml:"h:Username,omitempty"`                  // Within the domain specified by Domain, Identifies the user that is requesting access to the network. MaxLen=128
	Password                        string                          `xml:"h:Password,omitempty"`                  // The password associated with the user identified by Username and Domain. MaxLen=32
	Domain                          string                          `xml:"h:Domain,omitempty"`                    // The domain within which Username is unique. The Domain string shouldn't contain the suffix, so the user name (Domain\user) will be correct. If the Domain string contains a suffix (e.g. Domain = intel.com), the user trying to authenticate will be of the form intel.com\user (instead of intel\user) and thus authentication will fail. MaxLen=128
	ProtectedAccessCredential       []int                           `xml:"h:ProtectedAccessCredential,omitempty"` // A credential used by the supplicant and AAA server to establish a mutually authenticated encrypted tunnel for confidential user authentication. This field is relevant for EAP-FAST only. It is not required if the server is configured for "PAC provisioning".
	PACPassword                     string                          `xml:"h:PACPassword,omitempty"`               // Optional password to extract the PAC (Protected Access Credential)information from the PAC data. MaxLen=256
	ClientCertificate               string                          `xml:"h:ClientCertificate,omitempty"`         // The client certificate that should be used by the profile. The client certificate should be specified in the Put request while configuring the profile. This will delete the existing instance of AMT_8021xCredentialContext that represents the client certificate, and create a new instance if a client certificate EPR is provided. This property will never be returned in Get response.
	ServerCertificateIssue          string                          `xml:"h:ServerCertificateIssue,omitempty"`    // The trusted root CA that should be used while verifying the server certificate. The root certificate should be specified in the Put request while configuring the profile. This will delete the existing instance of AMT_8021xCredentialContext that represents the trusted root certificate, and create a new instance if a trusted root certificate EPR is provided. This property will never be returned in Get response. This field is optional. If not defined, AMT looks for a matching root certidicate in its repository.
	PxeTimeout                      int                             `xml:"h:PxeTimeout,omitempty"`                // Timeout in seconds, in which the Intel® AMT will hold an authenticated 802.1X session. During the defined period, Intel® AMT manages the 802.1X negotiation while a PXE boot takes place. After the timeout, control of the negotiation passes to the host. The maximum value is 86400 seconds (one day). A value of 0 disables the feature. If you do not set a profile, the value of PxeTimeout is 0. If you set a profile without specifying a value for PxeTimeout, the firmware sets it to 120.
}

// Identifies the authentication protocol used to authenticate the access requestor to the AAA server.
//
// ValueMap={0, 1, 2, 3, 4, 5, 6}
//
// Values={TLS, TTLS_MSCHAPv2, PEAP_MSCHAPv2, EAP_GTC, EAPFAST_MSCHAPv2, EAPFAST_GTC, EAPFAST_TLS}.
type AuthenticationProtocol int

// Determines the comparison algorithm used between the ServerCertificateName value and the subject name field of the certificate presented by the AAA server. This field is mandatory if ServerCertificateName is defined.
//
// ValueMap={0, 1}
//
// Values={FullName, DomainSuffix}.
type ServerCertificateNameComparison int
