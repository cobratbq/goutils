// SPDX-License-Identifier: AGPL-3.0-or-later

package time

import (
	"strconv"
	"time"

	strconv_ "github.com/cobratbq/goutils/std/strconv"
)

func NanotimeAsStringHex(t time.Time) string {
	return strconv.FormatInt(t.UnixNano(), strconv_.HexadecimalBase)
}
