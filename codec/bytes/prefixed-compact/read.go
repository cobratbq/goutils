// SPDX-License-Identifier: LGPL-3.0-only

// TODO at some point implement read-functions for reading from io.Reader
package prefixed

import (
	"bytes"
	"io"

	"github.com/cobratbq/goutils/std/errors"
	io_ "github.com/cobratbq/goutils/std/io"
)

type CompositeType uint8

const (
	// singular, value = value, i.e. bytes
	TYPE_BYTES CompositeType = iota
	// singular, key-value = key-value-pair
	TYPE_KEYVALUE
	// multiple, value = sequence of values (ordered by virtue of position in data)
	TYPE_SEQUENCE
	// multiple, key-value = map of key-value-pairs (ordered by virtue of position in data)
	TYPE_MAP
)

type Header struct {
	Vtype      CompositeType
	Size       uint16
	Terminated bool
}

func ReadHeader(in io.Reader) (Header, error) {
	var err error
	var b byte
	if b, err = io_.ReadByte(in); err != nil {
		return Header{}, err
	}
	var vtype CompositeType
	if b&FLAG_KEYVALUE == FLAG_KEYVALUE {
		vtype |= 1
	}
	if b&FLAG_MULTIPLICITY == FLAG_MULTIPLICITY {
		vtype |= 2
	}
	var term = b&FLAG_TERMINATION == FLAG_TERMINATION
	var size = uint16(b & MASK_SIZEBITS)
	if b&FLAG_HEADERSIZE == 0 {
		return Header{vtype, size, term}, nil
	}
	if b, err = io_.ReadByte(in); err != nil {
		return Header{}, err
	}
	size <<= 8
	size |= uint16(b)
	size += SIZE_2BYTE_OFFSET
	return Header{vtype, size, term}, nil
}

func readOrCopyHeader(in io.Reader, _hdr *Header) (Header, error) {
	if _hdr != nil {
		return *_hdr, nil
	}
	return ReadHeader(in)
}

// ParseHeader reads the 1-byte or 2-byte header from input-data
// - data: input-data
func ParseHeader(data []byte) (uint, Header) {
	var in = bytes.NewReader(data)
	hdr, err := ReadHeader(in)
	if err != nil {
		return 0, Header{}
	}
	return uint(in.Size() - int64(in.Len())), hdr
}

func ReadBytes(in io.Reader, _hdr *Header) (Bytes, error) {
	var h Header
	var err error
	if h, err = readOrCopyHeader(in, _hdr); err != nil {
		return nil, err
	} else if h.Vtype != TYPE_BYTES {
		return nil, errors.ErrIllegal
	}
	var b bytes.Buffer
	for {
		if _, err = io.CopyN(&b, in, int64(h.Size)); err != nil {
			return nil, err
		}
		if h.Terminated {
			return Bytes(bytes.Clone(b.Bytes())), nil
		}
		if h, err = ReadHeader(in); err != nil {
			return nil, err
		} else if h.Vtype != TYPE_BYTES {
			return nil, errors.ErrIllegal
		}
	}
}

// ParseBytes reads plain bytes.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
func ParseBytes(data []byte, _hdr *Header) (uint, Bytes) {
	var in = bytes.NewReader(data)
	v, err := ReadBytes(in, _hdr)
	if err != nil {
		return 0, nil
	}
	return uint(in.Size() - int64(in.Len())), v
}

// FIXME support non-terminated key-entry
func ReadKeyValue(in io.Reader, _hdr *Header) (*KeyValue, error) {
	var h Header
	var err error
	if h, err = readOrCopyHeader(in, _hdr); err != nil {
		return nil, err
	} else if h.Vtype != TYPE_KEYVALUE {
		return nil, errors.ErrIllegal
	}
	var key []byte
	if key, err = io_.ReadN(in, uint(h.Size)); err != nil {
		return nil, err
	}
	var val Value
	if val, err = ReadValue(in); err != nil {
		return nil, err
	}
	return &KeyValue{K: string(key), V: val}, nil
}

// ParseKeyValue reads the key-value from input-data.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
func ParseKeyValue(data []byte, _hdr *Header) (uint, *KeyValue) {
	var in = bytes.NewReader(data)
	v, err := ReadKeyValue(in, _hdr)
	if err != nil {
		return 0, nil
	}
	return uint(in.Size() - int64(in.Len())), v
}

func ReadSequence(in io.Reader, _hdr *Header) (SequenceValue, error) {
	var h Header
	var err error
	if h, err = readOrCopyHeader(in, _hdr); err != nil {
		return nil, err
	} else if h.Vtype != TYPE_SEQUENCE {
		return nil, errors.ErrIllegal
	}
	var entries = make([]Value, 0, h.Size)
	for {
		for i := uint16(0); i < h.Size; i++ {
			entry, err := ReadValue(in)
			if err != nil {
				return nil, err
			}
			entries = append(entries, entry)
		}
		if h.Terminated {
			return entries, nil
		}
		if h, err = ReadHeader(in); err != nil {
			return nil, err
		} else if h.Vtype != TYPE_SEQUENCE {
			return nil, errors.ErrIllegal
		}
	}
}

// ParseSequence reads a sequence-value from input-data.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
func ParseSequence(data []byte, _hdr *Header) (uint, SequenceValue) {
	var in = bytes.NewReader(data)
	v, err := ReadSequence(in, _hdr)
	if err != nil {
		return 0, nil
	}
	return uint(in.Size() - int64(in.Len())), v
}

// TODO map assumes distinct keys, hence count is exact number of map entries.
func ReadMap(in io.Reader, _hdr *Header) (MapValue, error) {
	var h Header
	var err error
	if h, err = readOrCopyHeader(in, _hdr); err != nil {
		return nil, err
	} else if h.Vtype != TYPE_MAP {
		return nil, errors.ErrIllegal
	}
	entries := make(map[string]Value, h.Size)
	var v *KeyValue
	for {
		for i := uint16(0); i < h.Size; i++ {
			if v, err = ReadKeyValue(in, nil); err != nil {
				return nil, err
			}
			entries[v.K] = v.V
		}
		if h.Terminated {
			return entries, nil
		}
		if h, err = ReadHeader(in); err != nil {
			return nil, err
		} else if h.Vtype != TYPE_MAP {
			return nil, errors.ErrIllegal
		}
	}
}

// ParseMap reads a map-value from input-data.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
func ParseMap(data []byte, _hdr *Header) (uint, MapValue) {
	var in = bytes.NewReader(data)
	v, err := ReadMap(in, _hdr)
	if err != nil {
		return 0, nil
	}
	return uint(in.Size() - int64(in.Len())), v
}

func ReadValue(in io.Reader) (Value, error) {
	var err error
	var h Header
	h, err = ReadHeader(in)
	if err != nil {
		return nil, err
	}
	switch h.Vtype {
	case TYPE_BYTES:
		return ReadBytes(in, &h)
	case TYPE_KEYVALUE:
		return ReadKeyValue(in, &h)
	case TYPE_SEQUENCE:
		return ReadSequence(in, &h)
	case TYPE_MAP:
		return ReadMap(in, &h)
	default:
		panic("BUG: should not be reached")
	}
}

// ParseValue reads a value of any type from input-data.
// - data: input-data
// TODO currently borrows data from input-array when constructing types, i.e. no cloning.
// TODO future: add support for custom mapping of type-to-readFunction mapping for custom types
func ParseValue(data []byte) (uint, Value) {
	var in = bytes.NewReader(data)
	v, err := ReadValue(in)
	if err != nil {
		return 0, nil
	}
	return uint(in.Size() - int64(in.Len())), v
}
