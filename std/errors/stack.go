// SPDX-License-Identifier: LGPL-3.0-or-later

package errors

import (
	"runtime/debug"

	"github.com/cobratbq/goutils/assert"
)

func Stacktrace(cause error) error {
	assert.Required(cause, "Stacktrace expects to wrap an error but got nil")
	return stack{cause: cause, stack: debug.Stack()}
}

type stack struct {
	cause error
	stack []byte
}

func (s stack) Error() string {
	return s.cause.Error() + "\n" + string(s.stack)
}

func (s stack) Unwrap() error {
	return s.cause
}

func (s stack) Stack() []byte {
	return s.stack
}
