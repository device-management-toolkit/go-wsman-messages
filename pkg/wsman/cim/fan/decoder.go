/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package fan

const (
	// CIMFan is the WSMAN resource class name for CIM_Fan. It appears as the
	// final segment of the ResourceURI on every request this package emits.
	CIMFan string = "CIM_Fan"

	// ValueNotFound is the string returned by the enum String() helpers when
	// the underlying integer value is not present in the value map.
	ValueNotFound string = "Value not found in map"
)

// OperationalStatus values for CIM_Fan.OperationalStatus. Many enumeration
// values are self-explanatory; the less obvious ones are documented on the
// constant.
const (
	// OperationalStatusUnknown — status cannot currently be determined.
	OperationalStatusUnknown OperationalStatus = iota
	// OperationalStatusOther — status is vendor-specific.
	OperationalStatusOther
	// OperationalStatusOK — the element is functioning normally.
	OperationalStatusOK
	// OperationalStatusDegraded — the element is functioning but not at full capability.
	OperationalStatusDegraded
	// OperationalStatusStressed — the element is functioning but needs attention (e.g. overload, overheated).
	OperationalStatusStressed
	// OperationalStatusPredictiveFailure — the element is functioning nominally but predicting a failure in the near future.
	OperationalStatusPredictiveFailure
	// OperationalStatusError — the element has reported an error.
	OperationalStatusError
	// OperationalStatusNonRecoverableError — the element has failed and cannot recover.
	OperationalStatusNonRecoverableError
	// OperationalStatusStarting — the element is starting up.
	OperationalStatusStarting
	// OperationalStatusStopping — the element is shutting down.
	OperationalStatusStopping
	// OperationalStatusStopped — the element has been cleanly stopped.
	OperationalStatusStopped
	// OperationalStatusInService — the element is being configured, maintained, cleaned, or otherwise administered.
	OperationalStatusInService
	// OperationalStatusNoContact — the monitoring system has knowledge of the element but has never been able to communicate with it.
	OperationalStatusNoContact
	// OperationalStatusLostCommunication — the element is known to exist and was reachable in the past but is currently unreachable.
	OperationalStatusLostCommunication
	// OperationalStatusAborted — the element stopped abruptly; its state and configuration may need updating.
	OperationalStatusAborted
	// OperationalStatusDormant — the element is inactive or quiesced.
	OperationalStatusDormant
	// OperationalStatusSupportingEntityInError — the element itself may be OK but another element it depends on is in error.
	OperationalStatusSupportingEntityInError
	// OperationalStatusCompleted — the element has completed its operation; combine with OK / Error / Degraded for the outcome.
	OperationalStatusCompleted
	// OperationalStatusPowerMode — additional power-model information is available via the associated PowerManagementService.
	OperationalStatusPowerMode
	// OperationalStatusRelocating — the element is being relocated.
	OperationalStatusRelocating
)

// operationalStatusMap maps OperationalStatus to its human-readable label.
var operationalStatusMap = map[OperationalStatus]string{
	OperationalStatusUnknown:                 "Unknown",
	OperationalStatusOther:                   "Other",
	OperationalStatusOK:                      "OK",
	OperationalStatusDegraded:                "Degraded",
	OperationalStatusStressed:                "Stressed",
	OperationalStatusPredictiveFailure:       "Predictive Failure",
	OperationalStatusError:                   "Error",
	OperationalStatusNonRecoverableError:     "Non-Recoverable Error",
	OperationalStatusStarting:                "Starting",
	OperationalStatusStopping:                "Stopping",
	OperationalStatusStopped:                 "Stopped",
	OperationalStatusInService:               "In Service",
	OperationalStatusNoContact:               "No Contact",
	OperationalStatusLostCommunication:       "Lost Communication",
	OperationalStatusAborted:                 "Aborted",
	OperationalStatusDormant:                 "Dormant",
	OperationalStatusSupportingEntityInError: "Supporting Entity In Error",
	OperationalStatusCompleted:               "Completed",
	OperationalStatusPowerMode:               "Power Mode",
	OperationalStatusRelocating:              "Relocating",
}

