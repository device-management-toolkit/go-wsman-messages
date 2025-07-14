/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package http

const (
	ValueNotFound       = "Value not found in map"
	IPSHTTPProxyService = "IPS_HTTPProxyService"

	AddProxyAccessPoint = "AddProxyAccessPoint"
)

// InfoFormatToString is a map for converting InfoFormat values to their string representations.
var InfoFormatToString = map[InfoFormat]string{
	InfoFormatIPv4: "IPv4 Address",
	InfoFormatIPv6: "IPv6 Address",
	InfoFormatFQDN: "FQDN",
}

// String returns a human-readable string representation of the InfoFormat enumeration.
func (i InfoFormat) String() string {
	if value, exists := InfoFormatToString[i]; exists {
		return value
	}

	return ValueNotFound
}

// ReturnValueToString is a map for converting return values to their string representations.
var ReturnValueToString = map[int]string{
	PTStatusSuccess:          "PT_STATUS_SUCCESS",
	PTStatusInternalError:    "PT_STATUS_INTERNAL_ERROR",
	PTStatusNotPermitted:     "PT_STATUS_NOT_PERMITTED",
	PTStatusMaxLimitReached:  "PT_STATUS_MAX_LIMIT_REACHED",
	PTStatusInvalidParameter: "PT_STATUS_INVALID_PARAMETER",
	PTStatusDuplicate:        "PT_STATUS_DUPLICATE",
}

// GetReturnValueString returns a human-readable string representation of the return value.
func GetReturnValueString(returnValue int) string {
	if value, exists := ReturnValueToString[returnValue]; exists {
		return value
	}

	return ValueNotFound
}
