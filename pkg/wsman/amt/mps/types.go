/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mps

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type UsernamePassword struct {
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
		GetResponse       MPSUsernamePasswordResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
	}

	MPSUsernamePasswordResponse struct {
		XMLName    xml.Name `xml:"AMT_MPSUsernamePassword"`
		InstanceID string   `xml:"InstanceID,omitempty"` // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. In order to ensure uniqueness within the NameSpace, the value of InstanceID SHOULD be constructed using the following 'preferred' algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon ':', and where <OrgID> MUST include a copyrighted, trademarked or otherwise unique name that is owned by the business entity creating/defining the InstanceID, or is a registered ID that is assigned to the business entity by a recognized global authority. (This is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> MUST NOT contain a colon (':'). When using this algorithm, the first colon to appear in InstanceID MUST appear between <OrgID> and <LocalID>. <LocalID> is chosen by the organizational entity and SHOULD not be re-used to identify different underlying (real-world) elements. If the above 'preferred' algorithm is not used, the defining entity MUST assure that the resultant InstanceID is not re-used across any InstanceIDs produced by this or other providers for this instance's NameSpace. For DMTF defined instances, the 'preferred' algorithm MUST be used with the <OrgID> set to 'CIM'.	RemoteID   string   `xml:"RemoteID,omitempty"` //
		RemoteID   string   `xml:"RemoteID,omitempty"`   // RemoteID is the name by which the principal is known at the remote secret key authentication service.
		Secret     string   `xml:"Secret,omitempty"`     // The secret known by the principal. This property is write-only.
		Algorithm  string   `xml:"Algorithm,omitempty"`  // No AMT SDK Documentation
		Protocol   string   `xml:"Protocol,omitempty"`   // No AMT SDK Documentation
	}

	PullResponse struct {
		XMLName                  xml.Name                      `xml:"PullResponse"`
		MPSUsernamePasswordItems []MPSUsernamePasswordResponse `xml:"Items>AMT_MPSUsernamePassword"`
	}
)

// INPUTS
// Request Types.
type (
	MPSUsernamePasswordRequest struct {
		XMLName    xml.Name `xml:"h:AMT_MPSUsernamePassword"`
		H          string   `xml:"xmlns:h,attr"`
		InstanceID string   `xml:"h:InstanceID,omitempty"` // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. In order to ensure uniqueness within the NameSpace, the value of InstanceID SHOULD be constructed using the following 'preferred' algorithm: <OrgID>:<LocalID> Where <OrgID> and <LocalID> are separated by a colon ':', and where <OrgID> MUST include a copyrighted, trademarked or otherwise unique name that is owned by the business entity creating/defining the InstanceID, or is a registered ID that is assigned to the business entity by a recognized global authority. (This is similar to the <Schema Name>_<Class Name> structure of Schema class names.) In addition, to ensure uniqueness, <OrgID> MUST NOT contain a colon (':'). When using this algorithm, the first colon to appear in InstanceID MUST appear between <OrgID> and <LocalID>. <LocalID> is chosen by the organizational entity and SHOULD not be re-used to identify different underlying (real-world) elements. If the above 'preferred' algorithm is not used, the defining entity MUST assure that the resultant InstanceID is not re-used across any InstanceIDs produced by this or other providers for this instance's NameSpace. For DMTF defined instances, the 'preferred' algorithm MUST be used with the <OrgID> set to 'CIM'.	RemoteID   string   `xml:"RemoteID,omitempty"` //
		RemoteID   string   `xml:"h:RemoteID,omitempty"`   // RemoteID is the name by which the principal is known at the remote secret key authentication service.
		Secret     string   `xml:"h:Secret,omitempty"`     // The secret known by the principal. This property is write-only.
		Algorithm  string   `xml:"h:Algorithm,omitempty"`  // No AMT SDK Documentation
		Protocol   string   `xml:"h:Protocol,omitempty"`   // No AMT SDK Documentation
	}
)
