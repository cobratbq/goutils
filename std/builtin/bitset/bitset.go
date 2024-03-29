// SPDX-License-Identifier: LGPL-3.0-only

// bitset is a package that provides a bit-wise set implementation for indexes/positions.
// Essentially, a bit is 1 if present, or 0 if absent. These implementation can be used to
// efficiently store the minimal information of 1 bit to keep track of e.g. presence for large
// numbers.
package bitset

import "github.com/cobratbq/goutils/types"

const LimbLength = types.UintSize

// Calculate calculates the size of a backing array such that it can store at least the specified
// capacity. (Rounded up to the next full limb. See `LimbLength`)
func Calculate(capacity uint) int {
	size := int(capacity / LimbLength)
	if capacity%LimbLength != 0 {
		size++
	}
	return size
}

// Len returns the length of the set in number of (available) bits.
func Len(bitset []uint) int {
	return len(bitset) * LimbLength
}

// Bit returns the bit value for specified index.
func Bit(bitset []uint, idx uint) bool {
	limb, bit := loc(idx)
	return bitset[limb]&bit == bit
}

// InsertMany inserts multiple indices into the bitset.
func InsertMany(bitset []uint, idxs ...uint) {
	for _, idx := range idxs {
		Insert(bitset, idx)
	}
}

// Insert sets the bit for specified index.
func Insert(bitset []uint, idx uint) {
	limb, bit := loc(idx)
	bitset[limb] |= bit
}

// RemoveMany removes multiple indices from the bitset.
func RemoveMany(bitset []uint, idxs ...uint) {
	for _, idx := range idxs {
		Remove(bitset, idx)
	}
}

// Remove removes the bit for specified index.
func Remove(bitset []uint, idx uint) {
	limb, bit := loc(idx)
	bitset[limb] &^= bit
}

// Clear sets all bits to 0. This implementation is not ideal as it has to operate on a slice of
// arbitrary length. It assigns 0 to each individual cell in the slice. If `bitset` is backed by an
// array that is directly accessible, consider assigning it an empty array of same (total) size.
// This approach enjoys better compile-time support.
//
//	var mybitset [40]uint
//	mybitset = [40]uint{}
func Clear(bitset []uint) {
	for i := 0; i < len(bitset); i++ {
		bitset[i] = 0
	}
}

func loc(idx uint) (uint, uint) {
	return idx / LimbLength, 1 << (idx % LimbLength)
}
