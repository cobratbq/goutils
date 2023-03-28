// SPDX-License-Identifier: AGPL-3.0-or-later

package math

import (
	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/types"
)

// Difference computes an absolute value that's the difference between `a` and `b`.
func Difference[N types.Integer](a, b N) N {
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
func LCM[N types.UnsignedInteger](a, b N) N {
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
func GCD[N types.UnsignedInteger](a, b N) N {
	assert.Require(a > 0, "Non-zero value a required for this method of GCD calculation.")
	assert.Require(b > 0, "Non-zero value b required for this method of GCD calculation.")
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
func AbsInt[N types.SignedInteger](n N) N {
	if n < 0 {
		return -n
	}
	return n
}

// Max returns the maximum of two values.
func Max[N types.Integer](x, y N) N {
	if x > y {
		return x
	}
	return y
}

// MaxN returns the maximum of vararg provided number of values. At least one value must be provided
// or the function will panic.
func MaxN[N types.Integer](x ...N) N {
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

// Min returns the minimum of two values.
func Min[N types.Integer](x, y N) N {
	if x < y {
		return x
	}
	return y
}

// MinN returns the minimum of vararg provided number of values.
func MinN[N types.Integer](x ...N) N {
	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return min
}

// Sign returns the sign of the provided value: `1` for positive value, `0` for zero, `-1` for
// negative value.
func Sign[N types.Integer](x N) int {
	if x > 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	return -1
}

// PowUint calculates a power of a given base for an arbitrary base value.
//
// NOTE: this function is not optimized for very large calculations and care should be taken to
// choose a data-type with sufficient space to contain the result.
func PowUint[U types.UnsignedInteger](base, power U) U {
	if power == 0 {
		return 1
	}
	result := base
	for i := U(1); i < power; i++ {
		result *= base
	}
	return result
}
