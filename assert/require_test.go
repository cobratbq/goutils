// SPDX-License-Identifier: LGPL-3.0-or-later
package assert

import (
	"errors"
	"testing"
)

func TestRequireFailed(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected non-nil result from recovery.")
		}
	}()
	Require(false, "")
	t.FailNow()
}

func TestRequireSuccess(t *testing.T) {
	Require(true, "Uh oh!")
}

func TestRequireNonNilNil(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected non-nil result from recovery.")
		}
	}()
	RequireNonNil(nil, "This should fail ...")
	t.FailNow()
}

func TestRequireNonNil(t *testing.T) {
	RequireNonNil("Hello", "World!")
}

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

func TestUnreachable(t *testing.T) {
	defer func() {
		recover()
	}()
	Unreachable()
	t.FailNow()
}

func TestUnsupported(t *testing.T) {
	defer func() {
		recover()
	}()
	Unsupported("To be implemented")
	t.FailNow()
}
