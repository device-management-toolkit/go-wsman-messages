/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package environmentdetection

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
		XMLName           xml.Name `xml:"Body"`
		GetAndPutResponse EnvironmentDetectionSettingDataResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}
	EnvironmentDetectionSettingDataResponse struct {
		XMLName                    xml.Name           `xml:"AMT_EnvironmentDetectionSettingData"`
		ElementName                string             `xml:"ElementName,omitempty"`                // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID                 string             `xml:"InstanceID,omitempty"`                 // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>. <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance. For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
		DetectionAlgorithm         DetectionAlgorithm `xml:"DetectionAlgorithm,omitempty"`         // Specifies which algorithm is used to determine whether the system is in its intranet environment or in the Internet environment. Currently, AMT supports only the "Local Domains" type.
		DetectionStrings           []string           `xml:"DetectionStrings,omitempty"`           // An array of strings used in the environment detection algorithm. If DetectionAlgorithm is "Local Domains", DetectionStrings contains a set of local domain strings. If DetectionAlgorithm is "Remote URLs", then DetectionStrings contains a set of remote URLs.
		DetectionIPv6LocalPrefixes []string           `xml:"DetectionIPv6LocalPrefixes,omitempty"` // Can be used for environment detection in IPv6 networks that do not configure the DNS suffix via DHCP. The format is: "XXXX:XXXX:XXXX:XXXX/Y" where Y is the prefix length, the XXXX:XXXX:XXXX:XXXX part can include zeros compression (e.g. 0:0:0 or ::) formats.
	}
	PullResponse struct {
		XMLName                              xml.Name                                  `xml:"PullResponse"`
		EnvironmentDetectionSettingDataItems []EnvironmentDetectionSettingDataResponse `xml:"Items>AMT_EnvironmentDetectionSettingData"`
	}
)

// INPUTS
// Request Types.
type EnvironmentDetectionSettingDataRequest struct {
	XMLName                    xml.Name           `xml:"h:AMT_EnvironmentDetectionSettingData"`
	H                          string             `xml:"xmlns:h,attr"`
	ElementName                string             `xml:"h:ElementName"`                          // Required. The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
	InstanceID                 string             `xml:"h:InstanceID"`                           // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. To ensure uniqueness within the NameSpace, the value of InstanceID should be constructed using the following "preferred" algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon (:), and where <OrgID> must include a copyrighted, trademarked, or otherwise unique name that is owned by the business entity that is creating or defining the InstanceID or that is a registered ID assigned to the business entity by a recognized global authority. (This requirement is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> must not contain a colon (:). When using this algorithm, the first colon to appear in InstanceID must appear between <OrgID> and <LocalID>. <LocalID> is chosen by the business entity and should not be reused to identify different underlying (real-world) elements. If the above "preferred" algorithm is not used, the defining entity must assure that the resulting InstanceID is not reused across any InstanceIDs produced by this or other providers for the NameSpace of this instance. For DMTF-defined instances, the "preferred" algorithm must be used with the <OrgID> set to CIM.
	DetectionAlgorithm         DetectionAlgorithm `xml:"h:DetectionAlgorithm"`                   // Required. Specifies which algorithm is used to determine whether the system is in its intranet environment or in the Internet environment. Currently, AMT supports only the "Local Domains" type.
	DetectionStrings           []string           `xml:"h:DetectionStrings,omitempty"`           // An array of strings used in the environment detection algorithm. If DetectionAlgorithm is "Local Domains", DetectionStrings contains a set of local domain strings. If DetectionAlgorithm is "Remote URLs", then DetectionStrings contains a set of remote URLs.
	DetectionIPv6LocalPrefixes []string           `xml:"h:DetectionIPv6LocalPrefixes,omitempty"` // Can be used for environment detection in IPv6 networks that do not configure the DNS suffix via DHCP. The format is: "XXXX:XXXX:XXXX:XXXX/Y" where Y is the prefix length, the XXXX:XXXX:XXXX:XXXX part can include zeros compression (e.g. 0:0:0 or ::) formats.
}

// Specifies which algorithm is used to determine whether the system is in its intranet environment or in the Internet environment.
//
// Currently, AMT supports only the "Local Domains" type.
//
// ValueMap={0, 1}
//
// Values={Local Domains, Remote URLs}.
type DetectionAlgorithm int
