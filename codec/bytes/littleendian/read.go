// SPDX-License-Identifier: LGPL-3.0-only

package littleendian

import (
	"io"

	io_ "github.com/cobratbq/goutils/std/io"
)

func ReadUint8(in io.Reader) (uint8, error) {
	return io_.ReadByte(in)
}

func MustReadUint8(in io.Reader) uint8 {
	return io_.MustReadByte(in)
}

func ReadUint16(in io.Reader) (uint16, error) {
	var b [2]byte
	if _, err := io.ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return Uint16(b[0], b[1]), nil
}

func MustReadUint16(in io.Reader) uint16 {
	var b [2]byte
	io_.MustReadBytes(in, b[:])
	return Uint16(b[0], b[1])
}

func Uint16(b0, b1 byte) uint16 {
	return uint16(b1)<<8 + uint16(b0)
}

func ReadUint32(in io.Reader) (uint32, error) {
	var b [4]byte
	if _, err := io.ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return Uint32(b[0], b[1], b[2], b[3]), nil
}

func MustReadUint32(in io.Reader) uint32 {
	var b [4]byte
	io_.MustReadBytes(in, b[:])
	return Uint32(b[0], b[1], b[2], b[3])
}

func Uint32(b0, b1, b2, b3 byte) uint32 {
	return uint32(b3)<<24 + uint32(b2)<<16 + uint32(b1)<<8 + uint32(b0)
}

func ReadUint64(in io.Reader) (uint64, error) {
	var b [8]byte
	if _, err := io.ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return Uint64(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]), nil
}

func MustReadUint64(in io.Reader) uint64 {
	var b [8]byte
	io_.MustReadBytes(in, b[:])
	return Uint64(b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7])
}

func Uint64(b0, b1, b2, b3, b4, b5, b6, b7 byte) uint64 {
	return uint64(b7)<<56 + uint64(b6)<<48 + uint64(b5)<<40 + uint64(b4)<<32 +
		uint64(b3)<<24 + uint64(b2)<<16 + uint64(b1)<<8 + uint64(b0)
}
