package math

import (
	"testing"

	"github.com/cobratbq/goutils/std/builtin"
	t_ "github.com/cobratbq/goutils/std/testing"
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
		{[]int{builtin.MaxInt, builtin.MinInt}, builtin.MaxInt},
		{[]int{builtin.MaxInt - 3, builtin.MaxInt - 1, builtin.MaxInt - 2, builtin.MaxInt, builtin.MaxInt - 6}, builtin.MaxInt},
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
