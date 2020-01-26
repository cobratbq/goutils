package strconv

import (
	"strconv"

	"github.com/cobratbq/goutils/std/errors"
)

// MustParseInt parses a string for an integer value of at most specified
// bitsize. Success is assumed and the function will panic on error.
func MustParseInt(v string, base, bitsize int) int64 {
	num, err := strconv.ParseInt(v, base, bitsize)
	errors.RequireSuccess(err, "illegal string representation of int: %v")
	return num
}
