// SPDX-License-Identifier: LGPL-3.0-or-later
package io

import (
	"bytes"
	"io"
	"testing"
)

func TestMustReadAllNil(t *testing.T) {
	defer func() { recover() }()
	MustReadAll(nil)
	t.FailNow()
}

func TestMustReadAll(t *testing.T) {
	r := bytes.NewReader([]byte("hello world"))
	if !bytes.Equal(MustReadAll(r), []byte("hello world")) {
		t.FailNow()
	}
}

func TestMustReadAllBadReader(t *testing.T) {
	defer func() { recover() }()
	MustReadAll(&badreader{})
	t.FailNow()
}

type badreader struct{}

func (*badreader) Read(b []byte) (n int, err error) {
	return 2, io.ErrUnexpectedEOF
}

func TestDiscardNil(t *testing.T) {
	defer func() { recover() }()
	Discard(nil)
	t.FailNow()
}

func TestDiscard(t *testing.T) {
	r := bytes.NewReader([]byte("hello world"))
	if Discard(r) != 11 {
		t.FailNow()
	}
}
