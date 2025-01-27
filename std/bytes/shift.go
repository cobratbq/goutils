package bytes

import (
	"github.com/cobratbq/goutils/assert"
)

func ShiftLeft(slice []byte, n int) {
	assert.NonNegative(n)
	for n >= 8 {
		for i := range slice {
			if i == 0 {
				continue
			}
			slice[i-1] = slice[i]
		}
		slice[len(slice)-1] = 0
		n -= 8
	}
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
