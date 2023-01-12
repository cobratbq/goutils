// SPDX-License-Identifier: AGPL-3.0-or-later

package math

import (
	"testing"

	t_ "github.com/cobratbq/goutils/std/testing"
	"github.com/cobratbq/goutils/types"
)

func TestMaxInts(t *testing.T) {
	testdata := []struct {
		x, y, result int
	}{
		{0, 0, 0},
		{-1, -1, -1},
		{-1, 0, 0},
		{-1, 1, 1},
		{99, -99, 99},
	}
	for _, d := range testdata {
		t_.Equal(t, d.result, Max(d.x, d.y))
		t_.Equal(t, d.result, Max(d.y, d.x))
		t_.Equal(t, Max(d.x, d.y), Max(d.y, d.x))
	}
}

func TestMaxNInts(t *testing.T) {
	testdata := []struct {
		v      []int
		result int
	}{
		{[]int{0, 0, 0, 0, 0, 0, 0}, 0},
		{[]int{-1, -1, -1}, -1},
		{[]int{-1, 0}, 0},
		{[]int{-1, 1}, 1},
		{[]int{99, -99}, 99},
		{[]int{types.MaxInt, types.MinInt}, types.MaxInt},
		{[]int{types.MaxInt - 3, types.MaxInt - 1, types.MaxInt - 2, types.MaxInt, types.MaxInt - 6}, types.MaxInt},
	}
	for _, d := range testdata {
		t_.Equal(t, d.result, MaxN(d.v...))
		t_.Equal(t, d.result, MaxN(d.v...))
		t_.Equal(t, MaxN(d.v...), MaxN(d.v...))
	}
}

func TestMaxNIntsEmpty(t *testing.T) {
	defer t_.RequirePanic(t)
	MaxN([]int{}...)
}

func TestMaxUints(t *testing.T) {
	testdataInt := []struct {
		x, y, result uint
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{0, 99, 99},
		{1, 2, 2},
	}
	for _, d := range testdataInt {
		t_.Equal(t, d.result, Max(d.x, d.y))
		t_.Equal(t, d.result, Max(d.y, d.x))
		t_.Equal(t, Max(d.x, d.y), Max(d.y, d.x))
	}
}

func TestMinNEmpty(t *testing.T) {
	defer t_.RequirePanic(t)
	MinN([]int{}...)
}
