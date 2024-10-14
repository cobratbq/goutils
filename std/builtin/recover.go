// SPDX-License-Identifier: LGPL-3.0-only

package builtin

import (
	"os"
	"runtime/debug"

	"github.com/cobratbq/goutils/std/log"
)

// RecoverLogged recovers from panic and logs the payload.
// This function must be _deferred_ to be useful.
func RecoverLogged(format string) {
	if m := recover(); m != nil {
		log.Errorf(format, m)
	}
}

// RecoverLoggedExit recovers from panic and logs the payload, same as RecoverLogged, except that calls
// `os.Exit(code)` after recovering a non-nil (panic-)value to abnormally terminate the program after logging
// the unhandled panic.
// This function must be _deferred_ to be useful.
func RecoverLoggedExit(code int, format string) {
	if m := recover(); m != nil {
		log.Errorf(format, m)
		os.Exit(code)
	}
}

// RecoverLoggedStackTrace calls `recover()` and in case the goroutine was indeed panicking, it writes the
// error message and a stack trace to log.
// This function must be _deferred_ to be useful.
func RecoverLoggedStackTrace(format string) {
	stacktrace := debug.Stack()
	if v := recover(); v != nil {
		log.Errorf(format, v)
		log.Debugln(string(stacktrace))
	}
}

// RecoverLoggedStackTraceExit calls `recover()` and in case the goroutine was indeed panicking, it writes the
// error message and a stack trace to log. After logging the panic, exit with specified exit-code.
// This function must be _deferred_ to be useful.
func RecoverLoggedStackTraceExit(code int, format string) {
	stacktrace := debug.Stack()
	if v := recover(); v != nil {
		log.Errorf(format, v)
		log.Debugln(string(stacktrace))
		os.Exit(code)
	}
}
