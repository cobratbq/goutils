// SPDX-License-Identifier: AGPL-3.0-or-later

package maps

import "github.com/cobratbq/goutils/assert"

// ContainsAll checks whether all specified keys are present in the map. Returns `true` if all
// exist, or `false` otherwise.
func ContainsAll[K comparable, V any](map_ map[K]V, keys ...K) bool {
	for _, k := range keys {
		if _, found := map_[k]; !found {
			return false
		}
	}
	return true
}

// ContainsAny checks whether any of the specified keys is present in the map. Returns `true` if any
// one occurrence is found, and returns on the first find. If no key is found, `false` is returned.
func ContainsAny[K comparable, V any](map_ map[K]V, keys ...K) bool {
	for _, k := range keys {
		if _, found := map_[k]; found {
			return true
		}
	}
	return false
}

// ContainsKey checks a map for the specified key.
func Contains[K comparable, V any](map_ map[K]V, key K) bool {
	_, ok := map_[key]
	return ok
}

// DuplicateMap duplicates the map only. A shallow copy of map entries into a new map of equal size.
func Duplicate[K comparable, V any](src map[K]V) map[K]V {
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
func ReduceKeys[K comparable, V any, R any](input map[K]V, initial R, reduce func(R, K) R) R {
	v := initial
	for k := range input {
		v = reduce(v, k)
	}
	return v
}

// TransformMap transforms both keys and values of a map into the output types for keys and values.
// TransformMap assumes correct operation of the transformation function `f`. It will allow
// overlapping keys in the output map, possibly resulting in loss of values.
func Transform[KIN, KOUT comparable, VIN, VOUT any](input map[KIN]VIN,
	transform func(KIN, VIN) (KOUT, VOUT)) map[KOUT]VOUT {

	output := make(map[KOUT]VOUT, 0)
	for kin, vin := range input {
		kout, vout := transform(kin, vin)
		output[kout] = vout
	}
	return output
}

// TransformMapKeyType transforms an input map into an output map, using different types for keys.
// Given that only keys are transformed, this implementation will assume that destination key types
// will not overlap. If the transformation maps to the same key more than once, execution will
// panic. This prevents losing values by overlapping destination keys.
func TransformKeyType[KIN comparable, KOUT comparable, V any](input map[KIN]V,
	transform func(KIN, V) KOUT) map[KOUT]V {

	output := make(map[KOUT]V, len(input))
	for kIn, value := range input {
		kOut := transform(kIn, value)
		output[kOut] = value
	}
	// detect loss of data (values) through multiple mappings onto the same KOUT (key in output map)
	assert.Equal(len(input), len(output))
	return output
}

// TransformMapValueType transforms an input map into an output map, using different types for
// values.
func TransformValueType[K comparable, VIN any, VOUT any](input map[K]VIN,
	transform func(K, VIN) VOUT) map[K]VOUT {

	output := make(map[K]VOUT, len(input))
	for k, vin := range input {
		output[k] = transform(k, vin)
	}
	return output
}

// FilterMap filters a map according to the provided filter, returning a new map containing the
// filtered result.
func Filter[K comparable, V any](input map[K]V, filter func(K, V) bool) map[K]V {
	filtered := make(map[K]V, 0)
	for k, v := range input {
		if filter(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}

// MergeMap merges `src` map into `dst`. It requires all keys to be distinct. MergeMap will panic if
// a key is present in both maps. MergeMapFunc can be used if such conflict resolution is needed.
func Merge[K comparable, V any](dst, src map[K]V) {
	for k, v := range src {
		assert.False(Contains(dst, k))
		dst[k] = v
	}
}

// MergeMapFunc merges two distinct maps into one destination map, freshly created. In case a key
// exists in both maps, func `conflict` is called for conflict resolution. It will return the
// desired value, which can be determined based on provided key and the original values from both
// maps.
func MergeFunc[K comparable, V any](dst, src map[K]V, conflict func(K, V, V) V) {
	for k, v2 := range src {
		if v1, present := dst[k]; present {
			dst[k] = conflict(k, v1, v2)
		} else {
			dst[k] = v2
		}
	}
}

// MapKeysSubset checks, `O(n)` for `n` entries, if all keys of `subset` map are present in `set` map. Values are not
// considered.
func KeySubset[K comparable, V any](set map[K]V, subset map[K]V) bool {
	for k := range subset {
		if _, ok := set[k]; !ok {
			return false
		}
	}
	return true
}
