/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package sensor

// CIMSensor is the WSMAN resource class name for CIM_Sensor instances. Combined
// with the CIM schema base URI it forms the resource URI that every envelope
// sent by this package targets. ValueNotFound is the sentinel returned by the
// String() methods below when an integer falls outside the documented range.
const (
	CIMSensor     string = "CIM_Sensor"
	ValueNotFound string = "Value not found in map"
)

// OperationalStatus values defined by CIM_ManagedSystemElement.OperationalStatus.
// AMT typically reports a single value at index 0; the remaining entries are
// reserved for future use. Per the DMTF model, "Stressed" means the element is
// functioning but needs attention (e.g. overheat); "Predictive Failure" means
// it is functioning nominally but predicting a future failure; "Stopped"
// implies a clean stop and "Aborted" implies an abrupt one; and a value
// combined with "Completed" indicates whether a finished operation passed,
// failed, or degraded.
const (
	OperationalStatusUnknown OperationalStatus = iota
	OperationalStatusOther
	OperationalStatusOK
	OperationalStatusDegraded
	OperationalStatusStressed
	OperationalStatusPredictiveFailure
	OperationalStatusError
	OperationalStatusNonRecoverableError
	OperationalStatusStarting
	OperationalStatusStopping
	OperationalStatusStopped
	OperationalStatusInService
	OperationalStatusNoContact
	OperationalStatusLostCommunication
	OperationalStatusAborted
	OperationalStatusDormant
	OperationalStatusSupportingEntityInError
	OperationalStatusCompleted
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

// String returns a human-readable string representation of the OperationalStatus enumeration.
func (e OperationalStatus) String() string {
	if s, ok := operationalStatusMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// HealthState values defined by CIM_ManagedSystemElement.HealthState. The
// 0–30 continuum reports the element's own health, not its subcomponents:
// 0 = unable to report, 5 = healthy and within normal operating parameters,
// 10 = working but degraded (e.g. recoverable errors), 15–25 = increasing
// severity of functional impairment, 30 = completely failed with no recovery.
// Values in 32768–65535 are vendor-specific.
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

// String returns a human-readable string representation of the HealthState enumeration.
func (e HealthState) String() string {
	if s, ok := healthStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// EnabledState values defined by CIM_EnabledLogicalElement.EnabledState. The
// enumeration covers both the enabled/disabled steady states and the
// transitions between them: Enabled (2) executes commands and queues new
// requests; Disabled (3) drops new requests; ShuttingDown (4) and Starting
// (10) are transient. NotApplicable (5) is what AMT reports for hardware
// sensors that cannot be enabled or disabled. EnabledButOffline (6) may
// complete in-flight work but drops new requests; Deferred (8) queues them;
// Quiesce (9) runs in a restricted mode.
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

// String returns a human-readable string representation of the EnabledState enumeration.
func (e EnabledState) String() string {
	if s, ok := enabledStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// RequestedState values defined by CIM_EnabledLogicalElement.RequestedState.
// This is the last requested or desired EnabledState, independent of how the
// request was issued. NoChange (5) is deprecated in favour of Unknown (0).
// Reboot (10) means ShutDown followed by Enabled; Reset (11) means Disabled
// followed by Enabled. ShutDown (4) is an orderly transition to Disabled
// (possibly including power removal); Disabled (3) is an immediate disable.
// NotApplicable (12) is reported when the element does not track the last
// requested state — AMT-side sensors commonly report this.
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

// String returns a human-readable string representation of the RequestedState enumeration.
func (e RequestedState) String() string {
	if s, ok := requestedStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// SensorType values defined by CIM_Sensor.SensorType. A Temperature sensor
// reads environmental temperature; Voltage and Current report electrical
// readings; a Tachometer measures rotational speed (e.g. fan RPM). A Counter
// is a general numeric sensor that can be cleared but never decreases. Switch
// has states like Open/Close, Lock has Locked/Unlocked. Humidity, SmokeDetection,
// AirFlow, Pressure, and Intrusion report the equivalent environmental
// characteristics. Presence reports whether a PhysicalElement is present;
// PowerConsumption and PowerProduction report instantaneous power.
//
// On Intel® AMT, only Temperature, Voltage, and Current map from the ASF
// alert tables (ASF_ALERTDATA.EventSensorType 01h/02h/03h respectively); any
// other ASF event type is reported as Unknown (0).
const (
	SensorTypeUnknown SensorType = iota
	SensorTypeOther
	SensorTypeTemperature
	SensorTypeVoltage
	SensorTypeCurrent
	SensorTypeTachometer
	SensorTypeCounter
	SensorTypeSwitch
	SensorTypeLock
	SensorTypeHumidity
	SensorTypeSmokeDetection
	SensorTypePresence
	SensorTypeAirFlow
	SensorTypePowerConsumption
	SensorTypePowerProduction
	SensorTypePressure
	SensorTypeIntrusion
)

// sensorTypeMap is a map of the SensorType enumeration.
var sensorTypeMap = map[SensorType]string{
	SensorTypeUnknown:          "Unknown",
	SensorTypeOther:            "Other",
	SensorTypeTemperature:      "Temperature",
	SensorTypeVoltage:          "Voltage",
	SensorTypeCurrent:          "Current",
	SensorTypeTachometer:       "Tachometer",
	SensorTypeCounter:          "Counter",
	SensorTypeSwitch:           "Switch",
	SensorTypeLock:             "Lock",
	SensorTypeHumidity:         "Humidity",
	SensorTypeSmokeDetection:   "SmokeDetection",
	SensorTypePresence:         "Presence",
	SensorTypeAirFlow:          "AirFlow",
	SensorTypePowerConsumption: "PowerConsumption",
	SensorTypePowerProduction:  "PowerProduction",
	SensorTypePressure:         "Pressure",
	SensorTypeIntrusion:        "Intrusion",
}

// String returns a human-readable string representation of the SensorType enumeration.
func (e SensorType) String() string {
	if s, ok := sensorTypeMap[e]; ok {
		return s
	}

	return ValueNotFound
}
