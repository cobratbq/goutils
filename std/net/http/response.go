// SPDX-License-Identifier: LGPL-3.0-only

package http

import (
	"net/http"
	"strings"
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

// RespondMethodNotAllowed responds with status-code 405 Method Not Allowed with list of allowed methods.
//
// allowed: allowed methods, such as "CONNECT", "GET", etc. (or constants such as `http.MethodConnect`)
func RespondMethodNotAllowed(resp http.ResponseWriter, allowed []string, body []byte) (int, error) {
	resp.Header().Set("Allow", strings.Join(allowed, ", "))
	resp.WriteHeader(http.StatusMethodNotAllowed)
	return resp.Write(body)
}
