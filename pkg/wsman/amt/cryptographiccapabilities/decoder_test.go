/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package cryptographiccapabilities

import "testing"

func TestResponseJSON(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: CryptographicCapabilities{
				ElementName:          "Intel(r) AMT: Cryptographic Capabilities",
				HardwareAcceleration: 1,
				InstanceID:           "Intel(r) AMT: Cryptographic Capabilities 0",
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
			GetResponse: CryptographicCapabilities{
				ElementName:          "Intel(r) AMT: Cryptographic Capabilities",
				HardwareAcceleration: 1,
				InstanceID:           "Intel(r) AMT: Cryptographic Capabilities 0",
			},
		},
	}

	yaml := response.YAML()
	if yaml == "" {
		t.Error("Expected YAML output, but got empty string")
	}
}
