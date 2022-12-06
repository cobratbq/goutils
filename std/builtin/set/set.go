// set is a map with unit-values (`struct{}{}“). These values do not take up additional space. This
// allows the map to be used as a set, indicating whether an element is present (as a key) or not.
package set

// Insert inserts an element into the map (with the unit-value).
func Insert[K comparable](set map[K]struct{}, e K) {
	set[e] = struct{}{}
}

// Remove removes an element if present in the map.
func Remove[K comparable](set map[K]struct{}, e K) {
	delete(set, e)
}