// SPDX-License-Identifier: LGPL-3.0-or-later
package errors

import (
	"errors"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestStringError(t *testing.T) {
	var err = NewStringError("hello world!")
	var err2 = NewStringError("hello world!")
	assert.Equal(t, err.Error(), "hello world!")
	assert.Nil(t, errors.Unwrap(err))
	assert.True(t, errors.Is(err, err))
	assert.Unequal(t, err, err2)
	assert.False(t, errors.Is(err2, err))
}

func TestUintError(t *testing.T) {
	var err = NewUintError(99)
	var err2 = NewUintError(99)
	assert.Equal(t, err.Error(), "99")
	assert.Nil(t, errors.Unwrap(err))
	assert.True(t, errors.Is(err, err))
	assert.Unequal(t, err, err2)
	assert.False(t, errors.Is(err2, err))
}

func TestIntError(t *testing.T) {
	var err = NewIntError(-42)
	var err2 = NewIntError(-42)
	assert.Equal(t, err.Error(), "-42")
	assert.Nil(t, errors.Unwrap(err))
	assert.True(t, errors.Is(err, err))
	assert.Unequal(t, err, err2)
	assert.False(t, errors.Is(err2, err))
}
