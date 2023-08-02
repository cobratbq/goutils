// SPDX-License-Identifier: LGPL-3.0-only

package modular

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestMod(t *testing.T) {
	testdata := []struct{ modulus, in, out int }{
		{1, 0, 0},
		{1, 1, 0},
		{2, 0, 0},
		{2, 1, 1},
		{2, 2, 0},
		{2, 3, 1},
		{2, 4, 0},
		{2, 5, 1},
		{4, 0, 0},
		{4, 1, 1},
		{4, 2, 2},
		{4, 3, 3},
		{4, 4, 0},
		{4, 5, 1},
		{4, 6, 2},
		{4, 7, 3},
		{4, 8, 0},
		{10, 0, 0},
		{10, 1, 1},
		{2, 3, 1},
		{2, -1, 1},
		{2, -3, 1},
		{3, 8, 2},
	}
	for _, d := range testdata {
		assert.Equal(t, d.out, Reduce(d.in, d.modulus))
	}
}

func TestModZero(t *testing.T) {
	defer assert.RequirePanic(t)
	Reduce(1, 0)
	t.FailNow()
}

func TestModNegative(t *testing.T) {
	defer assert.RequirePanic(t)
	Reduce(1, -1)
	t.FailNow()
}
