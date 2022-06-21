package errors

import "errors"

// Context wraps an error to provide additional context information.
// TODO consider defining different variations of context: with-message, with-key-values-pairs, ...
func Context(cause error, message string) error {
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
