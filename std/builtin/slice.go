// SPDX-License-Identifier: AGPL-3.0-or-later

package builtin

import (
	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/builtin/multiset"
	"github.com/cobratbq/goutils/std/builtin/set"
)

// Contains checks if provided value is present anywhere in the slice.
func Contains[E comparable](slice []E, value E) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsFunc checks any element in the slice satisfies func `test`.
func ContainsFunc[E any](slice []E, test func(E) bool) bool {
	for _, v := range slice {
		if test(v) {
			return true
		}
	}
	return false
}

// TODO Index and IndexFunc are duplicates of `golang.org/x/exp/slices`. Keep for convenience or remove for duplicate?
// Index looks for a value linearly in the slice and returns its index if found, or -1 otherwise.
func Index[E comparable](data []E, value E) int {
	for i, v := range data {
		if v == value {
			return i
		}
	}
	return -1
}

// IndexFunc looks for an element linearly in the slice and returns its index if found, or -1
// otherwise. Func `test` is used to test if the value is found.
func IndexFunc[E any](data []E, test func(E) bool) int {
	for i, v := range data {
		if test(v) {
			return i
		}
	}
	return -1
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
func TransformSlice[I any, O any](input []I, transform func(I) O) []O {
	output := make([]O, 0, len(input))
	for _, in := range input {
		output = append(output, transform(in))
	}
	assert.Equal(len(output), cap(output))
	assert.Equal(len(input), len(output))
	return output
}

// ConvertSliceToMap transforms a slice with data into a map. It assumes that there
// FIXME consider how to deal with duplicate keys. The transform function should be able to assume no overlapping values. Otherwise input should be sanitized first.
func ConvertSliceToMap[E any, K comparable, V any](input []E, transform func(int, E) (K, V)) map[K]V {
	output := make(map[K]V)
	for i, e := range input {
		k, v := transform(i, e)
		// FIXME not considering duplicate elements in input
		output[k] = v
	}
	return output
}

// TransformSliceToMapKeys assumes non-overlapping map keys, meaning that there will be the same
// number of keys in the output as there are entries in the input slice. This assumption exists to
// be able to detect loss of information, due to faulty logic.
// TODO consider changing this to a "MergeSliceIntoMapKeys" that does not create the map itself and provides mutating logic.
// TODO consider renaming to ConvertSliceToMapKeys
func TransformSliceToMapKeys[K comparable, V any](input []K, transform func(int, K) V) map[K]V {
	output := make(map[K]V, len(input))
	for i, k := range input {
		// FIXME not considering duplicate elements in input
		output[k] = transform(i, k)
	}
	assert.Equal(len(input), len(output))
	return output
}

// SummarizeSliceElementsCount summarizes the contents of the slice as a multiset/bag containing
// each distinct element with a count for the number of occurrences.
// FIXME reconsider name for something shorter.
func SummarizeSliceDistinctElementCount[E comparable](data []E) map[E]uint {
	counts := make(map[E]uint)
	for _, e := range data {
		multiset.Insert(counts, e)
	}
	return counts
}

// SummarizeSliceElements summarizes the contents of the slice as a set containing each distinct
// element present.
// FIXME reconsider name for something shorter.
func SummarizeSliceDistinctElements[E comparable](data []E) map[E]struct{} {
	counts := make(map[E]struct{})
	for _, e := range data {
		set.Insert(counts, e)
	}
	return counts
}

// FilterSlice takes a slice `input` and a function `filter`. If `filter` returns true, the value is
// preserved. If `filter` returns false, the value is dropped.
func FilterSlice[E any](input []E, filter func(E) bool) []E {
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
func ReduceSlice[E any, V any](input []E, initial V, reduce func(V, E) V) V {
	v := initial
	for _, e := range input {
		v = reduce(v, e)
	}
	return v
}

// UpdateSlice updates all elements of a slice using the provided `update` func. Elements are passed
// in in isolation, therefore the update logic must operate on individual elements.
// TODO consider renaming to `UpdateElements` or something to reflect that this function operates on the slice's elements.
func UpdateSlice[E any](input []E, update func(E) E) {
	for i := 0; i < len(input); i++ {
		input[i] = update(input[i])
	}
}

// MoveElementTo moves an element at given index `from` in `input` to index `to`.
func MoveElementTo[E any](input []E, from, to int) {
	assert.Require(to >= 0 && to < len(input), "Invalid `to` index specified.")
	MoveElementN(input, from, to-from)
}

// MoveElementN moves an element any number of positions in a slice by repeated performing
// in-place swaps of elements.
func MoveElementN[E any](input []E, idx int, n int) {
	assert.Require(idx >= 0 && idx < len(input),
		"Invalid index value `idx` provided for input slice.")
	assert.Require(idx+n >= 0 && idx+n < len(input),
		"Invalid movement number `n` specified for input and starting index.")
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		input[idx+i], input[idx+i+1] = input[idx+i+1], input[idx+i]
	}
	for i := 0; i > n; i-- {
		input[idx+i], input[idx+i-1] = input[idx+i-1], input[idx+i]
	}
}

// Any iterates over elements in the slice and tests if they satisfy `test`. Result is returned upon
// first element found, and will iterate over all elements if none satisfy the condition.
func Any[E any](input []E, test func(int, E) bool) bool {
	for idx, e := range input {
		if test(idx, e) {
			return true
		}
	}
	return false
}
