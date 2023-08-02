// SPDX-License-Identifier: LGPL-3.0-only

// modular provides arithmetic functions for modular arithmetic. This helps to avoid gotchas of Go
// remainder operator such that modulo may (unintentionally) return negative values, which is often
// not desirable.
// (Some of) these functions are not perfectly efficient. The priority is in avoiding mistakes. For
// very efficient modular arithmetic you will want to use the remainder operator and perform modulo
// operations only if strictly necessary within your value-range.
//
// References:
// - <https://en.wikipedia.org/wiki/Modular_arithmetic>
package modular

import (
	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/types"
)

// Increment increments a value mod `modulus`.
func Increment[T types.Integer](value, modulus T) T {
	assert.Require(value >= 0 && value < modulus, "Provided value is not within the modular domain")
	return (value + 1) % modulus
}

// IncrementN increments a value by `n` mod `modulus`.
func IncrementN[T types.Integer](value, n, modulus T) T {
	assert.Require(value >= 0 && value < modulus, "Provided value is not within the modular domain")
	return (value + n) % modulus
}

// Decrement decrements a value mod `modulus`. It will ensure a non-negative value.
func Decrement[T types.Integer](value, modulus T) T {
	assert.Require(value >= 0 && value < modulus, "Provided value is not within the modular domain")
	// TODO consider branching to avoid always adding modulo, risking overflowing.
	return (value + modulus - 1) % modulus
}

// DecrementN decrements a value by `n` mod `modulus`. It will ensure a non-negative value.
func DecrementN[T types.Integer](value, n, modulus T) T {
	assert.Require(value >= 0 && value < modulus, "Provided value is not within the modular domain")
	// TODO consider branching to avoid always adding modulo, risking overflowing.
	return (value + modulus - n) % modulus
}

// Reduce performs modular reduction operation on `value` using `modulus`. `modulus` must be a
// positive integer value.
func Reduce[T types.Integer](value, modulus T) T {
	assert.Require(modulus > 0, "The modulus must be a positive integer.")
	if value >= 0 {
		return value % modulus
	}
	return (value % modulus) + modulus
}
