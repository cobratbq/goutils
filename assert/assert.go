// SPDX-License-Identifier: LGPL-3.0-or-later

// assert provides various assertion functions that can be used to confirm certain conditions such
// that these conditions are guaranteed true afterwards. These functions are particularly useful to
// catch unexpected and unsupported use cases, without having to litter the code with if-statements.
// Assertions may be placeholders for use cases that will later be supported, or they may indicate
// failure conditions that will not or cannot ever be supported, or cannot even occur. Some use
// cases or possible error conditions are illusions created by the type-system, for example when a
// function implements an interface but will never fail the operation.
// Regardless, assertions allow you to (subtly) handle cases and failure conditions that you do not
// handle otherwise.
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
