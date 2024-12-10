// SPDX-License-Identifier: LGPL-3.0-only

//go:build !disable_assert

package assert

import (
	"maps"
	"slices"

	"github.com/cobratbq/goutils/std/log"
	"github.com/cobratbq/goutils/types"
)

func False(v bool) {
	if v {
		log.TracelnDepth(1, "assert.False:", v)
		panic("assertion failed: False")
	}
}

func True(v bool) {
	if !v {
		log.TracelnDepth(1, "assert.True:", v)
		panic("assertion failed: True")
	}
}

func Any[T comparable](actual T, values ...T) T {
	for _, v := range values {
		if actual == v {
			return actual
		}
	}
	log.TracelnDepth(1, "assert.Any:", actual)
	panic("assertion failed: expected one of specified values")
}

func Equal[T comparable](v1, v2 T) {
	if v1 != v2 {
		log.TracelnDepth(1, "assert.Equal:", v1, v2)
		panic("assertion failed: Equal")
	}
}

func EqualSlices[S ~[]E, E comparable](s1, s2 S) {
	if !slices.Equal(s1, s2) {
		log.TracelnDepth(1, "assert.EqualSlices:", s1, s2)
		panic("assertion failed: EqualSlices")
	}
}

func EqualMaps[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) {
	if !maps.Equal(m1, m2) {
		log.TracefDepth(1, "assert.EqualMaps: %+v %+v", m1, m2)
		panic("assertion failed: EqualMaps")
	}
}

func Unequal[T comparable](v1, v2 T) {
	if v1 == v2 {
		log.TracelnDepth(1, "assert.Unequal:", v1, v2)
		panic("assertion failed: Unequal")
	}
}

func Zero[T types.Integer](v T) T {
	if v == 0 {
		return v
	}
	log.TracelnDepth(1, "assert.Zero:", v)
	panic("assertion failed: zero value")
}

func Positive[T types.Integer](v T) T {
	if v > 0 {
		return v
	}
	log.TracelnDepth(1, "assert.Positive:", v)
	panic("assertion failed: positive value")
}

func NonNegative[T types.SignedInteger](v T) T {
	if v >= 0 {
		return v
	}
	log.TracelnDepth(1, "assert.NonNegative:", v)
	panic("assertion failed: non-negative value")
}

func Negative[T types.SignedInteger](v T) T {
	if v < 0 {
		return v
	}
	log.TracelnDepth(1, "assert.Negative:", v)
	panic("assertion failed: negative value")
}

func NonPositive[T types.Integer](v T) T {
	if v <= 0 {
		return v
	}
	log.TracelnDepth(1, "assert.NonPositive:", v)
	panic("assertion failed: non-positive value")
}

func EmptyString(s string) {
	if len(s) == 0 {
		return
	}
	log.TracelnDepth(1, "assert.EmptyString:", s)
	panic("assertion failed: string is not empty/blank")
}

func EmptySlice[E any](collection []E) {
	if len(collection) == 0 {
		return
	}
	log.TracelnDepth(1, "assert.EmptySlice:", len(collection), "elements")
	panic("assertion failed: slice is not empty")
}

func EmptyMap[K comparable, V any](collection map[K]V) {
	if len(collection) == 0 {
		return
	}
	log.TracelnDepth(1, "assert.EmptyMap:", len(collection), "entries")
	panic("assertion failed: map is not empty")
}

func AtLeast[C types.Ordered](least C, value C) C {
	if value >= least {
		return value
	}
	log.TracelnDepth(1, "assert.AtLeast:", value)
	panic("assertion failed: value below (inclusive) lower bound")
}

func AtMost[C types.Ordered](most C, value C) C {
	if value <= most {
		return value
	}
	log.TracelnDepth(1, "assert.AtMost:", value)
	panic("assertion failed: value above (inclusive) upper bound")
}
