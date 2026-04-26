/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hdr8021filter

import "testing"

func TestResponseJSON(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: Hdr8021Filter{
				ElementName:   "Hdr 802.1 Filter",
				VLANPriority:  1,
				VLANID:        10,
				FilterEnabled: true,
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
			GetResponse: Hdr8021Filter{
				ElementName:   "Hdr 802.1 Filter",
				VLANPriority:  1,
				VLANID:        10,
				FilterEnabled: true,
			},
		},
	}

	yaml := response.YAML()
	if yaml == "" {
		t.Error("Expected YAML output, but got empty string")
	}
}
