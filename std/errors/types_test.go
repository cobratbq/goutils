// SPDX-License-Identifier: AGPL-3.0-or-later

package errors

import (
	"errors"
	"testing"
	"unsafe"

	"github.com/cobratbq/goutils/std/log"
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
	// TODO unfortunately, this test seems to be a bit flaky. Running the test after clearing the test cache (`go clean -testcache`) will most-of-the-time result in a successful test. However incidentally there may be a spurious allocation. To be fixed ...
	type httpStatusCodeError struct {
		UintError
	}
	// Assume that the root error is predefined in a package variable.
	var errBase = NewUintError(401)
	// Now benchmark the wrapped use of the error, ensuring no allocations are performed. This test
	// ensure that properly wrapping existing root errors are possible without additional
	// allocations. This allows extending errors without suffering unexpected, subtle disadvantages.
	report := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var errForbidden = &httpStatusCodeError{*errBase}
			// assert that error-wrapping is effective without overhead
			assert.True(b, error(errBase) != error(errForbidden))
			assert.Equal(b, unsafe.Sizeof(*errBase), unsafe.Sizeof(*errForbidden))
		}
	})
	// Verify that there is no overhead in memory allocations.
	log.Debugln("MemAllocs")
	assert.Equal(t, report.MemAllocs, 0)
	log.Debugln("AllocsPerOp")
	assert.Equal(t, report.AllocsPerOp(), 0)
	log.Debugln("AllocatedByte")
	assert.Equal(t, report.AllocedBytesPerOp(), 0)
}
