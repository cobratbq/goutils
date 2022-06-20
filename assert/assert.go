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

func Equal[T comparable](v1, v2 T) {
	if v1 != v2 {
		panic("assertion failed: Equal")
	}
}

// Expect checks for error and either panics on error, or passes through result.
func Expect[T any](result T, err error) T {
	Success(err, "unexpected failure encountered")
	return result
}
