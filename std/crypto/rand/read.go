// SPDX-License-Identifier: LGPL-3.0-only

package rand

import (
	"crypto/rand"

	"github.com/cobratbq/goutils/assert"
)

// RandomizeBytes reads `len(dst)` random bytes, then returns `dst`.
// It will panic on any kind of failure reading random bytes.
func RandomizeBytes(dst []byte) []byte {
	MustReadBytes(dst)
	return dst
}

// MustReadBytes reads random bytes into dst and fails if anything out of the
// ordinary happens.
func MustReadBytes(dst []byte) {
	_, err := rand.Read(dst)
	// rand.Read(...) api specifies that there will always be an error if
	// `num bytes read < len(dst)`.
	assert.Success(err, "failed to read random bytes")
}
