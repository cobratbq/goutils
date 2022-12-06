// Multiset is a map with uint-values that represents a set of elements with their respective number
// of elements. Functions `Increment` and `Decrement` will increase or decrease the count of
// elements of that key, and if the count reaches zero, it will delete the element from the map
// (multiset). This allows testing the map with `len(multiset)` to test current size.
//
// The multiset uses `uint`-based counts such that the multiset adapts to the capacity of the
// architecture on which it is executed.
//
// Invariant: all elements have a count strictly larger than 0. (I.e. elements that reach 0 must be
// deleted from the multiset.)
package multiset

import "github.com/cobratbq/goutils/assert"

// Increment increments the count of an element, inserting the element if necessary.
func Increment[K comparable](multiset map[K]uint, e K) {
	multiset[e] += 1
}

// IncrementN increments the count of an element with n, inserting the element if necessary.
func IncrementN[K comparable](multiset map[K]uint, e K, n uint) {
	multiset[e] += n
}

// Decrement decrements the presence count of an element, deleting the element from the map
// (multiset) if the count reaches 0. By deleting the element, `len(...)` may be used to accurately
// measure how many different elements are currently present. Go's property of assuming zero-value
// default allows transparently querying non-existing element and receiving an accurate count of 0.
func Decrement[K comparable](multiset map[K]uint, e K) {
	multiset[e] -= 1
	if multiset[e] == 0 {
		delete(multiset, e)
	}
}

// DecrementN decrements the count of an element by `n`. The element count must be at least `n`,
// otherwise the function panics.
func DecrementN[K comparable](multiset map[K]uint, e K, n uint) {
	assert.Require(n <= multiset[e], "Decrement count cannot be bigger than count in multiset.")
	multiset[e] -= n
	if multiset[e] == 0 {
		delete(multiset, e)
	}
}

// Sum the numbers of each element in the multiset into a total count.
func Count[K comparable](multiset map[K]uint) uint {
	count := uint(0)
	for _, n := range multiset {
		count += n
	}
	return count
}
