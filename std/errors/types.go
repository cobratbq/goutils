// SPDX-License-Identifier: LGPL-3.0-or-later

package errors

import "strconv"

// StringError as a base type for const errors.
//
// This type is intended to be used as replacement for errors.New(..) from std, such that you can
// define an error as const (constant). The idea being that the "root" error type is just the type
// and the circumstances within which the error occurs are dictated by any number of contexts
// wrapped around the root error.
type StringError struct{ v string }

func NewStringError(msg string) *StringError {
	return &StringError{msg}
}

func (e *StringError) Error() string {
	return string(e.v)
}

// UintError as a base type for const errors.
//
// Similar to StringError, this type can be used to declare const errors. This type is based on uint
// therefore is most suitable for errors that are signaled through a numeric code, such as with
// HTTP-like protocols.
type UintError struct{ v uint }

func NewUintError(value uint) *UintError {
	return &UintError{value}
}

func (e *UintError) Error() string {
	return strconv.FormatUint(uint64(e.v), 10)
}

// IntError as a base type for const errors.
//
// Similar to StringError, this type can be used to declare const errors. This type is based on int,
// therefore most suitable for errors that are signaled through a numeric code, such as with
// HTTP-like protocols.
type IntError struct{ v int }

func NewIntError(v int) *IntError {
	return &IntError{v}
}

func (e *IntError) Error() string {
	return strconv.FormatInt(int64(e.v), 10)
}
