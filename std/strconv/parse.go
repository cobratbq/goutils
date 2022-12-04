// SPDX-License-Identifier: LGPL-3.0-or-later

package strconv

import (
	"strconv"

	"github.com/cobratbq/goutils/assert"
)

// MustParseInt parses a string for an integer value of at most specified
// bitsize. Success is assumed and the function will panic on error.
func MustParseInt(v string, base, bitsize int) int64 {
	return assert.Unwrap(strconv.ParseInt(v, base, bitsize))
}

// MustParseUint parses a string for an unsigned integer value of at most specified bitsize. Success
// is assumed and the function will panic on error.
func MustParseUint(v string, base, bitsize int) uint64 {
	return assert.Unwrap(strconv.ParseUint(v, base, bitsize))
}
