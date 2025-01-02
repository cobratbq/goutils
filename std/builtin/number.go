// SPDX-License-Identifier: LGPL-3.0-only

package builtin

import "github.com/cobratbq/goutils/types"

// Negate negates the provided value.
func Negate[T types.Signed](v T) T {
	return -v
}

// Zero tests whether provided value `v` is zero.
func Zero[T types.Number](v T) bool {
	return v == 0
}

// NonZero tests whether provided value `v` is non-zero.
func NonZero[T types.Number](v T) bool {
	return v != 0
}

// Add is a trivial func. It does not protect its boundaries, e.g. overflowing.
// It can be referenced, for example in `Reduce`.
func Add[T types.Number](a, b T) T {
	return a + b
}

// Multiply is a trivial func. It does not protect its boundaries, e.g. overflowing.
// It can be referenced, for example in `Reduce`.
func Multiply[T types.Number](a, b T) T {
	return a * b
}
