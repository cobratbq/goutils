// SPDX-License-Identifier: LGPL-3.0-only

package builtin

import "github.com/cobratbq/goutils/std/log"

// DiscardUntilClosed drains a channel until it is closed.
func DiscardUntilClosed[T any](recv <-chan T) {
	log.TracelnDepth(1, "Discarding from channel until it is closed…")
	for _, ok := <-recv; ok; _, ok = <-recv {
	}
}
