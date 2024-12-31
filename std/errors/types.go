// SPDX-License-Identifier: LGPL-3.0-only

package errors

import (
	"strconv"

	strconv_ "github.com/cobratbq/goutils/std/strconv"
)

/* Design notes:
 *
 * - Struct-with-private-field provides immutability of the error instances.
 * - Use of structs, or more specifically struct-pointers, provide unique identities for sentinel values.
 * - No additional fields for context or "user messages", such that sentinel errors can be constructed, i.e.
 *   the instance itself identifies a particular error. Only one instance exists, pre-constructed.
 * - Context must be provided by wrapping the appropriate sentinel error. (See `Context(err, message)`.)
 * - Sentinel errors are used to signal one specific error(-case). To abstract away from specific errors, e.g.
 *   that cross an architectural boundary, or to aggregate multiple errors together, use
 *   `Aggregate(cause, message, errors...)`, s.t. one matches on `cause` but the message shows how many and
 *   which particular errors were aggregated e.g. as indication of failed multi-stage processing.
 */

// StringError as a base type for const errors.
//
// This type is intended to be used as replacement for errors.New(..) from std, such that you can
// define an error as const (constant). The idea being that the "root" error type is just the type
// and the circumstances within which the error occurs are dictated by any number of contexts
// wrapped around the root error.
type StringError struct{ v string }

// NewStringError creates a new string-based error instance. If used as-is, the pointer that is
// returned is uniquely-identifying and therefore immediately useable. In case the StringError is a
// basis for a custom error type, the pointer can be dereferenced to include the value itself in the
// newly defined struct-type, for efficiency.
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

// NewUintError creates a new uint-based error instance. If used as-is, the pointer that is returned
// is uniquely-identifying and therefore immediately useable. In case UintError is a basis for a
// custom error type, the pointer can be dereferenced as to include the value itself in the newly
// defined struct-type, for efficiency.
func NewUintError(value uint) *UintError {
	return &UintError{value}
}

func (e *UintError) Error() string {
	return strconv.FormatUint(uint64(e.v), strconv_.DecimalBase)
}

// IntError as a base type for const errors.
//
// Similar to StringError, this type can be used to declare const errors. This type is based on int,
// therefore most suitable for errors that are signaled through a numeric code, such as with
// HTTP-like protocols.
type IntError struct{ v int }

// NewIntError creates a new int-based error instance. If used as-is, the pointer that is returned
// is uniquely-identifying and therefore immediately useable. In case the IntError is a basis for a
// custom error type, the pointer can be dereferenced to include the value itself in the newly
// defined struct-type, for efficiency.
func NewIntError(v int) *IntError {
	return &IntError{v}
}

func (e *IntError) Error() string {
	return strconv.FormatInt(int64(e.v), strconv_.DecimalBase)
}
