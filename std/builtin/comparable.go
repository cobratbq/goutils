// SPDX-License-Identifier: LGPL-3.0-only

package builtin

import "github.com/cobratbq/goutils/types"

// EqualsFixedOf creates a closure that tests any provided value against a fixed value.
func EqualsFixedOf[T comparable](fixed T) func(T) bool {
	return func(v T) bool {
		return v == fixed
	}
}

// EqualsAny matches the specified value with any of the provided `matches` values. It returns
// `true` if it is any of the provided matches, or `false` if none match.
func EqualsAny[T comparable](value T, matches ...T) bool {
	for _, m := range matches {
		if value == m {
			return true
		}
	}
	return false
}

// EqualsAnyOf creates a closure that tests any provided value against a list of matches.
func EqualsAnyOf[T comparable](matches ...T) func(v T) bool {
	return func(v T) bool {
		return EqualsAny(v, matches...)
	}
}

func EqualT[T types.Equaler[T]](a, b T) bool {
	return a.Equal(b)
}
