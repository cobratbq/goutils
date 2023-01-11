// SPDX-License-Identifier: AGPL-3.0-or-later

package hex

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestHexEncode(t *testing.T) {
	entries := map[byte]byte{
		0:  '0',
		1:  '1',
		2:  '2',
		3:  '3',
		4:  '4',
		5:  '5',
		6:  '6',
		7:  '7',
		8:  '8',
		9:  '9',
		10: 'a',
		11: 'b',
		12: 'c',
		13: 'd',
		14: 'e',
		15: 'f',
	}
	for s, d := range entries {
		assert.Equal(t, d, HexEncodeChar(s))
	}
}

func TestHexEncodeOOB(t *testing.T) {
	defer assert.RequirePanic(t)
	_ = HexEncodeChar(16)
	t.FailNow()
}

func TestHexDecode(t *testing.T) {
	entries := map[byte]byte{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'a': 10,
		'b': 11,
		'c': 12,
		'd': 13,
		'e': 14,
		'f': 15,
	}
	for s, d := range entries {
		assert.Equal(t, d, HexDecodeChar(s))
	}
}

func TestHexDecodeInvalid(t *testing.T) {
	defer assert.RequirePanic(t)
	HexDecodeChar('z')
	t.FailNow()
}
