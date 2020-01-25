package rand

import (
	"bytes"
	"testing"

	"github.com/cobratbq/goutils/std/builtin"
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
	builtin.Require(result == nil, "expected to get nil returned")
}

func TestRandomizeBytes(t *testing.T) {
	var original [20]byte
	var dst [20]byte
	copy(original[:], dst[:])
	result := RandomizeBytes(dst[:])
	builtin.Require(bytes.Equal(result, dst[:]), "expected to receive randomized bytes")
	builtin.Require(!bytes.Equal(result, original[:]), "expected different bytes from start")
}
