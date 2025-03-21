/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

const (
	AMTSetupAndConfigurationService string = "AMT_SetupAndConfigurationService"
	CommitChanges                   string = "CommitChanges"
	Unprovision                     string = "Unprovision"
	PartialUnprovision              string = "PartialUnprovision"
	SetMEBxPassword                 string = "SetMEBxPassword"
	GetUUID                         string = "GetUuid"
	ValueNotFound                   string = "Value not found in map"
)

const (
	RequestedStateEnabled RequestedState = iota + 2
	RequestedStateDisabled
	RequestedStateShutDown
	RequestedStateNoChange
	RequestedStateOffline
	RequestedStateTest
	RequestedStateDeferred
	RequestedStateQuiesce
	RequestedStateReboot
	RequestedStateReset
	RequestedStateNotApplicable
	RequestedStateUnknown RequestedState = 0
)

// requestedStateToString is a map of RequestedState values to their string representations.
var requestedStateToString = map[RequestedState]string{
	RequestedStateEnabled:       "RequestedStateEnabled",
	RequestedStateDisabled:      "RequestedStateDisabled",
	RequestedStateShutDown:      "RequestedStateShutDown",
	RequestedStateNoChange:      "RequestedStateNoChange",
	RequestedStateOffline:       "RequestedStateOffline",
	RequestedStateTest:          "RequestedStateTest",
	RequestedStateDeferred:      "RequestedStateDeferred",
	RequestedStateQuiesce:       "RequestedStateQuiesce",
	RequestedStateReboot:        "RequestedStateReboot",
	RequestedStateReset:         "RequestedStateReset",
	RequestedStateNotApplicable: "RequestedStateNotApplicable",
	RequestedStateUnknown:       "RequestedStateUnknown",
}

// String returns the string representation of the RequestedState value.
func (r RequestedState) String() string {
	if value, exists := requestedStateToString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

// enabledStateToString is a map of EnabledState values to their string representations.
var enabledStateToString = map[EnabledState]string{
	EnabledStateUnknown:           "EnabledStateUnknown",
	EnabledStateOther:             "EnabledStateOther",
	EnabledStateEnabled:           "EnabledStateEnabled",
	EnabledStateDisabled:          "EnabledStateDisabled",
	EnabledStateShuttingDown:      "EnabledStateShuttingDown",
	EnabledStateNotApplicable:     "EnabledStateNotApplicable",
	EnabledStateEnabledButOffline: "EnabledStateEnabledButOffline",
	EnabledStateInTest:            "EnabledStateInTest",
	EnabledStateDeferred:          "EnabledStateDeferred",
	EnabledStateQuiesce:           "EnabledStateQuiesce",
	EnabledStateStarting:          "EnabledStateStarting",
}

// String returns the string representation of the EnabledState value.
func (e EnabledState) String() string {
	if value, exists := enabledStateToString[e]; exists {
		return value
	}

	return ValueNotFound
}

const (
	AdminControlMode  ProvisioningModeValue  = 1
	ClientControlMode ProvisioningModeValue  = 4
	PreProvisioning   ProvisioningStateValue = 0
	InProvisioning    ProvisioningStateValue = 1
	PostProvisioning  ProvisioningStateValue = 2
)

// provisioningModeToString is a map of ProvisioningMode values to their string representations.
var provisioningModeToString = map[ProvisioningModeValue]string{
	AdminControlMode:  "AdminControlMode",
	ClientControlMode: "ClientControlMode",
}

// String returns the string representation of the ProvisioningMode value.
func (p ProvisioningModeValue) String() string {
	if value, exists := provisioningModeToString[p]; exists {
		return value
	}

	return ValueNotFound
}

// provisioningStateToString is a map of ProvisioningState values to their string representations.
var provisioningStateToString = map[ProvisioningStateValue]string{
	PreProvisioning:  "PreProvisioning",
	InProvisioning:   "InProvisioning",
	PostProvisioning: "PostProvisioning",
}

// String returns the string representation of the ProvisioningState value.
func (p ProvisioningStateValue) String() string {
	if value, exists := provisioningStateToString[p]; exists {
		return value
	}

	return ValueNotFound
}

const (
	CoupledPasswordModel PasswordModelValue = iota
	SeparatePasswordModel
	SeparateHashPasswordModel
)

// PasswordModelToString is a map of PasswordModel values to their string representations.
var PasswordModelToString = map[PasswordModelValue]string{
	CoupledPasswordModel:      "CoupledPasswordModel",
	SeparatePasswordModel:     "SeparatePasswordModel",
	SeparateHashPasswordModel: "SeparateHashPasswordModel",
}

// String returns the string representation of the PasswordModel value.
func (p PasswordModelValue) String() string {
	if value, exists := PasswordModelToString[p]; exists {
		return value
	}

	return ValueNotFound
}

const (
	ReturnValueSuccess                 ReturnValue = 0
	ReturnValueInternalError           ReturnValue = 1
	ReturnValueNotPermitted            ReturnValue = 16
	ReturnValueInvalidParameter        ReturnValue = 36
	ReturnValueFlashWriteLimitExceeded ReturnValue = 38
	ReturnValueInvalidPassword         ReturnValue = 2054
	ReturnValueBlockingComponent       ReturnValue = 2076
)

var returnValueToString = map[ReturnValue]string{
	ReturnValueSuccess:                 "ReturnValueSuccess",
	ReturnValueInternalError:           "ReturnValueInternalError",
	ReturnValueNotPermitted:            "ReturnValueNotPermitted",
	ReturnValueInvalidParameter:        "ReturnValueInvalidParameter",
	ReturnValueFlashWriteLimitExceeded: "ReturnValueFlashWriteLimitExceeded",
	ReturnValueInvalidPassword:         "ReturnValueInvalidPassword",
	ReturnValueBlockingComponent:       "ReturnValueBlockingComponent",
}

// String returns the string representation of the ReturnValue value.
func (r ReturnValue) String() string {
	if value, exists := returnValueToString[r]; exists {
		return value
	}

	return ValueNotFound
}
