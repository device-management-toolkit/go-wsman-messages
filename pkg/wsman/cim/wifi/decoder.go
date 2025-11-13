/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import "strings"

const (
	CIMWiFiEndpoint         string = "CIM_WiFiEndpoint"
	CIMWiFiEndpointSettings string = "CIM_WiFiEndpointSettings"
	CIMWiFiPort             string = "CIM_WiFiPort"
	ValueNotFound           string = "Value not found in map"
)

const (
	AuthenticationMethodOther AuthenticationMethod = iota + 1
	AuthenticationMethodOpenSystem
	AuthenticationMethodSharedKey
	AuthenticationMethodWPAPSK
	AuthenticationMethodWPAIEEE8021x
	AuthenticationMethodWPA2PSK
	AuthenticationMethodWPA2IEEE8021x
	AuthenticationMethodWPA3SAE AuthenticationMethod = 32768
	AuthenticationMethodWPA3OWE AuthenticationMethod = 32769
)

// authenticationMethodMap is a map of the AuthenticationMethod enumeration.
var authenticationMethodMap = map[AuthenticationMethod]string{
	AuthenticationMethodOther:         "Other",
	AuthenticationMethodOpenSystem:    "OpenSystem",
	AuthenticationMethodSharedKey:     "SharedKey",
	AuthenticationMethodWPAPSK:        "WPAPSK",
	AuthenticationMethodWPAIEEE8021x:  "WPAIEEE8021x",
	AuthenticationMethodWPA2PSK:       "WPA2PSK",
	AuthenticationMethodWPA2IEEE8021x: "WPA2IEEE8021x",
	AuthenticationMethodWPA3SAE:       "WPA3SAE",
	AuthenticationMethodWPA3OWE:       "WPA3OWE",
}

// authenticationMethodReverseMap is a reverse lookup map for AuthenticationMethod enumeration.
var authenticationMethodReverseMap = map[string]AuthenticationMethod{
	"OTHER":         AuthenticationMethodOther,
	"OPENSYSTEM":    AuthenticationMethodOpenSystem,
	"SHAREDKEY":     AuthenticationMethodSharedKey,
	"WPAPSK":        AuthenticationMethodWPAPSK,
	"WPAIEEE8021X":  AuthenticationMethodWPAIEEE8021x,
	"WPA2PSK":       AuthenticationMethodWPA2PSK,
	"WPA2IEEE8021X": AuthenticationMethodWPA2IEEE8021x,
	"WPA3SAE":       AuthenticationMethodWPA3SAE,
	"WPA3OWE":       AuthenticationMethodWPA3OWE,
}

