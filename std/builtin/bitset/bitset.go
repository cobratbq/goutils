package bitset

import (
	"math"

	"github.com/cobratbq/goutils/std/builtin"
)

const LimbLength = builtin.UintSize
const LimbMask = LimbLength - 1

var numShifts = uint(math.Log2(LimbLength))

// Cap returns the capacity of the set in (available) bits.
func Cap(bitset []uint) int {
	return len(bitset) * LimbLength
}

// Bit returns the bit value for specified index.
func Bit(bitset []uint, idx uint) bool {
	limb, bit := loc(idx)
	return bitset[limb]&bit == bit
}

// Insert sets the bit for specified index.
func Insert(bitset []uint, idx uint) {
	limb, bit := loc(idx)
	bitset[limb] |= bit
}

// Remove removes the bit for specified index.
func Remove(bitset []uint, idx uint) {
	limb, bit := loc(idx)
	bitset[limb] &^= bit
}

// Clear sets all bits to 0.
func Clear(bitset []uint) {
	for i := 0; i < len(bitset); i++ {
		bitset[i] = 0
	}
}

func loc(idx uint) (uint, uint) {
	// TODO "optimization" that isn't based on measurements. So needs to be double-checked.
	return idx >> numShifts, 1 << (idx & LimbMask)
	//return idx / LimbLength, 1 << (idx % LimbLength)
}
