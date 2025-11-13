/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"
	"strings"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type (
	SettingData struct {
		base message.Base
	}
	CredentialContext struct {
		base message.Base
	}
)

// OUTPUT
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                               xml.Name                           `xml:"Body"`
		SettingDataGetAndPutResponse          SettingDataResponse                `xml:"AMT_TLSSettingData"`
		CredentialContextGetResponse          CredentialContextResponse          `xml:"AMT_TLSCredentialContext"`
		CredentialContextCreateResponse       CredentialContextCreateResponse    `xml:"ResourceCreated"`
		ProtocolEndpointCollectionGetResponse ProtocolEndpointCollectionResponse `xml:"AMT_TLSProtocolEndpointCollection"`
		EnumerateResponse                     common.EnumerateResponse
		PullResponse                          PullResponse
	}
	SettingDataResponse struct {
		XMLName                       xml.Name `xml:"AMT_TLSSettingData"`
		ElementName                   string   `xml:"ElementName,omitempty"`         // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID                    string   `xml:"InstanceID,omitempty"`          // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID>	Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>.	<LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance.	For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		MutualAuthentication          bool     `xml:"MutualAuthentication"`          // Administrator-settable property that determines whether or not mutual authentication is used at the TLS layer is used on the associated service access point. If False, then only the server authenticates itself at the TLS layer. Use of Mutual Authentication on the local interface is deprecated in Release 6.0. The feature will be removed in a future release. This property is only visible / usable for users of ADMIN_SECURITY_ADMINISTRATION realm. This property must be supplied if Enabled property is True.
		Enabled                       bool     `xml:"Enabled"`                       // Administrator-settable property that determines whether or not TLS is used on the associated service access point.
		TrustedCN                     []string `xml:"TrustedCN,omitempty"`           // An array of strings, used to validate the CN subfield of the subject field in X.509 certificates presented to Intel® AMT in the TLS handshake. This value must comply with the requirements of RFC 1035.
		AcceptNonSecureConnections    bool     `xml:"AcceptNonSecureConnections"`    // This setting defines once TLS is enabled and configured whether non-secure EOI/WSMAN connections are still accepted by FW on ports 16992 and 623. If AcceptNonSecureConnections is set to TRUE then non-secure connections are still accepted. If set to FALSE then non-secure connections are rejected. This setting may be set per interface for the local and network interfaces. AMT_TLSSettingData.AcceptNonSecureConnections may only be modified for the remote interface. It is a read-only property for the local interface instance.
		NonSecureConnectionsSupported *bool    `xml:"NonSecureConnectionsSupported"` // Indicates the removal of support for the non-TLS WS-MAN ports for the remote interface. Available starting Intel CSME 16.1 firmware on Raptor Lake platforms. If this read-only field exists and its value is True, changing the value of the AcceptNonSecureConnections field is allowed only for the local interface. Note that this class and field can be accessed locally as well as remotely. Invoking the AMT_TLSSettingData.Put() command on the remote instance with AcceptNonSecureConnections set to True will fail with error code AMT_STATUS_NOT_PERMITTED. Setting AMT_TLSSettingData.Enabled to False will also fail for the remote interface.
	}
	SelectorResponse struct {
		XMLName           xml.Name                  `xml:"Selector,omitempty"`
		Name              string                    `xml:"Name,attr,omitempty"`
		Text              string                    `xml:",chardata"`
		EndpointReference EndpointReferenceResponse `xml:"EndpointReference,omitempty"`
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
		XMLName                 xml.Name                        `xml:"AMT_TLSCredentialContext"`
		ElementInContext        ElementInContextResponse        `xml:"ElementInContext"`
		ElementProvidingContext ElementProvidingContextResponse `xml:"ElementProvidingContext"`
	}
	ProtocolEndpointCollectionResponse struct {
		XMLName     xml.Name `xml:"AMT_TLSProtocolEndpointCollection"`
		ElementName string   `xml:"ElementName"`
	}
	PullResponse struct {
		XMLName                         xml.Name                             `xml:"PullResponse"`
		SettingDataItems                []SettingDataResponse                `xml:"Items>AMT_TLSSettingData"`
		ProtocolEndpointCollectionItems []ProtocolEndpointCollectionResponse `xml:"Items>AMT_TLSProtocolEndpointCollection"`
		CredentialContextItems          []CredentialContextResponse          `xml:"Items>AMT_TLSCredentialContext"`
	}
	CredentialContextCreateResponse struct {
		XMLName             xml.Name                    `xml:"ResourceCreated"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
	EndpointReferenceResponse struct {
		XMLName             xml.Name                    `xml:"EndpointReference"`
		Address             string                      `xml:"Address,omitempty"`
		ReferenceParameters ReferenceParametersResponse `xml:"ReferenceParameters,omitempty"`
	}
)

// UnmarshalXML trims insignificant whitespace in SelectorResponse.Text while decoding.
func (s *SelectorResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias SelectorResponse

	var aux Alias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	aux.Text = strings.TrimSpace(aux.Text)
	*s = SelectorResponse(aux)

	return nil
}

type (
	SettingDataRequest struct {
		XMLName                       xml.Name `xml:"h:AMT_TLSSettingData"`
		H                             string   `xml:"xmlns:h,attr"`
		ElementName                   string   `xml:"h:ElementName,omitempty"`         // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID                    string   `xml:"h:InstanceID,omitempty"`          // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID>	Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>.	<LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance.	For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		MutualAuthentication          bool     `xml:"h:MutualAuthentication"`          // Administrator-settable property that determines whether or not mutual authentication is used at the TLS layer is used on the associated service access point. If False, then only the server authenticates itself at the TLS layer. Use of Mutual Authentication on the local interface is deprecated in Release 6.0. The feature will be removed in a future release. This property is only visible / usable for users of ADMIN_SECURITY_ADMINISTRATION realm. This property must be supplied if Enabled property is True.
		Enabled                       bool     `xml:"h:Enabled"`                       // Administrator-settable property that determines whether or not TLS is used on the associated service access point.
		TrustedCN                     []string `xml:"h:TrustedCN,omitempty"`           // An array of strings, used to validate the CN subfield of the subject field in X.509 certificates presented to Intel® AMT in the TLS handshake. This value must comply with the requirements of RFC 1035.
		AcceptNonSecureConnections    bool     `xml:"h:AcceptNonSecureConnections"`    // This setting defines once TLS is enabled and configured whether non-secure EOI/WSMAN connections are still accepted by FW on ports 16992 and 623. If AcceptNonSecureConnections is set to TRUE then non-secure connections are still accepted. If set to FALSE then non-secure connections are rejected. This setting may be set per interface for the local and network interfaces. AMT_TLSSettingData.AcceptNonSecureConnections may only be modified for the remote interface. It is a read-only property for the local interface instance.
		NonSecureConnectionsSupported bool     `xml:"h:NonSecureConnectionsSupported"` // Indicates the removal of support for the non-TLS WS-MAN ports for the remote interface. Available starting Intel CSME 16.1 firmware on Raptor Lake platforms. If this read-only field exists and its value is True, changing the value of the AcceptNonSecureConnections field is allowed only for the local interface. Note that this class and field can be accessed locally as well as remotely. Invoking the AMT_TLSSettingData.Put() command on the remote instance with AcceptNonSecureConnections set to True will fail with error code AMT_STATUS_NOT_PERMITTED. Setting AMT_TLSSettingData.Enabled to False will also fail for the remote interface.
	}
)
