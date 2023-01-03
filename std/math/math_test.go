package math

import (
	"testing"

	t_ "github.com/cobratbq/goutils/std/testing"
)

func TestMaxInts(t *testing.T) {
	testdataInt := []struct {
		x, y, result int
	}{
		{0, 0, 0},
		{-1, -1, -1},
		{-1, 0, 0},
		{-1, 1, 1},
	}
	for _, d := range testdataInt {
		t_.Equal(t, d.result, Max(d.x, d.y))
		t_.Equal(t, d.result, Max(d.y, d.x))
		t_.Equal(t, Max(d.x, d.y), Max(d.y, d.x))
	}
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
