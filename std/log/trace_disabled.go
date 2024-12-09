// SPDX-License-Identifier: LGPL-3.0-only

//go:build !tracelog

package log

import "io"

func Tracing() bool {
	return false
}

func Traceln(args ...any) {}

func TracelnDepth(depth uint, args ...any) {}

func Tracef(format string, args ...any) {}

func TracefDepth(depth uint, format string, args ...any) {}

func TracelnSlice[T any](prefix string, data []T) {}

func TracelnSliceAsString(prefix string, data [][]byte) {}

func TracelnMap[K comparable, V any](prefix string, data map[K]V) {}

func TraceReport(assert bool, format string, args ...any) {}

func TraceFlags() int { return 0 }

func SetTraceFlags(flags int) {}

func SetTraceOutput(output io.Writer) {}
