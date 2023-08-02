// SPDX-License-Identifier: LGPL-3.0-only

package builtin

import (
	"log"
	"runtime/debug"
)

// RecoverLogged recovers from panic and logs the payload.
// This function must be _deferred_ to be useful.
func RecoverLogged(format string) {
	if m := recover(); m != nil {
		log.Printf(format, m)
	}
}

// RecoverLoggedStackTrace calls `recover()` and in case the goroutine was
// indeed panicking, it writes the error message and a stack trace to
// os.Stderr.
// This function must be _deferred_ to be useful.
func RecoverLoggedStackTrace(format string) {
	stacktrace := debug.Stack()
	if v := recover(); v != nil {
		log.Printf(format, v)
		log.Writer().Write(stacktrace)
	}
}
