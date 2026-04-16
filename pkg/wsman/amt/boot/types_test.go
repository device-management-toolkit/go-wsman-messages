/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"testing"
)

func TestRPEParameterTypeValues(t *testing.T) {
	tests := []struct {
		name      string
		paramType ParameterType
		expected  ParameterType
	}{
		{name: "RPE_DEVICE_BITMASK", paramType: RPE_DEVICE_BITMASK, expected: 1},
		{name: "RPE_PSID", paramType: RPE_PSID, expected: 10},
		{name: "RPE_SSD_MASTER_PASSWORD", paramType: RPE_SSD_MASTER_PASSWORD, expected: 20},
		{name: "RPE_OEM_PARAMETER", paramType: RPE_OEM_PARAMETER, expected: 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.paramType != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, tt.paramType, tt.expected)
			}
		})
	}
}

func TestRPEMaxSizes(t *testing.T) {
	allRPETypes := []ParameterType{
		RPE_DEVICE_BITMASK,
		RPE_PSID,
		RPE_SSD_MASTER_PASSWORD,
		RPE_OEM_PARAMETER,
	}

	for _, paramType := range allRPETypes {
		if _, ok := RPEMaxSizes[paramType]; !ok {
			t.Errorf("RPEMaxSizes missing entry for ParameterType %v", paramType)
		}
	}

	expectedSizes := map[ParameterType]int{
		RPE_DEVICE_BITMASK:      4,
		RPE_PSID:                64,
		RPE_SSD_MASTER_PASSWORD: 64,
		RPE_OEM_PARAMETER:       500,
	}

	for paramType, size := range expectedSizes {
		if RPEMaxSizes[paramType] != size {
			t.Errorf("RPEMaxSizes[%v] = %v, want %v", paramType, RPEMaxSizes[paramType], size)
		}
	}
}

func TestRPEParameterNames(t *testing.T) {
	allRPETypes := []ParameterType{
		RPE_DEVICE_BITMASK,
		RPE_PSID,
		RPE_SSD_MASTER_PASSWORD,
		RPE_OEM_PARAMETER,
	}

	for _, paramType := range allRPETypes {
		name, ok := RPEParameterNames[paramType]
		if !ok {
			t.Errorf("RPEParameterNames missing entry for ParameterType %v", paramType)

			continue
		}

		if name == "" {
			t.Errorf("RPEParameterNames[%v] is empty", paramType)
		}
	}
}

func TestRPEParameterDetails(t *testing.T) {
	allRPETypes := []ParameterType{
		RPE_DEVICE_BITMASK,
		RPE_PSID,
		RPE_SSD_MASTER_PASSWORD,
		RPE_OEM_PARAMETER,
	}

	for _, paramType := range allRPETypes {
		detail, ok := RPEParameterDetails[paramType]
		if !ok {
			t.Errorf("RPEParameterDetails missing entry for ParameterType %v", paramType)

			continue
		}

		if detail.Comment == "" {
			t.Errorf("RPEParameterDetails[%v].Comment is empty", paramType)
		}
	}

	// RPE_DEVICE_BITMASK is the only mandatory parameter
	if !RPEParameterDetails[RPE_DEVICE_BITMASK].Mandatory {
		t.Errorf("RPEParameterDetails[RPE_DEVICE_BITMASK].Mandatory should be true")
	}

	for _, paramType := range []ParameterType{RPE_PSID, RPE_SSD_MASTER_PASSWORD, RPE_OEM_PARAMETER} {
		if RPEParameterDetails[paramType].Mandatory {
			t.Errorf("RPEParameterDetails[%v].Mandatory should be false", paramType)
		}
	}

	// RPE_DEVICE_BITMASK is the only parameter used in unconfigure
	if !RPEParameterDetails[RPE_DEVICE_BITMASK].UsedInUnconfigure {
		t.Errorf("RPEParameterDetails[RPE_DEVICE_BITMASK].UsedInUnconfigure should be true")
	}

	for _, paramType := range []ParameterType{RPE_PSID, RPE_SSD_MASTER_PASSWORD, RPE_OEM_PARAMETER} {
		if RPEParameterDetails[paramType].UsedInUnconfigure {
			t.Errorf("RPEParameterDetails[%v].UsedInUnconfigure should be false", paramType)
		}
	}
}
