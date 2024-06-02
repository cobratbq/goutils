// SPDX-License-Identifier: LGPL-3.0-only

package log

import (
	"fmt"
	"log"
	"os"
)

var debuglog = log.New(os.Stderr, "[debug] ", log.Ltime|log.LUTC|log.Lmicroseconds|log.Lmsgprefix)

const calldepth = 2

// Debug writes output to os.Stderr with prefix 'debug'.
func Debug(line string) {
	debuglog.Output(calldepth, fmt.Sprintln(line))
}

// Debugln writes a line to os.Stderr with prefix `debug`, then ends with newline.
func Debugln(args ...any) {
	debuglog.Output(calldepth, fmt.Sprintln(args...))
}

// DebuglnSlice prints each entry in `data` on a new line. Every debug-line is prefixed with `prefix`.
func DebuglnSlice[T any](prefix string, data []T) {
	for i, e := range data {
		debuglog.Output(calldepth, fmt.Sprintln(prefix, i, "->", e))
	}
}

// DebuglnSliceAsString prints each entry in `data` on a new line. Every debug-line is prefixed with `prefix`.
// Each line of data (bytes), are printed as ANSI characters, thus converted to a string.
func DebuglnSliceAsString(prefix string, data [][]byte) {
	for i, e := range data {
		debuglog.Output(calldepth, fmt.Sprintln(prefix, i, "->", string(e)))
	}
}

// DebuglnMap prints each entry in `data` (every key in the map) on a new line. Every debug-line is prefixed
// with `prefix`.
func DebuglnMap[K comparable, V any](prefix string, data map[K]V) {
	for k, v := range data {
		debuglog.Output(calldepth, fmt.Sprintln(prefix, k, "->", v))
	}
}

// Debugf writes a line to os.Stderr with prefix 'debug', using fmt formatting options.
func Debugf(format string, args ...any) {
	debuglog.Output(calldepth, fmt.Sprintf(format+"\n", args...))
}

// DebugReport logs a debug-level message in case the assertion does not hold.
func DebugReport(assert bool, format string, args ...any) {
	if !assert {
		debuglog.Output(calldepth, fmt.Sprintf("Failed assertion: "+format+"\n", args...))
	}
}

var infolog = log.New(os.Stderr, " [info] ", log.Ltime|log.LUTC|log.Lmicroseconds|log.Lmsgprefix)

// Info writes a line to os.Stderr with prefix 'info'.
func Info(line string) {
	infolog.Output(calldepth, fmt.Sprintln(line))
}

// Infoln writes a line to os.Stderr with prefix 'info', closing with newline.
func Infoln(args ...any) {
	infolog.Output(calldepth, fmt.Sprintln(args...))
}

// Info writes a line to os.Stderr with prefix 'info'.
func Infof(format string, args ...any) {
	infolog.Output(calldepth, fmt.Sprintf(format+"\n", args...))
}

var warnlog = log.New(os.Stderr, " [warn] ", log.Ltime|log.LUTC|log.Lmicroseconds|log.Lmsgprefix)

// Warn writes a line to os.Stderr with prefix 'warn'.
func Warn(line string) {
	warnlog.Output(calldepth, fmt.Sprintln(line))
}

// Warnln writes a line to os.Stderr with prefix 'warn', closing with newline.
func Warnln(args ...any) {
	warnlog.Output(calldepth, fmt.Sprintln(args...))
}

// Warn writes a line to os.Stderr with prefix 'warn'.
func Warnf(format string, args ...any) {
	warnlog.Output(calldepth, fmt.Sprintf(format+"\n", args...))
}

var errorlog = log.New(os.Stderr, "[ERROR] ", log.Ltime|log.LUTC|log.Lmicroseconds|log.Lmsgprefix)

// Error writes a line to os.Stderr with prefix 'ERROR'.
func Error(line string) {
	errorlog.Output(calldepth, fmt.Sprintln(line))
}

// Errorln writes a line to os.Stderr with prefix 'ERROR', closing with newline.
func Errorln(args ...any) {
	errorlog.Output(calldepth, fmt.Sprintln(args...))
}

// Errorf writes a line to os.Stderr with prefix 'ERROR', using fmt formatting options.
func Errorf(format string, args ...any) {
	errorlog.Output(calldepth, fmt.Sprintf(format+"\n", args...))
}

func Flags() int {
	return debuglog.Flags()
}

func SetFlags(flags int) {
	debuglog.SetFlags(flags)
	infolog.SetFlags(flags)
	warnlog.SetFlags(flags)
	errorlog.SetFlags(flags)
}
