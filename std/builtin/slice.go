// SPDX-License-Identifier: LGPL-3.0-or-later

package builtin

import "github.com/cobratbq/goutils/assert"

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Duplicate[T any](src []T) []T {
	d := make([]T, len(src))
	assert.Equal(copy(d, src), len(src))
	return d
}
