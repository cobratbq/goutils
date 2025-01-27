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
	if skip := n / 8; skip > 0 {
		for i := range slice {
			if i < skip {
				continue
			}
			slice[i-skip] = slice[i]
		}
		slice[len(slice)-1] = 0
		n -= skip * 8
	}
	assert.AtMost(7, n)
	if n == 0 {
		return
	}
	for i := range slice {
		if i > 0 {
			slice[i-1] |= byte(slice[i]&byte(0b11111111<<(8-n))) >> (8 - n)
		}
		slice[i] <<= n
	}
}
