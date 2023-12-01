// SPDX-License-Identifier: LGPL-3.0-only

package digit

func IsDigit(in byte) bool {
	// There are only 10 possibilities. Let's just test each one, that way even for obscure
	// character encodings where other characters are in-between these seemingly consecutive chars
	// it will not pose a problem.
	return in == '0' || in == '1' || in == '2' || in == '3' || in == '4' || in == '5' ||
		in == '6' || in == '7' || in == '8' || in == '9'
}

func IsDigitRune(in rune) bool {
	// There are only 10 possibilities. Let's just test each one, that way even for obscure
	// character encodings where other characters are in-between these seemingly consecutive chars
	// it will not pose a problem.
	return in == '0' || in == '1' || in == '2' || in == '3' || in == '4' || in == '5' ||
		in == '6' || in == '7' || in == '8' || in == '9'
}

// FindDigit finds the digits and their location in provided data.
func FindDigit(data []byte) map[int]uint8 {
	finds := make(map[int]uint8)
	for i := 0; i < len(data); i++ {
		if IsDigit(data[i]) {
			finds[i] = DecodeDigit(data[i])
		}
	}
	return finds
}

// DecodeDigit decodes a byte-representation of a digit into a unsigned integer value.
func DecodeDigit(in byte) uint8 {
	switch in {
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
	default:
		panic("Illegal symbol: not a digit")
	}
}

// EncodeDigit encodes a integer value as a digit-character.
func EncodeDigit(in uint8) byte {
	switch in {
	case 0:
		return '0'
	case 1:
		return '1'
	case 2:
		return '2'
	case 3:
		return '3'
	case 4:
		return '4'
	case 5:
		return '5'
	case 6:
		return '6'
	case 7:
		return '7'
	case 8:
		return '8'
	case 9:
		return '9'
	default:
		panic("Illegal value")
	}
}
