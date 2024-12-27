// SPDX-License-Identifier: LGPL-3.0-only

package prefixed

import (
	"bytes"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestVerify(t *testing.T) {
	err := Verify(bytes.NewReader([]byte{}))
	assert.Nil(t, err)
}
