// SPDX-License-Identifier: LGPL-3.0-only

package os

import (
	"testing"
)

func TestWorkingDir(t *testing.T) {
	dir := WorkingDir()
	if dir == "" {
		t.FailNow()
	}
}