// String returns a human-readable string representation of the AuthenticationMethod enumeration.
func (e AuthenticationMethod) String() string {
	if s, ok := authenticationMethodMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ParseAuthenticationMethod returns the AuthenticationMethod enumeration value for a given string.
// The comparison is case-insensitive.
func ParseAuthenticationMethod(s string) (AuthenticationMethod, bool) {
	if method, ok := authenticationMethodReverseMap[strings.ToUpper(s)]; ok {
		return method, true
	}

	return 0, false
}

const (
	BSSTypeUnknown        BSSType = 0
	BSSTypeIndependent    BSSType = 2
	BSSTypeInfrastructure BSSType = 3
)

// bssTypeMap is a map of the BSSType enumeration.
var bssTypeMap = map[BSSType]string{
	BSSTypeUnknown:        "Unknown",
	BSSTypeIndependent:    "Independent",
	BSSTypeInfrastructure: "Infrastructure",
}

// bssTypeReverseMap is a reverse lookup map for BSSType enumeration.
var bssTypeReverseMap = map[string]BSSType{
	"UNKNOWN":        BSSTypeUnknown,
	"INDEPENDENT":    BSSTypeIndependent,
	"INFRASTRUCTURE": BSSTypeInfrastructure,
}

// String returns a human-readable string representation of the BSSType enumeration.
func (e BSSType) String() string {
	if s, ok := bssTypeMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ParseBSSType returns the BSSType enumeration value for a given string.
// The comparison is case-insensitive.
func ParseBSSType(s string) (BSSType, bool) {
	if bssType, ok := bssTypeReverseMap[strings.ToUpper(s)]; ok {
		return bssType, true
	}

	return 0, false
}

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
	EncryptionMethodOther EncryptionMethod = iota + 1
	EncryptionMethodWEP
	EncryptionMethodTKIP
	EncryptionMethodCCMP
	EncryptionMethodNone
)

// encryptionMethodMap is a map of the EncryptionMethod enumeration.
var encryptionMethodMap = map[EncryptionMethod]string{
	EncryptionMethodOther: "Other",
	EncryptionMethodWEP:   "WEP",
	EncryptionMethodTKIP:  "TKIP",
	EncryptionMethodCCMP:  "CCMP",
	EncryptionMethodNone:  "None",
}

// encryptionMethodReverseMap is a reverse lookup map for EncryptionMethod enumeration.
var encryptionMethodReverseMap = map[string]EncryptionMethod{
	"OTHER": EncryptionMethodOther,
	"WEP":   EncryptionMethodWEP,
	"TKIP":  EncryptionMethodTKIP,
	"CCMP":  EncryptionMethodCCMP,
	"NONE":  EncryptionMethodNone,
}

// String returns a human-readable string representation of the EncryptionMethod enumeration.
func (e EncryptionMethod) String() string {
	if s, ok := encryptionMethodMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ParseEncryptionMethod returns the EncryptionMethod enumeration value for a given string.
// The comparison is case-insensitive.
func ParseEncryptionMethod(s string) (EncryptionMethod, bool) {
	if method, ok := encryptionMethodReverseMap[strings.ToUpper(s)]; ok {
		return method, true
	}

	return 0, false
}

const (
	HealthStateUnknown             HealthState = 0
	HealthStateOK                  HealthState = 5
	HealthStateDegraded            HealthState = 10
	HealthStateMinorFailure        HealthState = 15
	HealthStateMajorFailure        HealthState = 20
	HealthStateCriticalFailure     HealthState = 25
	HealthStateNonRecoverableError HealthState = 30
)

// healthStateMap is a map of the HealthState enumeration.
var healthStateMap = map[HealthState]string{
	HealthStateUnknown:             "Unknown",
	HealthStateOK:                  "OK",
	HealthStateDegraded:            "Degraded",
	HealthStateMinorFailure:        "MinorFailure",
	HealthStateMajorFailure:        "MajorFailure",
	HealthStateCriticalFailure:     "CriticalFailure",
	HealthStateNonRecoverableError: "NonRecoverableError",
}

// healthStateReverseMap is a reverse lookup map for HealthState enumeration.
var healthStateReverseMap = map[string]HealthState{
	"UNKNOWN":             HealthStateUnknown,
	"OK":                  HealthStateOK,
	"DEGRADED":            HealthStateDegraded,
	"MINORFAILURE":        HealthStateMinorFailure,
	"MAJORFAILURE":        HealthStateMajorFailure,
	"CRITICALFAILURE":     HealthStateCriticalFailure,
	"NONRECOVERABLEERROR": HealthStateNonRecoverableError,
}

// String returns a human-readable string representation of the HealthState enumeration.
func (e HealthState) String() string {
	if s, ok := healthStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ParseHealthState returns the HealthState enumeration value for a given string.
// The comparison is case-insensitive.
func ParseHealthState(s string) (HealthState, bool) {
	if state, ok := healthStateReverseMap[strings.ToUpper(s)]; ok {
		return state, true
	}

	return 0, false
}

const (
	LinkTechnologyUnknown LinkTechnology = iota
	LinkTechnologyOther
	LinkTechnologyEthernet
	LinkTechnologyIB
	LinkTechnologyFC
	LinkTechnologyFDDI
	LinkTechnologyATM
	LinkTechnologyTokenRing
	LinkTechnologyFrameRelay
	LinkTechnologyInfrared
	LinkTechnologyBlueTooth
	LinkTechnologyWirelessLAN
)

// linkTechnologyMap is a map of the LinkTechnology enumeration.
var linkTechnologyMap = map[LinkTechnology]string{
	LinkTechnologyUnknown:     "Unknown",
	LinkTechnologyOther:       "Other",
	LinkTechnologyEthernet:    "Ethernet",
	LinkTechnologyIB:          "IB",
	LinkTechnologyFC:          "FC",
	LinkTechnologyFDDI:        "FDDI",
	LinkTechnologyATM:         "ATM",
	LinkTechnologyTokenRing:   "TokenRing",
	LinkTechnologyFrameRelay:  "FrameRelay",
	LinkTechnologyInfrared:    "Infrared",
	LinkTechnologyBlueTooth:   "BlueTooth",
	LinkTechnologyWirelessLAN: "WirelessLAN",
}

// linkTechnologyReverseMap is a reverse lookup map for LinkTechnology enumeration.
var linkTechnologyReverseMap = map[string]LinkTechnology{
	"UNKNOWN":     LinkTechnologyUnknown,
	"OTHER":       LinkTechnologyOther,
	"ETHERNET":    LinkTechnologyEthernet,
	"IB":          LinkTechnologyIB,
	"FC":          LinkTechnologyFC,
	"FDDI":        LinkTechnologyFDDI,
	"ATM":         LinkTechnologyATM,
	"TOKENRING":   LinkTechnologyTokenRing,
	"FRAMERELAY":  LinkTechnologyFrameRelay,
	"INFRARED":    LinkTechnologyInfrared,
	"BLUETOOTH":   LinkTechnologyBlueTooth,
	"WIRELESSLAN": LinkTechnologyWirelessLAN,
}

// String returns a human-readable string representation of the LinkTechnology enumeration.
func (e LinkTechnology) String() string {
	if s, ok := linkTechnologyMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ParseLinkTechnology returns the LinkTechnology enumeration value for a given string.
// The comparison is case-insensitive.
func ParseLinkTechnology(s string) (LinkTechnology, bool) {
	if tech, ok := linkTechnologyReverseMap[strings.ToUpper(s)]; ok {
		return tech, true
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

const (
	CompletedWithNoError ReturnValue = iota
	NotSupported
	UnknownOrUnspecifiedError
	CannotCompleteWithinTimeoutPeriod
	Failed
	InvalidParameter
	InUse
	MethodParametersCheckedJobStarted ReturnValue = 4096
	InvalidStateTransition            ReturnValue = 4097
	UseOfTimeoutParameterNotSupported ReturnValue = 4098
	Busy                              ReturnValue = 4099
)

// returnValueMap is a map of the ReturnValue enumeration.
var returnValueMap = map[ReturnValue]string{
	CompletedWithNoError:              "CompletedWithNoError",
	NotSupported:                      "NotSupported",
	UnknownOrUnspecifiedError:         "UnknownOrUnspecifiedError",
	CannotCompleteWithinTimeoutPeriod: "CannotCompleteWithinTimeoutPeriod",
	Failed:                            "Failed",
	InvalidParameter:                  "InvalidParameter",
	InUse:                             "InUse",
	MethodParametersCheckedJobStarted: "MethodParametersCheckedJobStarted",
	InvalidStateTransition:            "InvalidStateTransition",
	UseOfTimeoutParameterNotSupported: "UseOfTimeoutParameterNotSupported",
	Busy:                              "Busy",
}

// returnValueReverseMap is a reverse lookup map for ReturnValue enumeration.
var returnValueReverseMap = map[string]ReturnValue{
	"COMPLETEDWITHNOERROR":              CompletedWithNoError,
	"NOTSUPPORTED":                      NotSupported,
	"UNKNOWNORUNSPECIFIEDERROR":         UnknownOrUnspecifiedError,
	"CANNOTCOMPLETEWITHINTIMEOUTPERIOD": CannotCompleteWithinTimeoutPeriod,
	"FAILED":                            Failed,
	"INVALIDPARAMETER":                  InvalidParameter,
	"INUSE":                             InUse,
	"METHODPARAMETERSCHECKEDJOBSTARTED": MethodParametersCheckedJobStarted,
	"INVALIDSTATETRANSITION":            InvalidStateTransition,
	"USEOFTIMEOUTPARAMETERNOTSUPPORTED": UseOfTimeoutParameterNotSupported,
	"BUSY":                              Busy,
}

// String returns a human-readable string representation of the ReturnValue enumeration.
func (e ReturnValue) String() string {
	if s, ok := returnValueMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ParseReturnValue returns the ReturnValue enumeration value for a given string.
// The comparison is case-insensitive.
func ParseReturnValue(s string) (ReturnValue, bool) {
	if value, ok := returnValueReverseMap[strings.ToUpper(s)]; ok {
		return value, true
	}

	return 0, false
}

const (
	PortTypeUnknown PortType = 0
	PortTypeOther   PortType = 1
	PortType80211a  PortType = 70
	PortType80211b  PortType = 71
	PortType80211g  PortType = 72
	PortType80211n  PortType = 73
)

// portTypeMap is a map of the PortType enumeration.
var portTypeMap = map[PortType]string{
	PortTypeUnknown: "Unknown",
	PortTypeOther:   "Other",
	PortType80211a:  "802.11a",
	PortType80211b:  "802.11b",
	PortType80211g:  "802.11g",
	PortType80211n:  "802.11n",
}

// portTypeReverseMap is a reverse lookup map for PortType enumeration.
var portTypeReverseMap = map[string]PortType{
	"UNKNOWN": PortTypeUnknown,
	"OTHER":   PortTypeOther,
	"802.11A": PortType80211a,
	"802.11B": PortType80211b,
	"802.11G": PortType80211g,
	"802.11N": PortType80211n,
}

// String returns a human-readable string representation of the PortType enumeration.
func (e PortType) String() string {
	if s, ok := portTypeMap[e]; ok {
		return s
	}

	return ValueNotFound
}

// ParsePortType returns the PortType enumeration value for a given string.
// The comparison is case-insensitive.
func ParsePortType(s string) (PortType, bool) {
	if portType, ok := portTypeReverseMap[strings.ToUpper(s)]; ok {
		return portType, true
	}

	return 0, false
}
