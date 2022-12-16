package math

import "github.com/cobratbq/goutils/assert"

// LCM_uint is a unoptimized version of the Least/Lowest Common Multiple.
//
// Using the Greatest Common Divisor, following [1]:
// `lcm(a,b) = |ab| / gcd(a,b)`
//
// [1]: <https://en.wikipedia.org/wiki/Least_common_multiple>
func LCM_uint(a, b uint) uint {
	return (a * b) / GCD_uint(a, b)
}

// GCD_uint is a unoptimized, inefficient version of the Greatest Common Divisor.
//
// Using Euclidian algorithm, following [1]:
// `gcd(a, b) == gcd(b, a mod b)` for `a > b`. Perform repeatedly until `gcd(x, 0)`, at which point
// `x` is greatest common divisor for `gcd(a,b)`.
//
// [1]: <https://en.wikipedia.org/wiki/Greatest_common_divisor>
func GCD_uint(a, b uint) uint {
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
