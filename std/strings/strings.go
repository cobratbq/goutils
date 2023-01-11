// SPDX-License-Identifier: AGPL-3.0-or-later

package strings

import "strings"

// AnyPrefix tests for any of series of prefixes.
func AnyPrefix(s string, prefixes ...string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
