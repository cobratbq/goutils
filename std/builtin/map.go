// SPDX-License-Identifier: LGPL-3.0-or-later

package builtin

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
