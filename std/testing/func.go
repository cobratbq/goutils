// SPDX-License-Identifier: AGPL-3.0-or-later

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
