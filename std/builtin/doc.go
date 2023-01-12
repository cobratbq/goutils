// SPDX-License-Identifier: AGPL-3.0-or-later

// builtin contains additional utility functions for built-in types of the Go programming languages.
// This package contains the utils that apply to types and functions that are always immediately
// available.
//
// TODO consider adding function for "compacting" a map, i.e. copying content over into new map such that there is no memory wasted after intense use.
// TODO consider if we need to move slices and maps utils out to prevent `builtin` from growing too big.
package builtin
