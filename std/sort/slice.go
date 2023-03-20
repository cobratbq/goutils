// SPDX-License-Identifier: AGPL-3.0-or-later

package sort

import (
	"sort"

	"github.com/cobratbq/goutils/types"
)

// Slice is the generic variant for general slice sorting.
func Slice[E any](vals []E, less func(i, j int) bool) {
	sort.Slice(vals, less)
}

// Number sorts slice of any numeric type (signed/unsigned, integer/float).
func Number[E types.Number](vals []E) {
	Ordered(vals)
}

// String sorts a slice of any type rooted in `string`.
func String[E ~string](vals []E) {
	Ordered(vals)
}

// Ordered sorts in-place a slice of any ordered-type elements, including any numbers and strings.
func Ordered[E types.Ordered](vals []E) {
	sort.Slice(vals, func(i, j int) bool { return vals[i] < vals[j] })
}

// IsSorted checks if any ordered-typed slice is sorted.
func IsSorted[E types.Ordered](vals []E) bool {
	for i := 1; i < len(vals); i++ {
		if vals[i-1] > vals[i] {
			return false
		}
	}
	return true
}
