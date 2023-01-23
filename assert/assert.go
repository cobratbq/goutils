// SPDX-License-Identifier: AGPL-3.0-or-later

//go:build !disable_assert

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

func Any[T comparable](actual T, values ...T) {
	for _, v := range values {
		if actual == v {
			return
		}
	}
	panic("assertion failed: expected one of specified values")
}

func Equal[T comparable](v1, v2 T) {
	if v1 != v2 {
		panic("assertion failed: Equal")
	}
}

func Unequal[T comparable](v1, v2 T) {
	if v1 == v2 {
		panic("assertion failed: Unequal")
	}
}
