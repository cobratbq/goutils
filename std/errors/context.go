package errors

import (
	"errors"
	"runtime/debug"
)

func ContextStacktrace(cause error, message string) context {
	return Context(stack{cause: cause, stack: debug.Stack()}, message)
}

// Context wraps an error to provide additional context information.
// TODO consider defining different variations of context: with-message, with-key-values-pairs, ...
func Context(cause error, message string) context {
	// TODO should we return the pointer instead?
	return context{cause: cause, message: message}
}

type context struct {
	cause   error
	message string
}

func (c context) Error() string {
	return c.message + ": " + c.cause.Error()
}

func (c context) Unwrap() error {
	return c.cause
}

// TODO think about how we want to unwind and print stack trace, because wrapping multiple contexts means that we might produce multiple stack traces, of which the inner-most one is most accurate.

type stack struct {
	cause error
	stack []byte
}

func (s stack) Error() string {
	// FIXME continue here
	panic("to be implemented")
}

func (s stack) Unwrap() error {
	return s.cause
}

func (s stack) Stack() []byte {
	return s.stack
}

// Is repeatedly unwraps an error and compares to target on each unwrapping.
// Is uses the implementation from std/errors.
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// Unwrap unwraps an error if `Unwrap() error` exists, or returns nil otherwise.
// Unwrap uses the implementation from std/errors.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
