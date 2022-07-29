// SPDX-License-Identifier: GPL-3.0-or-later

package assert

import (
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestTrue(t *testing.T) {
	True(true)
}

func TestTruePanics(t *testing.T) {
	defer assert.RequirePanic(t)
	True(false)
	t.FailNow()
}

func TestFalse(t *testing.T) {
	False(false)
}

func TestFalsePanics(t *testing.T) {
	defer assert.RequirePanic(t)
	False(true)
}
