// SPDX-License-Identifier: LGPL-3.0-only

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
	var errMyError = NewUintError(0)
	assert.Equal(t, "0", errMyError.Error())
	err := Stacktrace(errMyError)
	lines := strings.Split(err.Error(), "\n")
	trace1 := strings.Join(lines[1:len(lines)-1], "\n")
	assert.Equal(t, "Error: 0", lines[0])
	err = Context(err, "Failed to increment from zero")
	lines = strings.Split(err.Error(), "\n")
	trace2 := strings.Join(lines[1:len(lines)-1], "\n")
	assert.Equal(t, "Failed to increment from zero: Error: 0", lines[0])
	// Ensure that stack-traces themselves are unaffected.
	assert.Equal(t, trace1, trace2)
}
