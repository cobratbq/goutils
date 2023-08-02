// SPDX-License-Identifier: LGPL-3.0-only

package sort

import (
	"testing"
)

func TestStringSetNil(t *testing.T) {
	if sorted := StringSet(nil); len(sorted) != 0 {
		t.FailNow()
	}
}

func TestStringSetSingle(t *testing.T) {
	set := make(map[string]struct{}, 1)
	set["hello"] = struct{}{}
	sorted := StringSet(set)
	if len(sorted) != 1 {
		t.FailNow()
	}
	if sorted[0] != "hello" {
		t.FailNow()
	}
}

func TestStringSetSomeStrings(t *testing.T) {
	set := map[string]struct{}{
		"citrus": {},
		"banana": {},
		"pear":   {},
		"apple":  {},
	}
	sorted := StringSet(set)
	if len(sorted) != 4 {
		t.Fail()
	}
	if sorted[0] != "apple" || sorted[1] != "banana" || sorted[2] != "citrus" || sorted[3] != "pear" {
		t.Fail()
	}
}
