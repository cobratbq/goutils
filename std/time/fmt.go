// SPDX-License-Identifier: LGPL-3.0-only

package time

import (
	"strconv"
	"time"

	strconv_ "github.com/cobratbq/goutils/std/strconv"
)

func NanotimeAsStringHex(t time.Time) string {
	return strconv.FormatInt(t.UnixNano(), strconv_.HexadecimalBase)
}
