// SPDX-License-Identifier: LGPL-3.0-or-later
package errors

import "errors"

// Is repeatedly unwraps an error and compares to target on each unwrapping.
// Is uses the implementation from std/errors.
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// Stack extracts the first stacktrace encountered in a wrapped error, or nil if no stack is
// present/found. It is assumed that, generally, at most one stacktrace is present.
func Stack(err error) []byte {
	for ; err != nil; err = errors.Unwrap(err) {
		if trace, ok := err.(stack); ok {
			return trace.Stack()
		}
	}
	return nil
}

// Unwrap unwraps an error if `Unwrap() error` exists, or returns nil otherwise.
// Unwrap uses the implementation from std/errors.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
