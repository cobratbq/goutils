// SPDX-License-Identifier: LGPL-3.0-only

package prefixed

import (
	"bytes"
	"io"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/codec/bytes/bigendian"
	"github.com/cobratbq/goutils/std/builtin/maps"
	"github.com/cobratbq/goutils/std/builtin/slices"
	io_ "github.com/cobratbq/goutils/std/io"
	"github.com/cobratbq/goutils/std/math"
)

// TODO there is going to be nested wrapping of _out into CountingWriter with recursive calls to various value-types. This is probably not ideal. :-P

func writeHeader(out io.Writer, count int, flags byte) (int, error) {
	assert.Equal(0, flags&FLAG_HEADERSIZE)
	if uint(count) <= SIZE_1BYTE_MAX {
		return out.Write([]byte{byte(count) | flags})
	} else if uint(count) <= SIZE_2BYTE_MAX {
		var header = bigendian.FromUint16(uint16(count) - SIZE_2BYTE_OFFSET)
		header[0] |= flags | FLAG_HEADERSIZE
		return out.Write(header[:])
	} else {
		panic("BUG: illegal count, unsupported by encoding.")
	}
}

// WriteRaw writes raw bytes, usually a byte-array with prefixed header byte(s).
// Type-flags need to be provided.
func WriteRaw(_out io.Writer, data []byte, typeflags byte) (int64, error) {
	var err error
	out := io_.NewCountingWriter(_out)
	if len(data) <= int(SIZE_2BYTE_MAX) {
		if _, err = writeHeader(&out, len(data), typeflags|FLAG_TERMINATION); err != nil {
			return out.Cum, err
		}
		if _, err = out.Write(data); err != nil {
			return out.Cum, err
		}
	} else {
		dataHead := data[:SIZE_2BYTE_MAX]
		if _, err = writeHeader(&out, len(dataHead), typeflags); err != nil {
			return out.Cum, err
		}
		if _, err = out.Write(dataHead); err != nil {
			return out.Cum, err
		}
		if _, err = WriteRaw(&out, data[SIZE_2BYTE_MAX:], typeflags); err != nil {
			return out.Cum, err
		}
	}
	return out.Cum, nil
}

// Type 0, 0 (singular, plain value) for any length.
type Bytes []byte

// Len returns the number of bytes in the value.
func (v Bytes) Len() int {
	return len(v)
}

// Equal tests equality between this Bytes-value and other value.
func (v Bytes) Equal(other Value) bool {
	if o, ok := other.(Bytes); ok {
		return bytes.Equal(v, o)
	}
	return false
}

func (v Bytes) WriteTo(out io.Writer) (int64, error) {
	return WriteRaw(out, v, 0)
}

// WriteBytes is a one-shot function call for writing a raw bytes as an encoded Bytes-value.
func WriteBytes(out io.Writer, data []byte) (int64, error) {
	return Bytes(data).WriteTo(out)
}

// Type 0, 1 (singular, key-value-pair).
//
// Write a key and corresponding value. Syntactically enforced key with value, requiring no assumptions on
// conventions.
//
// These "labeled" values prove valuable when a number of properties need to be checked, e.g. a version or
// identifier before one can decide on the type of (composite) encoded data structure. The "labeled" value is
// a guarantee that there is a identifier with corresponding value.
//
// Note: data-type `string` chosen here for convenience, because Go swaps between string and []byte without
// active conversion necessary.
type KeyValue struct {
	K string
	V Value
}

// Equal tests if this key-value pair is equal to other.
func (v *KeyValue) Equal(other Value) bool {
	if o, ok := other.(*KeyValue); ok {
		return v.K == o.K && v.V.Equal(o.V)
	}
	return false
}

// Len returns the size of the key in the key-value pair.
func (v *KeyValue) Len() int {
	return len(v.K)
}

func (v *KeyValue) WriteTo(_out io.Writer) (int64, error) {
	var err error
	out := io_.NewCountingWriter(_out)
	if _, err = WriteRaw(&out, []byte(v.K), FLAG_KEYVALUE); err != nil {
		return out.Cum, err
	}
	_, err = v.V.WriteTo(&out)
	return out.Cum, err
}

