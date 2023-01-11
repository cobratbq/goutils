// SPDX-License-Identifier: AGPL-3.0-or-later

package sort

import "sort"

func Slice[E any](vals []E, less func(i, j int) bool) {
	sort.Slice(vals, less)
}
