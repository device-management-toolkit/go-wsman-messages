/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import "testing"

func TestAuthenticationMethod_String(t *testing.T) {
	tests := []struct {
		state    AuthenticationMethod
		expected string
	}{
		{AuthenticationMethodOther, "Other"},
		{AuthenticationMethodOpenSystem, "OpenSystem"},
		{AuthenticationMethodSharedKey, "SharedKey"},
		{AuthenticationMethodWPAPSK, "WPAPSK"},
		{AuthenticationMethodWPAIEEE8021x, "WPAIEEE8021x"},
		{AuthenticationMethodWPA2PSK, "WPA2PSK"},
		{AuthenticationMethodWPA2IEEE8021x, "WPA2IEEE8021x"},
		{AuthenticationMethodWPA3SAE, "WPA3SAE"},
		{AuthenticationMethodWPA3OWE, "WPA3OWE"},
		{AuthenticationMethod(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestBSSType_String(t *testing.T) {
	tests := []struct {
		state    BSSType
		expected string
	}{
		{BSSTypeUnknown, "Unknown"},
		{BSSTypeIndependent, "Independent"},
		{BSSTypeInfrastructure, "Infrastructure"},
		{BSSType(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestEnabledState_String(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
		{EnabledStateWifiDisabled, "WifiDisabled"},
		{EnabledStateWifiEnabledS0, "WifiEnabledS0"},
		{EnabledStateWifiEnabledS0SxAC, "WifiEnabledS0SxAC"},
		{EnabledState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestEncryptionMethod_String(t *testing.T) {
	tests := []struct {
		state    EncryptionMethod
		expected string
	}{
		{EncryptionMethodOther, "Other"},
		{EncryptionMethodWEP, "WEP"},
		{EncryptionMethodTKIP, "TKIP"},
		{EncryptionMethodCCMP, "CCMP"},
		{EncryptionMethodNone, "None"},
		{EncryptionMethod(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestHealthState_String(t *testing.T) {
	tests := []struct {
		state    HealthState
		expected string
	}{
		{HealthStateOK, "OK"},
		{HealthStateDegraded, "Degraded"},
		{HealthStateMinorFailure, "MinorFailure"},
		{HealthStateMajorFailure, "MajorFailure"},
		{HealthStateCriticalFailure, "CriticalFailure"},
		{HealthStateNonRecoverableError, "NonRecoverableError"},
		{HealthState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestLinkTechnology_String(t *testing.T) {
	tests := []struct {
		state    LinkTechnology
		expected string
	}{
		{LinkTechnologyUnknown, "Unknown"},
		{LinkTechnologyOther, "Other"},
		{LinkTechnologyEthernet, "Ethernet"},
		{LinkTechnologyIB, "IB"},
		{LinkTechnologyFC, "FC"},
		{LinkTechnologyFDDI, "FDDI"},
		{LinkTechnologyATM, "ATM"},
		{LinkTechnologyTokenRing, "TokenRing"},
		{LinkTechnologyFrameRelay, "FrameRelay"},
		{LinkTechnologyInfrared, "Infrared"},
		{LinkTechnologyBlueTooth, "BlueTooth"},
		{LinkTechnologyWirelessLAN, "WirelessLAN"},
		{LinkTechnology(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestRequestedState_String(t *testing.T) {
	tests := []struct {
		state    RequestedState
		expected string
	}{
		{RequestedStateWifiDisabled, "WifiDisabled"},
		{RequestedStateWifiEnabledS0, "WifiEnabledS0"},
		{RequestedStateWifiEnabledS0SxAC, "WifiEnabledS0SxAC"},
		{RequestedState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestReturnValue_String(t *testing.T) {
	tests := []struct {
		state    ReturnValue
		expected string
	}{
		{CompletedWithNoError, "CompletedWithNoError"},
		{NotSupported, "NotSupported"},
		{UnknownOrUnspecifiedError, "UnknownOrUnspecifiedError"},
		{CannotCompleteWithinTimeoutPeriod, "CannotCompleteWithinTimeoutPeriod"},
		{Failed, "Failed"},
		{InvalidParameter, "InvalidParameter"},
		{InUse, "InUse"},
		{MethodParametersCheckedJobStarted, "MethodParametersCheckedJobStarted"},
		{InvalidStateTransition, "InvalidStateTransition"},
		{UseOfTimeoutParameterNotSupported, "UseOfTimeoutParameterNotSupported"},
		{Busy, "Busy"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestPortType_String(t *testing.T) {
	tests := []struct {
		state    PortType
		expected string
	}{
		{PortTypeUnknown, "Unknown"},
		{PortTypeOther, "Other"},
		{PortType80211a, "802.11a"},
		{PortType80211b, "802.11b"},
		{PortType80211g, "802.11g"},
		{PortType80211n, "802.11n"},
		{PortType(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestParseEncryptionMethod(t *testing.T) {
	tests := []struct {
		input    string
		expected EncryptionMethod
		success  bool
	}{
		{"Other", EncryptionMethodOther, true},
		{"WEP", EncryptionMethodWEP, true},
		{"TKIP", EncryptionMethodTKIP, true},
		{"CCMP", EncryptionMethodCCMP, true},
		{"None", EncryptionMethodNone, true},
		{"wep", EncryptionMethodWEP, true},      // case insensitive
		{"tkip", EncryptionMethodTKIP, true},    // case insensitive
		{"invalid", EncryptionMethod(0), false}, // invalid
	}

	for _, test := range tests {
		result, ok := ParseEncryptionMethod(test.input)
		if ok != test.success {
			t.Errorf("For input %s, expected success %v but got %v", test.input, test.success, ok)
		}

		if result != test.expected {
			t.Errorf("For input %s, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
