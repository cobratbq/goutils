package rand

import (
	"crypto/rand"

	"github.com/cobratbq/goutils/std/errors"
)

// MustReadBytes reads random bytes into dst and fails if anything out of the
// ordinary happens.
func MustReadBytes(dst []byte) {
	_, err := rand.Read(dst)
	// rand.Read(...) api specifies that there will always be an error if
	// `num bytes read < len(dst)`.
	errors.RequireSuccess(err, "failed to read random bytes: %+v")
}
