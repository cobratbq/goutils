// SPDX-License-Identifier: LGPL-3.0-only

package bitops

import (
	"bytes"
	"testing"

	"github.com/cobratbq/goutils/std/log"
	assert "github.com/cobratbq/goutils/std/testing"
)

func TestOps(t *testing.T) {
	// TODO could use more testdata, more variations
	entries := []struct {
		a   []byte
		b   []byte
		xor []byte
		or  []byte
		and []byte
	}{
		{a: []byte{0b11110000}, b: []byte{0b00001111}, xor: []byte{0b11111111}, or: []byte{0b11111111}, and: []byte{0b00000000}},
		{a: []byte{0b11110000}, b: []byte{0b11111111}, xor: []byte{0b00001111}, or: []byte{0b11111111}, and: []byte{0b11110000}},
		{a: []byte{0b11110000, 0b00000000, 0b11111111}, b: []byte{0b00001111, 0b11111111, 0b11111111}, xor: []byte{0b11111111, 0b11111111, 0b00000000}, or: []byte{0b11111111, 0b11111111, 0b11111111}, and: []byte{0b00000000, 0b00000000, 0b11111111}},
	}
	var buf, out []byte
	for e := range entries {
		// Bitwise XOR operation
		out = Xor(entries[e].a, entries[e].b)
		assert.SlicesEqual(t, entries[e].xor, out)
		buf = bytes.Clone(entries[e].a)
		XorInto(buf, entries[e].b)
		assert.SlicesEqual(t, entries[e].xor, buf)
		assert.SlicesEqual(t, out, buf)

		// Bitwise OR operation
		out = Or(entries[e].a, entries[e].b)
		assert.SlicesEqual(t, entries[e].or, out)
		buf = bytes.Clone(entries[e].a)
		OrInto(buf, entries[e].b)
		assert.SlicesEqual(t, entries[e].or, buf)
		assert.SlicesEqual(t, out, buf)

		// Bitwise AND operation
		out = And(entries[e].a, entries[e].b)
		assert.SlicesEqual(t, entries[e].and, out)
		buf = bytes.Clone(entries[e].a)
		AndInto(buf, entries[e].b)
		assert.SlicesEqual(t, entries[e].and, buf)
		assert.SlicesEqual(t, out, buf)
	}
}

func TestShiftLeftNegativeN(t *testing.T) {
	defer assert.RequirePanic(t)
	ShiftLeft([]byte{1}, -1)
	t.FailNow()
}

func TestShiftLeft(t *testing.T) {
	entries := []struct {
		data     []byte
		n        int
		expected []byte
	}{
		{[]byte{}, 0, []byte{}},
		{[]byte{}, 1, []byte{}},
		{[]byte{}, 8, []byte{}},
		{[]byte{}, 9, []byte{}},
		{[]byte{0}, 0, []byte{0}},
		{[]byte{1}, 0, []byte{1}},
		{[]byte{1}, 1, []byte{2}},
		{[]byte{1}, 2, []byte{4}},
		{[]byte{1}, 3, []byte{8}},
		{[]byte{1}, 4, []byte{16}},
		{[]byte{1}, 5, []byte{32}},
		{[]byte{1}, 6, []byte{64}},
		{[]byte{1}, 7, []byte{128}},
		{[]byte{1}, 8, []byte{0}},
		{[]byte{0b10101011}, 0, []byte{0b10101011}},
		{[]byte{0b10101011}, 1, []byte{0b01010110}},
		{[]byte{0b10101011}, 2, []byte{0b10101100}},
		{[]byte{0b10101011}, 3, []byte{0b01011000}},
		{[]byte{0b10101011}, 4, []byte{0b10110000}},
		{[]byte{0b10101011}, 5, []byte{0b01100000}},
		{[]byte{0b10101011}, 6, []byte{0b11000000}},
		{[]byte{0b10101011}, 7, []byte{0b10000000}},
		{[]byte{0b10101011}, 8, []byte{0b00000000}},
		{[]byte{0b10101011}, 9, []byte{0b00000000}},
		{[]byte{0b10101011}, 999, []byte{0b00000000}},
		{[]byte{0, 0xff}, 0, []byte{0b00000000, 0b11111111}},
		{[]byte{0, 0xff}, 1, []byte{0b00000001, 0b11111110}},
		{[]byte{0, 0xff}, 2, []byte{0b00000011, 0b11111100}},
		{[]byte{0, 0xff}, 3, []byte{0b00000111, 0b11111000}},
		{[]byte{0, 0xff}, 4, []byte{0b00001111, 0b11110000}},
		{[]byte{0, 0xff}, 5, []byte{0b00011111, 0b11100000}},
		{[]byte{0, 0xff}, 6, []byte{0b00111111, 0b11000000}},
		{[]byte{0, 0xff}, 7, []byte{0b01111111, 0b10000000}},
		{[]byte{0, 0xff}, 8, []byte{0b11111111, 0b00000000}},
		{[]byte{0, 0xff}, 9, []byte{0b11111110, 0b00000000}},
		{[]byte{0, 0xff}, 10, []byte{0b11111100, 0b00000000}},
		{[]byte{0, 0xff}, 11, []byte{0b11111000, 0b00000000}},
		{[]byte{0, 0xff}, 12, []byte{0b11110000, 0b00000000}},
		{[]byte{0, 0xff}, 13, []byte{0b11100000, 0b00000000}},
		{[]byte{0, 0xff}, 14, []byte{0b11000000, 0b00000000}},
		{[]byte{0, 0xff}, 15, []byte{0b10000000, 0b00000000}},
		{[]byte{0, 0xff}, 16, []byte{0b00000000, 0b00000000}},
		{[]byte{0, 0xff}, 17, []byte{0b00000000, 0b00000000}},
		{[]byte{0, 0xff}, 18, []byte{0b00000000, 0b00000000}},
		{[]byte{0, 0, 0xff}, 15, []byte{0x7f, 0x80, 0}},
		{[]byte{0, 0, 0xff}, 16, []byte{0xff, 0, 0}},
		{[]byte{0, 0, 0xff}, 17, []byte{0xfe, 0, 0}},
	}
	for e := range entries {
		t.Log("Entry", e)
		entry := bytes.Clone(entries[e].data)
		ShiftLeft(entry, entries[e].n)
		assert.SlicesEqual(t, entries[e].expected, entry)
	}
}

