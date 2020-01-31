package builtin

import (
	"sort"
	"testing"
)

func TestExtractKeysNilMap(t *testing.T) {
	defer func() {
		recover()
	}()
	ExtractStringKeys(nil)
	t.FailNow()
}

func TestExtractKeysNonString(t *testing.T) {
	defer func() {
		recover()
	}()
	m := map[int]func(){
		0:  func() {},
		-1: func() {},
	}
	ExtractStringKeys(m)
	t.FailNow()
}

func TestExtractKeysFuncMap(t *testing.T) {
	m := map[string]func(){
		"hello": func() {},
		"world": func() {},
	}
	keys := ExtractStringKeys(m)
	if len(keys) != 2 {
		t.Fail()
	}
	sort.Strings(keys)
	if keys[0] != "hello" || keys[1] != "world" {
		t.Fail()
	}
}
