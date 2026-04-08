/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package eventlogentry

import "testing"

func TestResponseJSON(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: EventLogEntry{
				ElementName: "Event Log Entry",
				RecordID:    1,
				EventType:   15,
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
			GetResponse: EventLogEntry{
				ElementName: "Event Log Entry",
				RecordID:    1,
				EventType:   15,
			},
		},
	}

	yaml := response.YAML()
	if yaml == "" {
		t.Error("Expected YAML output, but got empty string")
	}
}
