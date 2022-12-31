// SPDX-License-Identifier: LGPL-3.0-or-later

package builtin

import "github.com/cobratbq/goutils/assert"

// Contains checks if provided value is present anywhere in the slice.
func Contains[E comparable](slice []E, value E) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Duplicate copies the provided source slice to new slice of same size. Copying is a shallow copy
// operation.
func Duplicate[T any](src []T) []T {
	d := make([]T, len(src))
	assert.Equal(copy(d, src), len(src))
	return d
}

// MapSlice maps a slice of data-type `I` to a function `I -> O` and returns a result slice of
// data-type `O“.
// The result slice is immediately allocated with equal capacity to prevent further allocations.
func TransformSlice[I any, O any](input []I, transform func(in I) O) []O {
	output := make([]O, 0, len(input))
	for _, in := range input {
		output = append(output, transform(in))
	}
	return output
}

// TransformSliceToMapKeys assumes non-overlapping map keys, meaning that there will be the same
// number of keys in the output as there are entries in the input slice. This assumption exists to
// be able to detect loss of information, due to faulty logic.
// TODO consider changing this to a "MergeSliceIntoMapKeys" that does not create the map itself and provides mutating logic.
func TransformSliceToMapKeys[K comparable, V any](input []K,
	transform func(index int, element K) V) map[K]V {

	output := make(map[K]V, len(input))
	for i, k := range input {
		output[k] = transform(i, k)
	}
	assert.Equal(len(input), len(output))
	return output
}

// FilterSlice takes a slice `input` and a function `filter`. If `filter` returns true, the value is
// preserved. If `filter` returns false, the value is dropped.
func FilterSlice[E any](input []E, filter func(e E) bool) []E {
	filtered := make([]E, 0)
	for _, e := range input {
		if filter(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// ReduceSlice reduces a slice `input` to a single aggregate value of type `V`, using `initial V` as
// starting value. Function `reduce` defines exactly how `V` is determined with each entry.
func ReduceSlice[E any, V any](input []E, initial V, reduce func(v V, e E) V) V {
	v := initial
	for _, e := range input {
		v = reduce(v, e)
	}
	return v
}

// UpdateSlice updates all elements of a slice using the provided `update` func. Elements are passed
// in in isolation, therefore the update logic must operate on individual elements.
// TODO consider renaming to `UpdateElements` or something to reflect that this function operates on the slice's elements.
func UpdateSlice[E any](input []E, update func(e E) E) {
	for i := 0; i < len(input); i++ {
		input[i] = update(input[i])
	}
}

// Any iterates over elements in the slice and tests if they satisfy `test`. Result is returned upon
// first element found, and will iterate over all elements if none satisfy the condition.
func Any[E any](input []E, test func(idx int, e E) bool) bool {
	for idx, e := range input {
		if test(idx, e) {
			return true
		}
	}
	return false
}
