// SPDX-License-Identifier: LGPL-3.0-only

package url

import (
	"net/url"
	"path"
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

// ExtractFilename extracts the filename component from the URL path. Unlike `path.Base`, this does not
// consider the last subdirectory of a path with trailing slash as the base. It returns empty string instead.
func ExtractFilename(url *url.URL) string {
	if url.Path[len(url.Path)-1] == '/' {
		// `path.Base` will take the last section even if there is a trailing slash, to ensure we have a file-
		// name, check for trailing slash first.
		return ""
	}
	return path.Base(url.Path)
}

// ExtractFilenameFromLocation extracts the filename component from an URL location string. Unlike
// `path.Base`, this does not consider the last subdirectory of a path with trailing slash as the base. It
// returns empty string instead.
func ExtractFilenameFromLocation(location string) (string, error) {
	u, err := url.Parse(location)
	if err != nil {
		return "", errors.ErrIllegal
	}
	return ExtractFilename(u), nil
}
