package prefixed

import (
	"io"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/codec/bytes/bigendian"
	"github.com/cobratbq/goutils/std/math"
)

// FIXME do error handling
// FIXME properly update count
// WriteRaw writes raw bytes, usually a byte-array with prefixed header byte(s).
// Type-flags need to be provided.
func WriteRaw(out io.Writer, data []byte, typeflags byte) (int64, error) {
	var count int
	if len(data) <= int(SIZE_1BYTE_MAX) {
		out.Write([]byte{byte(len(data)) | typeflags | FLAG_TERMINATION})
		out.Write(data)
		return int64(count), nil
	} else if len(data) <= int(SIZE_2BYTE_MAX) {
		var header = bigendian.FromUint16(uint16(len(data)) - 1)
		header[0] |= typeflags | FLAG_TERMINATION | FLAG_HEADERSIZE
		out.Write(header)
		out.Write(data)
		return int64(count), nil
	} else {
		dataHead := data[:SIZE_2BYTE_MAX]
		var header = bigendian.FromUint16(uint16(len(dataHead)) - 1)
		header[0] |= typeflags | FLAG_HEADERSIZE
		out.Write(header)
		out.Write(dataHead)
		var countNext int64
		countNext, _ = WriteRaw(out, data[SIZE_2BYTE_MAX:], typeflags)
		return int64(count) + countNext, nil
	}
}

// Type 0, 0 (singular, plain value) for any length.
type Bytes []byte

func (v Bytes) WriteTo(out io.Writer) (int64, error) {
	return WriteRaw(out, v, 0)
}

// Type 0, 1 (singular, key-value-pair) for any length.
// Write a key and corresponding value. Syntactically enforced key with value, requiring no assumptions on
// conventions. These "labeled" values provide valuable when a number of properties need to be checked, e.g.
// a version or identifier before one can decide on the type of (composite) encoded data structure.
// FIXME KeyValue is probably better described as "labeled" value, i.e. a value with an accompanying label that provides an informal indication of meaning.
type KeyValue struct {
	K string
	V Value
}

// FIXME add proper error handling
// FIXME properly update count
func (v KeyValue) WriteTo(out io.Writer) (int64, error) {
	WriteRaw(out, []byte(v.K), FLAG_KEYVALUE)
	return v.V.WriteTo(out)
}

// Type 1, 0 (multiple, plain value) for any length.
type SequenceValue []Value

// FIXME add proper error handling
// FIXME properly update count
func (v SequenceValue) WriteTo(out io.Writer) (int64, error) {
	var count int
	if len(v) <= int(SIZE_1BYTE_MAX) {
		out.Write([]byte{byte(len(v)) | FLAG_TERMINATION | FLAG_MULTIPLICITY})
		for _, e := range v {
			e.WriteTo(out)
		}
		return int64(count), nil
	} else if len(v) <= int(SIZE_2BYTE_MAX) {
		header := bigendian.FromUint16(uint16(len(v)) - 1)
		header[0] |= FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_HEADERSIZE
		out.Write(header)
		for _, e := range v {
			e.WriteTo(out)
		}
		return int64(count), nil
	} else {
		subset := v[:SIZE_2BYTE_MAX]
		header := bigendian.FromUint16(uint16(len(subset)) - 1)
		header[0] |= FLAG_MULTIPLICITY | FLAG_HEADERSIZE
		out.Write(header)
		for _, e := range subset {
			e.WriteTo(out)
		}
		var countNext int64
		countNext, _ = SequenceValue(v[SIZE_2BYTE_MAX:]).WriteTo(out)
		return int64(count) + countNext, nil
	}
}

// Type 1, 1 (multiple, key-value-pairs) for any length.
type MapValue map[string]Value

// FIXME add proper error handling
// FIXME properly update count
func (v MapValue) WriteTo(out io.Writer) (int64, error) {
	var count int64
	var total = uint(len(v))
	if total <= SIZE_1BYTE_MAX {
		out.Write([]byte{byte(total) | FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE})
		for key, value := range v {
			WriteRaw(out, []byte(key), FLAG_KEYVALUE)
			value.WriteTo(out)
		}
	} else if total <= SIZE_2BYTE_MAX {
		header := bigendian.FromUint16(uint16(total) - 1)
		header[0] |= FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE | FLAG_HEADERSIZE
		out.Write(header)
		for key, value := range v {
			WriteRaw(out, []byte(key), FLAG_KEYVALUE)
			value.WriteTo(out)
		}
	} else {
		var cum, part uint
		for key, value := range v {
			if part == 0 {
				// Whenever part==0, start a new batch, i.e. new map-value with its own items.
				part = math.Min(total-cum, SIZE_2BYTE_MAX)
				assert.Positive(part)
				if part <= SIZE_1BYTE_MAX {
					out.Write([]byte{byte(part) | FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE})
				} else if part <= SIZE_2BYTE_MAX {
					header := bigendian.FromUint16(uint16(part) - 1)
					header[0] |= FLAG_MULTIPLICITY | FLAG_KEYVALUE | FLAG_HEADERSIZE
					if part < SIZE_2BYTE_MAX {
						header[0] |= FLAG_TERMINATION
					}
					out.Write(header)
				} else {
					panic("BUG: we should have selected at most SIZE_2BYTE_MAX for part")
				}
			}
			WriteRaw(out, []byte(key), FLAG_KEYVALUE)
			value.WriteTo(out)
			cum, part = cum+1, part-1
		}
		assert.Equal(0, part)
		assert.Equal(total, cum)
	}
	return count, nil
}
