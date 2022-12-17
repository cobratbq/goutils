// SPDX-License-Identifier: LGPL-3.0-or-later

package digit

import (
	"testing"

	"github.com/cobratbq/goutils/assert"
)

func TestReadWriteDigit(t *testing.T) {
	for i := uint8(0); i < 10; i++ {
		assert.Equal(i, ReadDigit(WriteDigit(i)))
	}
	for c := byte('0'); c <= '9'; c++ {
		assert.Equal(c, WriteDigit(ReadDigit(c)))
	}
}
