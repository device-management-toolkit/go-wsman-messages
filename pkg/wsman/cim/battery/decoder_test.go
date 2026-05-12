/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package battery

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
			t.Errorf("OperationalStatus(%d).String() = %q, want %q", test.state, got, test.expected)
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
			t.Errorf("HealthState(%d).String() = %q, want %q", test.state, got, test.expected)
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
			t.Errorf("EnabledState(%d).String() = %q, want %q", test.state, got, test.expected)
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
			t.Errorf("RequestedState(%d).String() = %q, want %q", test.state, got, test.expected)
		}
	}
}

func TestBatteryStatus_String(t *testing.T) {
	tests := []struct {
		state    BatteryStatus
		expected string
	}{
		{BatteryStatusOther, "Other"},
		{BatteryStatusUnknown, "Unknown"},
		{BatteryStatusFullyCharged, "FullyCharged"},
		{BatteryStatusLow, "Low"},
		{BatteryStatusCritical, "Critical"},
		{BatteryStatusCharging, "Charging"},
		{BatteryStatusChargingAndHigh, "ChargingAndHigh"},
		{BatteryStatusChargingAndLow, "ChargingAndLow"},
		{BatteryStatusChargingAndCritical, "ChargingAndCritical"},
		{BatteryStatusUndefined, "Undefined"},
		{BatteryStatusPartiallyCharged, "PartiallyCharged"},
		{BatteryStatusLearning, "Learning"},
		{BatteryStatusOvercharged, "Overcharged"},
		{BatteryStatus(999), ValueNotFound},
	}

	for _, test := range tests {
		if got := test.state.String(); got != test.expected {
			t.Errorf("BatteryStatus(%d).String() = %q, want %q", test.state, got, test.expected)
		}
	}
}

func TestChemistry_String(t *testing.T) {
	tests := []struct {
		state    Chemistry
		expected string
	}{
		{ChemistryOther, "Other"},
		{ChemistryUnknown, "Unknown"},
		{ChemistryLeadAcid, "LeadAcid"},
		{ChemistryNickelCadmium, "NickelCadmium"},
		{ChemistryNickelMetalHydride, "NickelMetalHydride"},
		{ChemistryLithiumIon, "LithiumIon"},
		{ChemistryZincAir, "ZincAir"},
		{ChemistryLithiumPolymer, "LithiumPolymer"},
		{Chemistry(999), ValueNotFound},
	}

	for _, test := range tests {
		if got := test.state.String(); got != test.expected {
			t.Errorf("Chemistry(%d).String() = %q, want %q", test.state, got, test.expected)
		}
	}
}

func TestChargingStatus_String(t *testing.T) {
	tests := []struct {
		state    ChargingStatus
		expected string
	}{
		{ChargingStatusUnknown, "Unknown"},
		{ChargingStatusCharging, "Charging"},
		{ChargingStatusDischarging, "Discharging"},
		{ChargingStatusIdle, "Idle"},
		{ChargingStatus(999), ValueNotFound},
	}

	for _, test := range tests {
		if got := test.state.String(); got != test.expected {
			t.Errorf("ChargingStatus(%d).String() = %q, want %q", test.state, got, test.expected)
		}
	}
}
