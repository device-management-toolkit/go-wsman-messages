/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
)

const AMT_RemoteAccessPolicyRule = "AMT_RemoteAccessPolicyRule"

type RemoteAccessPolicyRule struct {
	Trigger        Trigger
	TunnelLifeTime int
	ExtendedData   string
}

type Trigger uint8

const (
	UserInitiated Trigger = iota
	Alert
	Periodic
	HomeProvisioning
)

type PolicyRule struct {
	base wsman.Base
}

func NewRemoteAccessPolicyRule(wsmanMessageCreator *wsman.WSManMessageCreator) PolicyRule {
	return PolicyRule{
		base: wsman.NewBase(wsmanMessageCreator, AMT_RemoteAccessPolicyRule),
	}
}

// Get retrieves the representation of the instance
func (RemoteAccessPolicyRule PolicyRule) Get() string {
	return RemoteAccessPolicyRule.base.Get(nil)
}
func (RemoteAccessPolicyRule PolicyRule) Enumerate() string {
	return RemoteAccessPolicyRule.base.Enumerate()
}
func (RemoteAccessPolicyRule PolicyRule) Pull(enumerationContext string) string {
	return RemoteAccessPolicyRule.base.Pull(enumerationContext)
}
func (RemoteAccessPolicyRule PolicyRule) Put(remoteAccessPolicyRule RemoteAccessPolicyRule) string {
	return RemoteAccessPolicyRule.base.Put(remoteAccessPolicyRule, false, nil)
}
func (RemoteAccessPolicyRule PolicyRule) Delete(handle string) string {
	selector := wsman.Selector{Name: "PolicyRuleName", Value: handle}
	return RemoteAccessPolicyRule.base.Delete(selector)
}
