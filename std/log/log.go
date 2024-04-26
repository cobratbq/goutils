// SPDX-License-Identifier: LGPL-3.0-only

package log

import (
	"log"
)

// Debug writes output to os.Stderr with prefix 'debug'.
func Debug(line string) {
	log.Println("[debug] " + line)
}

// Debugln writes a line to os.Stderr with prefix `debug`, then ends with newline.
func Debugln(args ...any) {
	log.Println(append([]any{"[debug]"}, args...)...)
}

// DebuglnSlice prints each entry in `data` on a new line. Every debug-line is prefixed with `prefix`.
func DebuglnSlice[T any](prefix string, data []T) {
	for i, e := range data {
		Debugln(prefix, i, e)
	}
}

// DebuglnMap prints each entry in `data` (every key in the map) on a new line. Every debug-line is prefixed
// with `prefix`.
func DebuglnMap[K comparable, V any](prefix string, data map[K]V) {
	for k, v := range data {
		Debugln(prefix, k, v)
	}
}

// Debugf writes a line to os.Stderr with prefix 'debug', using fmt formatting options.
func Debugf(format string, args ...any) {
	log.Printf("[debug] "+format+"\n", args...)
}

// Info writes a line to os.Stderr with prefix 'info'.
func Info(line string) {
	log.Println("[info] " + line)
}

func Infoln(args ...any) {
	log.Println(append([]any{"[info]"}, args...)...)
}

// Info writes a line to os.Stderr with prefix 'info'.
func Infof(format string, args ...any) {
	log.Printf("[info] "+format+"\n", args...)
}

// Warn writes a line to os.Stderr with prefix 'warn'.
func Warn(line string) {
	log.Println("[warn] " + line)
}

func Warnln(args ...any) {
	log.Println(append([]any{"[warn]"}, args...)...)
}

// Warn writes a line to os.Stderr with prefix 'warn'.
func Warnf(format string, args ...any) {
	log.Printf("[warn] "+format+"\n", args...)
}

// Error writes a line to os.Stderr with prefix 'ERROR'.
func Error(line string) {
	log.Println("ERROR: " + line)
}

func Errorln(args ...any) {
	log.Println(append([]any{"ERROR:"}, args...)...)
}

// Errorf writes a line to os.Stderr with prefix 'ERROR', using fmt formatting options.
func Errorf(format string, args ...any) {
	log.Printf("ERROR: "+format+"\n", args...)
}
