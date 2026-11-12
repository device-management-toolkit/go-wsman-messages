/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSidString(t *testing.T) {
	sid := string([]byte{1, 2, 0, 0, 0, 0, 0, 5, 21, 0, 0, 0, 1, 0, 0, 0})

	assert.Equal(t, "S-1-5-21-1", GetSidString(sid))
}

func TestGetSidStringShortInput(t *testing.T) {
	assert.NotPanics(t, func() {
		assert.Equal(t, "", GetSidString("\x01\x02"))
	})
}
