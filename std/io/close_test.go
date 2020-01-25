package io

import (
	"errors"
	"io/ioutil"
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

type closefailer struct{}

func (closefailer) Close() error {
	return errors.New("bad shit happened")
}