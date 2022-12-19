// SPDX-License-Identifier: LGPL-3.0-or-later

package math

import (
	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/builtin"
)

// Difference computes an absolute value that's the difference between `a` and `b`.
func Difference[N builtin.Number](a, b N) N {
	if a < b {
		return b - a
	}
	return a - b
}

// LCM is a unoptimized version of the Least/Lowest Common Multiple.
//
// Using the Greatest Common Divisor, following [1]:
// `lcm(a,b) = |ab| / gcd(a,b)`
//
// [1]: <https://en.wikipedia.org/wiki/Least_common_multiple>
func LCM[N builtin.UnsignedNumber](a, b N) N {
	// NOTE: first divide `b` by `GCD` to prevent multiplication from greatly increasing the
	// intermediate result with risk of overflowing.
	return a * (b / GCD(a, b))
}

// GCD is a unoptimized, inefficient version of the Greatest Common Divisor.
//
// Using Euclidian algorithm, following [1]:
// `gcd(a, b) == gcd(b, a mod b)` for `a > b`. Perform repeatedly until `gcd(x, 0)`, at which point
// `x` is greatest common divisor for `gcd(a,b)`.
//
// [1]: <https://en.wikipedia.org/wiki/Greatest_common_divisor>
func GCD[N builtin.UnsignedNumber](a, b N) N {
	assert.Require(a > 0, "Non-zero values required for this method of GCD calculation.")
	assert.Require(b > 0, "Non-zero values required for this method of GCD calculation.")
	if a == b {
		return a
	}
	if b > a {
		// Make sure a is the larger of the two uints
		b = b % a
	}
	// Using Euclidian algorithm: `gcd(a, b) == gcd(a-b, b)` for `a > b`
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

// AbsInt determines the absolute value of provided integer.
func AbsInt[N builtin.SignedNumber](n N) N {
	if n < 0 {
		return -n
	}
	return n
}

// Max returns the maximum of two values.
func Max[N builtin.Number](x, y N) N {
	if x > y {
		return x
	}
	return y
}

// Min returns the minimum of two values.
func Min[N builtin.Number](x, y N) N {
	if x < y {
		return x
	}
	return y
}

// Sign returns the sign of the provided value: `1` for positive value, `0` for zero, `-1` for
// negative value.
func Sign[N builtin.Number](x N) int {
	if x > 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	return -1
}
