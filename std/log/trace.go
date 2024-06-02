// SPDX-License-Identifier: LGPL-3.0-only

//go:build enable_trace

package log

import (
	"fmt"
	"log"
	"os"
)

var tracelog = log.New(os.Stderr, "[trace] ", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)

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
		tracelog.Output(calldepth, fmt.Sprintln(prefix, k, "->", v))
	}
}

func TraceReportln(assert bool, message string) {
	if !assert {
		tracelog.Output(calldepth, fmt.Sprintln("Failed assertion:", message))
	}
}
