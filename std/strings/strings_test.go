// SPDX-License-Identifier: LGPL-3.0-only

package strings

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestFindOverlapping(t *testing.T) {
	text := "nonono"
	//count := FindOverlapping(text, "nono")
	//assert.Equal(t, 3, len(count))
	count := FindNonOverlapping(text, "nono")
	assert.Equal(t, 1, len(count))
}
