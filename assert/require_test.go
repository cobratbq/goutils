// SPDX-License-Identifier: AGPL-3.0-or-later

//go:build !disable_assert

package assert

import (
	"errors"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestRequireFailed(t *testing.T) {
	defer assert.RequirePanic(t)
	Require(false, "")
	t.FailNow()
}

func TestRequireSuccess(t *testing.T) {
	Require(true, "Uh oh!")
}

func TestRequireNonNilNil(t *testing.T) {
	defer assert.RequirePanic(t)
	Required(nil, "This should fail ...")
	t.FailNow()
}

func TestRequireNonNil(t *testing.T) {
	Required("Hello", "World!")
}

func TestRequireSuccessFails(t *testing.T) {
	defer assert.RequirePanic(t)
	RequireSuccess(errors.New("bad stuff happened"), "We expected to receive no error!")
	t.FailNow()
}

func TestRequireSuccessSucceeds(t *testing.T) {
	RequireSuccess(nil, "Everything should have been fine!")
}

func TestRequireSuccessFormatSpecifier(t *testing.T) {
	defer assert.RequirePanic(t)
	RequireSuccess(errors.New("Test"), "BUG: this message is part of the panic: %+v")
	t.FailNow()
}

func TestUnreachable(t *testing.T) {
	defer assert.RequireRecover(t)
	Unreachable()
	t.FailNow()
}

func TestUnsupported(t *testing.T) {
	defer assert.RequireRecover(t)
	Unsupported("To be implemented")
	t.FailNow()
}
