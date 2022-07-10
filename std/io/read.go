// SPDX-License-Identifier: LGPL-3.0-or-later
package io

import (
	"io"
	"io/ioutil"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/errors"
)

const ErrIncompleteRead errors.StringError = "incomplete read"

func ReadUint8(in io.Reader) (uint8, error) {
	var b [1]byte
	if err := ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return b[0], nil
}

func MustReadUint8(in io.Reader) uint8 {
	var b [1]byte
	MustReadBytes(in, b[:])
	return b[0]
}

func ReadUint16(in io.Reader) (uint16, error) {
	var b [2]byte
	if err := ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return uint16(b[0])<<8 + uint16(b[1]), nil
}

func MustReadUint16(in io.Reader) uint16 {
	var b [2]byte
	MustReadBytes(in, b[:])
	return uint16(b[0])<<8 + uint16(b[1])
}

func ReadUint32(in io.Reader) (uint32, error) {
	var b [4]byte
	if err := ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return uint32(b[0])<<24 + uint32(b[1])<<16 + uint32(b[2])<<8 + uint32(b[3]), nil
}

func MustReadUint32(in io.Reader) uint32 {
	var b [4]byte
	MustReadBytes(in, b[:])
	return uint32(b[0])<<24 + uint32(b[1])<<16 + uint32(b[2])<<8 + uint32(b[3])
}

func ReadUint64(in io.Reader) (uint64, error) {
	var b [8]byte
	if err := ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return uint64(b[0])<<56 + uint64(b[1])<<48 + uint64(b[2])<<40 + uint64(b[3])<<32 +
		uint64(b[4])<<24 + uint64(b[5])<<16 + uint64(b[6])<<8 + uint64(b[7]), nil
}

func MustReadUint64(in io.Reader) uint64 {
	var b [8]byte
	MustReadBytes(in, b[:])
	return uint64(b[0])<<56 + uint64(b[1])<<48 + uint64(b[2])<<40 + uint64(b[3])<<32 +
		uint64(b[4])<<24 + uint64(b[5])<<16 + uint64(b[6])<<8 + uint64(b[7])
}

func ReadFull(in io.Reader, dst []byte) error {
	n, err := in.Read(dst)
	if err != nil {
		return err
	}
	if n < len(dst) {
		return ErrIncompleteRead
	}
	return nil
}

// MustReadAll reads all data from reader and panics in case an error occurs.
func MustReadAll(r io.Reader) []byte {
	data, err := ioutil.ReadAll(r)
	assert.Success(err, "Failed to read all data from reader: %+v")
	return data
}

// MustReadBytes reads bytes into dst and fails if anything out of the ordinary happens.
func MustReadBytes(in io.Reader, dst []byte) {
	n, err := in.Read(dst)
	assert.Success(err, "failed to read random bytes: %+v")
	assert.Equal(n, len(dst))
}

// Discard reads remaining data from reader and discards it. Any possible
// errors in the process are ignored. Returns nr of bytes written, thus
// discarded.
func Discard(r io.Reader) int64 {
	n, _ := io.Copy(io.Discard, r)
	return n
}
