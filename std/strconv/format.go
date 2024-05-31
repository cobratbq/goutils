package strconv

import (
	"strconv"

	"github.com/cobratbq/goutils/types"
)

func FormatInt[N types.SignedInteger](value N, base int) string {
	return strconv.FormatInt(int64(value), base)
}

func FormatIntDecimal[N types.SignedInteger](value N) string {
	return strconv.FormatInt(int64(value), DecimalBase)
}

func FormatUint[N types.UnsignedInteger](value N, base int) string {
	return strconv.FormatUint(uint64(value), base)
}

func FormatUintDecimal[N types.UnsignedInteger](value N) string {
	return strconv.FormatUint(uint64(value), DecimalBase)
}
