// SPDX-License-Identifier: LGPL-3.0-or-later
package strconv

import (
	"testing"
)

func TestMustParseIntEmptyString(t *testing.T) {
	defer func() {
		recover()
	}()
	MustParseInt("", 10, 64)
	t.FailNow()
}

func TestMustParseIntIllegalString(t *testing.T) {
	defer func() {
		recover()
	}()
	MustParseInt("abcdefg", 10, 64)
	t.FailNow()
}

func TestMustParseIntZero(t *testing.T) {
	if MustParseInt("0", 10, 64) != 0 {
		t.FailNow()
	}
}