func TestShiftRight(t *testing.T) {
	entries := []struct {
		expected []byte
		n        int
		data     []byte
	}{
		{[]byte{}, 0, []byte{}},
		{[]byte{}, 1, []byte{}},
		{[]byte{}, 8, []byte{}},
		{[]byte{}, 9, []byte{}},
		{[]byte{0}, 0, []byte{0}},
		{[]byte{1}, 0, []byte{1}},
		{[]byte{1}, 1, []byte{2}},
		{[]byte{1}, 2, []byte{4}},
		{[]byte{1}, 3, []byte{8}},
		{[]byte{1}, 4, []byte{16}},
		{[]byte{1}, 5, []byte{32}},
		{[]byte{1}, 6, []byte{64}},
		{[]byte{1}, 7, []byte{128}},
		{[]byte{0}, 8, []byte{0}},
		{[]byte{0b10101011}, 0, []byte{0b10101011}},
		{[]byte{0b00101011}, 1, []byte{0b01010110}},
		{[]byte{0b00101011}, 2, []byte{0b10101100}},
		{[]byte{0b00001011}, 3, []byte{0b01011000}},
		{[]byte{0b00001011}, 4, []byte{0b10110000}},
		{[]byte{0b00000011}, 5, []byte{0b01100000}},
		{[]byte{0b00000011}, 6, []byte{0b11000000}},
		{[]byte{0b00000001}, 7, []byte{0b10000000}},
		{[]byte{0b00000000}, 8, []byte{0b00000000}},
		{[]byte{0b00000000}, 9, []byte{0b00000000}},
		{[]byte{0b00000000}, 999, []byte{0b00000000}},
		{[]byte{0, 0xff}, 0, []byte{0b00000000, 0b11111111}},
		{[]byte{0, 0xff}, 1, []byte{0b00000001, 0b11111110}},
		{[]byte{0, 0xff}, 2, []byte{0b00000011, 0b11111100}},
		{[]byte{0, 0xff}, 3, []byte{0b00000111, 0b11111000}},
		{[]byte{0, 0xff}, 4, []byte{0b00001111, 0b11110000}},
		{[]byte{0, 0xff}, 5, []byte{0b00011111, 0b11100000}},
		{[]byte{0, 0xff}, 6, []byte{0b00111111, 0b11000000}},
		{[]byte{0, 0xff}, 7, []byte{0b01111111, 0b10000000}},
		{[]byte{0, 0xff}, 8, []byte{0b11111111, 0b00000000}},
		{[]byte{0, 0b01111111}, 9, []byte{0b11111110, 0b00000000}},
		{[]byte{0, 0b00111111}, 10, []byte{0b11111100, 0b00000000}},
		{[]byte{0, 0b00011111}, 11, []byte{0b11111000, 0b00000000}},
		{[]byte{0, 0b00001111}, 12, []byte{0b11110000, 0b00000000}},
		{[]byte{0, 0b00000111}, 13, []byte{0b11100000, 0b00000000}},
		{[]byte{0, 0b00000011}, 14, []byte{0b11000000, 0b00000000}},
		{[]byte{0, 0b00000001}, 15, []byte{0b10000000, 0b00000000}},
		{[]byte{0, 0b00000000}, 16, []byte{0b00000000, 0b00000000}},
		{[]byte{0, 0b00000000}, 17, []byte{0b00000000, 0b00000000}},
		{[]byte{0, 0b00000000}, 18, []byte{0b00000000, 0b00000000}},
		{[]byte{0, 0, 0xff}, 15, []byte{0x7f, 0x80, 0}},
		{[]byte{0, 0, 0xff}, 16, []byte{0xff, 0, 0}},
		{[]byte{0, 0, 0x7f}, 17, []byte{0xfe, 0, 0}},
	}
	for e := range entries {
		t.Log("Entry", e)
		entry := bytes.Clone(entries[e].data)
		ShiftRight(entry, entries[e].n)
		assert.SlicesEqual(t, entries[e].expected, entry)
	}
}

