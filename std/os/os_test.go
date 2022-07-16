// SPDX-License-Identifier: LGPL-3.0-or-later

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
