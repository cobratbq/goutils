// SPDX-License-Identifier: LGPL-3.0-only

package multiset

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestOperationsn(t *testing.T) {
	empty := make(map[int]uint, 0)
	testdata := []struct {
		a, b, union, intersection, sum, difference map[int]uint
	}{
		{empty, empty, empty, empty, empty, empty},
		{map[int]uint{1: 12, 2: 24, 3: 36}, map[int]uint{1: 12, 2: 24, 3: 36},
			map[int]uint{1: 12, 2: 24, 3: 36}, map[int]uint{1: 12, 2: 24, 3: 36}, map[int]uint{1: 24, 2: 48, 3: 72}, empty},
		{map[int]uint{1: 12, 2: 24, 3: 36, 4: 8}, map[int]uint{1: 16, 2: 10, 4: 8},
			map[int]uint{1: 16, 2: 24, 3: 36, 4: 8}, map[int]uint{1: 12, 2: 10, 4: 8}, map[int]uint{1: 28, 2: 34, 3: 36, 4: 16}, map[int]uint{2: 14, 3: 36}},
	}
	for _, d := range testdata {
		assert.True(t, Equal(d.union, Union(d.a, d.b)))
		assert.True(t, Equal(d.intersection, Intersection(d.a, d.b)))
		assert.True(t, Equal(d.sum, Sum(d.a, d.b)))
		assert.True(t, Equal(d.difference, Difference(d.a, d.b)))
	}
}

func TestDisjointSumIsUnion(t *testing.T) {
	empty := make(map[int]uint, 0)
	testdata := []struct {
		a, b map[int]uint
	}{
		{empty, empty},
		{map[int]uint{1: 12, 2: 24}, map[int]uint{3: 36, 4: 3, 5: 1}},
	}
	for _, d := range testdata {
		assert.True(t, Equal(Union(d.a, d.b), Sum(d.a, d.b)))
	}
}

func TestDisjointIntersectionIsEmpty(t *testing.T) {
	empty := make(map[int]uint, 0)
	testdata := []struct {
		a, b map[int]uint
	}{
		{empty, empty},
		{map[int]uint{1: 12, 2: 24}, empty},
		{empty, map[int]uint{1: 12, 2: 24}},
		{map[int]uint{1: 12, 2: 24}, map[int]uint{3: 36, 4: 3}},
	}
	for _, d := range testdata {
		assert.True(t, Equal(Intersection(d.a, d.b), empty))
	}
}
