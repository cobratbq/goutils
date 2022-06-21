// SPDX-License-Identifier: LGPL-3.0-or-later
package io

import (
	"io"
	"io/ioutil"

	"github.com/cobratbq/goutils/assert"
)

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
