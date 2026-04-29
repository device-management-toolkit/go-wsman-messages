/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package systempowerscheme

import "testing"

func TestResponseJSON(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: SystemPowerScheme{
				Description: "Mobile: ON in S0",
				ElementName: "Intel(r) AMT Power Scheme",
				InstanceID:  "SCHEME 0",
				SchemeGUID:  "djmXEQtWUEOIcJgS85G1YA==",
			},
		},
	}

	json := response.JSON()
	if json == "" {
		t.Error("Expected JSON output, but got empty string")
	}
}

func TestResponseYAML(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: SystemPowerScheme{
				Description: "Mobile: ON in S0",
				ElementName: "Intel(r) AMT Power Scheme",
				InstanceID:  "SCHEME 0",
				SchemeGUID:  "djmXEQtWUEOIcJgS85G1YA==",
			},
		},
	}

	yaml := response.YAML()
	if yaml == "" {
		t.Error("Expected YAML output, but got empty string")
	}
}
