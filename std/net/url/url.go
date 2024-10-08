// SPDX-License-Identifier: LGPL-3.0-only

package url

import (
	"net/url"
	"strings"

	"github.com/cobratbq/goutils/std/errors"
)

var ErrSchemeMissing = errors.NewStringError("scheme missing")
var ErrSchemeUnknown = errors.NewStringError("scheme unknown")

func PortOrProtocolDefault(uri *url.URL) (string, error) {
	port := uri.Port()
	if port != "" {
		return port, nil
	}
	scheme := strings.ToLower(uri.Scheme)
	switch scheme {
	case "":
		return "", ErrSchemeMissing
	case "http":
		return "80", nil
	case "https":
		return "443", nil
	default:
		return "", ErrSchemeUnknown
	}
}

func DeriveOrigin(url *url.URL, allowPort bool) (string, error) {
	var origin string
	switch url.Scheme {
	case "http", "ws":
		origin = "http://"
	case "https", "wss":
		origin = "https://"
	default:
		return "", errors.ErrIllegal
	}
	if allowPort {
		return origin + url.Host, nil
	} else {
		return origin + url.Hostname(), nil
	}
}
