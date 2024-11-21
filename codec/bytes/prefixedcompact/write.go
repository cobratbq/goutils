package prefixed

import (
	"io"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/codec/bytes/bigendian"
	io_ "github.com/cobratbq/goutils/std/io"
	"github.com/cobratbq/goutils/std/math"
)

// FIXME there is going to be nested wrapping of _out into CountingWriter with recursive calls to various value-types. This is probably not ideal. :-P

// WriteRaw writes raw bytes, usually a byte-array with prefixed header byte(s).
// Type-flags need to be provided.
func WriteRaw(_out io.Writer, data []byte, typeflags byte) (int64, error) {
	var err error
	out := io_.NewCountingWriter(_out)
	if len(data) <= int(SIZE_1BYTE_MAX) {
		if _, err = out.Write([]byte{byte(len(data)) | typeflags | FLAG_TERMINATION}); err != nil {
			return out.Cum, err
		}
		if _, err = out.Write(data); err != nil {
			return out.Cum, err
		}
	} else if len(data) <= int(SIZE_2BYTE_MAX) {
		var header = bigendian.FromUint16(uint16(len(data)) - SIZE_2BYTE_OFFSET)
		header[0] |= typeflags | FLAG_TERMINATION | FLAG_HEADERSIZE
		if _, err = out.Write(header); err != nil {
			return out.Cum, err
		}
		if _, err = out.Write(data); err != nil {
			return out.Cum, err
		}
	} else {
		dataHead := data[:SIZE_2BYTE_MAX]
		var header = bigendian.FromUint16(uint16(len(dataHead)) - SIZE_2BYTE_OFFSET)
		header[0] |= typeflags | FLAG_HEADERSIZE
		if _, err = out.Write(header); err != nil {
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

func (v Bytes) WriteTo(out io.Writer) (int64, error) {
	return WriteRaw(out, v, 0)
}

// Type 0, 1 (singular, key-value-pair).
// Write a key and corresponding value. Syntactically enforced key with value, requiring no assumptions on
// conventions. These "labeled" values prove valuable when a number of properties need to be checked, e.g.
// a version or identifier before one can decide on the type of (composite) encoded data structure. The
// "labeled" value is a guarantee that there is a identifier with corresponding value.
type KeyValue struct {
	K string
	V Value
}

func (v KeyValue) WriteTo(_out io.Writer) (int64, error) {
	var err error
	out := io_.NewCountingWriter(_out)
	if _, err = WriteRaw(&out, []byte(v.K), FLAG_KEYVALUE); err != nil {
		return out.Cum, err
	}
	_, err = v.V.WriteTo(&out)
	return out.Cum, err
}

// Type 1, 0 (multiple, plain value) for any length.
type SequenceValue []Value

func (v SequenceValue) WriteTo(_out io.Writer) (int64, error) {
	var err error
	out := io_.NewCountingWriter(_out)
	if len(v) <= int(SIZE_1BYTE_MAX) {
		if _, err = out.Write([]byte{byte(len(v)) | FLAG_TERMINATION | FLAG_MULTIPLICITY}); err != nil {
			return out.Cum, err
		}
		for _, e := range v {
			if _, err = e.WriteTo(&out); err != nil {
				return out.Cum, err
			}
		}
	} else if len(v) <= int(SIZE_2BYTE_MAX) {
		header := bigendian.FromUint16(uint16(len(v)) - SIZE_2BYTE_OFFSET)
		header[0] |= FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_HEADERSIZE
		if _, err = out.Write(header); err != nil {
			return out.Cum, err
		}
		for _, e := range v {
			if _, err = e.WriteTo(&out); err != nil {
				return out.Cum, err
			}
		}
	} else {
		subset := v[:SIZE_2BYTE_MAX]
		header := bigendian.FromUint16(uint16(len(subset)) - SIZE_2BYTE_OFFSET)
		header[0] |= FLAG_MULTIPLICITY | FLAG_HEADERSIZE
		if _, err = out.Write(header); err != nil {
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

// Type 1, 1 (multiple, key-value-pairs) for any length.
type MapValue map[string]Value

func (v MapValue) WriteTo(_out io.Writer) (int64, error) {
	var err error
	out := io_.NewCountingWriter(_out)
	var total = uint(len(v))
	if total <= SIZE_1BYTE_MAX {
		if _, err = out.Write([]byte{byte(total) | FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE}); err != nil {
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
	} else if total <= SIZE_2BYTE_MAX {
		header := bigendian.FromUint16(uint16(total) - SIZE_2BYTE_OFFSET)
		header[0] |= FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE | FLAG_HEADERSIZE
		if _, err = out.Write(header); err != nil {
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
		var processed, part uint
		for key, value := range v {
			if part == 0 {
				// Whenever part==0, start a new batch, i.e. new map-value with its own items.
				part = math.Min(total-processed, SIZE_2BYTE_MAX)
				assert.AtMost(total, processed+part)
				assert.Positive(part)
				if part <= SIZE_1BYTE_MAX {
					if _, err = out.Write([]byte{byte(part) | FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE}); err != nil {
						return out.Cum, err
					}
				} else if part <= SIZE_2BYTE_MAX {
					header := bigendian.FromUint16(uint16(part) - SIZE_2BYTE_OFFSET)
					header[0] |= FLAG_MULTIPLICITY | FLAG_KEYVALUE | FLAG_HEADERSIZE
					if processed+part == total {
						header[0] |= FLAG_TERMINATION
					}
					if _, err = out.Write(header); err != nil {
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
