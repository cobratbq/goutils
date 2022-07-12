// SPDX-License-Identifier: LGPL-3.0-or-later
package http

import (
	"net/http"
	"time"

	"github.com/cobratbq/goutils/std/errors"
)

var ErrHeaderMissing = errors.NewStringError(`'Date' header not available in response`)

func ExtractResponseHeaderDate(resp *http.Response) (time.Time, error) {
	dates, ok := resp.Header["Date"]
	if !ok || len(dates) <= 0 {
		return time.Time{}, ErrHeaderMissing
	}
	return time.Parse(time.RFC1123, dates[0])
}
