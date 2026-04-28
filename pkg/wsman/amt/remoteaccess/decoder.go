/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess

// INPUTS constants

const (
	AMTRemoteAccessCapabilities       string = "AMT_RemoteAccessCapabilities"
	AMTRemoteAccessPolicyAppliesToMPS string = "AMT_RemoteAccessPolicyAppliesToMPS"
	AMTRemoteAccessPolicyRule         string = "AMT_RemoteAccessPolicyRule"
	AMTRemoteAccessService            string = "AMT_RemoteAccessService"
	AddMps                            string = "AddMpServer"
	AddRemoteAccessPolicyRule         string = "AddRemoteAccessPolicyRule"
	ValueNotFound                     string = "Value not found in map"
)

const (
	IPv4Address MPServerInfoFormat = 3
	IPv6Address MPServerInfoFormat = 4
	FQDN        MPServerInfoFormat = 201
)

const (
	MutualAuthentication           MPServerAuthMethod = 1
	UsernamePasswordAuthentication MPServerAuthMethod = 2
)

const (
	UserInitiated Trigger = iota
	Alert
	Periodic
	HomeProvisioning
)

const (
	PolicyDecisionStrategyFirstMatching PolicyDecisionStrategy = 1
	PolicyDecisionStrategyAll           PolicyDecisionStrategy = 2
)

const (
	TriggerUserInitiated Trigger = iota
	TriggerAlert
	TriggerPeriodic
	TriggerHomeProvisioning
)

// triggerToString is a map of trigger values to their string representation.
var triggerToString = map[Trigger]string{
	TriggerUserInitiated:    "UserInitiated",
	TriggerAlert:            "Alert",
	TriggerPeriodic:         "Periodic",
	TriggerHomeProvisioning: "HomeProvisioning",
}

// String returns the string representation of the Trigger value.
func (t Trigger) String() string {
	if value, exists := triggerToString[t]; exists {
		return value
	}

	return ValueNotFound
}

const (
	ExternalMPS MPSType = iota
	InternalMPS
	BothMPS
)

// mpsTypeToString is a map of MPSType values to their string representation.
var mpsTypeToString = map[MPSType]string{
	ExternalMPS: "ExternalMPS",
	InternalMPS: "InternalMPS",
	BothMPS:     "BothMPS",
}

// String returns the string representation of the MPSType value.
func (m MPSType) String() string {
	if value, exists := mpsTypeToString[m]; exists {
		return value
	}

	return ValueNotFound
}

const (
	ReturnValueSuccess                 ReturnValue = 0
	ReturnValueInternalError           ReturnValue = 1
	ReturnValueNotPermitted            ReturnValue = 16
	ReturnValueMaxLimitReached         ReturnValue = 23
	ReturnValueInvalidParameter        ReturnValue = 36
	ReturnValueFlashWriteLimitExceeded ReturnValue = 38
	ReturnValueDuplicate               ReturnValue = 2058
)

// returnValueToString is a map of ReturnValue values to their string representation.
var returnValueToString = map[ReturnValue]string{
	ReturnValueSuccess:                 "Success",
	ReturnValueInternalError:           "InternalError",
	ReturnValueNotPermitted:            "NotPermitted",
	ReturnValueMaxLimitReached:         "MaxLimitReached",
	ReturnValueInvalidParameter:        "InvalidParameter",
	ReturnValueFlashWriteLimitExceeded: "FlashWriteLimitExceeded",
	ReturnValueDuplicate:               "Duplicate",
}

// String returns the string representation of the ReturnValue value.
func (r ReturnValue) String() string {
	if value, exists := returnValueToString[r]; exists {
		return value
	}

	return ValueNotFound
}
