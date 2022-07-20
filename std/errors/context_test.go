// SPDX-License-Identifier: LGPL-3.0-or-later

package errors

import (
	"errors"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestContext(t *testing.T) {
	var err = NewStringError("I am the error used for testing")
	singleWrap := Context(err, "Singly-wrapped context message")
	assert.Equal(t, "Singly-wrapped context message: I am the error used for testing",
		singleWrap.Error())
	assert.True(t, err == errors.Unwrap(singleWrap))
	assert.True(t, Is(singleWrap, err))
	doubleWrap := Context(singleWrap, "Doubly-wrapped")
	assert.Equal(t, "Doubly-wrapped: Singly-wrapped context message: I am the error used for testing",
		doubleWrap.Error())
	assert.True(t, singleWrap == errors.Unwrap(doubleWrap))
	assert.True(t, err == errors.Unwrap(errors.Unwrap(doubleWrap)))
	assert.True(t, Is(doubleWrap, err))
}

func TestContextNilError(t *testing.T) {
	defer assert.RequirePanic(t)
	Context(nil, "providing context to nil")
	t.FailNow()
}
