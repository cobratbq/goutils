// SPDX-License-Identifier: AGPL-3.0-or-later

//go:build disable_assert

package assert

func False(expected bool) {}

func True(expected bool) {}

func Any[T comparable](actual T, values ...T) {}

func Equal[T comparable](v1, v2 T) {}

func Unqual[T comparable](v1, v2 T) {}
