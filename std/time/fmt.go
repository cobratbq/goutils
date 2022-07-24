// SPDX-License-Identifier: LGPL-3.0-or-later

package time

import (
	"strconv"
	"time"
)

func NanotimeAsStringHex(t time.Time) string {
	return strconv.FormatInt(t.UnixNano(), 16)
}
