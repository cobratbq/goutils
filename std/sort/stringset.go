// SPDX-License-Identifier: AGPL-3.0-or-later

package sort

import (
	"sort"

	"github.com/cobratbq/goutils/std/builtin"
)

// StringSet sorts the keys in a set (a map[string]struct{}) in
// lexicographical order.
func StringSet(set map[string]struct{}) []string {
	keys := builtin.ExtractKeys(set)
	sort.Strings(keys)
	return keys
}
