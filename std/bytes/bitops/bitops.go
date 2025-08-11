// SPDX-License-Identifier: LGPL-3.0-only

package bitops

import (
	"unsafe"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/types"
)

// XorInto, XOR the contents of two byte-slices of same length, write result into first parameter.
// `inout` will contain `inout[i] ^ b[i]`, for all `i`.
func XorInto[T types.Integer](inout, b []T) {
	assert.Equal(len(inout), len(b))
	for i := range inout {
		inout[i] ^= b[i]
	}
}

// Xor, XOR the contents of `a ^ b`.
// Returns the result, leaving `a` and `b` unmodified.
func Xor[T types.Integer](a, b []T) []T {
	assert.Equal(len(a), len(b))
	out := make([]T, len(a))
	for i := range out {
		out[i] = a[i] ^ b[i]
	}
	return out
}

// AndInto, AND the contents of two byte-slices of same length, write result into first parameter.
// `inout` will contain `inout[i] & b[i]`, for all `i`.
func AndInto[T types.Integer](inout, b []T) {
	assert.Equal(len(inout), len(b))
	for i := range inout {
		inout[i] &= b[i]
	}
}

// And, AND the contents: `a & b`.
// Returns the result, leaving `a` and `b` unmodified.
func And[T types.Integer](a, b []T) []T {
	assert.Equal(len(a), len(b))
	out := make([]T, len(a))
	for i := range out {
		out[i] = a[i] & b[i]
	}
	return out
}

// OrInto, OR the contents of two byte-slices of same length, write result into first parameter.
// `inout` will contain `inout[i] | b[i]`, for all `i`.
func OrInto[T types.Integer](inout, b []T) {
	assert.Equal(len(inout), len(b))
	for i := range inout {
		inout[i] |= b[i]
	}
}

// Or, OR the contents: `a | b`.
// Returns the result, leaving `a` and `b` unmodified.
func Or[T types.Integer](a, b []T) []T {
	assert.Equal(len(a), len(b))
	out := make([]T, len(a))
	for i := range out {
		out[i] = a[i] | b[i]
	}
	return out
}

// RotateLeft rotates (as opposed to shifting) the bits through the data-type's memory region.
//
// note: in most cases, it would be better to take the single line of code and apply it yourself.
func RotateLeft[T types.Integer](a T, n uint) T {
	memsize := uint(unsafe.Sizeof(a)) * 8
	if n %= memsize; n == 0 {
		return a
	}
	return a<<n | a>>(memsize-n)
}

// RotateRight rotates (as opposed to shifting) the bits through the data-type's memory region.
//
// note: in most cases, it would be better to take the single line of code and apply it yourself.
func RotateRight[T types.Integer](a T, n uint) T {
	memsize := uint(unsafe.Sizeof(a)) * 8
	if n %= memsize; n == 0 {
		return a
	}
	return a>>n | a<<(memsize-n)
}

// ShiftLeft shifts contents of byte-slice `n` bits left, for n >= 0.
func ShiftLeft(slice []byte, n int) {
	assert.NonNegative(n)
	if len(slice) == 0 || n == 0 {
		return
	}
	// Skipping/shifting full bytes
	if skip := n / 8; skip > 0 {
		for i := range slice {
			if i < skip {
				continue
			}
			slice[i-skip] = slice[i]
		}
		// TODO is this properly implemented to zero the right-most full bytes after shifting?
		for i := 0; i < skip && len(slice)-1-i >= 0; i++ {
			slice[len(slice)-1-i] = 0
		}
		n -= skip * 8
	}
	assert.AtMost(7, n)
	if n == 0 {
		return
	}
	// Shifting remaining bits
	for i := range slice {
		if i > 0 {
			slice[i-1] |= byte(slice[i]&byte(0b11111111<<(8-n))) >> (8 - n)
		}
		slice[i] <<= n
	}
}

// ShiftRight shifts contents of byte-slice `n` bits right, for n >= 0.
func ShiftRight(slice []byte, n int) {
	assert.NonNegative(n)
	if len(slice) == 0 || n == 0 {
		return
	}
	// Skipping/shifting full bytes
	if skip := n / 8; skip > 0 {
		for i := len(slice) - 1; i >= skip; i-- {
			slice[i] = slice[i-skip]
		}
		for i := 0; i < skip && i < len(slice); i++ {
			slice[i] = 0
		}
		n -= skip * 8
	}
	assert.AtMost(7, n)
	if n == 0 {
		return
	}
	// Shifting remaining bits
	for i := len(slice) - 1; i >= 0; i-- {
		slice[i] >>= n
		if i > 0 {
			slice[i] |= byte(slice[i-1]&byte(0b11111111>>(8-n))) << (8 - n)
		}
	}
}
