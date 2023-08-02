// SPDX-License-Identifier: LGPL-3.0-only

//go:build disable_assert

package assert

// Success checks that err is nil. If the error is non-nil, it will panic. `message` can have '%v'
// format specifier so that it can be substituted with the error message.
func Success(err error, message string) {}

// Required checks if provided value is nil, if so panics with provided message.
func Required(val any, message string) {}

// Require check required condition and panics if condition does not hold.
func Require(condition bool, message string) {}

// Unreachable panics with a message to indicate this should not be happening.
// TODO how useful is this? Go compiler/static analysis cannot detect that this function call is terminal
func Unreachable() {}

// Unsupported panics with the provided message in order to signal for an unsupported case.
func Unsupported(message string) {}
