package errors

import (
	"runtime/debug"

	"github.com/cobratbq/goutils/assert"
)

func Stacktrace(cause error) error {
	assert.Required(cause, "Stacktrace is expected to wrap an error")
	return stack{cause: cause, stack: debug.Stack()}
}

// TODO think about how we want to unwind and print stack trace, because wrapping multiple contexts means that we might produce multiple stack traces, of which the inner-most one is most accurate.

type stack struct {
	cause error
	stack []byte
}

func (s stack) Error() string {
	// FIXME should we, and how, include the stacktrace in the error string?
	return s.cause.Error() + "\n" + string(s.stack)
}

func (s stack) Unwrap() error {
	return s.cause
}

func (s stack) Stack() []byte {
	return s.stack
}
