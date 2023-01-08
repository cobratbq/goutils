// SPDX-License-Identifier: LGPL-3.0-or-later

package errors

import (
	"github.com/cobratbq/goutils/assert"
)

// Aggregate creates a context for a root cause, based on any number of errors.
// TODO consider if we need to embed the actual errors. This touches on a tangential consideration: how to handle errors crossing abstraction boundaries. Do we want to make them available, i.e. leak the abstraction?
func Aggregate(cause error, message string, errs ...error) error {
	return Context(cause, message+" (["+JoinMessages(errs, "],[")+"])")
}

// Context wraps an error to provide additional context information.
// TODO consider defining different variations of context: with-message, with-key-values-pairs, ... We can consider this basic context type as key-value pair with key 'message'. Then as we extract key-value pairs, we can include base contexts.
func Context(cause error, message string) error {
	assert.Required(cause, "Context expects to wrap an error but got nil")
	return context{cause: cause, message: message}
}

type context struct {
	cause   error
	message string
}

func (c context) Error() string {
	return c.cause.Error() + ": " + c.message
}

func (c context) Unwrap() error {
	return c.cause
}
