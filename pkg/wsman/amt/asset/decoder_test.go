/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package asset

import "testing"

func TestResponseJSON(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{
				AssetTableItems: []AssetTable{
					{
						ElementName:   "Intel(r) AMT Asset Table",
						InstanceID:    "1",
						TableType:     131,
						TableTypeInfo: "SMbios",
						TableData:     "CgAAAEAAAAABAAAAPwAAAAAAAAAAAAAA+ACCUQAAAAAp4AAAAQAQAEwIGQAAAAAA/gD//wAAAAAAAAAA5gEAAHZQcm8AAAAA",
					},
				},
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
			PullResponse: PullResponse{
				AssetTableItems: []AssetTable{
					{
						ElementName:   "Intel(r) AMT Asset Table",
						InstanceID:    "1",
						TableType:     130,
						TableTypeInfo: "SMbios",
						TableData:     "CwAAABQAAAABAAAAAQEBAQGl7wLAPwEARQAABA==",
					},
				},
			},
		},
	}

	yaml := response.YAML()
	if yaml == "" {
		t.Error("Expected YAML output, but got empty string")
	}
}
