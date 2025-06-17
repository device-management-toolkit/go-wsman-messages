/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type PolicyRule struct {
	base.WSManService[Response]
}

// NewPolicyRuleWithClient instantiates a new PolicyRule.
func NewPolicyRuleWithClient(wsmanMessageCreator *message.WSManMessageCreator, clientPolicy client.WSMan) PolicyRule {
	return PolicyRule{
		base.NewService[Response](wsmanMessageCreator, AMTRemoteAccessPolicyRule, clientPolicy),
	}
}

// Delete removes a the specified instance.
func (policyRule PolicyRule) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "PolicyRuleName", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.Base.Delete(selector),
		},
	}
	// send the message to AMT
	err = policyRule.Base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}
