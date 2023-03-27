//go:build !enable_trace

package log

func Traceln(args ...any) {}

func Tracef(format string, args ...any) {}
