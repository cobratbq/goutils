package coord

import "github.com/cobratbq/goutils/assert"

func Encode2D(length, first, second uint) uint {
	assert.Require(length > 0, "Length must be greater than zero")
	assert.Require(first < length, "First dimension coordinate cannot be larger than length")
	return second*length + first
}

func Decode2D(length, index uint) (uint, uint) {
	return index % length, index / length
}

func Encode2DInt(length, first, second int) uint {
	assert.Require(length > 0, "Length must be greater than zero")
	assert.Require(first < length, "First dimension coordinate cannot be larger than length")
	assert.Require(first >= 0, "First dimension component cannot be negative")
	assert.Require(second >= 0, "Second dimension component cannot be negative")
	return uint(second)*uint(length) + uint(first)
}

func Decode2DInt(length int, index uint) (int, int) {
	assert.Require(length > 0, "Length must be greater than zero")
	return int(index % uint(length)), int(index / uint(length))
}
