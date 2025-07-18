// SPDX-License-Identifier: LGPL-3.0-only

//go:build tracelog

package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

var tracelog = log.New(os.Stderr, "\033[1;30m[trace]\033[0m ", log.Ltime|log.LUTC|log.Lmicroseconds|log.Lshortfile)

func Tracing() bool {
	return true
}

// Traceln (if `enable_trace`) logs the provided arguments. Included in the logging output is the
// file name and line number of the immediate caller.
func Traceln(args ...any) {
	tracelog.Output(calldepth, fmt.Sprintln(args...))
}

// TracelnDepth (if `enable_trace`) logs the provided arguments. Included in the logging output is
// the file name and line number of the immediate caller. The caller to be included is modified with
// the `depth` parameter. By default the immediate caller is logged.
func TracelnDepth(depth uint, args ...any) {
	tracelog.Output(calldepth+int(depth), fmt.Sprintln(args...))
}

// Tracef (if `enable_trace`) logs the provided arguments in specified format. Included in the
// logging output is the file name and line number of the immediate caller.
func Tracef(format string, args ...any) {
	tracelog.Output(calldepth, fmt.Sprintf(format, args...))
}

// TracefDepth (if `enable_trace`) logs the provided arguments in specified format. Included in the
// logging output is the file name and line number of the (in)direct caller, as modified with
// parameter `depth`. (Default value 0 means the immediate caller.)
func TracefDepth(depth uint, format string, args ...any) {
	tracelog.Output(calldepth+int(depth), fmt.Sprintf(format, args...))
}

// TracelnSlice prints each entry in `data` on a new line. Every trace-line is prefixed with `prefix`.
func TracelnSlice[T any](prefix string, data []T) {
	for i, e := range data {
		tracelog.Output(calldepth, fmt.Sprintln(prefix, i, "->", e))
	}
}

// TracelnSliceAsString prints each entry in `data` on a new line. Every trace-line is prefixed with `prefix`.
// Each line of data (bytes), are printed as ANSI characters, thus converted to a string.
func TracelnSliceAsString(prefix string, data [][]byte) {
	for i, e := range data {
		tracelog.Output(calldepth, fmt.Sprintln(prefix, i, "->", string(e)))
	}
}

// TracelnMap prints each entry in `data` (every key in the map) on a new line. Every trace-line is prefixed
// with `prefix`.
func TracelnMap[K comparable, V any](prefix string, data map[K]V) {
	for k, v := range data {
		tracelog.Output(calldepth, fmt.Sprintf("%v: %+v -> %+v", prefix, k, v))
	}
}

// TraceReport logs a trace-message in case the assertion does not hold.
func TraceReport(assert bool, format string, args ...any) {
	if !assert {
		tracelog.Output(calldepth, fmt.Sprintf("Failed assertion: "+format+"\n", args...))
	}
}

// TraceFlags returns the flags set for trace-logging. (If trace-logging is enabled.)
func TraceFlags() int {
	return tracelog.Flags()
}

// SetTraceFlags sets the flags for trace-logging. (If trace-logging is enabled.)
func SetTraceFlags(flags int) {
	tracelog.SetFlags(flags)
}

// SetTraceOutput sets the output writer for trace-logging. (If trace-logging is enabled.)
func SetTraceOutput(output io.Writer) {
	tracelog.SetOutput(output)
}
