// SPDX-License-Identifier: AGPL-3.0-or-later

package slices

import (
	"testing"

	t_ "github.com/cobratbq/goutils/std/testing"
)

func TestUniformDimensions2D(t *testing.T) {
	t_.True(t, UniformDimensions2D([][]uint{}))
	t_.True(t, UniformDimensions2D([][]uint{{}, {}, {}}))
	t_.True(t, UniformDimensions2D([][]uint{{1}, {2}, {3}}))
	t_.False(t, UniformDimensions2D([][]uint{{1}, {2}, {}}))
	t_.False(t, UniformDimensions2D([][]uint{{1, 3, 4}, {2, 3}, {}}))
}
