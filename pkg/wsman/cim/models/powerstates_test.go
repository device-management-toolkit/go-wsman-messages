/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package models

import "testing"

func TestAvailableRequestedPowerStates_String(t *testing.T) {
	tests := []struct {
		state    AvailableRequestedPowerStates
		expected string
	}{
		{AvailableRequestedPowerStatesOther, "Other"},
		{AvailableRequestedPowerStatesOn, "On"},
		{AvailableRequestedPowerStatesSleepLight, "SleepLight"},
		{AvailableRequestedPowerStatesSleepDeep, "SleepDeep"},
		{AvailableRequestedPowerStatesPowerCycleSoft, "PowerCycleSoft"},
		{AvailableRequestedPowerStatesOffHard, "OffHard"},
		{AvailableRequestedPowerStatesHibernate, "Hibernate"},
		{AvailableRequestedPowerStatesOffSoft, "OffSoft"},
		{AvailableRequestedPowerStatesPowerCycleHard, "PowerCycleHard"},
		{AvailableRequestedPowerStatesMasterBusReset, "MasterBusReset"},
		{AvailableRequestedPowerStatesDiagnosticInterrupt, "DiagnosticInterrupt"},
		{AvailableRequestedPowerStatesPowerOffSoftGraceful, "PowerOffSoftGraceful"},
		{AvailableRequestedPowerStatesPowerOffHardGraceful, "PowerOffHardGraceful"},
		{AvailableRequestedPowerStatesMasterBusResetGraceful, "MasterBusResetGraceful"},
		{AvailableRequestedPowerStatesPowerCycleSoftGraceful, "PowerCycleSoftGraceful"},
		{AvailableRequestedPowerStatesPowerCycleHardGraceful, "PowerCycleHardGraceful"},
		{AvailableRequestedPowerStates(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPowerState_String(t *testing.T) {
	tests := []struct {
		state    PowerState
		expected string
	}{
		{PowerStateOther, "Other"},
		{PowerStateOn, "On"},
		{PowerStateSleepLight, "SleepLight"},
		{PowerStateSleepDeep, "SleepDeep"},
		{PowerStatePowerCycleSoft, "PowerCycleSoft"},
		{PowerStateOffHard, "OffHard"},
		{PowerStateHibernate, "Hibernate"},
		{PowerStateOffSoft, "OffSoft"},
		{PowerStatePowerCycleHard, "PowerCycleHard"},
		{PowerStateMasterBusReset, "MasterBusReset"},
		{PowerStateDiagnosticInterruptNMI, "DiagnosticInterruptNMI"},
		{PowerStatePowerOffSoftGraceful, "PowerOffSoftGraceful"},
		{PowerStatePowerOffHardGraceful, "PowerOffHardGraceful"},
		{PowerStateMasterBusResetGraceful, "MasterBusResetGraceful"},
		{PowerStatePowerCycleSoftGraceful, "PowerCycleSoftGraceful"},
		{PowerStatePowerCycleHardGraceful, "PowerCycleHardGraceful"},
		{PowerStateDiagnosticInterruptINIT, "DiagnosticInterruptINIT"},
		{PowerState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestRequestedPowerState_String(t *testing.T) {
	tests := []struct {
		state    RequestedPowerState
		expected string
	}{
		{RequestedPowerStateUnknown, "Unknown"},
		{RequestedPowerStateOther, "Other"},
		{RequestedPowerStateOn, "On"},
		{RequestedPowerStateSleepLight, "SleepLight"},
		{RequestedPowerStateSleepDeep, "SleepDeep"},
		{RequestedPowerStatePowerCycleSoft, "PowerCycleSoft"},
		{RequestedPowerStateOffHard, "OffHard"},
		{RequestedPowerStateHibernate, "Hibernate"},
		{RequestedPowerStateOffSoft, "OffSoft"},
		{RequestedPowerStatePowerCycleHard, "PowerCycleHard"},
		{RequestedPowerStateMasterBusReset, "MasterBusReset"},
		{RequestedPowerStateDiagnosticInterruptNMI, "DiagnosticInterruptNMI"},
		{RequestedPowerStateNotApplicable, "NotApplicable"},
		{RequestedPowerStatePowerOffSoftGraceful, "PowerOffSoftGraceful"},
		{RequestedPowerStatePowerOffHardGraceful, "PowerOffHardGraceful"},
		{RequestedPowerStateMasterBusResetGraceful, "MasterBusResetGraceful"},
		{RequestedPowerStatePowerCycleSoftGraceful, "PowerCycleSoftGraceful"},
		{RequestedPowerStatePowerCycleHardGraceful, "PowerCycleHardGraceful"},
		{RequestedPowerStateDiagnosticInterruptINIT, "DiagnosticInterruptINIT"},
		{RequestedPowerState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestTransitioningToPowerState_String(t *testing.T) {
	tests := []struct {
		state    TransitioningToPowerState
		expected string
	}{
		{TransitioningToPowerStateOther, "Other"},
		{TransitioningToPowerStateOn, "On"},
		{TransitioningToPowerStateSleepLight, "SleepLight"},
		{TransitioningToPowerStateSleepDeep, "SleepDeep"},
		{TransitioningToPowerStatePowerCycleSoft, "PowerCycleSoft"},
		{TransitioningToPowerStateOffHard, "OffHard"},
		{TransitioningToPowerStateHibernate, "Hibernate"},
		{TransitioningToPowerStateOffSoft, "OffSoft"},
		{TransitioningToPowerStatePowerCycleHard, "PowerCycleHard"},
		{TransitioningToPowerStateMasterBusReset, "MasterBusReset"},
		{TransitioningToPowerStateDiagnosticInterruptNMI, "DiagnosticInterruptNMI"},
		{TransitioningToPowerStatePowerOffSoftGraceful, "PowerOffSoftGraceful"},
		{TransitioningToPowerStatePowerOffHardGraceful, "PowerOffHardGraceful"},
		{TransitioningToPowerStateMasterBusResetGraceful, "MasterBusResetGraceful"},
		{TransitioningToPowerStatePowerCycleSoftGraceful, "PowerCycleSoftGraceful"},
		{TransitioningToPowerStatePowerCycleHardGraceful, "PowerCycleHardGraceful"},
		{TransitioningToPowerStateDiagnosticInterruptINIT, "DiagnosticInterruptINIT"},
		{TransitioningToPowerStateNotApplicable, "NotApplicable"},
		{TransitioningToPowerStateNoChange, "NoChange"},
		{TransitioningToPowerState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
