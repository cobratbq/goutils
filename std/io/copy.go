// SPDX-License-Identifier: AGPL-3.0-or-later

package io

import (
	"io"
	"sync"

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

// Transfer may be called in a goroutine. It copies all content from one connection to the next.
// Errors are ignored. In case copying is interrupted, for whatever reason, the function finishes up
// and releases `wg`.
func Transfer(wg *sync.WaitGroup, dst io.Writer, src io.Reader) {
	defer wg.Done()
	// Skip all error handling, because we simply cannot distinguish between expected and unexpected
	// events. Logging this will only produce noise.
	_, _ = io.Copy(dst, src)
}
