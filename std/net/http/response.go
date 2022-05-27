package http

import (
	"errors"
	"net/http"
	"time"
)

var ErrDateHeaderMissing = errors.New(`'Date' header not available in response`)

func ExtractResponseHeaderDate(resp *http.Response) (time.Time, error) {
	dates, ok := resp.Header["Date"]
	if !ok || len(dates) <= 0 {
		return time.Time{}, ErrDateHeaderMissing
	}
	return time.Parse(time.RFC1123, dates[0])
}
