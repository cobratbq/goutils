// SPDX-License-Identifier: LGPL-3.0-or-later

// testing contains assertions for use in unit tests. I recommend aliasing the import to 'assert',
// so that you can call methods with `assert.Equal(t, a, b)` which makes it nicely readable. In
// addition, one would either use `assert` in production logic, or `std/testing` in unit tests.
package testing

import "testing"

func True(t testing.TB, v bool) {
	if !v {
		t.Error("Value is false")
	}
}

func False(t testing.TB, v bool) {
	if v {
		t.Error("Value is true")
	}
}

func Nil(t testing.TB, v interface{}) {
	if v != nil {
		t.Errorf("Value is not nil: %v", v)
	}
}

func NotNil(t testing.TB, v interface{}) {
	if v == nil {
		t.Error("Value is nil.")
	}
}

func Equal[T comparable](t testing.TB, a, b T) {
	if a == b {
		return
	}
	t.Errorf("Strings '%v' and '%v' should be equal", a, b)
}

func Unequal[T comparable](t testing.TB, a, b T) {
	if a == b {
		t.Errorf("Strings '%v' and '%v' should not be equal", a, b)
	}
}

func SliceContains[T comparable](t testing.TB, slice []T, elm T) {
	for _, v := range slice {
		if v == elm {
			return
		}
	}
	t.Errorf("Expected element '%v' to be present in slice.", elm)
}

func ElementPresent[K comparable](t testing.TB, set map[K]struct{}, key K) {
	KeyPresent(t, set, key)
}

func ElementAbsent[K comparable](t testing.TB, set map[K]struct{}, key K) {
	KeyAbsent(t, set, key)
}

func KeyPresent[K comparable, V any](t testing.TB, set map[K]V, key K) {
	if _, ok := set[key]; !ok {
		t.Errorf("Expected key '%v' to be present in map.", key)
	}
}

func KeyAbsent[K comparable, V any](t testing.TB, set map[K]V, key K) {
	if _, ok := set[key]; ok {
		t.Errorf("Expected key '%v' to be absent in map.", key)
	}
}

func ValuePresent[K comparable, V comparable](t testing.TB, map_ map[K]V, value V) {
	for _, v := range map_ {
		if v == value {
			return
		}
	}
	t.Errorf("Expected value '%v' to be present in map.", value)
}

func ValueAbsent[K comparable, V comparable](t testing.TB, map_ map[K]V, value V) {
	for _, v := range map_ {
		if v == value {
			t.Errorf("Expected value '%v' to be absent in map.", value)
		}
	}
}
