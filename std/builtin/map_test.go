// SPDX-License-Identifier: LGPL-3.0-or-later
package builtin

import (
	"sort"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestExtractKeysNilMap(t *testing.T) {
	ExtractKeys(map[string]string(nil))
}

func TestExtractKeysIntegers(t *testing.T) {
	m := map[int]func(){
		0:  func() {},
		-1: func() {},
	}
	k := ExtractKeys(m)
	assert.SliceContains(t, k, 0)
	assert.SliceContains(t, k, -1)
}

func TestExtractKeysFuncMap(t *testing.T) {
	m := map[string]func(){
		"hello": func() {},
		"world": func() {},
	}
	keys := ExtractKeys(m)
	if len(keys) != 2 {
		t.Errorf("Failed to extract 2 keys from map.")
	}
	sort.Strings(keys)
	if keys[0] != "hello" || keys[1] != "world" {
		t.Errorf("Failed to find keys at expected sorted positions.")
	}
}
