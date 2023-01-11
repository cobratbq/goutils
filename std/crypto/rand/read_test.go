// SPDX-License-Identifier: AGPL-3.0-or-later

package rand

import (
	"bytes"
	"testing"

	"github.com/cobratbq/goutils/assert"
)

func TestMustReadBytesNil(t *testing.T) {
	MustReadBytes(nil)
}

func TestMustReadBytes(t *testing.T) {
	var dst [20]byte
	MustReadBytes(dst[:])
}

func TestRandomizeNil(t *testing.T) {
	result := RandomizeBytes(nil)
	assert.Require(result == nil, "expected to get nil returned")
}

func TestRandomizeBytes(t *testing.T) {
	var original [20]byte
	var dst [20]byte
	copy(original[:], dst[:])
	result := RandomizeBytes(dst[:])
	assert.Require(bytes.Equal(result, dst[:]), "expected to receive randomized bytes")
	assert.Require(!bytes.Equal(result, original[:]), "expected different bytes from start")
}