// String returns a human-readable representation of the OperationalStatus
// value, or ValueNotFound if the value is not defined.
func (e OperationalStatus) String() string {
	if s, ok := operationalStatusMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// EnabledState values for CIM_Fan.EnabledState.
const (
	// EnabledStateUnknown — state cannot currently be determined.
	EnabledStateUnknown EnabledState = iota
	// EnabledStateOther — state is vendor-specific and described by CIM_EnabledLogicalElement.OtherEnabledState.
	EnabledStateOther
	// EnabledStateEnabled — the element is or could be executing commands and queues new requests.
	EnabledStateEnabled
	// EnabledStateDisabled — the element will not execute commands and drops new requests.
	EnabledStateDisabled
	// EnabledStateShuttingDown — the element is in the process of transitioning to Disabled.
	EnabledStateShuttingDown
	// EnabledStateNotApplicable — the element does not support being enabled or disabled. AMT reports this for fans.
	EnabledStateNotApplicable
	// EnabledStateEnabledButOffline — the element might be completing commands and will drop any new requests.
	EnabledStateEnabledButOffline
	// EnabledStateInTest — the element is in a test state.
	EnabledStateInTest
	// EnabledStateDeferred — the element might be completing commands but will queue any new requests.
	EnabledStateDeferred
	// EnabledStateQuiesce — the element is enabled but in a restricted mode.
	EnabledStateQuiesce
	// EnabledStateStarting — the element is in the process of transitioning to Enabled; new requests are queued.
	EnabledStateStarting
)

// enabledStateMap maps EnabledState to its human-readable label.
var enabledStateMap = map[EnabledState]string{
	EnabledStateUnknown:           "Unknown",
	EnabledStateOther:             "Other",
	EnabledStateEnabled:           "Enabled",
	EnabledStateDisabled:          "Disabled",
	EnabledStateShuttingDown:      "ShuttingDown",
	EnabledStateNotApplicable:     "NotApplicable",
	EnabledStateEnabledButOffline: "EnabledButOffline",
	EnabledStateInTest:            "InTest",
	EnabledStateDeferred:          "Deferred",
	EnabledStateQuiesce:           "Quiesce",
	EnabledStateStarting:          "Starting",
}

// String returns a human-readable representation of the EnabledState value, or
// ValueNotFound if the value is not defined.
func (e EnabledState) String() string {
	if s, ok := enabledStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// RequestedState values for CIM_Fan.RequestedState. The "Reboot" and "Reset"
// values build on EnabledState statuses: Reboot does a "Shut Down" and then
// moves to "Enabled"; Reset first "Disables" and then "Enables".
const (
	// RequestedStateUnknown — the last requested state is unknown.
	RequestedStateUnknown RequestedState = 0
	// RequestedStateEnabled — the element was last requested to be Enabled.
	RequestedStateEnabled RequestedState = 2
	// RequestedStateDisabled — the element was last requested to be Disabled (immediate disable, drop new requests).
	RequestedStateDisabled RequestedState = 3
	// RequestedStateShutDown — orderly transition to Disabled, possibly removing power.
	RequestedStateShutDown RequestedState = 4
	// RequestedStateNoChange — deprecated in favor of Unknown (0).
	RequestedStateNoChange RequestedState = 5
	// RequestedStateOffline — transition to EnabledButOffline.
	RequestedStateOffline RequestedState = 6
	// RequestedStateTest — transition to InTest.
	RequestedStateTest RequestedState = 7
	// RequestedStateDeferred — transition to Deferred.
	RequestedStateDeferred RequestedState = 8
	// RequestedStateQuiesce — transition to Quiesce.
	RequestedStateQuiesce RequestedState = 9
	// RequestedStateReboot — Shut Down then Enable.
	RequestedStateReboot RequestedState = 10
	// RequestedStateReset — Disable then Enable.
	RequestedStateReset RequestedState = 11
	// RequestedStateNotApplicable — the element does not track requested states. AMT reports this for fans.
	RequestedStateNotApplicable RequestedState = 12
)

// requestedStateMap maps RequestedState to its human-readable label.
var requestedStateMap = map[RequestedState]string{
	RequestedStateUnknown:       "Unknown",
	RequestedStateEnabled:       "Enabled",
	RequestedStateDisabled:      "Disabled",
	RequestedStateShutDown:      "ShutDown",
	RequestedStateNoChange:      "NoChange",
	RequestedStateOffline:       "Offline",
	RequestedStateTest:          "Test",
	RequestedStateDeferred:      "Deferred",
	RequestedStateQuiesce:       "Quiesce",
	RequestedStateReboot:        "Reboot",
	RequestedStateReset:         "Reset",
	RequestedStateNotApplicable: "NotApplicable",
}

// String returns a human-readable representation of the RequestedState value,
// or ValueNotFound if the value is not defined.
func (e RequestedState) String() string {
	if s, ok := requestedStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// HealthState values for CIM_Fan.HealthState. The continuum runs 0..30 where 5
// means entirely healthy and 30 means completely non-functional.
const (
	// HealthStateUnknown — the implementation cannot report on HealthState at this time.
	HealthStateUnknown HealthState = 0
	// HealthStateOK — fully functional and operating within normal parameters without error.
	HealthStateOK HealthState = 5
	// HealthStateDegraded — in working order with all functionality available, but not at the best of its abilities.
	HealthStateDegraded HealthState = 10
	// HealthStateMinorFailure — all functionality is available but some may be degraded.
	HealthStateMinorFailure HealthState = 15
	// HealthStateMajorFailure — the element is failing; some or all functionality is degraded or not working.
	HealthStateMajorFailure HealthState = 20
	// HealthStateCriticalFailure — the element is non-functional and recovery may not be possible.
	HealthStateCriticalFailure HealthState = 25
	// HealthStateNonRecoverableError — the element has completely failed and recovery is not possible.
	HealthStateNonRecoverableError HealthState = 30
)

// healthStateMap maps HealthState to its human-readable label.
var healthStateMap = map[HealthState]string{
	HealthStateUnknown:             "Unknown",
	HealthStateOK:                  "OK",
	HealthStateDegraded:            "Degraded",
	HealthStateMinorFailure:        "MinorFailure",
	HealthStateMajorFailure:        "MajorFailure",
	HealthStateCriticalFailure:     "CriticalFailure",
	HealthStateNonRecoverableError: "NonRecoverableError",
}

// String returns a human-readable representation of the HealthState value, or
// ValueNotFound if the value is not defined.
func (e HealthState) String() string {
	if s, ok := healthStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}
