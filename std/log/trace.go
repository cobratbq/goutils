//go:build enable_trace

package log

import (
	"fmt"
	"log"
	"os"
)

var tracelog = log.New(os.Stderr, "[TRACE] ", log.LstdFlags|log.Llongfile)

// Traceln (if `enable_trace`) logs the provided arguments. Included in the logging output is the
// file name and line number of the immediate caller.
func Traceln(args ...any) {
	tracelog.Output(2, fmt.Sprintln(args...))
}

// TracelnDepth (if `enable_trace`) logs the provided arguments. Included in the logging output is
// the file name and line number of the immediate caller. The caller to be included is modified with
// the `depth` parameter. By default the immediate caller is logged.
func TracelnDepth(depth uint, args ...any) {
	tracelog.Output(2+int(depth), fmt.Sprintln(args...))
}

// Tracef (if `enable_trace`) logs the provided arguments in specified format. Included in the
// logging output is the file name and line number of the immediate caller.
func Tracef(format string, args ...any) {
	tracelog.Output(2, fmt.Sprintf(format, args...))
}

// TracefDepth (if `enable_trace`) logs the provided arguments in specified format. Included in the
// logging output is the file name and line number of the (in)direct caller, as modified with
// parameter `depth`. (Default value 0 means the immediate caller.)
func TracefDepth(depth uint, format string, args ...any) {
	tracelog.Output(2+int(depth), fmt.Sprintf(format, args...))
}
