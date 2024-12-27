// SPDX-License-Identifier: LGPL-3.0-only

package io

import (
	"io"
	"sync"

	"github.com/cobratbq/goutils/std/log"
)

// CopyNoWarning copies contents from in to out. Do NOT log anything in case transfer interrupts.
func CopyNoWarning(out io.Writer, in io.Reader) int64 {
	n, err := io.Copy(out, in)
	log.TracelnDepth(1, "`CopyNoWarning` ignores error", err)
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

// Discard reads remaining data from reader and discards it.
//
// Returns nr of bytes read, thus discarded. An error is returned if encountered while reading from input.
func Discard(in io.Reader) (int64, error) {
	return io.Copy(io.Discard, in)
}

// DiscardN reads `n` bytes from `in` and discards them.
//
// The number of discarded bytes is returned, and an error is returned only if it is not successful in reading
// and discarding `n` bytes.
func DiscardN(in io.Reader, n int64) (int64, error) {
	return io.CopyN(io.Discard, in, n)
}

// Transfer may be called in a goroutine. It copies all content from one connection to the next.
// Errors are ignored. In case copying is interrupted, for whatever reason, the function finishes up
// and releases `wg`.
func Transfer(wg *sync.WaitGroup, dst io.Writer, src io.Reader) {
	defer wg.Done()
	// Skip all error handling, because we simply cannot distinguish between expected and unexpected
	// events. Logging this will only produce noise.
	_, _ = io.Copy(dst, src)
	// TODO if I include trace-logging here for `io.Copy` output, would that (ideally) be discarded for builds with `!enable_trace`?
}
