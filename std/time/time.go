// SPDX-License-Identifier: LGPL-3.0-only

package time

import "time"

func MiddleTimestamps(earlier, later time.Time) time.Time {
	delta := later.Sub(earlier)
	return earlier.Add(time.Duration(delta.Nanoseconds() / 2))
}