func TestRotations(t *testing.T) {
	log.Infoln("Testing with 8-bit unsinged integers…")
	entries8 := []struct {
		input uint8
		n     uint
		left  uint8
		right uint8
	}{
		{input: 0b10101010, n: 0, left: 0b10101010, right: 0b10101010},
		{input: 0b10101010, n: 1, left: 0b01010101, right: 0b01010101},
		{input: 0b10101010, n: 2, left: 0b10101010, right: 0b10101010},
		{input: 0b10101010, n: 3, left: 0b01010101, right: 0b01010101},
		{input: 0b10101010, n: 4, left: 0b10101010, right: 0b10101010},
		{input: 0b10101010, n: 5, left: 0b01010101, right: 0b01010101},
		{input: 0b10101010, n: 6, left: 0b10101010, right: 0b10101010},
		{input: 0b10101010, n: 7, left: 0b01010101, right: 0b01010101},
		{input: 0b10101010, n: 8, left: 0b10101010, right: 0b10101010},
		{input: 0b10101010, n: 9, left: 0b01010101, right: 0b01010101},
		{input: 0b11110000, n: 1, left: 0b11100001, right: 0b01111000},
		{input: 0b11110000, n: 2, left: 0b11000011, right: 0b00111100},
		{input: 0b11110000, n: 3, left: 0b10000111, right: 0b00011110},
		{input: 0b11110000, n: 4, left: 0b00001111, right: 0b00001111},
		{input: 0b11110000, n: 5, left: 0b00011110, right: 0b10000111},
		{input: 0b11110000, n: 6, left: 0b00111100, right: 0b11000011},
		{input: 0b11110000, n: 7, left: 0b01111000, right: 0b11100001},
	}
	for i := range entries8 {
		log.Infoln("Iteration", i)
		left := RotateLeft(entries8[i].input, entries8[i].n)
		assert.Equal(t, left, entries8[i].left)
		right := RotateRight(entries8[i].input, entries8[i].n)
		assert.Equal(t, right, entries8[i].right)
	}
	log.Infoln("Testing with 64-bit unsinged integers…")
	entries64 := []struct {
		input uint64
		n     uint
		left  uint64
		right uint64
	}{
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 0, left: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, right: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 64, left: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, right: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 128, left: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, right: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 1, left: 0b11111111_11111111_11111111_11111111_11111111_11111111_11110000_00000001, right: 0b01111111_11111111_11111111_11111111_11111111_11111111_11111100_00000000},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 2, left: 0b11111111_11111111_11111111_11111111_11111111_11111111_11100000_00000011, right: 0b00111111_11111111_11111111_11111111_11111111_11111111_11111110_00000000},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 3, left: 0b11111111_11111111_11111111_11111111_11111111_11111111_11000000_00000111, right: 0b00011111_11111111_11111111_11111111_11111111_11111111_11111111_00000000},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 7, left: 0b11111111_11111111_11111111_11111111_11111111_11111100_00000000_01111111, right: 0b00000001_11111111_11111111_11111111_11111111_11111111_11111111_11110000},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 8, left: 0b11111111_11111111_11111111_11111111_11111111_11111000_00000000_11111111, right: 0b00000000_11111111_11111111_11111111_11111111_11111111_11111111_11111000},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 9, left: 0b11111111_11111111_11111111_11111111_11111111_11110000_00000001_11111111, right: 0b00000000_01111111_11111111_11111111_11111111_11111111_11111111_11111100},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 16, left: 0b11111111_11111111_11111111_11111111_11111000_00000000_11111111_11111111, right: 0b11111000_00000000_11111111_11111111_11111111_11111111_11111111_11111111},
		{input: 0b11111111_11111111_11111111_11111111_11111111_11111111_11111000_00000000, n: 32, left: 0b11111111_11111111_11111000_00000000_11111111_11111111_11111111_11111111, right: 0b11111111_11111111_11111000_00000000_11111111_11111111_11111111_11111111},
	}
	for i := range entries64 {
		log.Infoln("Iteration", i)
		left := RotateLeft(entries64[i].input, entries64[i].n)
		assert.Equal(t, left, entries64[i].left)
		right := RotateRight(entries64[i].input, entries64[i].n)
		assert.Equal(t, right, entries64[i].right)
	}
}
