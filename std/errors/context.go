// SPDX-License-Identifier: LGPL-3.0-only

package errors

import (
	"github.com/cobratbq/goutils/assert"
)

// Aggregate creates a context for a root cause, with any number of additional errors.
func Aggregate(cause error, message string, errs ...error) error {
	assert.Required(cause, "Primary cause cannot be nil")
	return &contextN{message: message, causes: append([]error{cause}, errs...)}
}

type contextN struct {
	causes  []error
	message string
}

func (c *contextN) Error() string {
	return c.message + " ([" + JoinMessages(c.causes, "],[") + "])"
}

func (c *contextN) Unwrap() []error {
	return c.causes
}

// Context wraps an error to provide additional context information.
// TODO consider defining different variations of context: with-message, with-key-values-pairs, ... We can consider this basic context type as key-value pair with key 'message'. Then as we extract key-value pairs, we can include base contexts.
func Context(cause error, message string) error {
	assert.Required(cause, "Context expects to wrap an error but got nil")
	return &context{cause: cause, message: message}
}

type context struct {
	cause   error
	message string
}

func (c *context) Error() string {
	return c.message + ": " + c.cause.Error()
}

func (c *context) Unwrap() error {
	return c.cause
}
