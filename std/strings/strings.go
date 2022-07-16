// SPDX-License-Identifier: LGPL-3.0-or-later

package strings

import "strings"

func AnyPrefix(s string, prefixes ...string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
