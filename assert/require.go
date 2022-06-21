// SPDX-License-Identifier: LGPL-3.0-or-later
package assert

import (
	"fmt"
	"log"
)

// RequireSuccess checks that err is nil. If the error is non-nil, it will panic. `message` can have
// '%v' format specifier so that it can be substituted with the error message.
// Deprecated: `RequireSuccess` deprecated in favor of `Success`.
func RequireSuccess(err error, message string) {
	log.Println("assert.RequireSuccess is deprecated. Use assert.Success")
	Success(err, message)
}

// Success checks that err is nil. If the error is non-nil, it will panic. `message` can have '%v'
// format specifier so that it can be substituted with the error message.
func Success(err error, message string) {
	Require(err == nil, fmt.Sprintf(message, err))
}

// Require check required condition and panics if condition does not hold.
func Require(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

// RequireNonNil checks if provided value is nil, if so panics with provided message.
func RequireNonNil(val interface{}, message string) {
	if val == nil {
		panic(message)
	}
}

// Unreachable panics with a message to indicate this should not be happening.
func Unreachable() {
	panic("BUG: this code should not be reachable.")
}

// Unsupported panics with the provided message in order to signal for an unsupported case.
func Unsupported(message string) {
	panic(message)
}
