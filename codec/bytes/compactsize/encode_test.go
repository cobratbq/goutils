// SPDX-License-Identifier: LGPL-3.0-only

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

func TestEncodeIntoUint8Concat4(t *testing.T) {
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
		expected [3]byte
	}{
		{0, 1, [3]byte{0, 0, 0}},
		{7, 1, [3]byte{7, 0, 0}},
		{0xff, 3, [3]byte{0xfd, 0xff, 0}},
	}
	var buffer [3]byte
	for _, test := range testdata {
		buffer = [3]byte{}
		n := EncodeIntoUint8(buffer[:], test.value)
		assert.Equal(t, test.n, n)
		assert.Equal(t, test.expected, buffer)
	}
}

func TestEncodeIntoUint16(t *testing.T) {
	testdata := []struct {
		value    uint16
		n        uint
		expected [3]byte
	}{
		{0, 1, [3]byte{0, 0, 0}},
		{7, 1, [3]byte{7, 0, 0}},
		{0xff00, 3, [3]byte{0xfd, 0, 0xff}},
		{0xffff, 3, [3]byte{0xfd, 0xff, 0xff}},
	}
	var buffer [3]byte
	for _, test := range testdata {
		buffer = [3]byte{}
		n := EncodeIntoUint16(buffer[:], test.value)
		assert.Equal(t, test.n, n)
		assert.Equal(t, test.expected, buffer)
	}
}

func TestEncodeIntoUint32(t *testing.T) {
	testdata := []struct {
		value    uint32
		n        uint
		expected [5]byte
	}{
		{0, 1, [5]byte{0, 0, 0, 0, 0}},
		{7, 1, [5]byte{7, 0, 0, 0, 0}},
		{0xfd, 3, [5]byte{0xfd, 0xfd, 0}},
		{0xff00, 3, [5]byte{0xfd, 0, 0xff}},
		{0xffff, 3, [5]byte{0xfd, 0xff, 0xff}},
		{0xffff0000, 5, [5]byte{0xfe, 0, 0, 0xff, 0xff}},
		{0xffffffff, 5, [5]byte{0xfe, 0xff, 0xff, 0xff, 0xff}},
	}
	var buffer [5]byte
	for _, test := range testdata {
		buffer = [5]byte{}
		n := EncodeIntoUint32(buffer[:], test.value)
		assert.Equal(t, test.n, n)
		assert.Equal(t, test.expected, buffer)
	}
}

func TestEncodeIntoUint64(t *testing.T) {
	testdata := []struct {
		value    uint64
		n        uint
		expected [9]byte
	}{
		{0, 1, [9]byte{0, 0, 0, 0, 0}},
		{7, 1, [9]byte{7, 0, 0, 0, 0}},
		{0xfd, 3, [9]byte{0xfd, 0xfd, 0}},
		{0xff00, 3, [9]byte{0xfd, 0, 0xff}},
		{0xffff, 3, [9]byte{0xfd, 0xff, 0xff}},
		{0xffff0000, 5, [9]byte{0xfe, 0, 0, 0xff, 0xff}},
		{0xffffffff, 5, [9]byte{0xfe, 0xff, 0xff, 0xff, 0xff}},
		{0xffff000000000000, 9, [9]byte{0xff, 0, 0, 0, 0, 0, 0, 0xff, 0xff}},
		{0xffffffffffffffff, 9, [9]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
	}
	var buffer [9]byte
	for _, test := range testdata {
		buffer = [9]byte{}
		n := EncodeIntoUint64(buffer[:], test.value)
		assert.Equal(t, test.n, n)
		assert.Equal(t, test.expected, buffer)
	}
}
