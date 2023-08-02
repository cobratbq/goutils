// SPDX-License-Identifier: LGPL-3.0-only

//go:build disable_assert

package assert

func False(expected bool) {}

func True(expected bool) {}

func Any[T comparable](actual T, values ...T) {}

func Equal[T comparable](v1, v2 T) {}

func Unqual[T comparable](v1, v2 T) {}

func Positive[T types.SignedInteger](v T) {}

func NonNegative[T types.SignedInteger](v T) {}

func Negative[T types.SignedInteger](v T) {}

func NonPositive[T types.SignedInteger](v T) {}
