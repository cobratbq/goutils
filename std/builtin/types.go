// SPDX-License-Identifier: LGPL-3.0-or-later

package builtin

type Integer interface {
	UnsignedInteger | SignedInteger
}

type UnsignedInteger interface {
	uint8 | uint16 | uint32 | uint64 | uint
}

type SignedInteger interface {
	int8 | int16 | int32 | int64 | int
}

type Float interface {
	float32 | float64
}
