// SPDX-License-Identifier: LGPL-3.0-only

//go:build !disable_assert

package assert

import (
	"github.com/cobratbq/goutils/std/log"
	"github.com/cobratbq/goutils/types"
)

func False(expected bool) {
	if expected {
		panic("assertion failed: False")
	}
}

func True(expected bool) {
	if !expected {
		panic("assertion failed: True")
	}
}

func Any[T comparable](actual T, values ...T) {
	for _, v := range values {
		if actual == v {
			return
		}
	}
	log.Traceln("assert.Any:", actual)
	panic("assertion failed: expected one of specified values")
}

func Equal[T comparable](v1, v2 T) {
	if v1 != v2 {
		log.Traceln("assert.Equal:", v1, v2)
		panic("assertion failed: Equal")
	}
}

func Unequal[T comparable](v1, v2 T) {
	if v1 == v2 {
		log.Traceln("assert.Unequal:", v1, v2)
		panic("assertion failed: Unequal")
	}
}

func Zero[T types.Integer](v T) {
	if v == 0 {
		return
	}
	log.Traceln("assert.Zero:", v)
	panic("assertion failed: zero value")
}

func Positive[T types.Integer](v T) {
	if v > 0 {
		return
	}
	log.Traceln("assert.Positive:", v)
	panic("assertion failed: positive value")
}

func NonNegative[T types.SignedInteger](v T) {
	if v >= 0 {
		return
	}
	log.Traceln("assert.NonNegative:", v)
	panic("assertion failed: non-negative value")
}

func Negative[T types.SignedInteger](v T) {
	if v < 0 {
		return
	}
	log.Traceln("assert.Negative:", v)
	panic("assertion failed: negative value")
}

func NonPositive[T types.Integer](v T) {
	if v <= 0 {
		return
	}
	log.Traceln("assert.NonPositive:", v)
	panic("assertion failed: non-positive value")
}

// TODO is it possible to define a function `Empty` that accepts string, slice, map, etc. all at once, then perform a `len(c) == 0` check?

func EmptySlice[E any](collection []E) {
	if len(collection) > 0 {
		panic("assertion failed: slice is not empty")
	}
}

func EmptyMap[K comparable, V any](collection map[K]V) {
	if len(collection) > 0 {
		panic("assertion failed: map is not empty")
	}
}

func AtLeast[C types.Ordered](least C, value C) {
	if value < least {
		panic("assertion failed: value below (inclusive) lower bound")
	}
}

func AtMost[C types.Ordered](most C, value C) {
	if value > most {
		panic("assertion failed: value above (inclusive) upper bound")
	}
}
