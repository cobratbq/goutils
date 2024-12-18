// SPDX-License-Identifier: LGPL-3.0-only

package io

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestCloseDiscardedNil(t *testing.T) {
	defer assert.RequirePanic(t)
	CloseIgnored(nil)
	t.FailNow()
}

func TestCloseDiscardedGraceful(t *testing.T) {
	c := io.NopCloser(nil)
	CloseIgnored(c)
}

func TestCloseDiscardedFailure(t *testing.T) {
	CloseIgnored(&closefailer{})
}

func TestCloseLoggedNilCloser(t *testing.T) {
	defer assert.RequirePanic(t)
	CloseLogged(nil, "Uh oh!")
	t.FailNow()
}

func TestCloseLoggedSuccess(t *testing.T) {
	CloseLogged(io.NopCloser(nil), "failed to close no-op")
}

func TestCloseLoggedFailure(t *testing.T) {
	CloseLogged(&closefailer{}, "correctly failed to close: %+v")
}

func TestCloseLoggedErrClosedPipe(t *testing.T) {
	// FIXME cannot properly test that this does indeed log an error.
	CloseLogged(&closeclosedpipe{}, "correctly failed to close: %+v")
}

func TestCloseLoggedWithIgnoresErrClosedPipe(t *testing.T) {
	// FIXME cannot properly test that this does not log an error.
	CloseLoggedWithIgnores(&closeclosedpipe{}, "correctly failed to close: %+v", io.ErrClosedPipe)
}

func TestClosePanickedNil(t *testing.T) {
	defer assert.RequirePanic(t)
	ClosePanicked(nil, "failed to close nil: %+v")
	t.FailNow()
}

func TestClosePanickedSuccess(t *testing.T) {
	ClosePanicked(io.NopCloser(nil), "closing successfully, right ...")
}

func TestClosePanickedFailure(t *testing.T) {
	defer assert.RequirePanic(t)
	ClosePanicked(&closefailer{}, "failed to close: %+v")
	t.FailNow()
}

func TestClosePanickedErrClosedPipe(t *testing.T) {
	defer assert.RequirePanic(t)
	ClosePanicked(&closeclosedpipe{}, "failed to close: %+v")
	t.FailNow()
}

func TestClosePanickedWithIgnoresErrClosedPipe(t *testing.T) {
	ClosePanickedWithIgnores(&closeclosedpipe{}, "failed to close: %+v", io.ErrClosedPipe)
}

type closefailer struct{}

func (closefailer) Close() error {
	return errors.New("bad shit happened")
}

type closeclosedpipe struct{}

func (closeclosedpipe) Close() error {
	return io.ErrClosedPipe
}

func TestNopCloserRepeatedClose(t *testing.T) {
	closer := &NopCloser{os.Stderr}
	err := closer.Close()
	if err != nil {
		t.FailNow()
	}
	err = closer.Close()
	if err != nil {
		t.FailNow()
	}
	err = closer.Close()
	if err != nil {
		t.FailNow()
	}
}

func TestNopCloserReadPassthrough(t *testing.T) {
	b := bytes.NewBufferString("Hello world!")
	c := NopCloser{b}
	data := make([]byte, 12)
	c.Read(data)
	if !bytes.Equal([]byte("Hello world!"), data) {
		t.FailNow()
	}
}

func TestNopCloserWritePassthrough(t *testing.T) {
	b := bytes.NewBuffer(nil)
	c := NopCloser{b}
	c.Write([]byte("Hello world!"))
	if b.String() != "Hello world!" {
		t.FailNow()
	}
}
