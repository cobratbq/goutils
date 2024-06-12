// SPDX-License-Identifier: LGPL-3.0-only

// A concise error handling package that provides the mechanisms for creating a "root" error that is
// a basic instance of String/Uint/Int, and the ability to add stacktrace and context to the root
// error. Existing "root" errors already present in the Go standard library can be used equally
// well.
//
// Ideally, following properties are desired:
//  1. uniquely-identifying error instances
//  2. constant/read-only instances, i.e. can be addressed but not changed
//  3. composable/extendable into dedicated error type/series for specific use cases (without
//     overhead)
//
// TODO above three properties cannot be satisfied yet.
// TODO do we need predefined errors? Errors such as `os.ErrInvalid` are defined specifically for the filesystem use-case. However, we need generic errors representing the various classes of incorrectness.
package errors

import (
	"errors"
	"strings"
)

// ErrIllegal indicates an illegal/bad value.
var ErrIllegal = NewStringError("illegal value")

// ErrInternalState indicates a problem with internal state, or use while in incorrect state.
var ErrInternalState = NewStringError("illegal state")

// ErrFailure indicates that there was a processing failure.
var ErrFailure = NewStringError("failure during processing")

// ErrOverflow indicates that the operation causes an overflow of some kind.
var ErrOverflow = NewStringError("overflowing")

// ErrUnderflow indicates that stack is empty, either at present or after the operation.
var ErrUnderflow = NewStringError("underflowing")

// ErrUnsupported indicates that an error occurred because of an unsupported operation or value.
var ErrUnsupported = NewStringError("unsupported")

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

// JoinMessages joins the error messages of each error, using the provided separator.
func JoinMessages(errs []error, sep string) string {
	var buf strings.Builder
	for i, err := range errs {
		if i > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(err.Error())
	}
	return buf.String()
}
