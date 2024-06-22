// SPDX-License-Identifier: LGPL-3.0-only

// testing contains assertions for use in unit tests. I recommend aliasing the import to 'assert',
// so that you can call methods with `assert.Equal(t, a, b)` which makes it nicely readable. In
// addition, one would either use `assert` in production logic, or `std/testing` in unit tests.
package testing

import (
	"errors"
	"slices"
	"testing"
)

// StopOnFailure halts the test immediately if it failed and logs provided message to provide context.
func StopOnFailure(t testing.TB, message ...any) {
	t.Helper()
	if t.Failed() {
		t.Log(message...)
		t.FailNow()
	}
}

func True(t testing.TB, v bool) {
	t.Helper()
	if !v {
		t.Error("Value is false")
	}
}

func False(t testing.TB, v bool) {
	t.Helper()
	if v {
		t.Error("Value is true")
	}
}

func Nil(t testing.TB, v interface{}) {
	t.Helper()
	if v != nil {
		t.Errorf("Value is not nil: %v", v)
	}
}

func NotNil(t testing.TB, v interface{}) {
	t.Helper()
	if v == nil {
		t.Error("Value is nil.")
	}
}

func Equal[T comparable](t testing.TB, a, b T) {
	t.Helper()
	if a == b {
		return
	}
	t.Errorf("Strings '%v' and '%v' should be equal", a, b)
}

func Unequal[T comparable](t testing.TB, a, b T) {
	t.Helper()
	if a == b {
		t.Errorf("Strings '%v' and '%v' should not be equal", a, b)
	}
}

func SliceContains[T comparable](t testing.TB, slice []T, elm T) {
	// TODO consider renaming
	t.Helper()
	for _, v := range slice {
		if v == elm {
			return
		}
	}
	t.Errorf("Expected element '%v' to be present in slice.", elm)
}

func ElementPresent[K comparable](t testing.TB, set map[K]struct{}, key K) {
	// TODO consider removing
	t.Helper()
	KeyPresent(t, set, key)
}

func ElementAbsent[K comparable](t testing.TB, set map[K]struct{}, key K) {
	// TODO consider removing
	t.Helper()
	KeyAbsent(t, set, key)
}

func KeyPresent[K comparable, V any](t testing.TB, set map[K]V, key K) {
	t.Helper()
	if _, ok := set[key]; !ok {
		t.Errorf("Expected key '%v' to be present in map.", key)
	}
}

func KeyAbsent[K comparable, V any](t testing.TB, set map[K]V, key K) {
	t.Helper()
	if _, ok := set[key]; ok {
		t.Errorf("Expected key '%v' to be absent in map.", key)
	}
}

func ValuePresent[K comparable, V comparable](t testing.TB, map_ map[K]V, value V) {
	t.Helper()
	for _, v := range map_ {
		if v == value {
			return
		}
	}
	t.Errorf("Expected value '%v' to be present in map.", value)
}

func ValueAbsent[K comparable, V comparable](t testing.TB, map_ map[K]V, value V) {
	t.Helper()
	for _, v := range map_ {
		if v == value {
			t.Errorf("Expected value '%v' to be absent in map.", value)
		}
	}
}

func IsError(t testing.TB, cause, actual error) {
	t.Helper()
	if !errors.Is(actual, cause) {
		t.Errorf("Actual error '%v' does not have expected root-cause: %v", actual, cause)
	}
}

func SlicesEqual[E comparable](t testing.TB, expected, actual []E) {
	t.Helper()
	if !slices.Equal(expected, actual) {
		t.Errorf("Slices %v and %v are not equal.", expected, actual)
	}
}
