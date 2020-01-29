package rand

import (
	"math/rand"
)

// MathRandUint64NonZero generates a non-zero uint64 value. Due to the non-zero
// requirement, the distribution of random numbers is NOT perfectly uniform.
// TODO find nicer way to get random non-zero uint64.
func MathRandUint64NonZero() uint64 {
	r := rand.Uint64()
	if r == 0 {
		r++
	}
	return r
}
