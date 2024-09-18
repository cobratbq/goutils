// SPDX-License-Identifier: LGPL-3.0-only

package hex

import (
	"encoding/hex"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/builtin"
	"github.com/cobratbq/goutils/std/errors"
)

const index string = `0123456789abcdef`

// MustDecodeString decodes an arbitrary-length hex-encoded string.
//
// `encoded` must be a valid hexadecimal-encoded string, i.e. even-length and containing only hexadecimal
// characters.
func MustDecodeString(encoded string) []byte {
	return builtin.Expect(hex.DecodeString(encoded))
}

func Decode(encoded []byte) ([]byte, error) {
	assert.Equal(0, len(encoded)%2)
	var decoded = make([]byte, len(encoded)/2)
	var n int
	var err error
	if n, err = hex.Decode(decoded, encoded); err != nil {
		return nil, err
	}
	if n != len(decoded) {
		return nil, errors.ErrIllegal
	}
	return decoded, nil
}

func MustDecode(encoded []byte) []byte {
	return builtin.Expect(Decode(encoded))
}

// HexEncodeChars encodes a uint8 value into its two-symbol (lower-case) hexadecimal representation.
func HexEncodeChars(value uint8) (byte, byte) {
	return HexEncodeChar((value & 0xf0) >> 4), HexEncodeChar(value & 0x0f)
}

// HexEncodeChar encodes uint value into its (lower-case) hexadecimal representation [0-9a-f].
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
	case 'a', 'A':
		return 10
	case 'b', 'B':
		return 11
	case 'c', 'C':
		return 12
	case 'd', 'D':
		return 13
	case 'e', 'E':
		return 14
	case 'f', 'F':
		return 15
	default:
		panic("BUG: invalid character encountered, not part of hexadecimal system")
	}
}

func AllHexadecimal(data []byte) bool {
	for i := range data {
		switch data[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'A', 'b', 'B', 'c', 'C', 'd', 'D', 'e', 'E', 'f', 'F':
			continue
		default:
			return false
		}
	}
	return true
}
