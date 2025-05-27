/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// NewPolicyRuleWithClient instantiates a new PolicyRule.
func NewPolicyRuleWithClient(wsmanMessageCreator *message.WSManMessageCreator, clientPolicy client.WSMan) PolicyRule {
	return PolicyRule{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTRemoteAccessPolicyRule, clientPolicy),
	}
}

// Get retrieves the representation of the instance.
func (policyRule PolicyRule) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Get(nil),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (policyRule PolicyRule) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (policyRule PolicyRule) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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

// Put will change properties of the selected instance.
func (policyRule PolicyRule) Put(remoteAccessPolicyRule RemoteAccessPolicyRuleRequest) (response Response, err error) {
	remoteAccessPolicyRule.H = fmt.Sprintf("%s%s", message.AMTSchema, AMTRemoteAccessPolicyRule)
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Put(remoteAccessPolicyRule, false, nil),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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

// Delete removes a the specified instance.
func (policyRule PolicyRule) Delete(handle string) (response Response, err error) {
	selector := message.Selector{Name: "PolicyRuleName", Value: handle}
	response = Response{
		Message: &client.Message{
			XMLInput: policyRule.base.Delete(selector),
		},
	}
	// send the message to AMT
	err = policyRule.base.Execute(response.Message)
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
