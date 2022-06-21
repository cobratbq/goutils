// SPDX-License-Identifier: LGPL-3.0-or-later
package url

import (
	"net/url"
	"strings"

	"github.com/cobratbq/goutils/std/errors"
)

const ErrSchemeMissing errors.StringError = "scheme missing"
const ErrSchemeUnknown errors.StringError = "scheme unknown"

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
