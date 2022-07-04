package testing

import "testing"

func RequirePanic(t *testing.T) {
	if recover() != nil {
		return
	}
	t.Error("panic was expected")
}
