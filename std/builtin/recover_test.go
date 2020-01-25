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
