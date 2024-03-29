// SPDX-License-Identifier: LGPL-3.0-only

package types

import (
	"math"
	"math/bits"
	"strconv"
)

// IntSize is the bit-size of the data-type for the targeted platform.
const IntSize = strconv.IntSize

// MaxInt is the maximum int value for the targeted platform. Other constants for data-types can be
// found in package `math` package.
const MaxInt int = math.MaxInt

// MinInt is the minimum int value for the targeted platform. Other constants for data-types can be
// found in package `math` package.
const MinInt int = math.MinInt

// UintSize is the bit-size of the data-type for the targeted platform.
const UintSize = bits.UintSize

// MaxUint is the maximum uint value for the targeted platform. Other constants for data-types can
// be found in package `math` package.
const MaxUint uint = math.MaxUint
