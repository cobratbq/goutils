// SPDX-License-Identifier: LGPL-3.0-only

package slices

import (
	"bytes"
	"testing"

	"github.com/cobratbq/goutils/std/crypto/rand"
	assert "github.com/cobratbq/goutils/std/testing"
)

func TestSliceReversing(t *testing.T) {
	var random [33]byte
	rand.MustReadBytes(random[:])
	revved := Reversed(random[:])
	assert.False(t, bytes.Equal(revved, random[:]))
	revrevved := Reversed(revved)
	assert.SlicesEqual(t, random[:], revrevved)
	Reverse(random[:])
	assert.SlicesEqual(t, revved, random[:])
}

func TestUniformDimensions2D(t *testing.T) {
	assert.True(t, UniformDimensions2D([][]uint{}))
	assert.True(t, UniformDimensions2D([][]uint{{}, {}, {}}))
	assert.True(t, UniformDimensions2D([][]uint{{1}, {2}, {3}}))
	assert.False(t, UniformDimensions2D([][]uint{{1}, {2}, {}}))
	assert.False(t, UniformDimensions2D([][]uint{{1, 3, 4}, {2, 3}, {}}))
}
