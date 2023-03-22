// SPDX-License-Identifier: AGPL-3.0-or-later

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
