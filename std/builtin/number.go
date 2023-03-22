// SPDX-License-Identifier: AGPL-3.0-or-later

package builtin

import "github.com/cobratbq/goutils/types"

func Zero[T types.Number](v T) bool {
	return v == 0
}

func NonZero[T types.Number](v T) bool {
	return v != 0
}
