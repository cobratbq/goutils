//go:build !enable_trace

package log

func Traceln(args ...any) {}

func TracelnDepth(depth uint, args ...any) {}

func Tracef(format string, args ...any) {}

func TracefDepth(depth uint, format string, args ...any) {}
