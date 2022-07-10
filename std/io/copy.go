// SPDX-License-Identifier: LGPL-3.0-or-later
package io

import "io"

// Discard reads remaining data from reader and discards it. Any possible
// errors in the process are ignored. Returns nr of bytes written, thus
// discarded.
func Discard(r io.Reader) int64 {
	n, _ := io.Copy(io.Discard, r)
	return n
}
