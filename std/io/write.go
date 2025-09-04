// SPDX-License-Identifier: LGPL-3.0-only

package io

import "io"

// TODO make `Cum` (cumulative) public or private with accessor function?
type CountingWriter struct {
	out io.Writer
	// FIXME consider better name for field; something that is descriptive when used.
	Cum int64
}

var _ io.Writer = (*CountingWriter)(nil)
var _ io.ReaderFrom = (*CountingWriter)(nil)

func NewCountingWriter(out io.Writer) CountingWriter {
	return CountingWriter{
		out: out,
		Cum: 0,
	}
}

// Write writes `p` to the underlying output and counts the number of bytes written.
// Write follows the conventions of io.Writer. It will add up the number of bytes written, even if the write
// was incomplete and consequently an error returned.
func (w *CountingWriter) Write(p []byte) (int, error) {
	n, err := w.out.Write(p)
	// just add onto the cumulative; let caller decide whether the error can be ignored.
	w.Cum += int64(n)
	return n, err
}

// ReadFrom implements the io.ReaderFrom interface.
func (w *CountingWriter) ReadFrom(r io.Reader) (int64, error) {
	n, err := io.Copy(w.out, r)
	w.Cum += n
	return n, err
}
