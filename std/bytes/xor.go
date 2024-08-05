// SPDX-License-Identifier: LGPL-3.0-only

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
