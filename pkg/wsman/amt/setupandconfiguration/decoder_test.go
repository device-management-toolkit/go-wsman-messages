/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package setupandconfiguration

import (
	"testing"
)

func TestRequestedStateString(t *testing.T) {
	tests := []struct {
		state    RequestedState
		expected string
	}{
		{RequestedStateEnabled, "RequestedStateEnabled"},
		{RequestedStateDisabled, "RequestedStateDisabled"},
		{RequestedStateShutDown, "RequestedStateShutDown"},
		{RequestedStateNoChange, "RequestedStateNoChange"},
		{RequestedStateOffline, "RequestedStateOffline"},
		{RequestedStateTest, "RequestedStateTest"},
		{RequestedStateDeferred, "RequestedStateDeferred"},
		{RequestedStateQuiesce, "RequestedStateQuiesce"},
		{RequestedStateReboot, "RequestedStateReboot"},
		{RequestedStateReset, "RequestedStateReset"},
		{RequestedStateNotApplicable, "RequestedStateNotApplicable"},
		{RequestedStateUnknown, "RequestedStateUnknown"},
		{RequestedState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestEnabledStateString(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
		{EnabledStateUnknown, "EnabledStateUnknown"},
		{EnabledStateOther, "EnabledStateOther"},
		{EnabledStateEnabled, "EnabledStateEnabled"},
		{EnabledStateDisabled, "EnabledStateDisabled"},
		{EnabledStateShuttingDown, "EnabledStateShuttingDown"},
		{EnabledStateNotApplicable, "EnabledStateNotApplicable"},
		{EnabledStateEnabledButOffline, "EnabledStateEnabledButOffline"},
		{EnabledStateInTest, "EnabledStateInTest"},
		{EnabledStateDeferred, "EnabledStateDeferred"},
		{EnabledStateQuiesce, "EnabledStateQuiesce"},
		{EnabledStateStarting, "EnabledStateStarting"},
		{EnabledState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestProvisioningModeString(t *testing.T) {
	tests := []struct {
		mode     ProvisioningModeValue
		expected string
	}{
		{AdminControlMode, "AdminControlMode"},
		{ClientControlMode, "ClientControlMode"},
		{ProvisioningModeValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.mode.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestProvisioningStateString(t *testing.T) {
	tests := []struct {
		state    ProvisioningStateValue
		expected string
	}{
		{PreProvisioning, "PreProvisioning"},
		{InProvisioning, "InProvisioning"},
		{PostProvisioning, "PostProvisioning"},
		{ProvisioningStateValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPasswordModelString(t *testing.T) {
	tests := []struct {
		model    PasswordModelValue
		expected string
	}{
		{CoupledPasswordModel, "CoupledPasswordModel"},
		{SeparatePasswordModel, "SeparatePasswordModel"},
		{SeparateHashPasswordModel, "SeparateHashPasswordModel"},
		{PasswordModelValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.model.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestReturnValueString(t *testing.T) {
	tests := []struct {
		value    ReturnValue
		expected string
	}{
		{ReturnValueSuccess, "ReturnValueSuccess"},
		{ReturnValueInternalError, "ReturnValueInternalError"},
		{ReturnValueNotPermitted, "ReturnValueNotPermitted"},
		{ReturnValueInvalidParameter, "ReturnValueInvalidParameter"},
		{ReturnValueFlashWriteLimitExceeded, "ReturnValueFlashWriteLimitExceeded"},
		{ReturnValueInvalidPassword, "ReturnValueInvalidPassword"},
		{ReturnValueDataMissing, "ReturnValueDataMissing"},
		{ReturnValueBlockingComponent, "ReturnValueBlockingComponent"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.value.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
