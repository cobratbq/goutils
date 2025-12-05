// SPDX-License-Identifier: LGPL-3.0-only

package ranges

import (
	"github.com/cobratbq/goutils/std/builtin/set"
	"github.com/cobratbq/goutils/std/builtin/slices"
	"github.com/cobratbq/goutils/std/math"
	"github.com/cobratbq/goutils/types"
)

// Len returns the length of a range.
func Len[E types.Number](range_ [2]E) int {
	return int(range_[1]-range_[0]) + 1
}

// Contains checks whether a value is contained in a range. (inclusive)
func Contains[E types.Number](range_ [2]E, val E) bool {
	return val >= range_[0] && val <= range_[1]
}

// Overlaps verify if two ranges overlap
func Overlaps[E types.Number](r1, r2 [2]E) bool {
	return (r1[0] <= r2[0] && r1[1] >= r2[0]) || (r2[0] <= r1[0] && r2[1] >= r1[0])
}

// Merge merges two overlapping ranges. (Ranges must overlap.)
func Merge[E types.Number](r1, r2 [2]E) [2]E {
	if r1[0] <= r2[0] && r1[1] >= r2[0] {
		return [2]E{r1[0], math.Max(r1[1], r2[1])}
	}
	if r2[0] <= r1[0] && r2[1] >= r1[0] {
		return [2]E{r2[0], math.Max(r1[1], r2[1])}
	}
	panic("ranges do not overlap")
}

// MergeAnyOverlapping merges any overlapping ranges of a provided slice of ranges.
func MergeAnyOverlapping[E types.Number](ranges [][2]E) [][2]E {
	initial, reduced := [][2]E{}, ranges
	for len(reduced) != len(initial) {
		initial, reduced = reduced, [][2]E{}
		merged := make(map[[2]E]struct{})
	next:
		for r := range initial {
			if set.Contains(merged, initial[r]) {
				continue
			}
			for o := r + 1; o < len(initial); o++ {
				if Overlaps(initial[r], initial[o]) {
					reduced = append(reduced, Merge(initial[r], initial[o]))
					set.Insert(merged, initial[o])
					continue next
				}
			}
			reduced = append(reduced, initial[r])
		}
	}
	return reduced
}

// Expand expands a range into its individidual elements. (inclusive)
func Expand[E types.Number](range_ [2]E) []E {
	elements := make([]E, 0, Len(range_))
	for r := range_[0]; r <= range_[1]; r++ {
		elements = slices.ExtendOne(elements, r)
	}
	return elements
}
