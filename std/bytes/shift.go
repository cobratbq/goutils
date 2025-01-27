package bytes

import (
	"github.com/cobratbq/goutils/assert"
)

func ShiftLeft(slice []byte, n int) {
	assert.NonNegative(n)
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
