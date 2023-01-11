// SPDX-License-Identifier: AGPL-3.0-or-later

package time

import "time"

func MiddleTimestamps(earlier, later time.Time) time.Time {
	delta := later.Sub(earlier)
	return earlier.Add(time.Duration(delta.Nanoseconds() / 2))
}
