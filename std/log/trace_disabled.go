// SPDX-License-Identifier: LGPL-3.0-only

//go:build !enable_trace

package log

func Tracing() bool {
	return false
}

func Traceln(args ...any) {}

func TracelnDepth(depth uint, args ...any) {}

func Tracef(format string, args ...any) {}

func TracefDepth(depth uint, format string, args ...any) {}

func TracelnSliceAsString(prefix string, data [][]byte) {}

func TracelnMap[K comparable, V any](prefix string, data map[K]V) {}

func TraceReportln(assert bool, message string) {}
