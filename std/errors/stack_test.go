// SPDX-License-Identifier: AGPL-3.0-or-later

package errors

import (
	"strings"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestStacktraceNilError(t *testing.T) {
	defer assert.RequirePanic(t)
	Stacktrace(nil)
	t.FailNow()
}

func TestStacktrace(t *testing.T) {
	var ErrMyError = NewUintError(0)
	assert.Equal(t, "0", ErrMyError.Error())
	err := Stacktrace(ErrMyError)
	lines := strings.Split(err.Error(), "\n")
	trace1 := strings.Join(lines[0:len(lines)-1], "\n")
	assert.Equal(t, "Error: 0", lines[len(lines)-1])
	err = Context(err, "Failed to increment from zero")
	lines = strings.Split(err.Error(), "\n")
	trace2 := strings.Join(lines[0:len(lines)-1], "\n")
	assert.Equal(t, "Error: 0: Failed to increment from zero", lines[len(lines)-1])
	assert.Equal(t, trace1, trace2)
}
