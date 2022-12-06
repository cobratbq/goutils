// Multiset is a map with uint-values that represents a set of elements with their respective number
// of elements. Functions `Increment` and `Decrement` will increase or decrease the count of
// elements of that key, and if the count reaches zero, it will delete the element from the map
// (multiset). This allows testing the map with `len(multiset)` to test current size.
package multiset

// Increment increments the presence count of an element, inserting the element if necessary.
func Increment[K comparable](multiset map[K]uint, e K) {
	multiset[e] += 1
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

// Sum the numbers of each element in the multiset into a total count.
func Count[K comparable](multiset map[K]uint) uint {
	count := uint(0)
	for _, n := range multiset {
		count += n
	}
	return count
}
