// SPDX-License-Identifier: AGPL-3.0-or-later

package builtin

import (
	"math"
	"math/bits"
	"strconv"
)

// IntSize is the bit-size of the data-type for the targeted platform.
const IntSize = strconv.IntSize

// MaxInt is the maximum int value for the targeted platform. Other constants for data-types can be
// found in package `math` package.
const MaxInt = math.MaxInt

// MinInt is the minimum int value for the targeted platform. Other constants for data-types can be
// found in package `math` package.
const MinInt = math.MinInt

// UintSize is the bit-size of the data-type for the targeted platform.
const UintSize = bits.UintSize

// MaxUint is the maximum uint value for the targeted platform. Other constants for data-types can
// be found in package `math` package.
const MaxUint = math.MaxUint
