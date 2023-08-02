// SPDX-License-Identifier: LGPL-3.0-only

// set is a map of keys with unit-values (`struct{}{}`). These values do not take up additional
// space. This allows the map to be used as a set, indicating whether an element is present (as a
// key) or not. This is a basic utility. For highly specialized and optimized use cases, look for a
// suitable specialized implementation.
package set

import "github.com/cobratbq/goutils/std/builtin/maps"

// Create creates and initializes a new map[K]struct{} for use as a set. All provided elements will
// immediately be included in the set. Initialization assumes that elements are unique; immediately
// allocates room for `len(elements)` elements in the map.
func Create[K comparable](elements ...K) map[K]struct{} {
	set := make(map[K]struct{}, len(elements))
	for _, e := range elements {
		set[e] = struct{}{}
	}
	return set
}

// ContainsAll checks whether all elements of `b` are present in `a`.
func ContainsAll[K comparable](a, b map[K]struct{}) bool {
	if len(b) > len(a) {
		return false
	}
	for k := range b {
		if _, ok := a[k]; !ok {
			return false
		}
	}
	return true
}

// Contains tests for the presence of element `e` and returns true iff present.
func Contains[K comparable](set map[K]struct{}, e K) bool {
	_, ok := set[e]
	return ok
}

// Insert inserts an element into the map (with the unit-value).
func Insert[K comparable](set map[K]struct{}, e K) {
	set[e] = struct{}{}
}

// InsertMany allows inserting any number of elements provided through vararg.
func InsertMany[K comparable](set map[K]struct{}, elms ...K) {
	for _, e := range elms {
		set[e] = struct{}{}
	}
}

// Remove removes an element if present in the map.
func Remove[K comparable](set map[K]struct{}, e K) {
	delete(set, e)
}

// RemoveMany removes any number of elements as provided through vararg.
func RemoveMany[K comparable](set map[K]struct{}, elms ...K) {
	for _, e := range elms {
		delete(set, e)
	}
}

// Merge merges `src` into `dst`.
func Merge[K comparable](dst, src map[K]struct{}) {
	for k := range src {
		dst[k] = struct{}{}
	}
}

// Subtract updates `set` by removing any element present in `other`.
//
// See: <https://en.wikipedia.org/wiki/Set_(mathematics)#Basic_operations>
//
// FIXME unused, untested
func Subtract[K comparable](set, other map[K]struct{}) {
	for k := range other {
		// remove indiscriminately because it doesn't matter if the element isn't present anyways
		delete(set, k)
	}
}

// Union produces a set of all elements present in either `a` or `b` (or both).
// FIXME unused, untested
func Union[K comparable](a, b map[K]struct{}) map[K]struct{} {
	union := maps.Duplicate(a)
	maps.MergeFunc(union, b, mergeUnit[K])
	return union
}

// Intersection produces a set of all elements present in both `a` and `b`.
// FIXME unused, untested
func Intersection[K comparable](a, b map[K]struct{}) map[K]struct{} {
	intersection := make(map[K]struct{}, 0)
	// TODO consider checking and working with smallest set, instead of assuming `b`
	for k := range b {
		if _, present := a[k]; present {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

// Difference creates a new set containing all elements from `set` that are not present in `other`.
//
// See: <https://en.wikipedia.org/wiki/Set_(mathematics)#Basic_operations>
//
// FIXME unused, untested
func Difference[K comparable](set, other map[K]struct{}) map[K]struct{} {
	difference := make(map[K]struct{}, 0)
	for k := range set {
		if _, present := other[k]; present {
			continue
		}
		// remove indiscriminately because it doesn't matter if the element isn't present anyways
		difference[k] = struct{}{}
	}
	return difference
}

// SymmetricDifference produces the symmetric difference of `set` and `other`, by removing elements
// that are present in both sets, and keeping elements present in only one set.
//
// See: <https://en.wikipedia.org/wiki/Set_(mathematics)#Basic_operations>
//
// REMARK consider starting at size `0` to avoid excessive use when this function is performed on large sets.
// FIXME unused, untested
func SymmetricDifference[K comparable](set, other map[K]struct{}) map[K]struct{} {
	difference := maps.Duplicate(set)
	for k := range other {
		if _, present := difference[k]; present {
			delete(difference, k)
		} else {
			difference[k] = struct{}{}
		}
	}
	return difference
}

func mergeUnit[K comparable](k K, a, b struct{}) struct{} {
	return struct{}{}
}
