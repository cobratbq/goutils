package io

import (
	"github.com/cobratbq/goutils/std/errors"
	"io"
	"io/ioutil"
)

// MustReadAll reads all data from reader and panics in case an error occurs.
// FIXME write unit tests
func MustReadAll(r io.Reader) []byte {
	data, err := ioutil.ReadAll(r)
	errors.RequireSuccess(err, "Failed to read all data from reader: %+v")
	return data
}
