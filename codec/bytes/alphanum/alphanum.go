// SPDX-License-Identifier: LGPL-3.0-only

package alphanum

import (
	"github.com/cobratbq/goutils/codec/bytes/digit"
)

func IsAlphanum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || digit.IsDigit(c)
}
