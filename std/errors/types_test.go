package errors

import (
	"errors"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestStringError(t *testing.T) {
	const err StringError = "hello world!"
	assert.Equal(t, err.Error(), "hello world!")
	assert.Nil(t, errors.Unwrap(err))
	assert.True(t, errors.Is(err, err))
}

func TestUintError(t *testing.T) {
	const err UintError = 99
	assert.Equal(t, err.Error(), "99")
	assert.Nil(t, errors.Unwrap(err))
	assert.True(t, errors.Is(err, err))
}

func TestIntError(t *testing.T) {
	const err IntError = -42
	assert.Equal(t, err.Error(), "-42")
	assert.Nil(t, errors.Unwrap(err))
	assert.True(t, errors.Is(err, err))
}
