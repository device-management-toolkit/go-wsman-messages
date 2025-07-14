/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package http

import (
	"testing"
)

func TestInfoFormat_String(t *testing.T) {
	tests := []struct {
		state    InfoFormat
		expected string
	}{
		{InfoFormatIPv4, "IPv4 Address"},
		{InfoFormatIPv6, "IPv6 Address"},
		{InfoFormatFQDN, "FQDN"},
		{InfoFormat(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestGetReturnValueString(t *testing.T) {
	tests := []struct {
		returnValue int
		expected    string
	}{
		{PTStatusSuccess, "PT_STATUS_SUCCESS"},
		{PTStatusInternalError, "PT_STATUS_INTERNAL_ERROR"},
		{PTStatusNotPermitted, "PT_STATUS_NOT_PERMITTED"},
		{PTStatusMaxLimitReached, "PT_STATUS_MAX_LIMIT_REACHED"},
		{PTStatusInvalidParameter, "PT_STATUS_INVALID_PARAMETER"},
		{PTStatusDuplicate, "PT_STATUS_DUPLICATE"},
		{999, "Value not found in map"},
	}

	for _, test := range tests {
		result := GetReturnValueString(test.returnValue)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
