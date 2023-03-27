//go:build enable_trace

package log

import (
	"log"
	"os"
)

var tracelog = log.New(os.Stderr, "[TRACE]", log.Ldate|log.Ltime|log.Llongfile)

func Traceln(args ...any) {
	tracelog.Println(args...)
}

func Tracef(format string, args ...any) {
	tracelog.Printf(format, args...)
}
