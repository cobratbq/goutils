package compactsize

import "math"

func EncodeIntoUint8(dest []byte, value uint8) uint {
	if (value < 0xfd && len(dest) < 1) || (value >= 0xfd && len(dest) < 3) {
		return 0
	}
	if value < 0xfd {
		dest[0] = value
		return 1
	}
	dest[1] = value
	dest[2] = 0
	dest[0] = 0xfd
	return 3
}

func EncodeIntoUint16(dest []byte, value uint16) uint {
	if (value < 0xfd && len(dest) < 1) || (value >= 0xfd && len(dest) < 3) {
		return 0
	}
	if value < 0xfd {
		dest[0] = uint8(value)
		return 1
	}
	dest[1] = uint8(value)
	dest[2] = uint8(value >> 8)
	dest[0] = 0xfd
	return 3
}

func EncodeIntoUint32(dest []byte, value uint32) uint {
	if (value < 0xfd && len(dest) < 1) || (value >= 0xfd && value <= math.MaxUint16 && len(dest) < 3) || (value > math.MaxUint16 && len(dest) < 5) {
		return 0
	}
	if value < 0xfd {
		dest[0] = uint8(value)
		return 1
	}
	dest[1] = uint8(value)
	dest[2] = uint8(value >> 8)
	if value <= math.MaxUint16 {
		dest[0] = 0xfd
		return 3
	}
	dest[3] = uint8(value >> 16)
	dest[4] = uint8(value >> 24)
	dest[0] = 0xfe
	return 5
}

func EncodeIntoUint64(dest []byte, value uint64) uint {
	if (value < 0xfd && len(dest) < 1) || (value >= 0xfd && value <= math.MaxUint16 && len(dest) < 3) || (value > math.MaxUint16 && value <= math.MaxUint32 && len(dest) < 5) || (value > math.MaxUint32 && len(dest) < 9) {
		return 0
	}
	if value < 0xfd {
		dest[0] = uint8(value)
		return 1
	}
	dest[1] = uint8(value)
	dest[2] = uint8(value >> 8)
	if value <= math.MaxUint16 {
		dest[0] = 0xfd
		return 3
	}
	dest[3] = uint8(value >> 16)
	dest[4] = uint8(value >> 24)
	if value <= math.MaxUint32 {
		dest[0] = 0xfe
		return 5
	}
	dest[5] = uint8(value >> 32)
	dest[6] = uint8(value >> 40)
	dest[7] = uint8(value >> 48)
	dest[8] = uint8(value >> 56)
	dest[0] = 0xff
	return 9
}
