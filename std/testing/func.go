// SPDX-License-Identifier: LGPL-3.0-only

package testing

import "testing"

func RequireRecover(t testing.TB) {
	recover()
}

func RequirePanic(t testing.TB) {
	if recover() != nil {
		return
	}
	t.Error("panic was expected")
}
