// SPDX-License-Identifier: AGPL-3.0-or-later

package time

import (
	"time"
)

func TimeDeltaCorrectionFunc(systemTime, currentTime *time.Time) func() time.Time {
	return func() time.Time {
		return currentTime.Add(time.Since(*systemTime))
	}
}
