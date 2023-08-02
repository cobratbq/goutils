// SPDX-License-Identifier: LGPL-3.0-only

package digit

import (
	"testing"

	"github.com/cobratbq/goutils/assert"
)

func TestEncodeDecodeDigit(t *testing.T) {
	for i := uint8(0); i < 10; i++ {
		assert.Equal(i, DecodeDigit(EncodeDigit(i)))
	}
	for c := byte('0'); c <= '9'; c++ {
		assert.Equal(c, EncodeDigit(DecodeDigit(c)))
	}
}
