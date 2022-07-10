// SPDX-License-Identifier: LGPL-3.0-or-later
package io

import (
	"io"
	"io/ioutil"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/errors"
)

const ErrIncompleteRead errors.StringError = "incomplete read"

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
	assert.Success(err, "Failed to read sufficient bytes to fill destination: %+v")
	assert.Equal(n, len(dst))
}

func ReadAll(in io.Reader) ([]byte, error) {
	return ioutil.ReadAll(in)
}

// MustReadAll reads all data from reader and panics in case an error occurs.
func MustReadAll(in io.Reader) []byte {
	data, err := ioutil.ReadAll(in)
	assert.Success(err, "Failed to read all data from reader: %+v")
	return data
}

// MustReadBytes reads bytes into dst and fails if anything out of the ordinary happens.
func MustReadBytes(in io.Reader, dst []byte) {
	n, err := in.Read(dst)
	assert.Success(err, "failed to read random bytes: %+v")
	assert.Equal(n, len(dst))
}
