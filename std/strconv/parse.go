// SPDX-License-Identifier: AGPL-3.0-or-later

package strconv

import (
	"strconv"
	"unsafe"

	"github.com/cobratbq/goutils/encoding/bytes/digit"
	"github.com/cobratbq/goutils/std/builtin"
	"github.com/cobratbq/goutils/types"
)

// ParseConsecutiveDigits reads bytes from the line until a non-digit is found, then parses the
// previous bytes in order to return an unsigned integer and the number of bytes read. Note that
// produced values will always be unsigned, due to accepting digits only.
//
// If no digits are found, this is not considered an error. Instead, `0, 0` is returned to
// accurately represent the input.
//
// If parsing the digits produces an error, this error is returned.
func ParseConsecutiveDigits[T types.UnsignedInteger](line []byte) (T, int) {
	var i int
	for len(line) > i && digit.IsDigit(line[i]) {
		i++
	}
	if i == 0 {
		return 0, 0
	}
	// TODO ideally, we get a constant expression, but this is not possible within the function.
	unsigned := T(MustParseUint(string(line[:i]), DecimalBase, int(unsafe.Sizeof(T(0)))*8))
	return unsigned, i
}

// DecimalBase is the base for the decimal system.
const DecimalBase = 10

// MustParseInt parses a string for an integer value of at most specified
// bitsize. Success is assumed and the function will panic on error.
func MustParseInt(v string, base, bitsize int) int64 {
	return builtin.Expect(strconv.ParseInt(v, base, bitsize))
}

// MustParseUint parses a string for an unsigned integer value of at most specified bitsize. Success
// is assumed and the function will panic on error.
func MustParseUint(v string, base, bitsize int) uint64 {
	return builtin.Expect(strconv.ParseUint(v, base, bitsize))
}
