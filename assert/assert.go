// SPDX-License-Identifier: LGPL-3.0-or-later
package assert

func False(expected bool) {
	if expected {
		panic("assertion failed: False")
	}
}

func True(expected bool) {
	if !expected {
		panic("assertion failed: True")
	}
}

func Any[T comparable](actual T, values ...T) {
	for _, v := range values {
		if actual == v {
			return
		}
	}
	panic("assertion failed: expected one of specified values")
}

func Equal[T comparable](v1, v2 T) {
	if v1 != v2 {
		panic("assertion failed: Equal")
	}
}

// Expect checks for error and either panics on error, or passes through only the result.
// TODO should we do Expect with our without parameter for error message?
func Expect[T any](result T, err error) T {
	Success(err, "unexpected failure encountered")
	return result
}

// Expect checks for error and either panics on error, or passes through only the result.
func Expect2[T any, T2 any](result T, result2 T2, err error) (T, T2) {
	Success(err, "unexpected failure encountered")
	return result, result2
}

// Expect checks for error and either panics on error, or passes through only the result.
func Expect3[T, T2, T3 any](result T, result2 T2, result3 T3, err error) (T, T2, T3) {
	Success(err, "unexpected failure encountered")
	return result, result2, result3
}
