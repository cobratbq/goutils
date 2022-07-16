// SPDX-License-Identifier: LGPL-3.0-or-later

package bigendian

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
	if err := io_.ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return uint16(b[0])<<8 + uint16(b[1]), nil
}

func MustReadUint16(in io.Reader) uint16 {
	var b [2]byte
	io_.MustReadBytes(in, b[:])
	return uint16(b[0])<<8 + uint16(b[1])
}

func ReadUint32(in io.Reader) (uint32, error) {
	var b [4]byte
	if err := io_.ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return uint32(b[0])<<24 + uint32(b[1])<<16 + uint32(b[2])<<8 + uint32(b[3]), nil
}

func MustReadUint32(in io.Reader) uint32 {
	var b [4]byte
	io_.MustReadBytes(in, b[:])
	return uint32(b[0])<<24 + uint32(b[1])<<16 + uint32(b[2])<<8 + uint32(b[3])
}

func ReadUint64(in io.Reader) (uint64, error) {
	var b [8]byte
	if err := io_.ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return uint64(b[0])<<56 + uint64(b[1])<<48 + uint64(b[2])<<40 + uint64(b[3])<<32 +
		uint64(b[4])<<24 + uint64(b[5])<<16 + uint64(b[6])<<8 + uint64(b[7]), nil
}

func MustReadUint64(in io.Reader) uint64 {
	var b [8]byte
	io_.MustReadBytes(in, b[:])
	return uint64(b[0])<<56 + uint64(b[1])<<48 + uint64(b[2])<<40 + uint64(b[3])<<32 +
		uint64(b[4])<<24 + uint64(b[5])<<16 + uint64(b[6])<<8 + uint64(b[7])
}
