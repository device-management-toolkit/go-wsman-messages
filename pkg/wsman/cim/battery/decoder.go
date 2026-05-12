/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package battery

// Resource URI suffix and the placeholder string returned by every enum
// String() when an unrecognised value is encountered. The full WSMAN
// resource URI is built by prepending the CIM schema base URI from
// wsmantesting.CIMResourceURIBase.
const (
	CIMBattery    string = "CIM_Battery"
	ValueNotFound string = "Value not found in map"
)

// OperationalStatus values per the DMTF CIM operational-status
// enumeration (ValueMap 0..19 plus DMTF/vendor reserved ranges). The
// first entry of CIM_Battery.OperationalStatus is the primary status;
// "Completed" is meant to be combined with "OK", "Error", or "Degraded"
// to convey the outcome of a long-running operation.
const (
	OperationalStatusUnknown OperationalStatus = iota
	OperationalStatusOther
	OperationalStatusOK
	OperationalStatusDegraded
	// OperationalStatusStressed indicates the element is functioning
	// but needs attention (e.g. overload, overheated).
	OperationalStatusStressed
	// OperationalStatusPredictiveFailure indicates the element is
	// functioning nominally but is predicting a failure in the near
	// future.
	OperationalStatusPredictiveFailure
	OperationalStatusError
	OperationalStatusNonRecoverableError
	OperationalStatusStarting
	OperationalStatusStopping
	// OperationalStatusStopped is an orderly stop, in contrast to
	// "Aborted" which implies an abrupt stop.
	OperationalStatusStopped
	// OperationalStatusInService means the element is being configured,
	// maintained, cleaned, or otherwise administered.
	OperationalStatusInService
	// OperationalStatusNoContact means the monitoring system knows of
	// this element but has never been able to contact it.
	OperationalStatusNoContact
	// OperationalStatusLostCommunication means the element has been
	// contacted in the past but is currently unreachable.
	OperationalStatusLostCommunication
	OperationalStatusAborted
	// OperationalStatusDormant means the element is inactive or
	// quiesced.
	OperationalStatusDormant
	// OperationalStatusSupportingEntityInError means this element is
	// likely fine but a dependency is in error.
	OperationalStatusSupportingEntityInError
	// OperationalStatusCompleted should be combined with OK, Error, or
	// Degraded so clients can tell whether an operation finished
	// successfully.
	OperationalStatusCompleted
	// OperationalStatusPowerMode means the element has additional
	// power-model information available through the
	// AssociatedPowerManagementService association.
	OperationalStatusPowerMode
	OperationalStatusRelocating
)

// operationalStatusMap is a map of the OperationalStatus enumeration.
var operationalStatusMap = map[OperationalStatus]string{
	OperationalStatusUnknown:                 "Unknown",
	OperationalStatusOther:                   "Other",
	OperationalStatusOK:                      "OK",
	OperationalStatusDegraded:                "Degraded",
	OperationalStatusStressed:                "Stressed",
	OperationalStatusPredictiveFailure:       "PredictiveFailure",
	OperationalStatusError:                   "Error",
	OperationalStatusNonRecoverableError:     "NonRecoverableError",
	OperationalStatusStarting:                "Starting",
	OperationalStatusStopping:                "Stopping",
	OperationalStatusStopped:                 "Stopped",
	OperationalStatusInService:               "InService",
	OperationalStatusNoContact:               "NoContact",
	OperationalStatusLostCommunication:       "LostCommunication",
	OperationalStatusAborted:                 "Aborted",
	OperationalStatusDormant:                 "Dormant",
	OperationalStatusSupportingEntityInError: "SupportingEntityInError",
	OperationalStatusCompleted:               "Completed",
	OperationalStatusPowerMode:               "PowerMode",
	OperationalStatusRelocating:              "Relocating",
}

