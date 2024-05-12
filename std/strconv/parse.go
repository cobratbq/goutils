// SPDX-License-Identifier: LGPL-3.0-only

package strconv

import (
	"strconv"
	"unsafe"

	"github.com/cobratbq/goutils/codec/bytes/digit"
	"github.com/cobratbq/goutils/std/builtin"
	"github.com/cobratbq/goutils/types"
)

// ParseConsecutiveDigits reads bytes from the line until a non-digit is found, then parses the
// previous bytes in order to return an unsigned integer and the number of bytes read. Note that
// produced values will always be unsigned, due to accepting digits only. If no digits are found,
// `0, 0` is returned.
func ParseConsecutiveDigits[T types.UnsignedInteger](line []byte) (T, int) {
	var i int
	for len(line) > i && digit.IsDigit(line[i]) {
		i++
	}
	if i == 0 {
		return 0, 0
	}
	// TODO there seems to be no way to produce an error, except maybe for extremely long series of digits.
	return MustParseUint[T](string(line[:i]), DecimalBase), i
}

const OctalBase = 8
const DecimalBase = 10
const HexadecimalBase = 16

// MustParseInt parses a string for an integer value of at most specified bitsize. Success is
// assumed and the function will panic on error.
func MustParseInt[T types.SignedInteger](s string, base int) T {
	return builtin.Expect(ParseInt[T](s, base))
}

// MustParseBytesInt parses a byte-array for a signed integer value of at most specified bitsize. Success is
// assumed and the function will panic on error.
func MustParseBytesInt[T types.SignedInteger](data []byte, base int) T {
	return MustParseInt[T](string(data), base)
}

// MustParseBytesIntDecimal parses a byte-array for a signed decimal integer value of at most specified
// bitsize. Success is assumed and the function will panic on error.
func MustParseBytesIntDecimal[T types.SignedInteger](data []byte) T {
	return MustParseInt[T](string(data), DecimalBase)
}

// ParseInt is generics-enabled variant of `strconv.ParseUint`, which derives the bitsize from the
// specified parametric type.
func ParseInt[T types.SignedInteger](s string, base int) (T, error) {
	// TODO ideally, we get a constant expression, but this is not possible within the function.
	var bitsize = int(unsafe.Sizeof(T(0))) * 8
	result, err := strconv.ParseInt(s, base, bitsize)
	return T(result), err
}

// MustParseUintDecimal parses a string and expects to convert to a decimal type.
func MustParseUintDecimal[T types.UnsignedInteger](s string) T {
	return MustParseUint[T](s, DecimalBase)
}

// MustParseBytesUintDecimal parses a string represented as raw bytes for a unsigned decimal value. Success
// is assumed and the function will panic on error.
func MustParseBytesUintDecimal[T types.UnsignedInteger](data []byte) T {
	return MustParseUintDecimal[T](string(data))
}

// MustParseUint parses a string for an unsigned integer value of at most specified bitsize. Success
// is assumed and the function will panic on error.
func MustParseUint[T types.UnsignedInteger](s string, base int) T {
	return builtin.Expect(ParseUint[T](s, base))
}

// MustParseBytesUint parses a string represented as raw bytes for a unsigned integer value. Success
// is assumed and the function will panic on error.
func MustParseBytesUint[T types.UnsignedInteger](data []byte, base int) T {
	return MustParseUint[T](string(data), base)
}

// ParseUint is generics-enabled variant of `strconv.ParseUint`, which derives the bitsize from the
// specified parametric type.
func ParseUint[T types.UnsignedInteger](s string, base int) (T, error) {
	// TODO ideally, we get a constant expression, but this is not possible within the function.
	var bitsize = int(unsafe.Sizeof(T(0))) * 8
	result, err := strconv.ParseUint(s, base, bitsize)
	return T(result), err
}

// ParseUintBytes is generics-enabled variant of `strconv.ParseUint` that parses a byte-string and converts
// the data on-the-fly to string.
func ParseUintBytes[T types.UnsignedInteger](data []byte, base int) (T, error) {
	return ParseUint[T](string(data), base)
}
