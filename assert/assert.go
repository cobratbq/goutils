// SPDX-License-Identifier: LGPL-3.0-or-later
package assert

// TODO add Equal(a, b)

func False(expected bool) {
	True(!expected)
}

func True(expected bool) {
	if !expected {
		panic("assertion failed")
	}
}
