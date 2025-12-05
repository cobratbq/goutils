// SPDX-License-Identifier: LGPL-3.0-only

package sort

import "github.com/cobratbq/goutils/types"

// LessThan is less-than comparator for ordered types.
func LessThan[T types.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// GreaterThan is greater-than comparator for ordered types.
func GreaterThan[T types.Ordered](a, b T) int {
	if a > b {
		return -1
	}
	if a < b {
		return 1
	}
	return 0
}
