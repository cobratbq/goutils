// SPDX-License-Identifier: LGPL-3.0-or-later
package testing

import "testing"

func RequirePanic(t testing.TB) {
	if recover() != nil {
		return
	}
	t.Error("panic was expected")
}
