// SPDX-License-Identifier: LGPL-3.0-or-later
package assert

func False(expected bool) {
	if expected {
		panic("assertion failed: False")
	}
}

func True(expected bool) {
	if !expected {
		panic("assertion failed: True")
	}
}

func Equal[T comparable](v1, v2 T) {
	if v1 != v2 {
		panic("assertion failed: Equal")
	}
}
