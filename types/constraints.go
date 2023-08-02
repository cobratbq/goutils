// SPDX-License-Identifier: LGPL-3.0-only

// TODO `golang.org/x/exp/constraints` provides similar types, but is not yet part of the standard library.

// Source: <https://go.dev/ref/spec#Interface_types> <https://pkg.go.dev/golang.org/x/exp>

package types

type Ordered interface {
	Number | ~string
}

type Number interface {
	Integer | Float
}

type Integer interface {
	UnsignedInteger | SignedInteger
}

type UnsignedInteger interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr
}

type SignedInteger interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type Float interface {
	~float32 | ~float64
}
