// SPDX-License-Identifier: LGPL-3.0-only

package prefixed

import (
	"bytes"
	"testing"

	"github.com/cobratbq/goutils/std/crypto/rand"
	assert "github.com/cobratbq/goutils/std/testing"
)

func TestParseWrittenBytes(t *testing.T) {
	var b [6000]byte
	rand.MustReadBytes(b[:])
	var result bytes.Buffer
	var n int64
	var err error
	n, err = writeRaw(&result, b[:], 0)
	assert.Nil(t, err)
	assert.Equal(t, 6004, result.Len())
	assert.Equal(t, n, int64(result.Len()))
	raw := result.Bytes()
	var h Header
	var n2 uint
	n2, h = ParseHeader(raw)
	assert.Equal(t, 2, n2)
	assert.Equal(t, 0, h.Vtype)
	assert.Equal(t, 4096, h.Size)
	assert.Equal(t, false, h.Terminated)
	assert.SlicesEqual(t, b[:4096], raw[2:4098])
	n2, h = ParseHeader(raw[4098:])
	assert.Equal(t, 2, n2)
	assert.Equal(t, 0, h.Vtype)
	assert.Equal(t, 1904, h.Size)
	assert.Equal(t, true, h.Terminated)
	assert.SlicesEqual(t, b[4096:], raw[4100:])
}

func TestParseBytes(t *testing.T) {
	var testdata = []struct {
		value   Bytes
		encoded []byte
	}{
		{value: Bytes([]byte{}), encoded: []byte{0 | FLAG_TERMINATION}},
		{value: Bytes([]byte{}), encoded: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 | FLAG_TERMINATION}},
		{value: Bytes([]byte("Hello to all earthlings.")), encoded: []byte{0 | FLAG_TERMINATION | FLAG_HEADERSIZE, 23, 'H', 'e', 'l', 'l', 'o', ' ', 't', 'o', ' ', 'a', 'l', 'l', ' ', 'e', 'a', 'r', 't', 'h', 'l', 'i', 'n', 'g', 's', '.'}},
		{value: Bytes([]byte("ghost")), encoded: []byte{0, 0, 0, 0, 1, 'g', 1, 'h', 1, 'o', 1, 's', 1, 't', 0, 0, 0, 0, 0, 0, 0, 0, 0 | FLAG_TERMINATION}},
	}
	for i, d := range testdata {
		t.Log("Iteration:", i)
		n, v := ParseBytes(d.encoded, nil)
		assert.Equal(t, len(d.encoded), int(n))
		assert.EqualT[Value](t, &d.value, v)
	}

}

func TestParseKeyValue(t *testing.T) {
	var testdata = []struct {
		value   KeyValue
		encoded []byte
	}{
		{value: KeyValue{K: "", V: Bytes([]byte{})}, encoded: []byte{0 | FLAG_TERMINATION | FLAG_KEYVALUE, 0 | FLAG_TERMINATION}},
		{value: KeyValue{K: string([]byte{0}), V: Bytes([]byte{})}, encoded: []byte{1 | FLAG_TERMINATION | FLAG_KEYVALUE, 0, 0 | FLAG_TERMINATION}},
		{value: KeyValue{K: "Hello to all earthlings.", V: Bytes([]byte{})}, encoded: []byte{0 | FLAG_TERMINATION | FLAG_KEYVALUE | FLAG_HEADERSIZE, 23, 'H', 'e', 'l', 'l', 'o', ' ', 't', 'o', ' ', 'a', 'l', 'l', ' ', 'e', 'a', 'r', 't', 'h', 'l', 'i', 'n', 'g', 's', '.', 0 | FLAG_TERMINATION}},
		{value: KeyValue{K: "Hello to all earthlings.", V: Bytes([]byte{})}, encoded: []byte{5 | FLAG_KEYVALUE, 'H', 'e', 'l', 'l', 'o', 1 | FLAG_KEYVALUE, ' ', 2 | FLAG_KEYVALUE, 't', 'o', 1 | FLAG_KEYVALUE, ' ', 3 | FLAG_KEYVALUE, 'a', 'l', 'l', 1 | FLAG_KEYVALUE, ' ', 11 | FLAG_TERMINATION | FLAG_KEYVALUE, 'e', 'a', 'r', 't', 'h', 'l', 'i', 'n', 'g', 's', '.', 0 | FLAG_TERMINATION}},
	}
	for i, d := range testdata {
		t.Log("Iteration:", i)
		n, v := ParseKeyValue(d.encoded, nil)
		assert.Equal(t, len(d.encoded), int(n))
		assert.EqualT[Value](t, &d.value, v)
	}
}

func TestParseMap(t *testing.T) {
	var testdata = []struct {
		value   MapValue
		keys    []string
		encoded []byte
	}{
		{value: map[string]Value{}, encoded: []byte{0 | FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE}, keys: []string{}},
		{value: map[string]Value{"": Bytes([]byte{})}, encoded: []byte{1 | FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE, 0 | FLAG_TERMINATION | FLAG_KEYVALUE, 0 | FLAG_TERMINATION}, keys: []string{""}},
	}
	for i, d := range testdata {
		t.Log("Iteration:", i)
		n, v := ParseMap(d.encoded, nil)
		assert.Equal(t, uint(len(d.encoded)), n)
		assert.Equal(t, len(d.keys), len(v))
		assert.AllKeysPresent(t, v, d.keys)
		// FIXME could use a assert.EqualMaps
	}
}
