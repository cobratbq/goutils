// SPDX-License-Identifier: LGPL-3.0-or-later

package builtin

type Number interface {
	UnsignedNumber | SignedNumber
}

type SignedNumber interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedNumber interface {
	uint | uint8 | uint16 | uint32 | uint64
}
