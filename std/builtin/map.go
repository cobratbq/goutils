// SPDX-License-Identifier: LGPL-3.0-or-later

package builtin

import "github.com/cobratbq/goutils/assert"

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

// TransformMap transforms both keys and values of a map into the output types for keys and values.
// TransformMap assumes correct operation of the transformation function `f`. It will allow
// overlapping keys in the output map, possibly resulting in loss of values.
func TransformMap[KIN, KOUT comparable, VIN, VOUT any](input map[KIN]VIN,
	transform func(k KIN, v VIN) (KOUT, VOUT)) map[KOUT]VOUT {

	output := make(map[KOUT]VOUT, 0)
	for kin, vin := range input {
		kout, vout := transform(kin, vin)
		output[kout] = vout
	}
	return output
}

// TransformMapKeys transforms an input map into an output map, using different types for keys.
// Given that only keys are transformed, this implementation will assume that destination key types
// will not overlap. If the transformation maps to the same key more than once, execution will
// panic. This prevents losing values by overlapping destination keys.
func TransformMapKeys[KIN comparable, KOUT comparable, V any](input map[KIN]V,
	transform func(k KIN, v V) KOUT) map[KOUT]V {

	output := make(map[KOUT]V, len(input))
	for kIn, value := range input {
		kOut := transform(kIn, value)
		output[kOut] = value
	}
	// detect loss of data (values) through multiple mappings onto the same KOUT (key in output map)
	assert.Equal(len(input), len(output))
	return output
}

// TransformMapValues transforms an input map into an output map, using different types for values.
func TransformMapValues[K comparable, VIN any, VOUT any](input map[K]VIN,
	transform func(k K, vin VIN) VOUT) map[K]VOUT {

	output := make(map[K]VOUT, len(input))
	for k, vin := range input {
		output[k] = transform(k, vin)
	}
	return output
}
