package bytes

import "github.com/cobratbq/goutils/assert"

func Xor(inout, b []byte) {
	assert.Equal(len(inout), len(b))
	for i := range inout {
		inout[i] ^= b[i]
	}
}
