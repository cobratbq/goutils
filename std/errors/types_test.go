// SPDX-License-Identifier: LGPL-3.0-or-later

package errors

import (
	"errors"
	"testing"
	"unsafe"

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

func TestExpectWrappedErrorSameSize(t *testing.T) {
	var errBase = NewStringError("Hello world")
	var errWrapped = &struct{ StringError }{*NewStringError("Hello world")}
	assert.True(t, error(errBase) != error(errWrapped))
	assert.Equal(t, unsafe.Sizeof(*errBase), unsafe.Sizeof(*errWrapped))
}

func TestBenchmarkAllocationlessWrappedError(t *testing.T) {
	type httpStatusCodeError struct {
		UintError
	}
	report := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errBase = NewUintError(401)
			var errForbidden = &httpStatusCodeError{*errBase}
			assert.True(b, error(errBase) != error(errForbidden))
			assert.Equal(b, unsafe.Sizeof(*errBase), unsafe.Sizeof(*errForbidden))
		}
	})
	assert.Equal(t, report.MemAllocs, 0)
	assert.Equal(t, report.AllocsPerOp(), 0)
	assert.Equal(t, report.AllocedBytesPerOp(), 0)
}
