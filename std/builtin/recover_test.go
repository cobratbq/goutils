// SPDX-License-Identifier: AGPL-3.0-or-later

package builtin

import "testing"

func TestRecoverLogged(t *testing.T) {
	defer RecoverLogged("Recovering ... %+v")
	panic("Hello world!")
}

func TestRecoverLoggedCalledDirectly(t *testing.T) {
	RecoverLogged("Hello world!")
}

func TestRecoverLoggedNil(t *testing.T) {
	defer RecoverLogged("Recovering ... %+v")
	panic(nil)
}

func TestRecoverLoggedStackTraceNotPanicking(t *testing.T) {
	defer RecoverLoggedStackTrace("Hello world")
}

func TestRecoverLoggedStackTraceNilPanic(t *testing.T) {
	defer RecoverLoggedStackTrace("Hello world")
	panic(nil)
}

func TestRecoverLoggedStackTracePanicking(t *testing.T) {
	defer RecoverLoggedStackTrace("Hello world!")
	panic("Hello!")
}
