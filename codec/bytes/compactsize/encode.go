// SPDX-License-Identifier: LGPL-3.0-only

package compactsize

import (
	"math"

	"github.com/cobratbq/goutils/types"
)

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

func EncodeUint8(value uint8) []byte {
	var data [3]byte
	n := EncodeIntoUint8(data[:], value)
	return data[:n]
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

func EncodeUint16(value uint16) []byte {
	var data [3]byte
	n := EncodeIntoUint16(data[:], value)
	return data[:n]
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

func EncodeUint32(value uint32) []byte {
	var data [5]byte
	n := EncodeIntoUint32(data[:], value)
	return data[:n]
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

func EncodeUint64(value uint64) []byte {
	var data [9]byte
	n := EncodeIntoUint64(data[:], value)
	return data[:n]
}

func EncodeIntoUint(dest []byte, value uint) uint {
	if types.MaxUint == math.MaxUint32 {
		return EncodeIntoUint32(dest, uint32(value))
	}
	return EncodeIntoUint64(dest, uint64(value))
}

func EncodeUint(value uint) []byte {
	if types.MaxUint == math.MaxUint32 {
		return EncodeUint32(uint32(value))
	}
	return EncodeUint64(uint64(value))
}
