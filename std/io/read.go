package io

import (
	"github.com/cobratbq/goutils/std/errors"
	"io"
	"io/ioutil"
)

// MustReadAll reads all data from reader and panics in case an error occurs.
func MustReadAll(r io.Reader) []byte {
	data, err := ioutil.ReadAll(r)
	errors.RequireSuccess(err, "Failed to read all data from reader: %+v")
	return data
}

// Discard reads remaining data from reader and discards it. Any possible
// errors in the process are ignored. Returns nr of bytes written, thus
// discarded.
func Discard(r io.Reader) int64 {
	n, _ := io.Copy(ioutil.Discard, r)
	return n
}
