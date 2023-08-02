// SPDX-License-Identifier: LGPL-3.0-only

package slices

func ContainedIn[E comparable](data []E) func(element E) bool {
	return func(element E) bool {
		for _, e := range data {
			if e == element {
				return true
			}
		}
		return false
	}
}

func NotContainedIn[E comparable](data []E) func(e E) bool {
	return func(element E) bool {
		for _, e := range data {
			if e == element {
				return false
			}
		}
		return true
	}
}
