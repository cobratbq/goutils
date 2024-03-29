// SPDX-License-Identifier: LGPL-3.0-only

package slices

import (
	"os"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/builtin/multiset"
	"github.com/cobratbq/goutils/std/builtin/set"
	"github.com/cobratbq/goutils/std/errors"
)

// AppendCond appends to a slice if condition holds true.
func AppendCond[E any](cond bool, slice []E, value ...E) []E {
	if cond {
		return append(slice, value...)
	}
	return slice
}

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

// ForEach executes a closure for every value in `data`
func ForEach[E any](data []E, process func(e E)) {
	for _, e := range data {
		process(e)
	}
}

// ForEachIndexed executes a closure for every value in `data`
func ForEachIndexed[E any](data []E, process func(idx int, e E)) {
	for idx, e := range data {
		process(idx, e)
	}
}

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

// TransformIndexed maps a slice of data-type `I` to a function `idx, I -> O` and returns a result
// slice of data-type `O`.` The result slice is immediately allocated with equal capacity to
// minimize allocations.
func TransformIndexed[I, O any](input []I, transform func(int, I) O) []O {
	output := make([]O, 0, len(input))
	for idx, in := range input {
		output = append(output, transform(idx, in))
	}
	assert.Equal(len(output), cap(output))
	assert.Equal(len(input), len(output))
	return output
}

// Transform maps a slice of data-type `I` to a function `I -> O` and returns a result slice of
// data-type `O`. The result slice is immediately allocated with equal capacity to minimize
// allocations.
func Transform[I, O any](input []I, transform func(I) O) []O {
	output := make([]O, 0, len(input))
	for _, in := range input {
		output = append(output, transform(in))
	}
	assert.Equal(len(output), cap(output))
	assert.Equal(len(input), len(output))
	return output
}

