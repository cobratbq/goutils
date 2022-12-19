// SPDX-License-Identifier: LGPL-3.0-or-later

package math

import "github.com/cobratbq/goutils/assert"

type number interface {
	signedNumber | unsignedNumber
}

type signedNumber interface {
	int | int8 | int16 | int32 | int64
}

type unsignedNumber interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// LCM is a unoptimized version of the Least/Lowest Common Multiple.
//
// Using the Greatest Common Divisor, following [1]:
// `lcm(a,b) = |ab| / gcd(a,b)`
//
// [1]: <https://en.wikipedia.org/wiki/Least_common_multiple>
func LCM[N unsignedNumber](a, b N) N {
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
func GCD[N unsignedNumber](a, b N) N {
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
