// SPDX-License-Identifier: LGPL-3.0-only

package io

import (
	"io"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/errors"
)

var ErrIncompleteRead = errors.NewStringError("incomplete read")

func ReadByte(in io.Reader) (byte, error) {
	var b [1]byte
	if err := ReadFull(in, b[:]); err != nil {
		return 0, err
	}
	return b[0], nil
}

func MustReadByte(in io.Reader) byte {
	var b [1]byte
	MustReadBytes(in, b[:])
	return b[0]
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

func MustReadFull(in io.Reader, dst []byte) {
	n, err := in.Read(dst)
	assert.Success(err, "Failed to read sufficient bytes to fill destination")
	assert.Equal(n, len(dst))
}

func ReadAll(in io.Reader) ([]byte, error) {
	return io.ReadAll(in)
}

// MustReadAll reads all data from reader and panics in case an error occurs.
func MustReadAll(in io.Reader) []byte {
	data, err := io.ReadAll(in)
	assert.Success(err, "Failed to read all data from reader")
	return data
}

// MustReadBytes reads bytes into dst and fails if anything out of the ordinary happens.
func MustReadBytes(in io.Reader, dst []byte) {
	n, err := in.Read(dst)
	assert.Success(err, "failed to read random bytes")
	assert.Equal(n, len(dst))
}

// ReadExpect reads and checks if the read byte is the expected `next` byte.
func ReadExpect(next byte, in io.Reader) (bool, error) {
	var b [1]byte
	if n, err := io.ReadFull(in, b[:]); err != nil {
		return false, errors.Context(err, "Failed to read next byte")
	} else if n == 0 {
		return false, errors.Context(io.ErrShortWrite, "Failed to read next byte")
	}
	return b[0] == next, nil
}

// ReadUntil reads until a specific byte is encountered.
//
// All bytes read before the stop-byte are returned. If an error is encountered, everything read until the
// error is returned together with the error. Following the behavior of `io.ReadFull`, ReadUntil will return
// io.EOF if end-of-file is reached.
//
// For more sophisticated and more capable functions, use `bufio` (buffered-io). These functions are provided
// for one-off cases and cases where reading ahead is not desirable or not allowed.
func ReadUntil(in io.Reader, stop byte) ([]byte, error) {
	var buffer []byte
	var b [1]byte
	for {
		if _, err := io.ReadFull(in, b[:]); err != nil {
			return buffer, err
		}
		if b[0] == stop {
			return buffer, nil
		}
		buffer = append(buffer, b[0])
	}
}
