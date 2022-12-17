// SPDX-License-Identifier: LGPL-3.0-or-later

package coord

import "github.com/cobratbq/goutils/assert"

// Encode2D encodes a two-dimensional coordinate into an index number, e.g. for use in
// arrays/slices. Length (of the first dimension) must be strictly greater than 0.
func Encode2D(length, first, second int) int {
	assert.Require(length > 0, "Length must be greater than zero")
	assert.Require(first < length, "First dimension coordinate cannot be larger than length")
	assert.Require(first >= 0, "First dimension component cannot be negative")
	assert.Require(second >= 0, "Second dimension component cannot be negative")
	return second*length + first
}

// Decode2D decodes an index value back into a two-dimensional coordinate. Length must be strictly
// greater than 0.
func Decode2D(length, index int) (int, int) {
	assert.Require(length > 0, "Length must be greater than zero")
	assert.Require(index >= 0, "Index must be a non-negative value")
	return index % length, index / length
}

// Encode2DUint encodes a two-dimensional uint coordinate into an index number, e.g. for use in
// arrays/slices. Length (of the first dimension) must be strictly greater than 0.
func Encode2DUint(length, first, second uint) int {
	assert.Require(length > 0, "Length must be greater than zero")
	assert.Require(first < length, "First dimension coordinate cannot be larger than length")
	return int(second*length + first)
}

// Decode2DUint decodes an index value back into a unsigned two-dimensional coordinate. Length must
// be strictly greater than 0.
func Decode2DUint(length uint, index int) (uint, uint) {
	assert.Require(length > 0, "Length must be greater than zero")
	return uint(index) % length, uint(index) / length
}
