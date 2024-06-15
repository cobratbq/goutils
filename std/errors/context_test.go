// SPDX-License-Identifier: LGPL-3.0-only

package errors

import (
	"errors"
	"os"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestContext(t *testing.T) {
	var err = NewStringError("I am the error used for testing")
	singleWrap := Context(err, "First addition of context information")
	assert.Equal(t, "First addition of context information: I am the error used for testing",
		singleWrap.Error())
	assert.True(t, err == errors.Unwrap(singleWrap))
	assert.True(t, Is(singleWrap, err))
	doubleWrap := Context(singleWrap, "Second addition of context")
	assert.Equal(t, "Second addition of context: First addition of context information: I am the error used for testing",
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

func TestAggregateContext(t *testing.T) {
	nested1 := NewStringError("hello world failed big-time")
	nested2 := Context(NewUintError(500), "Server failure")
	nested3 := Context(os.ErrNotExist, "could not find unix socket connection for fancy plug-in")
	aggregate := Aggregate(os.ErrInvalid, "provided input is bad", nested1, nested2, nested3)
	assert.Equal(t, aggregate.Error(), "provided input is bad ([hello world failed big-time],[Server failure: 500],[could not find unix socket connection for fancy plug-in: file does not exist]): invalid argument")
}
