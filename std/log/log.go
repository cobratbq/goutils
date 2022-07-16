// SPDX-License-Identifier: LGPL-3.0-or-later

package log

import (
	"log"
)

// Debug writes a line to os.Stderr with prefix 'debug'.
func Debug(line string) {
	log.Println("[debug] " + line)
}

// Debugf writes a line to os.Stderr with prefix 'debug', using fmt formatting options.
func Debugf(format string, args ...any) {
	log.Printf("[debug] "+format+"\n", args)
}

// Info writes a line to os.Stderr with prefix 'info'.
func Info(line string) {
	log.Println("[info] " + line)
}

// Info writes a line to os.Stderr with prefix 'info'.
func Infof(format string, args ...any) {
	log.Printf("[info] "+format+"\n", args)
}

// Warn writes a line to os.Stderr with prefix 'warn'.
func Warn(line string) {
	log.Println("[warn] " + line)
}

// Warn writes a line to os.Stderr with prefix 'warn'.
func Warnf(format string, args ...any) {
	log.Printf("[warn] "+format+"\n", args)
}

// Error writes a line to os.Stderr with prefix 'ERROR'.
func Error(line string) {
	log.Println("ERROR: " + line)
}

// Errorf writes a line to os.Stderr with prefix 'ERROR', using fmt formatting options.
func Errorf(format string, args ...any) {
	log.Printf("ERROR: "+format+"\n", args...)
}
