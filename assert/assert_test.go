// SPDX-License-Identifier: GPL-3.0-or-later
package assert

import "testing"

func TestTrue(t *testing.T) {
	True(true)
}

func TestTruePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.FailNow()
		}
	}()
	True(false)
	t.FailNow()
}

func TestFalse(t *testing.T) {
	False(false)
}

func TestFalsePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.FailNow()
		}
	}()
	False(true)
}
