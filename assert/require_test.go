// SPDX-License-Identifier: LGPL-3.0-only

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

func TestSuccessFails(t *testing.T) {
	defer assert.RequirePanic(t)
	Success(errors.New("bad stuff happened"), "We expected to receive no error!")
	t.FailNow()
}

func TestSuccessSucceeds(t *testing.T) {
	Success(nil, "Everything should have been fine!")
}

func TestSuccessFormatSpecifier(t *testing.T) {
	defer assert.RequirePanic(t)
	Success(errors.New("Test"), "BUG: this message is part of the panic: %+v")
	t.FailNow()
}
