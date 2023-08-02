// SPDX-License-Identifier: LGPL-3.0-only

package set

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestContainsAll(t *testing.T) {
	testdata := []struct {
		a, b   map[uint]struct{}
		ab, ba bool
	}{
		{Create[uint](), Create[uint](), true, true},
		{Create[uint](), Create[uint](1, 3, 4, 6, 7), false, true},
		{Create[uint](1, 3, 4, 6, 7), Create[uint](1, 2, 3, 4, 5, 6, 7), false, true},
		{Create[uint](1, 3, 5), Create[uint](2), false, false},
		{Create[uint](1, 3, 5), Create[uint](1), true, false},
		{Create[uint](1, 3, 5), Create[uint](2, 4, 6), false, false},
	}
	for _, d := range testdata {
		assert.True(t, ContainsAll(d.a, d.a))
		assert.True(t, ContainsAll(d.b, d.b))
		assert.Equal(t, ContainsAll(d.a, d.b), d.ab)
		assert.Equal(t, ContainsAll(d.b, d.a), d.ba)
	}
}
