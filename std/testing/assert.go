package testing

import (
	"testing"
)

func Nil(t *testing.T, v interface{}) {
	if v != nil {
		t.Errorf("Value is not nil: %v", v)
	}
}

func NotNil(t *testing.T, v interface{}) {
	if v == nil {
		t.Error("Value is nil.")
	}
}

func Equal[T comparable](t *testing.T, a, b T) {
	if a == b {
		return
	}
	t.Errorf("Strings '%v' and '%v' should be equal", a, b)
}

func Unequal[T comparable](t *testing.T, a, b string) {
	if a == b {
		t.Errorf("Strings '%v' and '%v' should not be equal", a, b)
	}
}

func ElementPresent[K comparable](t *testing.T, set map[K]struct{}, key K) {
	if _, ok := set[key]; !ok {
		t.Errorf("Expected key '%v' to be present in set.", key)
	}
}

func ElementAbsent[K comparable](t *testing.T, set map[K]struct{}, key K) {
	if _, ok := set[key]; ok {
		t.Errorf("Expected key '%v' to be absent in set.", key)
	}
}

func KeyPresent[K comparable, V any](t *testing.T, set map[K]V, key K) {
	if _, ok := set[key]; !ok {
		t.Errorf("Expected key '%v' to be present in map.", key)
	}
}

func KeyAbsent[K comparable, V any](t *testing.T, set map[K]V, key K) {
	if _, ok := set[key]; ok {
		t.Errorf("Expected key '%v' to be absent in map.", key)
	}
}
