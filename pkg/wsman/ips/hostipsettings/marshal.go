/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostipsettings

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// JSON marshals the type into JSON format.
func (r *Response) JSON() string {
	jsonOutput, err := json.Marshal(r.Body)
	if err != nil {
		return ""
	}

	return string(jsonOutput)
}

// YAML marshals the type into YAML format.
func (r *Response) YAML() string {
	yamlOutput, err := yaml.Marshal(r.Body)
	if err != nil {
		return ""
	}

	return string(yamlOutput)
}
