// SPDX-License-Identifier: LGPL-3.0-or-later

// assert provides various assertion functions that can be used to confirm certain conditions such
// that these conditions are guaranteed true afterwards. These functions are particularly useful to
// catch unexpected and unsupported use cases, without having to litter the code with if-statements.
// Assertions may be placeholders for use cases that will later be supported, or they may indicate
// failure conditions that will not or cannot ever be supported, or cannot even occur. Some use
// cases or possible error conditions are illusions created by the type-system, for example when a
// function implements an interface but will never fail the operation.
// Regardless, assertions allow you to (subtly) handle cases and failure conditions that you do not
// handle otherwise.
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

// Unwrap checks for error and either panics on error, or passes through only the result.
func Unwrap[T any](result T, err error) T {
	Success(err, "unexpected failure")
	return result
}

// Unwrap2 checks for error and either panics on error, or passes through only the results.
func Unwrap2[T any, T2 any](result T, result2 T2, err error) (T, T2) {
	Success(err, "unexpected failure")
	return result, result2
}

// Unwrap3 checks for error and either panics on error, or passes through only the results.
func Unwrap3[T, T2, T3 any](result T, result2 T2, result3 T3, err error) (T, T2, T3) {
	Success(err, "unexpected failure")
	return result, result2, result3
}

// Unwrap4 checks for error and either panics on error, or passes through only the results.
func Unwrap4[T, T2, T3, T4 any](result T, result2 T2, result3 T3, result4 T4, err error) (T, T2, T3, T4) {
	Success(err, "unexpected failure")
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