// WriteKeyValue is a one-shot function for writing encoded key-value pair.
func WriteKeyValue(out io.Writer, key string, value Value) (int64, error) {
	return (&KeyValue{key, value}).WriteTo(out)
}

// Type 1, 0 (multiple, plain value) for any length.
type SequenceValue []Value

// Equal tests equality of this value with some other value.
func (v SequenceValue) Equal(other Value) bool {
	if o, ok := other.(SequenceValue); ok {
		return slices.EqualT(v, o)
	}
	return false
}

// Len returns the number of elements in a sequence-value.
func (v SequenceValue) Len() int {
	return len(v)
}

func (v SequenceValue) WriteTo(_out io.Writer) (int64, error) {
	var err error
	out := io_.NewCountingWriter(_out)
	if len(v) <= int(SIZE_2BYTE_MAX) {
		if _, err = writeHeader(&out, len(v), FLAG_TERMINATION|FLAG_MULTIPLICITY); err != nil {
			return out.Cum, err
		}
		for _, e := range v {
			if _, err = e.WriteTo(&out); err != nil {
				return out.Cum, err
			}
		}
	} else {
		subset := v[:SIZE_2BYTE_MAX]
		if _, err = writeHeader(&out, len(subset), FLAG_MULTIPLICITY); err != nil {
			return out.Cum, err
		}
		for _, e := range subset {
			if _, err = e.WriteTo(&out); err != nil {
				return out.Cum, err
			}
		}
		if _, err = SequenceValue(v[SIZE_2BYTE_MAX:]).WriteTo(&out); err != nil {
			return out.Cum, err
		}
	}
	return out.Cum, nil
}

// WriteSequence is a one-shot function for writing an encoded sequence-value.
func WriteSequence(out io.Writer, seq []Value) (int64, error) {
	return SequenceValue(seq).WriteTo(out)
}

// Type 1, 1 (multiple, key-value-pairs) for any length.
type MapValue map[string]Value

// Equal tests equality between this and another value.
func (v MapValue) Equal(other Value) bool {
	if o, ok := other.(MapValue); ok {
		return maps.EqualT(v, o)
	}
	return false
}

// Len returns the number of key-value pairs in a map-value.
func (v MapValue) Len() int {
	return len(v)
}

func (v MapValue) WriteTo(_out io.Writer) (int64, error) {
	var err error
	out := io_.NewCountingWriter(_out)
	if len(v) <= int(SIZE_2BYTE_MAX) {
		if _, err = writeHeader(&out, len(v), FLAG_TERMINATION|FLAG_KEYVALUE|FLAG_MULTIPLICITY); err != nil {
			return out.Cum, err
		}
		for key, value := range v {
			if _, err = WriteRaw(&out, []byte(key), FLAG_KEYVALUE); err != nil {
				return out.Cum, err
			}
			if _, err = value.WriteTo(&out); err != nil {
				return out.Cum, err
			}
		}
	} else {
		var total, processed, part = uint(len(v)), uint(0), uint(0)
		for key, value := range v {
			if part == 0 {
				// Whenever part==0, start a new batch, i.e. new map-value with its own items.
				part = math.Min(total-processed, SIZE_2BYTE_MAX)
				assert.AtMost(total, processed+part)
				assert.Positive(part)
				if part <= SIZE_2BYTE_MAX {
					flags := FLAG_KEYVALUE | FLAG_MULTIPLICITY
					if processed+part == total {
						flags |= FLAG_TERMINATION
					}
					if _, err = writeHeader(&out, int(part), flags); err != nil {
						return out.Cum, err
					}
				} else {
					panic("BUG: we should have selected at most SIZE_2BYTE_MAX for part")
				}
			}
			if _, err = WriteRaw(&out, []byte(key), FLAG_KEYVALUE); err != nil {
				return out.Cum, err
			}
			if _, err = value.WriteTo(&out); err != nil {
				return out.Cum, err
			}
			processed, part = processed+1, part-1
		}
		assert.Equal(0, part)
		assert.Equal(total, processed)
	}
	return out.Cum, nil
}

// WriteMap write a map as encoded map-value.
func WriteMap(out io.Writer, mapv map[string]Value) (int64, error) {
	return MapValue(mapv).WriteTo(out)
}