// String returns a human-readable string representation of the
// OperationalStatus enumeration. Unrecognised values return
// ValueNotFound.
func (e OperationalStatus) String() string {
	if s, ok := operationalStatusMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// HealthState values per the DMTF CIM health-state continuum. The
// schema reserves 32768–65535 for vendor-specific health values.
//
//	 0 — Unknown
//	 5 — OK            (fully functional)
//	10 — Degraded/Warning
//	15 — Minor Failure
//	20 — Major Failure
//	25 — Critical Failure
//	30 — Non-recoverable Error
const (
	HealthStateUnknown             HealthState = 0
	HealthStateOK                  HealthState = 5
	HealthStateDegradedWarning     HealthState = 10
	HealthStateMinorFailure        HealthState = 15
	HealthStateMajorFailure        HealthState = 20
	HealthStateCriticalFailure     HealthState = 25
	HealthStateNonRecoverableError HealthState = 30
)

// healthStateMap is a map of the HealthState enumeration.
var healthStateMap = map[HealthState]string{
	HealthStateUnknown:             "Unknown",
	HealthStateOK:                  "OK",
	HealthStateDegradedWarning:     "DegradedWarning",
	HealthStateMinorFailure:        "MinorFailure",
	HealthStateMajorFailure:        "MajorFailure",
	HealthStateCriticalFailure:     "CriticalFailure",
	HealthStateNonRecoverableError: "NonRecoverableError",
}

// String returns a human-readable string representation of the
// HealthState enumeration. Unrecognised values return ValueNotFound.
func (e HealthState) String() string {
	if s, ok := healthStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// EnabledState values per the DMTF CIM enabled-state enumeration.
// "Shutting Down" and "Starting" are transient states; compare against
// RequestedState to detect in-flight transitions.
const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	// EnabledStateEnabled — the element is, or could be, executing
	// commands and will accept new requests.
	EnabledStateEnabled
	// EnabledStateDisabled — the element will not execute commands and
	// drops any new requests.
	EnabledStateDisabled
	// EnabledStateShuttingDown — transitional: heading to Disabled.
	EnabledStateShuttingDown
	// EnabledStateNotApplicable — the element does not support being
	// enabled or disabled; RequestedState has no meaning in this state.
	EnabledStateNotApplicable
	// EnabledStateEnabledButOffline — may complete in-flight commands
	// but drops any new ones.
	EnabledStateEnabledButOffline
	EnabledStateInTest
	// EnabledStateDeferred — may complete in-flight commands and
	// queues any new ones.
	EnabledStateDeferred
	// EnabledStateQuiesce — enabled but in a restricted mode.
	EnabledStateQuiesce
	// EnabledStateStarting — transitional: heading to Enabled; new
	// requests are queued.
	EnabledStateStarting
)

// enabledStateMap is a map of the EnabledState enumeration.
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

// String returns a human-readable string representation of the
// EnabledState enumeration. Unrecognised values return ValueNotFound.
func (e EnabledState) String() string {
	if s, ok := enabledStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// RequestedState values per the DMTF CIM requested-state enumeration.
// Notes:
//   - 1 is intentionally skipped by the schema.
//   - "No Change" (5) is deprecated in favour of "Unknown" (0).
//   - "Reboot" = Shut Down then Enable; "Reset" = Disable then Enable.
//   - "Not Applicable" (12) is returned when knowledge of the last
//     requested state is not supported.
const (
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)

// requestedStateMap is a map of the RequestedState enumeration.
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

// String returns a human-readable string representation of the
// RequestedState enumeration. Unrecognised values return ValueNotFound.
func (e RequestedState) String() string {
	if s, ok := requestedStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// BatteryStatus values describe the battery charge status. Sourced from
// the DMTF Portable Battery MIF (mapping string MIF.DMTF|Portable
// Battery|003.14).
//
// Several values are deprecated by the CIM schema and are exposed here
// for completeness only:
//
//   - 6 Charging, 7 ChargingAndHigh, 8 ChargingAndLow,
//     9 ChargingAndCritical are deprecated in favour of ChargingStatus.
//   - 10 Undefined is deprecated in favour of 2 Unknown. In DMI it
//     historically meant "no battery installed" — in that case the
//     CIM_Battery instance should not be created at all.
const (
	BatteryStatusOther               BatteryStatus = 1
	BatteryStatusUnknown             BatteryStatus = 2
	BatteryStatusFullyCharged        BatteryStatus = 3
	BatteryStatusLow                 BatteryStatus = 4
	BatteryStatusCritical            BatteryStatus = 5
	BatteryStatusCharging            BatteryStatus = 6  // deprecated — use ChargingStatus
	BatteryStatusChargingAndHigh     BatteryStatus = 7  // deprecated — use ChargingStatus
	BatteryStatusChargingAndLow      BatteryStatus = 8  // deprecated — use ChargingStatus
	BatteryStatusChargingAndCritical BatteryStatus = 9  // deprecated — use ChargingStatus
	BatteryStatusUndefined           BatteryStatus = 10 // deprecated — use BatteryStatusUnknown
	BatteryStatusPartiallyCharged    BatteryStatus = 11
	BatteryStatusLearning            BatteryStatus = 12
	BatteryStatusOvercharged         BatteryStatus = 13
)

// batteryStatusMap is a map of the BatteryStatus enumeration.
var batteryStatusMap = map[BatteryStatus]string{
	BatteryStatusOther:               "Other",
	BatteryStatusUnknown:             "Unknown",
	BatteryStatusFullyCharged:        "FullyCharged",
	BatteryStatusLow:                 "Low",
	BatteryStatusCritical:            "Critical",
	BatteryStatusCharging:            "Charging",
	BatteryStatusChargingAndHigh:     "ChargingAndHigh",
	BatteryStatusChargingAndLow:      "ChargingAndLow",
	BatteryStatusChargingAndCritical: "ChargingAndCritical",
	BatteryStatusUndefined:           "Undefined",
	BatteryStatusPartiallyCharged:    "PartiallyCharged",
	BatteryStatusLearning:            "Learning",
	BatteryStatusOvercharged:         "Overcharged",
}

// String returns a human-readable string representation of the
// BatteryStatus enumeration. Unrecognised values return ValueNotFound.
func (e BatteryStatus) String() string {
	if s, ok := batteryStatusMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// Chemistry values describe the battery chemistry. Sourced from the
// DMTF Portable Battery MIF (mapping string MIF.DMTF|Portable
// Battery|003.7).
const (
	ChemistryOther              Chemistry = 1
	ChemistryUnknown            Chemistry = 2
	ChemistryLeadAcid           Chemistry = 3
	ChemistryNickelCadmium      Chemistry = 4
	ChemistryNickelMetalHydride Chemistry = 5
	ChemistryLithiumIon         Chemistry = 6
	ChemistryZincAir            Chemistry = 7
	ChemistryLithiumPolymer     Chemistry = 8
)

// chemistryMap is a map of the Chemistry enumeration.
var chemistryMap = map[Chemistry]string{
	ChemistryOther:              "Other",
	ChemistryUnknown:            "Unknown",
	ChemistryLeadAcid:           "LeadAcid",
	ChemistryNickelCadmium:      "NickelCadmium",
	ChemistryNickelMetalHydride: "NickelMetalHydride",
	ChemistryLithiumIon:         "LithiumIon",
	ChemistryZincAir:            "ZincAir",
	ChemistryLithiumPolymer:     "LithiumPolymer",
}

// String returns a human-readable string representation of the
// Chemistry enumeration. Unrecognised values return ValueNotFound.
func (e Chemistry) String() string {
	if s, ok := chemistryMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ChargingStatus values report whether the battery is currently
// charging, discharging, or idle. Note that 1 is intentionally skipped
// by the schema. Vendor-specific values use the 32768–65535 range.
const (
	ChargingStatusUnknown     ChargingStatus = 0
	ChargingStatusCharging    ChargingStatus = 2
	ChargingStatusDischarging ChargingStatus = 3
	// ChargingStatusIdle — the battery is neither charging nor
	// discharging.
	ChargingStatusIdle ChargingStatus = 4
)

// chargingStatusMap is a map of the ChargingStatus enumeration.
var chargingStatusMap = map[ChargingStatus]string{
	ChargingStatusUnknown:     "Unknown",
	ChargingStatusCharging:    "Charging",
	ChargingStatusDischarging: "Discharging",
	ChargingStatusIdle:        "Idle",
}

// String returns a human-readable string representation of the
// ChargingStatus enumeration. Unrecognised values return ValueNotFound.
func (e ChargingStatus) String() string {
	if s, ok := chargingStatusMap[e]; ok {
		return s
	}

	return ValueNotFound
}