// ConvertToMap transforms a slice with data into a map. It assumes that there
// FIXME consider how to deal with duplicate keys. The transform function should be able to assume no overlapping values. Otherwise input should be sanitized first.
func ConvertToMap[E any, K comparable, V any](input []E, transform func(int, E) (K, V)) map[K]V {
	output := make(map[K]V)
	for idx, e := range input {
		k, v := transform(idx, e)
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
func ConvertToMapKeys[K comparable, V any](input []K, transform func(int, K) V) map[K]V {
	output := make(map[K]V, len(input))
	for idx, k := range input {
		// FIXME not considering duplicate elements in input
		output[k] = transform(idx, k)
	}
	assert.Equal(len(input), len(output))
	return output
}

// DistinctElementCount summarizes the contents of the slice as a multiset/bag containing
// each distinct element with a count for the number of occurrences.
// FIXME reconsider name for something shorter.
func DistinctElementCount[E comparable](data []E) map[E]uint {
	counts := make(map[E]uint)
	for _, e := range data {
		multiset.Insert(counts, e)
	}
	return counts
}

// DistinctElements summarizes the contents of the slice as a set containing each distinct
// element present.
// FIXME reconsider name for something shorter.
func DistinctElements[E comparable](data []E) map[E]struct{} {
	counts := make(map[E]struct{})
	for _, e := range data {
		set.Insert(counts, e)
	}
	return counts
}

// Filter takes a slice `input` and a function `filter`. If `filter` returns true, the value is
// preserved. If `filter` returns false, the value is dropped.
func Filter[E any](input []E, filter func(E) bool) []E {
	filtered := make([]E, 0)
	for _, e := range input {
		if filter(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// FilterIndexed takes a slice `input` and a function `filter`. If `filter` returns true, the
// value is preserved. If `filter` returns false, the value is dropped.
func FilterIndexed[E any](input []E, filter func(int, E) bool) []E {
	filtered := make([]E, 0)
	for idx, e := range input {
		if filter(idx, e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// Reduce reduces a slice `input` to a single aggregate value of type `V`, using `initial V` as
// starting value. Function `reduce` defines exactly how `V` is determined with each entry.
func Reduce[E any, R any](input []E, initial R, reduce func(R, E) R) R {
	r := initial
	for _, e := range input {
		r = reduce(r, e)
	}
	return r
}

// ReduceIndexed reduces a slice `input` to a single aggregate value of type `V`, using
// `initial V` as starting value. Function `reduce` defines exactly how `V` is determined with each
// entry.
// ReduceIndexed uses a callback function that receives the slice index in addition to the value.
func ReduceIndexed[E any, R any](input []E, initial R, reduce func(R, int, E) R) R {
	r := initial
	for idx, e := range input {
		r = reduce(r, idx, e)
	}
	return r
}

// Update updates all elements of a slice using the provided `update` func. Elements are passed
// in in isolation, therefore the update logic must operate on individual elements.
// TODO consider renaming to `UpdateElements` or something to reflect that this function operates on the slice's elements.
func Update[E any](input []E, update func(int, E) E) {
	for i := 0; i < len(input); i++ {
		input[i] = update(i, input[i])
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

// All checks if all values in a slice satisfy `test`, that is `test` returns true.
func All[E any](input []E, test func(E) bool) bool {
	for _, e := range input {
		if !test(e) {
			return false
		}
	}
	return true
}

// Any iterates over elements in the slice and tests if they satisfy `test`. Result is returned upon
// first element found, and will iterate over all elements if none satisfy the condition.
func Any[E any](input []E, test func(E) bool) bool {
	for _, e := range input {
		if test(e) {
			return true
		}
	}
	return false
}

// All checks if all values in a slice satisfy `test`, that is `test` returns true.
func AllIndexed[E any](input []E, test func(int, E) bool) bool {
	for idx, e := range input {
		if !test(idx, e) {
			return false
		}
	}
	return true
}

// AnyIndexed iterates over elements in the slice and tests if they satisfy `test`. Result is returned upon
// first element found, and will iterate over all elements if none satisfy the condition.
func AnyIndexed[E any](input []E, test func(int, E) bool) bool {
	for idx, e := range input {
		if test(idx, e) {
			return true
		}
	}
	return false
}

// TODO consider defining `ErrInvalid` error independent of `os` package
func MiddleElement[E any](slice []E) (int, E, error) {
	if idx, err := MiddleIndex(slice); err == nil {
		return idx, slice[idx], nil
	} else {
		var zero E
		return 0, zero, err
	}
}

// MiddleIndex looks up the middle position in an odd-sized slice and returns its index. If
// even-sized, it returns `0` and an `os.ErrInvalid` with context.
// TODO this may be too much: consider moving to different `slice` package or something, akin to `set` or `multiset` I'd guess.
func MiddleIndex[E any](slice []E) (int, error) {
	if len(slice)%2 == 0 {
		return 0, errors.Context(os.ErrInvalid, "no middle element in even-sized slice")
	}
	return len(slice) / 2, nil
}

// Reversed creates a new slice with the contents of provided slice in reversed order.
func Reversed[E any](slice []E) []E {
	result := make([]E, 0, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}
	return result
}

// Empty3DFrom creates an empty 3D slice with dimensions from the provided prototype.
func Empty3DFrom[E, V any](prototype [][][]E, initial V) [][][]V {
	outer := make([][][]V, len(prototype))
	for i := 0; i < len(outer); i++ {
		outer[i] = Empty2DFrom(prototype[i], initial)
	}
	return outer
}

// Empty2DFrom creates an empty 2D slice with dimensions from the provided prototype.
func Empty2DFrom[E, V any](prototype [][]E, initial V) [][]V {
	outer := make([][]V, len(prototype))
	for i := 0; i < len(outer); i++ {
		outer[i] = EmptyFrom(prototype[i], initial)
	}
	return outer
}

// EmptyFrom creates an empty (multi-dimensional) slice with initial values `initial`.
func EmptyFrom[E, V any](prototype []E, initial V) []V {
	slice := make([]V, len(prototype))
	for i := 0; i < len(slice); i++ {
		slice[i] = initial
	}
	return slice
}

func UniformDimensions2D[E any](slice [][]E) bool {
	if len(slice) == 0 {
		return true
	}
	length := len(slice[0])
	for i := 1; i < len(slice); i++ {
		if len(slice[i]) != length {
			return false
		}
	}
	return true
}

// Create2D creates a fully-allocated 2D dynamic array.
func Create2D[E any](sizeX, sizeY uint, initial E) [][]E {
	outer := make([][]E, sizeY)
	for y := uint(0); y < sizeY; y++ {
		outer[y] = make([]E, sizeX)
	}
	return outer
}
