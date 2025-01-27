package bytes

import "github.com/cobratbq/goutils/assert"

// Xor the contents of two byte-slices of same length.
// `inout` will contain `inout[i] ^ b[i]`, for all `i`.
func Xor(inout, b []byte) {
	assert.Equal(len(inout), len(b))
	for i := range inout {
		inout[i] ^= b[i]
	}
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
