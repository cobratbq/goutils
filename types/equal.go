// SPDX-License-Identifier: LGPL-3.0-only

package types

// Equaler defines `Equal` for implementation-level equality testing, in case Go's native comparable or
// identity comparison does not suffice.
type Equaler[T any] interface {
	Equal(o T) bool
}
