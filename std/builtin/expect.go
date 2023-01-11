// SPDX-License-Identifier: AGPL-3.0-or-later

package builtin

import "github.com/cobratbq/goutils/assert"

// Expect checks for error and either panics on error, or passes through only the result.
func Expect[T any](result T, err error) T {
	assert.Success(err, "unexpected failure")
	return result
}

// Expect2 checks for error and either panics on error, or passes through only the results.
func Expect2[T any, T2 any](result T, result2 T2, err error) (T, T2) {
	assert.Success(err, "unexpected failure")
	return result, result2
}

// Expect3 checks for error and either panics on error, or passes through only the results.
func Expect3[T, T2, T3 any](result T, result2 T2, result3 T3, err error) (T, T2, T3) {
	assert.Success(err, "unexpected failure")
	return result, result2, result3
}

// Expect4 checks for error and either panics on error, or passes through only the results.
func Expect4[T, T2, T3, T4 any](result T, result2 T2, result3 T3, result4 T4, err error) (T, T2, T3, T4) {
	assert.Success(err, "unexpected failure")
	return result, result2, result3, result4
}

// Error drops the result, returning only the error
func Error[T any](result T, err error) error {
	return err
}

// Error2 drops two results, returning only the error
func Error2[T, T2 any](result T, result2 T2, err error) error {
	return err
}

// Error3 drops three results, returning only the error
func Error3[T, T2, T3 any](result T, result2 T2, result3 T3, err error) error {
	return err
}

// Error4 drops four results, returning only the error
func Error4[T, T2, T3, T4 any](result T, result2 T2, result3 T3, result4 T4, err error) error {
	return err
}
