// SPDX-License-Identifier: AGPL-3.0-or-later

// Multiset is a map with uint-values that represents a set of elements with their respective number
// of elements. Functions `Increment` and `Decrement` will increase or decrease the count of
// elements of that key, and if the count reaches zero, it will delete the element from the map
// (multiset). This allows testing the map with `len(multiset)` to test current size.
//
// The multiset uses `uint`-based counts such that the multiset adapts to the capacity of the
// architecture on which it is executed.
//
// These multiset functions are just functions that operate on a map. Therefore there is no caching
// or any other kind of optimization possible/available.
//
// INVARIANT: all elements have a count strictly larger than 0. (I.e. elements that reach 0 are
// deleted from the multiset.)
// FIXME try to set an upper-limit guard for `C` to ensure that multiset does not wrap around undetected.
package multiset

import (
	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/builtin"
	"github.com/cobratbq/goutils/std/builtin/maps"
	math_ "github.com/cobratbq/goutils/std/math"
	"github.com/cobratbq/goutils/types"
)

// Create creates and initializes a new map[K]C for use as multiset. All provided elements will
// immediately be included in the multiset, with initial value as specified in param `count`.
// Initialization assumes that elements are unique; immediately allocates room for `len(elements)`
// elements in the map.
func Create[K comparable, C types.UnsignedInteger](count C, elements ...K) map[K]C {
	assert.Require(count > 0, "The initial count must be larger than 0.")
	multiset := make(map[K]C, len(elements))
	for _, e := range elements {
		multiset[e] = count
	}
	return multiset
}

// Contains tests for presence of an element in a multiset-like map.
func Contains[K comparable, C types.UnsignedInteger](multiset map[K]C, e K) bool {
	_, ok := multiset[e]
	return ok
}

// Insert increments the count of an element, inserting the element if necessary.
func Insert[K comparable, C types.UnsignedInteger](multiset map[K]C, e K) {
	multiset[e] += 1
}

// InsertN increments the count of an element with n, inserting the element if necessary.
func InsertN[K comparable, C types.UnsignedInteger](multiset map[K]C, e K, n C) {
	multiset[e] += n
}

// Remove decrements the count of an element, deleting the element from the map (multiset) if the
// count reaches 0. By deleting the element, `len(...)` may be used to accurately measure how many
// different elements are currently present. Go's property of assuming zero-value default allows
// transparently querying non-existing element and receiving an accurate count of 0.
func Remove[K comparable, C types.UnsignedInteger](multiset map[K]C, e K) {
	multiset[e] -= 1
	if multiset[e] == 0 {
		delete(multiset, e)
	}
}

// RemoveN decrements the count of an element by `n`. The element count must be at least `n`,
// otherwise the function panics.
func RemoveN[K comparable, C types.UnsignedInteger](multiset map[K]C, e K, n C) {
	assert.Require(n <= multiset[e], "Decrement count cannot be bigger than count in multiset.")
	multiset[e] -= n
	if multiset[e] == 0 {
		delete(multiset, e)
	}
}

// Sum the numbers of each element in the multiset into a total count. (The number of unique
// elements can be acquired using `len(multiset)`.)
func Count[K comparable, C types.UnsignedInteger](multiset map[K]C) C {
	var count C
	for _, n := range multiset {
		count += n
	}
	return count
}

func Union[K comparable, C types.UnsignedInteger](a, b map[K]C) map[K]C {
	united := maps.Duplicate(a)
	maps.MergeValuesFunc(united, b, math_.Max[C])
	return united
}

func Intersection[K comparable, C types.UnsignedInteger](a, b map[K]C) map[K]C {
	intersect := maps.Duplicate(a)
	maps.MergeValuesFunc(intersect, b, math_.Min[C])
	return intersect
}

func Sum[K comparable, C types.UnsignedInteger](a, b map[K]C) map[K]C {
	sum := maps.Duplicate(a)
	maps.MergeValuesFunc(sum, b, builtin.Add[C])
	return sum
}

func Difference[K comparable, C types.UnsignedInteger](a, b map[K]C) map[K]C {
	difference := maps.Duplicate(a)
	maps.MergeValuesFunc(difference, b, func(v1, v2 C) C {
		if v1 < v2 {
			return 0
		}
		return v1 - v2
	})
	clean(difference)
	return difference
}

func clean[K comparable, C types.UnsignedInteger](multiset map[K]C) {
	for k, v := range multiset {
		if v <= 0 {
			delete(multiset, k)
		}
	}
}
