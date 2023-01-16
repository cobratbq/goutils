// SPDX-License-Identifier: AGPL-3.0-or-later

package strconv

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestMustParseIntEmptyString(t *testing.T) {
	defer assert.RequirePanic(t)
	MustParseInt[int64]("", 10)
	t.FailNow()
}

func TestMustParseIntIllegalString(t *testing.T) {
	defer assert.RequirePanic(t)
	MustParseInt[int64]("abcdefg", 10)
	t.FailNow()
}

func TestMustParseIntZero(t *testing.T) {
	assert.Equal(t, 0, MustParseInt[int64]("0", 10))
}

func TestParseConsecutiveDigitsNil(t *testing.T) {
	testdata := []struct {
		input []byte
		val   uint64
		n     int
	}{
		{nil, 0, 0},
		{[]byte{}, 0, 0},
		{[]byte("0"), 0, 1},
		{[]byte("9"), 9, 1},
		{[]byte("0000"), 0, 4},
		{[]byte("0000a"), 0, 4},
		{[]byte("999"), 999, 3},
		{[]byte("3a2b1c"), 3, 1},
		{[]byte("a2b1c"), 0, 0},
	}
	for _, d := range testdata {
		val, n := ParseConsecutiveDigits[uint64](d.input)
		assert.Equal(t, d.val, val)
		assert.Equal(t, d.n, n)
	}
}
