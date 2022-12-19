// SPDX-License-Identifier: LGPL-3.0-or-later

// set is a map with unit-values (`struct{}{}“). These values do not take up additional space. This
// allows the map to be used as a set, indicating whether an element is present (as a key) or not.
package set

// Contains tests for the presence of element `e` and returns true iff present.
func Contains[K comparable](set map[K]struct{}, e K) bool {
	_, ok := set[e]
	return ok
}

// Insert inserts an element into the map (with the unit-value).
func Insert[K comparable](set map[K]struct{}, e K) {
	set[e] = struct{}{}
}

// Remove removes an element if present in the map.
func Remove[K comparable](set map[K]struct{}, e K) {
	delete(set, e)
}
