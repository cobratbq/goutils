// SPDX-License-Identifier: LGPL-3.0-or-later

// bitset is a package that provides a bit-wise set implementation for indexes/positions.
// Essentially, a bit is 1 if present, or 0 if absent. These implementation can be used to
// efficiently store the minimal information of 1 bit to keep track of e.g. presence for large
// numbers.
package bitset

import (
	"github.com/cobratbq/goutils/std/builtin"
)

const LimbLength = builtin.UintSize

// Len returns the length of the set in number of (available) bits.
func Len(bitset []uint) int {
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
	return idx / LimbLength, 1 << (idx % LimbLength)
}
