// SPDX-License-Identifier: LGPL-3.0-or-later
package io

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestCloseDiscardedNil(t *testing.T) {
	defer func() {
		v := recover()
		if v == nil {
			t.Error("Expected to have an actual panic.")
		}
	}()
	CloseDiscarded(nil)
	t.FailNow()
}

func TestCloseDiscardedGraceful(t *testing.T) {
	c := ioutil.NopCloser(nil)
	CloseDiscarded(c)
}

func TestCloseDiscardedFailure(t *testing.T) {
	CloseDiscarded(&closefailer{})
}

func TestCloseLoggedNilCloser(t *testing.T) {
	defer func() {
		v := recover()
		if v == nil {
			t.Error("Expected to have an actual panic.")
		}
	}()
	CloseLogged(nil, "Uh oh!")
	t.FailNow()
}

func TestCloseLoggedSuccess(t *testing.T) {
	CloseLogged(ioutil.NopCloser(nil), "failed to close no-op")
}

func TestCloseLoggedFailure(t *testing.T) {
	CloseLogged(&closefailer{}, "correctly failed to close: %+v")
}

func TestCloseLoggedErrClosedPipe(t *testing.T) {
	CloseLogged(&closeclosedpipe{}, "correctly failed to close: %+v")
}

func TestClosePanickedNil(t *testing.T) {
	defer func() {
		v := recover()
		if v == nil {
			t.Error("Expected to have an actual panic.")
		}
	}()
	ClosePanicked(nil, "failed to close nil: %+v")
	t.FailNow()
}

func TestClosePanickedSuccess(t *testing.T) {
	ClosePanicked(ioutil.NopCloser(nil), "closing successfully, right ...")
}

func TestClosePanickedFailure(t *testing.T) {
	defer func() {
		v := recover()
		if v != "failed to close: bad shit happened" {
			t.Error("Expected to have an actual panic.")
		}
	}()
	ClosePanicked(&closefailer{}, "failed to close: %+v")
	t.FailNow()
}

func TestClosePanickedErrClosedPipe(t *testing.T) {
	ClosePanicked(&closeclosedpipe{}, "failed to close: %+v")
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
