// SPDX-License-Identifier: AGPL-3.0-or-later

package http

import (
	"net/http"
	"time"

	"github.com/cobratbq/goutils/std/errors"
)

var ErrHeaderMissing = errors.NewStringError(`header unavailable`)

func ExtractResponseHeaderDate(resp *http.Response) (time.Time, error) {
	dates, ok := resp.Header["Date"]
	if !ok || len(dates) <= 0 {
		return time.Time{}, errors.Context(ErrHeaderMissing, "'Date' header missing in response")
	}
	return time.Parse(time.RFC1123, dates[0])
}
