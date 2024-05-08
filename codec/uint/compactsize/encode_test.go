package compactsize

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestEncodeIntoUint8_1(t *testing.T) {
	buffer := [1]byte{0}
	n := EncodeIntoUint8(buffer[:], 0xff)
	assert.Equal(t, 0, n)
	assert.Equal(t, [1]byte{0}, buffer)
}

func TestEncodeIntoUint8concat4(t *testing.T) {
	buffer := [8]byte{0}
	idx := uint(0)
	idx += EncodeIntoUint8(buffer[idx:], 0xaa)
	idx += EncodeIntoUint8(buffer[idx:], 0xff)
	idx += EncodeIntoUint8(buffer[idx:], 0xbb)
	idx += EncodeIntoUint8(buffer[idx:], 0xfd)
	assert.Equal(t, 8, idx)
	assert.Equal(t, [8]byte{0xaa, 0xfd, 0xff, 0, 0xbb, 0xfd, 0xfd, 0}, buffer)
}

func TestEncodeIntoUint8(t *testing.T) {
	testdata := []struct {
		value    uint8
		n        uint
		expected [9]byte
	}{
		{0, 1, [9]byte{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{7, 1, [9]byte{7, 0, 0, 0, 0, 0, 0, 0, 0}},
		{0xff, 3, [9]byte{0xfd, 0xff, 0, 0, 0, 0, 0, 0, 0}},
	}
	var buffer [9]byte
	for _, test := range testdata {
		buffer = [9]byte{}
		n := EncodeIntoUint8(buffer[:], test.value)
		assert.Equal(t, test.n, n)
		assert.Equal(t, test.expected, buffer)
	}
}
