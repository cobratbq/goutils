// SPDX-License-Identifier: AGPL-3.0-or-later

package hex

const index string = `0123456789abcdef`

// HexEncodeChars encodes a uint8 value into its two-symbol hexadecimal representation.
func HexEncodeChars(value uint8) (byte, byte) {
	return HexEncodeChar((value & 0xf0) >> 4), HexEncodeChar(value & 0x0f)
}

// HexEncodeChar encodes uint value into its hexadecimal representation [0-9a-f].
func HexEncodeChar(value uint8) byte {
	return index[value]
}

// HexDecodeChars decodes two chars from hexadecimal representation into a uint8 value.
func HexDecodeChars(c0 byte, c1 byte) uint8 {
	return HexDecodeChar(c0)<<4 | HexDecodeChar(c1)
}

// HexDecodeChar decodes a hexadecimal symbol into its ordinal value, ranged [0,16).
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
