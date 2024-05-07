// TODO return 'int' for read number of bytes instead of `uint`? (Although, now we guarantee by data-type that the value is always non-negative.)
// FIXME write CompactSize encoding
// FIXME needs testing
package compactsize

// Decode next value in buffer into `uint8`.
//
// Returns value as uint8, and number of bytes read.
// In case of 0 bytes read, either the encoding is bad or the value is too big to fit in the data-type.
func DecodeUint8(buffer []byte) (uint8, uint) {
	if len(buffer) < 1 || (buffer[0] == 0xfd && buffer[2] > 0) || buffer[0] > 0xfd {
		return 0, 0
	}
	if buffer[0] < 0xfd {
		return buffer[0], 1
	}
	return buffer[1], 3
}

// Decode next value in buffer into `uint16`.
//
// Returns value as uint16, and number of bytes read.
// In case of 0 bytes read, either the encoding is bad or the value is too big to fit in the data-type.
func DecodeUint16(buffer []byte) (uint16, uint) {
	if len(buffer) < 1 || (buffer[0] == 0xfd && len(buffer) < 3) || buffer[0] > 0xfd {
		return 0, 0
	}
	if buffer[0] < 0xfd {
		return uint16(buffer[0]), 1
	}
	return uint16(buffer[1]) | (uint16(buffer[2]) << 8), 3
}

// Decode next value in buffer into `uint32`.
//
// Returns value as uint32, and number of bytes read.
// In case of 0 bytes read, either the encoding is bad or the value is too big to fit in the data-type.
func DecodeUint32(buffer []byte) (uint32, uint) {
	if len(buffer) < 1 || (buffer[0] == 0xfd && len(buffer) < 3) || (buffer[0] == 0xfe && len(buffer) < 5) || buffer[0] > 0xfe {
		return 0, 0
	}
	if buffer[0] < 0xfd {
		return uint32(buffer[0]), 1
	}
	var result uint32
	result += uint32(buffer[1])
	result += uint32(buffer[2]) << 8
	if buffer[0] == 0xfd {
		return result, 3
	}
	result += uint32(buffer[3]) << 16
	result += uint32(buffer[4]) << 24
	return result, 5
}

// Decode next value in buffer into `uint64`.
//
// Returns value as uint64, and number of bytes read.
// In case of 0 bytes read, either the encoding is bad or the value is too big to fit in the data-type.
func DecodeUint64(buffer []byte) (uint64, uint) {
	if len(buffer) < 1 || (buffer[0] == 0xfd && len(buffer) < 3) || (buffer[0] == 0xfe && len(buffer) < 5) || (buffer[0] == 0xff && len(buffer) < 9) {
		return 0, 0
	}
	if buffer[0] < 0xfd {
		return uint64(buffer[0]), 1
	}
	var result uint64
	result += uint64(buffer[1])
	result += uint64(buffer[2]) << 8
	if buffer[0] == 0xfd {
		return result, 3
	}
	result += uint64(buffer[3]) << 16
	result += uint64(buffer[4]) << 24
	if buffer[0] == 0xfe {
		return result, 5
	}
	result += uint64(buffer[5]) << 32
	result += uint64(buffer[6]) << 40
	result += uint64(buffer[7]) << 48
	result += uint64(buffer[8]) << 56
	return result, 9
}
