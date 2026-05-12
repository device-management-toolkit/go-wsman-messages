/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package sensor

import "testing"

func TestOperationalStatus_String(t *testing.T) {
	tests := []struct {
		state    OperationalStatus
		expected string
	}{
		{OperationalStatusUnknown, "Unknown"},
		{OperationalStatusOther, "Other"},
		{OperationalStatusOK, "OK"},
		{OperationalStatusDegraded, "Degraded"},
		{OperationalStatusStressed, "Stressed"},
		{OperationalStatusPredictiveFailure, "PredictiveFailure"},
		{OperationalStatusError, "Error"},
		{OperationalStatusNonRecoverableError, "NonRecoverableError"},
		{OperationalStatusStarting, "Starting"},
		{OperationalStatusStopping, "Stopping"},
		{OperationalStatusStopped, "Stopped"},
		{OperationalStatusInService, "InService"},
		{OperationalStatusNoContact, "NoContact"},
		{OperationalStatusLostCommunication, "LostCommunication"},
		{OperationalStatusAborted, "Aborted"},
		{OperationalStatusDormant, "Dormant"},
		{OperationalStatusSupportingEntityInError, "SupportingEntityInError"},
		{OperationalStatusCompleted, "Completed"},
		{OperationalStatusPowerMode, "PowerMode"},
		{OperationalStatusRelocating, "Relocating"},
		{OperationalStatus(999), ValueNotFound},
	}

	for _, test := range tests {
		if got := test.state.String(); got != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, got)
		}
	}
}

func TestHealthState_String(t *testing.T) {
	tests := []struct {
		state    HealthState
		expected string
	}{
		{HealthStateUnknown, "Unknown"},
		{HealthStateOK, "OK"},
		{HealthStateDegradedWarning, "DegradedWarning"},
		{HealthStateMinorFailure, "MinorFailure"},
		{HealthStateMajorFailure, "MajorFailure"},
		{HealthStateCriticalFailure, "CriticalFailure"},
		{HealthStateNonRecoverableError, "NonRecoverableError"},
		{HealthState(999), ValueNotFound},
	}

	for _, test := range tests {
		if got := test.state.String(); got != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, got)
		}
	}
}

func TestEnabledState_String(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
		{EnabledStateUnknown, "Unknown"},
		{EnabledStateOther, "Other"},
		{EnabledStateEnabled, "Enabled"},
		{EnabledStateDisabled, "Disabled"},
		{EnabledStateShuttingDown, "ShuttingDown"},
		{EnabledStateNotApplicable, "NotApplicable"},
		{EnabledStateEnabledButOffline, "EnabledButOffline"},
		{EnabledStateInTest, "InTest"},
		{EnabledStateDeferred, "Deferred"},
		{EnabledStateQuiesce, "Quiesce"},
		{EnabledStateStarting, "Starting"},
		{EnabledState(999), ValueNotFound},
	}

	for _, test := range tests {
		if got := test.state.String(); got != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, got)
		}
	}
}

func TestRequestedState_String(t *testing.T) {
	tests := []struct {
		state    RequestedState
		expected string
	}{
		{RequestedStateUnknown, "Unknown"},
		{RequestedStateEnabled, "Enabled"},
		{RequestedStateDisabled, "Disabled"},
		{RequestedStateShutDown, "ShutDown"},
		{RequestedStateNoChange, "NoChange"},
		{RequestedStateOffline, "Offline"},
		{RequestedStateTest, "Test"},
		{RequestedStateDeferred, "Deferred"},
		{RequestedStateQuiesce, "Quiesce"},
		{RequestedStateReboot, "Reboot"},
		{RequestedStateReset, "Reset"},
		{RequestedStateNotApplicable, "NotApplicable"},
		{RequestedState(999), ValueNotFound},
	}

	for _, test := range tests {
		if got := test.state.String(); got != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, got)
		}
	}
}

func TestSensorType_String(t *testing.T) {
	tests := []struct {
		state    SensorType
		expected string
	}{
		{SensorTypeUnknown, "Unknown"},
		{SensorTypeOther, "Other"},
		{SensorTypeTemperature, "Temperature"},
		{SensorTypeVoltage, "Voltage"},
		{SensorTypeCurrent, "Current"},
		{SensorTypeTachometer, "Tachometer"},
		{SensorTypeCounter, "Counter"},
		{SensorTypeSwitch, "Switch"},
		{SensorTypeLock, "Lock"},
		{SensorTypeHumidity, "Humidity"},
		{SensorTypeSmokeDetection, "SmokeDetection"},
		{SensorTypePresence, "Presence"},
		{SensorTypeAirFlow, "AirFlow"},
		{SensorTypePowerConsumption, "PowerConsumption"},
		{SensorTypePowerProduction, "PowerProduction"},
		{SensorTypePressure, "Pressure"},
		{SensorTypeIntrusion, "Intrusion"},
		{SensorType(999), ValueNotFound},
	}

	for _, test := range tests {
		if got := test.state.String(); got != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, got)
		}
	}
}
