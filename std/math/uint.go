package math

import "github.com/cobratbq/goutils/assert"

// LCM_uint is a unoptimized version of the Least/Lowest Common Multiple.
func LCM_uint(a, b uint) uint {
	return (a * b) / GCD_uint(a, b)
}

// GCD_uint is a unoptimized, inefficient version of the Greatest Common Divisor.
func GCD_uint(a, b uint) uint {
	assert.Require(a > 0, "Non-zero values required for this method of GCD calculation.")
	assert.Require(b > 0, "Non-zero values required for this method of GCD calculation.")
	if a == b {
		return a
	}
	if b > a {
		a, b = b, a
	}
	// Using Euclidian algorithm: `gcd(a, b) == gcd(a-b, b)` for `a > b`
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
