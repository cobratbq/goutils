// SPDX-License-Identifier: AGPL-3.0-or-later

package builtin

// EqualsAny matches the specified value with any of the provided `matches` values. It returns
// `true` if it is any of the provided matches, or `false` if none match.
func EqualsAny[T comparable](value T, matches ...T) bool {
	for _, m := range matches {
		if value == m {
			return true
		}
	}
	return false
}
