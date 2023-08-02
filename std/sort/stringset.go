// SPDX-License-Identifier: LGPL-3.0-only

package sort

import (
	"sort"

	"github.com/cobratbq/goutils/std/builtin/maps"
)

// StringSet sorts the keys in a set (a map[string]struct{}) in
// lexicographical order.
// TODO add parametric type
func StringSet(set map[string]struct{}) []string {
	keys := maps.ExtractKeys(set)
	sort.Strings(keys)
	return keys
}
