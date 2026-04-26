/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiendpoint

import "strings"

const (
	CIMWiFiEndpoint string = "CIM_WiFiEndpoint"
	ValueNotFound   string = "Value not found in map"
)

const (
	EnabledStateWifiDisabled      EnabledState = 3
	EnabledStateWifiEnabledS0     EnabledState = 32768
	EnabledStateWifiEnabledS0SxAC EnabledState = 32769
)

// enabledStateMap is a map of the EnabledState enumeration.
var enabledStateMap = map[EnabledState]string{
	EnabledStateWifiDisabled:      "WifiDisabled",
	EnabledStateWifiEnabledS0:     "WifiEnabledS0",
	EnabledStateWifiEnabledS0SxAC: "WifiEnabledS0SxAC",
}

// enabledStateReverseMap is a reverse lookup map for EnabledState enumeration.
var enabledStateReverseMap = map[string]EnabledState{
	"WIFIDISABLED":      EnabledStateWifiDisabled,
	"WIFIENABLEDS0":     EnabledStateWifiEnabledS0,
	"WIFIENABLEDS0SXAC": EnabledStateWifiEnabledS0SxAC,
}

// String returns a human-readable string representation of the EnabledState enumeration.
func (e EnabledState) String() string {
	if s, ok := enabledStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ParseEnabledState returns the EnabledState enumeration value for a given string.
// The comparison is case-insensitive.
func ParseEnabledState(s string) (EnabledState, bool) {
	if state, ok := enabledStateReverseMap[strings.ToUpper(s)]; ok {
		return state, true
	}

	return 0, false
}

const (
	RequestedStateWifiDisabled      RequestedState = 3
	RequestedStateWifiEnabledS0     RequestedState = 32768
	RequestedStateWifiEnabledS0SxAC RequestedState = 32769
)

// requestedStateMap is a map of the RequestedState enumeration.
var requestedStateMap = map[RequestedState]string{
	RequestedStateWifiDisabled:      "WifiDisabled",
	RequestedStateWifiEnabledS0:     "WifiEnabledS0",
	RequestedStateWifiEnabledS0SxAC: "WifiEnabledS0SxAC",
}

// requestedStateReverseMap is a reverse lookup map for RequestedState enumeration.
var requestedStateReverseMap = map[string]RequestedState{
	"WIFIDISABLED":      RequestedStateWifiDisabled,
	"WIFIENABLEDS0":     RequestedStateWifiEnabledS0,
	"WIFIENABLEDS0SXAC": RequestedStateWifiEnabledS0SxAC,
}

// String returns a human-readable string representation of the RequestedState enumeration.
func (e RequestedState) String() string {
	if s, ok := requestedStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ParseRequestedState returns the RequestedState enumeration value for a given string.
// The comparison is case-insensitive.
func ParseRequestedState(s string) (RequestedState, bool) {
	if state, ok := requestedStateReverseMap[strings.ToUpper(s)]; ok {
		return state, true
	}

	return 0, false
}
