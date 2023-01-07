// SPDX-License-Identifier: LGPL-3.0-or-later

package errors

import (
	"errors"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestContext(t *testing.T) {
	var err = NewStringError("I am the error used for testing")
	singleWrap := Context(err, "First addition of context information")
	assert.Equal(t, "I am the error used for testing: First addition of context information",
		singleWrap.Error())
	assert.True(t, err == errors.Unwrap(singleWrap))
	assert.True(t, Is(singleWrap, err))
	doubleWrap := Context(singleWrap, "Second addition of context")
	assert.Equal(t, "I am the error used for testing: First addition of context information: Second addition of context",
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
