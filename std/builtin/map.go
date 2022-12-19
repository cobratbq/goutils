// SPDX-License-Identifier: LGPL-3.0-or-later

package builtin

// DuplicateMap duplicates the map only. A shallow copy of map entries into a new map of equal size.
func DuplicateMap[K comparable, V any](src map[K]V) map[K]V {
	dst := make(map[K]V, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// ExtractKeys extracts the keys from a map.
func ExtractKeys[K comparable, V any](map_ map[K]V) []K {
	keys := make([]K, 0, len(map_))
	for k := range map_ {
		keys = append(keys, k)
	}
	return keys
}

// ExtractValues extracts the values from a map.
func ExtractValues[K comparable, V any](map_ map[K]V) []V {
	vals := make([]V, 0, len(map_))
	for _, v := range map_ {
		vals = append(vals, v)
	}
	return vals
}

// ReduceMapKeys uses provided reduction function to reduce keys into a single resulting value.
func ReduceMapKeys[K comparable, V any, R any](input map[K]V, initial R, reduce func(v R, k K) R) R {
	v := initial
	for k, _ := range input {
		v = reduce(v, k)
	}
	return v
}
