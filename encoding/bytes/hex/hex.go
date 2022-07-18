// SPDX-License-Identifier: LGPL-3.0-or-later

package hex

import "github.com/cobratbq/goutils/assert"

const index string = `0123456789abcdef`
const indexLen uint = 16

func HexEncodeChars(value uint8) (byte, byte) {
	return HexEncodeChar((value & 0xf0) >> 4), HexEncodeChar(value & 0x0f)
}

func HexEncodeChar(value uint8) byte {
	assert.Require(value < uint8(indexLen), "Hexadecimal system contains only 16 values.")
	return index[value]
}

func HexDecodeChars(c0 byte, c1 byte) uint8 {
	return HexDecodeChar(c0)<<4 | HexDecodeChar(c1)
}

func HexDecodeChar(c byte) uint8 {
	switch c {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'a':
		return 10
	case 'b':
		return 11
	case 'c':
		return 12
	case 'd':
		return 13
	case 'e':
		return 14
	case 'f':
		return 15
	default:
		panic("BUG: invalid character encountered, not part of hexadecimal system")
	}
}
