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

// Merge merges `src` into `dst`.
// FIXME consider naming: change to Join or Union or something?
func Merge[K comparable](dst, src map[K]struct{}) {
	for k := range src {
		dst[k] = struct{}{}
	}
}

// MergeSlice merges `slice` into `set`.
// FIXME consider naming to follow function above in naming convention.
func MergeSlice[K comparable](set map[K]struct{}, slice []K) {
	for _, e := range slice {
		set[e] = struct{}{}
	}
}
