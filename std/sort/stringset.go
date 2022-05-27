// SPDX-License-Identifier: LGPL-3.0-or-later
package sort

import (
	"sort"
)

// StringSet sorts the keys in a set (a map[string]struct{}) in
// lexicographical order.
func StringSet(set map[string]struct{}) []string {
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
