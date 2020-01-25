package rand

import (
	"testing"
)

func TestMustReadBytesNil(t *testing.T) {
	MustReadBytes(nil)
}

func TestMustReadBytes(t *testing.T) {
	var dst [20]byte
	MustReadBytes(dst[:])
}
