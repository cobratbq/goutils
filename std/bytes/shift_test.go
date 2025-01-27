package bytes

import (
	"bytes"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

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
