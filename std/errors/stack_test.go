// SPDX-License-Identifier: LGPL-3.0-or-later
package errors

import (
	"strings"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestStacktraceNil(t *testing.T) {
	defer assert.RequirePanic(t)
	Stacktrace(nil)
	t.Fatal("Expected to panic because of nil")
}

func TestStacktrace(t *testing.T) {
	const ErrMyError = UintError(0)
	assert.Equal(t, "0", ErrMyError.Error())
	err := Stacktrace(ErrMyError)
	var firstLine, trace1 string
	firstLine, trace1, _ = strings.Cut(err.Error(), "\n")
	assert.Equal(t, "0", firstLine)
	err = Context(err, "Failed to increment from zero")
	var trace2 string
	firstLine, trace2, _ = strings.Cut(err.Error(), "\n")
	assert.Equal(t, "Failed to increment from zero: 0", firstLine)
	assert.Equal(t, trace1, trace2)
}
