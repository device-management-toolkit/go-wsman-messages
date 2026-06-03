/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiportconfiguration

// INPUTS constants

const (
	AMTWiFiPortConfigurationService string = "AMT_WiFiPortConfigurationService"
	AddWiFiSettings                 string = "AddWiFiSettings"
	UpdateWiFiSettings              string = "UpdateWiFiSettings"
	ValueNotFound                   string = "Value not found in map"
)

const (
	RelaxedPolicy NoHostCsmeSoftwarePolicy = iota
	AggressivePolicy
	Reserved
)

// noHostCsmeSoftwarePolicyToString is a map of NoHostCsmeSoftwarePolicy values to their string representations.
var noHostCsmeSoftwarePolicyToString = map[NoHostCsmeSoftwarePolicy]string{
	RelaxedPolicy:    "RelaxedPolicy",
	AggressivePolicy: "AggressivePolicy",
}

// String returns the string representation of the NoHostCsmeSoftwarePolicy value.
func (n NoHostCsmeSoftwarePolicy) String() string {
	if value, exists := noHostCsmeSoftwarePolicyToString[n]; exists {
		return value
	}

	return ValueNotFound
}

const (
	HealthStateUnknown             HealthState = 0
	HealthStateOK                  HealthState = 5
	HealthStateDegradedWarning     HealthState = 10
	HealthStateMinorFailure        HealthState = 15
	HealthStateMajorFailure        HealthState = 20
	HealthStateCriticalFailure     HealthState = 25
	HealthStateNonRecoverableError HealthState = 30
)

// healthStateToString is a map of HealthState values to their string representations.
var healthStateToString = map[HealthState]string{
	HealthStateUnknown:             "HealthStateUnknown",
	HealthStateOK:                  "HealthStateOK",
	HealthStateDegradedWarning:     "HealthStateDegradedWarning",
	HealthStateMinorFailure:        "HealthStateMinorFailure",
	HealthStateMajorFailure:        "HealthStateMajorFailure",
	HealthStateCriticalFailure:     "HealthStateCriticalFailure",
	HealthStateNonRecoverableError: "HealthStateNonRecoverableError",
}

// String returns the string representation of the HealthState value.
func (h HealthState) String() string {
	if value, exists := healthStateToString[h]; exists {
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
	LocalSyncDisabled                      LocalProfileSynchronizationEnabled = 0
	LocalUserProfileSynchronizationEnabled LocalProfileSynchronizationEnabled = 1
	UnrestrictedSync                       LocalProfileSynchronizationEnabled = 3
)

// localProfileSynchronizationEnabledToString is a map of LocalProfileSynchronizationEnabled to string.
var localProfileSynchronizationEnabledToString = map[LocalProfileSynchronizationEnabled]string{
	LocalSyncDisabled:                      "LocalSyncDisabled",
	LocalUserProfileSynchronizationEnabled: "LocalUserProfileSynchronizationEnabled",
	UnrestrictedSync:                       "UnrestrictedSync",
}

// String returns the string representation of the LocalProfileSynchronizationEnabled value.
func (l LocalProfileSynchronizationEnabled) String() string {
	if value, exists := localProfileSynchronizationEnabledToString[l]; exists {
		return value
	}

	return ValueNotFound
}

const (
	ReturnValueCompletedNoError ReturnValue = iota
	ReturnValueNotSupported
	ReturnValueFailed
	ReturnValueInvalidParameter
	ReturnValueInvalidReference
)

// returnValueToString is a map of ReturnValue values to their string representations.
var returnValueToString = map[ReturnValue]string{
	ReturnValueCompletedNoError: "ReturnValueCompletedNoError",
	ReturnValueNotSupported:     "ReturnValueNotSupported",
	ReturnValueFailed:           "ReturnValueFailed",
	ReturnValueInvalidParameter: "ReturnValueInvalidParameter",
	ReturnValueInvalidReference: "ReturnValueInvalidReference",
}

// String returns the string representation of the ReturnValue value.
func (r ReturnValue) String() string {
	if value, exists := returnValueToString[r]; exists {
		return value
	}

	return ValueNotFound
}
