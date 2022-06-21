package time

import (
	"strconv"
	"time"
)

func NanotimeAsHexString(t time.Time) string {
	return strconv.FormatInt(t.UnixNano(), 16)
}
