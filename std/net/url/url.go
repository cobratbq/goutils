// SPDX-License-Identifier: LGPL-3.0-or-later
package url

import (
	"errors"
	"net/url"
	"strings"
)

var ErrSchemeMissing error = errors.New("scheme missing")
var ErrSchemeUnknown error = errors.New("scheme unknown")

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
