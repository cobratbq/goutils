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

func TestVerifyBasicMap(t *testing.T) {
	var b bytes.Buffer
	_, err := WriteMap(&b, map[string]Value{"test1": Bytes("hello"), "test2": Bytes("world")})
	assert.Nil(t, err)
	err = Verify(&b)
	assert.Nil(t, err)
	assert.Equal(t, 0, b.Len())
}

func TestVerifyMixedMap(t *testing.T) {
	var b bytes.Buffer
	_, err := WriteMap(&b, map[string]Value{
		"test1": Bytes("hello"),
		"test2": &KeyValue{"id", Bytes("world")},
		"test3": SequenceValue{Bytes("a"), Bytes("b")}})
	assert.Nil(t, err)
	err = Verify(&b)
	assert.Nil(t, err)
	assert.Equal(t, 0, b.Len())
}
