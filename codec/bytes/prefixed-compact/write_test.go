// SPDX-License-Identifier: LGPL-3.0-only

package prefixed

import (
	"bytes"
	"testing"

	"github.com/cobratbq/goutils/std/crypto/rand"
	assert "github.com/cobratbq/goutils/std/testing"
)

func TestWriteRaw(t *testing.T) {
	var testdata = []struct {
		data     []byte
		flags    uint8
		encoded  []byte
		expected error
	}{
		{data: []byte{}, flags: 0, encoded: []byte{FLAG_TERMINATION}, expected: nil},
		{data: []byte{'1'}, flags: 0, encoded: []byte{1 | FLAG_TERMINATION, '1'}, expected: nil},
		{data: []byte{'1'}, flags: 0, encoded: []byte{1 | FLAG_TERMINATION, '1'}, expected: nil},
		{data: []byte{'a', 'b', 'c'}, flags: 0, encoded: []byte{3 | FLAG_TERMINATION, 'a', 'b', 'c'}, expected: nil},
		{data: []byte("Hello my beautiful friends!"), flags: 0, encoded: []byte{0 | FLAG_TERMINATION | FLAG_HEADERSIZE, 26, 'H', 'e', 'l', 'l', 'o', ' ', 'm', 'y', ' ', 'b', 'e', 'a', 'u', 't', 'i', 'f', 'u', 'l', ' ', 'f', 'r', 'i', 'e', 'n', 'd', 's', '!'}, expected: nil},
	}
	var buffer bytes.Buffer
	var n int64
	var err error
	for i, d := range testdata {
		buffer.Reset()
		t.Log("Iteration:", i)
		n, err = WriteRaw(&buffer, d.data, d.flags)
		assert.IsError(t, d.expected, err)
		assert.Equal(t, len(d.encoded), int(n))
		assert.SlicesEqual(t, d.encoded, buffer.Bytes())
	}
}

func TestWriteKeyValue(t *testing.T) {
	var testdata = []struct {
		value    KeyValue
		encoded  []byte
		expected error
	}{
		{value: KeyValue{K: "", V: Bytes([]byte{})}, encoded: []byte{0 | FLAG_TERMINATION | FLAG_KEYVALUE, 0 | FLAG_TERMINATION}, expected: nil},
		{value: KeyValue{K: string([]byte{0}), V: Bytes([]byte{})}, encoded: []byte{1 | FLAG_TERMINATION | FLAG_KEYVALUE, 0, 0 | FLAG_TERMINATION}, expected: nil},
		{value: KeyValue{K: "Hello to all earthlings.", V: Bytes([]byte{})}, encoded: []byte{0 | FLAG_TERMINATION | FLAG_KEYVALUE | FLAG_HEADERSIZE, 23, 'H', 'e', 'l', 'l', 'o', ' ', 't', 'o', ' ', 'a', 'l', 'l', ' ', 'e', 'a', 'r', 't', 'h', 'l', 'i', 'n', 'g', 's', '.', 0 | FLAG_TERMINATION}, expected: nil},
	}
	var b bytes.Buffer
	var n int64
	var err error
	for i, d := range testdata {
		b.Reset()
		t.Log("Iteration:", i)
		n, err = d.value.WriteTo(&b)
		assert.IsError(t, d.expected, err)
		assert.Equal(t, len(d.encoded), b.Len())
		assert.Equal(t, n, int64(b.Len()))
		raw := b.Bytes()
		assert.SlicesEqual(t, d.encoded, raw)
	}
}

func TestWriteSequence(t *testing.T) {
	var testdata = []struct {
		value    SequenceValue
		encoded  []byte
		expected error
	}{
		{value: SequenceValue([]Value{}), encoded: []byte{0 | FLAG_MULTIPLICITY | FLAG_TERMINATION}, expected: nil},
		{value: SequenceValue([]Value{Bytes([]byte{})}), encoded: []byte{1 | FLAG_MULTIPLICITY | FLAG_TERMINATION, 0 | FLAG_TERMINATION}, expected: nil},
	}
	var b bytes.Buffer
	var n int64
	var err error
	for i, d := range testdata {
		b.Reset()
		t.Log("Iteration:", i)
		n, err = d.value.WriteTo(&b)
		assert.IsError(t, d.expected, err)
		assert.Equal(t, len(d.encoded), b.Len())
		assert.Equal(t, n, int64(b.Len()))
		raw := b.Bytes()
		assert.SlicesEqual(t, d.encoded, raw)
	}
}

func TestWriteMap(t *testing.T) {
	var testdata = []struct {
		value    MapValue
		encoded  []byte
		expected error
	}{
		{value: map[string]Value{}, encoded: []byte{0 | FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE}, expected: nil},
		{value: map[string]Value{"": Bytes([]byte{})}, encoded: []byte{1 | FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE, 0 | FLAG_TERMINATION | FLAG_KEYVALUE, 0 | FLAG_TERMINATION}, expected: nil},
	}
	var b bytes.Buffer
	var n int64
	var err error
	for i, d := range testdata {
		b.Reset()
		t.Log("Iteration:", i)
		n, err = d.value.WriteTo(&b)
		assert.IsError(t, d.expected, err)
		assert.Equal(t, len(d.encoded), b.Len())
		assert.Equal(t, n, int64(b.Len()))
		raw := b.Bytes()
		assert.SlicesEqual(t, d.encoded, raw)
	}
}

func TestWriteRawVeryLarge(t *testing.T) {
	var b [6000]byte
	rand.MustReadBytes(b[:])
	var result bytes.Buffer
	var n int64
	var err error
	n, err = WriteRaw(&result, b[:], 0)
	assert.Nil(t, err)
	assert.Equal(t, 6004, n)
	assert.Equal(t, int(n), result.Len())
	raw := result.Bytes()
	assert.Equal(t, 15|FLAG_HEADERSIZE, raw[0])
	assert.Equal(t, 0xff, raw[1])
	assert.Equal(t, 0x07|FLAG_TERMINATION|FLAG_HEADERSIZE, raw[4098])
	assert.Equal(t, 0x6f, raw[4099])
}
