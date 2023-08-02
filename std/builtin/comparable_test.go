// SPDX-License-Identifier: LGPL-3.0-only

package builtin

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestEqualsAny(t *testing.T) {
	assert.True(t, EqualsAny(2, 0, 1, 2))
	assert.False(t, EqualsAny(99, 0, 1, 2))
}
