// SPDX-License-Identifier: LGPL-3.0-or-later

package rand

import (
	"testing"
)

func TestUint64NonZero(t *testing.T) {
	v1 := Uint64NonZero()
	v2 := Uint64NonZero()
	v3 := Uint64NonZero()
	if v1 == v2 || v2 == v3 || v1 == v3 {
		t.FailNow()
	}
}

func TestUint64NonZeroNotZero(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := Uint64NonZero()
		if v == 0 {
			t.Fail()
		}
	}
}
