// SPDX-License-Identifier: LGPL-3.0-only

package prefixed

import (
	"bytes"
	"testing"

	"github.com/cobratbq/goutils/std/crypto/rand"
	assert "github.com/cobratbq/goutils/std/testing"
)

func TestReadBytes(t *testing.T) {
	//var testdata = []struct {
	//	encoded []byte
	//	data    []byte
	//	flags   uint8
	//	error   error
	//}{
	//	{encoded: []byte{FLAG_TERMINATION}, data: []byte{}, flags: 0, error: io.ErrIncompleteRead},
	//}
	//var buffer bytes.Buffer
	//var pos, n uint
	//var h Header
	//var err error
	//for i, d := range testdata {
	//	var restored []byte
	//	buffer.Reset()
	//	t.Log("Iteration:", i)
	//	n, h = ReadHeader(d.encoded)
	//	assert.IsError(t, d.error, err)
	//	assert.Equal(t, len(d.encoded), int(n))
	//	assert.SlicesEqual(t, d.encoded, buffer.Bytes())
	//	if d.error != nil {
	//		continue
	//	}
	//	buffer.Write(d.encoded[n : n+uint(h.Size)])
	//	pos += n + uint(h.Size)
	//	n, h = ReadHeader(d.encoded[pos:])
	//	assert.IsError(t, d.error, err)
	//	assert.Equal(t, len(d.encoded), int(n))
	//	assert.SlicesEqual(t, d.encoded, buffer.Bytes())
	//	if d.error != nil {
	//		continue
	//	}
	//	buffer.Write(d.encoded[n : n+uint(h.Size)])
	//}
}

func TestReadWrittenBytes(t *testing.T) {
	var b [6000]byte
	rand.MustReadBytes(b[:])
	var result bytes.Buffer
	var n int64
	var err error
	n, err = WriteRaw(&result, b[:], 0)
	assert.Nil(t, err)
	assert.Equal(t, 6004, result.Len())
	assert.Equal(t, n, int64(result.Len()))
	raw := result.Bytes()
	var h Header
	var n2 uint
	n2, h = ReadHeader(raw)
	assert.Equal(t, 2, n2)
	assert.Equal(t, 0, h.Vtype)
	assert.Equal(t, 4096, h.Size)
	assert.Equal(t, false, h.Terminated)
	assert.SlicesEqual(t, b[:4096], raw[2:4098])
	n2, h = ReadHeader(raw[4098:])
	assert.Equal(t, 2, n2)
	assert.Equal(t, 0, h.Vtype)
	assert.Equal(t, 1904, h.Size)
	assert.Equal(t, true, h.Terminated)
	assert.SlicesEqual(t, b[4096:], raw[4100:])
}
