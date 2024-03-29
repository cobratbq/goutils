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
