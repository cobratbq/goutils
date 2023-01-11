// SPDX-License-Identifier: AGPL-3.0-or-later

package time

import (
	"strconv"
	"time"
)

func NanotimeAsStringHex(t time.Time) string {
	return strconv.FormatInt(t.UnixNano(), 16)
}
