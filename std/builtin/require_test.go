package builtin

import (
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
