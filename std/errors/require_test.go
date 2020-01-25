package errors

import (
	"errors"
	"testing"
)

func TestRequireSuccessFails(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected non-nil result from recovery.")
		}
	}()
	RequireSuccess(errors.New("bad stuff happened"), "We expected to receive no error!")
	t.FailNow()
}

func TestRequireSuccessSucceeds(t *testing.T) {
	RequireSuccess(nil, "Everything should have been fine!")
}

func TestRequireSuccessFormatSpecifier(t *testing.T) {
	defer func() {
		m := recover().(string)
		if m != "BUG: this message is part of the panic: Test" {
			t.FailNow()
		}
	}()
	RequireSuccess(errors.New("Test"), "BUG: this message is part of the panic: %+v")
	t.FailNow()
}
