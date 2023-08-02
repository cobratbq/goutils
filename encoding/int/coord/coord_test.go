// SPDX-License-Identifier: LGPL-3.0-only

package coord

import (
	"testing"

	"github.com/cobratbq/goutils/assert"
	t_ "github.com/cobratbq/goutils/std/testing"
)

func TestEncode2D(t *testing.T) {
	testdata := []struct {
		Length int
		X      int
		Y      int
		Result int
	}{
		{1, 0, 0, 0},
		{2, 1, 0, 1},
		{2, 1, 1, 3},
		{10, 5, 3, 35},
	}

	for _, e := range testdata {
		// verify Encode2D
		assert.Equal(e.Result, Encode2D(e.Length, e.X, e.Y))
		// verify Decode2D
		testX, testY := Decode2D(e.Length, e.Result)
		assert.Equal(e.X, testX)
		assert.Equal(e.Y, testY)
		// verify repeated conversion (Encode2D-Decode2D-Encode2D)
		tempIndex := Encode2D(e.Length, e.X, e.Y)
		tempX, tempY := Decode2D(e.Length, tempIndex)
		index := Encode2D(e.Length, tempX, tempY)
		assert.Equal(e.Result, index)
	}
}

func TestEncode2DXTooLarge(t *testing.T) {
	defer t_.RequirePanic(t)
	Encode2D(1, 1, 1)
}

func TestDecode2DInvalidLength(t *testing.T) {
	defer t_.RequirePanic(t)
	Decode2D(0, 0)
}

func TestDecode2DUintInvalidLength(t *testing.T) {
	defer t_.RequirePanic(t)
	Decode2DUint(0, 0)
}
