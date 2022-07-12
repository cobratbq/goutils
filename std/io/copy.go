// SPDX-License-Identifier: LGPL-3.0-or-later
package io

import (
	"io"

	"github.com/cobratbq/goutils/std/log"
)

// CopyNoWarning copies contents from in to out. Do NOT log anything in case transfer interrupts.
func CopyNoWarning(out io.Writer, in io.Reader) int64 {
	n, _ := io.Copy(out, in)
	return n
}

// CopyWithWarning copies contents from in to out. It logs a warning in case transfer interrupts.
func CopyWithWarning(out io.Writer, in io.Reader) int64 {
	var n int64
	var err error
	if n, err = io.Copy(out, in); err != nil {
		log.Warnf("Failed to copy all content (copied %v bytes): %v", n, err.Error())
	}
	return n
}

// Discard reads remaining data from reader and discards it. Any possible
// errors in the process are ignored. Returns nr of bytes written, thus
// discarded.
func Discard(r io.Reader) int64 {
	n, _ := io.Copy(io.Discard, r)
	return n
}
