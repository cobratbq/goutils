// SPDX-License-Identifier: LGPL-3.0-or-later

package builtin

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
