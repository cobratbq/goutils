// SPDX-License-Identifier: LGPL-3.0-or-later

package builtin

import "github.com/cobratbq/goutils/assert"

// Contains checks if provided value is present anywhere in the slice.
func Contains[T comparable](slice []T, value T) bool {
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
// TODO consider renaming because 'Map' may refer to data-type or "mapping" operation, which is confusing.
func MapSlice[I any, O any](input []I, f func(in I) O) []O {
	output := make([]O, 0, len(input))
	for _, in := range input {
		output = append(output, f(in))
	}
	return output
}

// FilterSlice takes a slice `input` and a function `filter`. If `filter` returns true, the value is
// preserved. If `filter` returns false, the value is dropped.
func FilterSlice[E any](input []E, filter func(e E) bool) []E {
	output := make([]E, 0)
	for _, e := range input {
		if filter(e) {
			output = append(output, e)
		}
	}
	return output
}
