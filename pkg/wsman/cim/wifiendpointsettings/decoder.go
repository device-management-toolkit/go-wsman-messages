/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifiendpointsettings

import "strings"

const (
	CIMWiFiEndpointSettings string = "CIM_WiFiEndpointSettings"
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
