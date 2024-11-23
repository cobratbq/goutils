// SPDX-License-Identifier: LGPL-3.0-only

package prefixed

import (
	"bytes"
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

// FIXME redo Read... implementations now that encoding/writing has matured

type Header struct {
	Vtype      CompositeType
	Size       uint16
	Terminated bool
}

// ReadHeader reads the 1-byte or 2-byte header from input-data
// - data: input-data
func ReadHeader(data []byte) (uint, Header) {
	if len(data) < 1 {
		return 0, Header{}
	}
	var vtype CompositeType
	if data[0]&FLAG_KEYVALUE == FLAG_KEYVALUE {
		vtype |= 1
	}
	if data[0]&FLAG_MULTIPLICITY == FLAG_MULTIPLICITY {
		vtype |= 2
	}
	var term = data[0]&FLAG_TERMINATION == FLAG_TERMINATION
	var size = uint16(data[0] & MASK_SIZEBITS)
	if data[0]&FLAG_HEADERSIZE == 0 {
		return 1, Header{vtype, size, term}
	}
	if len(data) < 2 {
		return 0, Header{}
	}
	size <<= 8
	size |= uint16(data[1])
	size += SIZE_2BYTE_OFFSET
	return 2, Header{vtype, size, term}
}

func readOrCopyHeader(data []byte, _hdr *Header) (uint, Header) {
	if _hdr != nil {
		return 0, *_hdr
	}
	return ReadHeader(data)
}

// ReadBytes reads plain bytes.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
// FIXME check/redo size-checks, especially inside the loops
func ReadBytes(data []byte, _hdr *Header) (uint, Bytes) {
	if len(data) < 1 {
		return 0, nil
	}
	var n uint
	var h Header
	if n, h = readOrCopyHeader(data, _hdr); (_hdr == nil && n == 0) || h.Vtype != TYPE_BYTES {
		return 0, nil
	}
	var pos = n
	var b bytes.Buffer
	for {
		if len(data[pos:]) < int(h.Size) {
			return 0, nil
		}
		b.Write(data[pos : pos+uint(h.Size)])
		pos += uint(h.Size)
		if h.Terminated {
			return pos, Bytes(bytes.Clone(b.Bytes()))
		}
		if n, h = ReadHeader(data[pos:]); n == 0 || h.Vtype != TYPE_BYTES {
			return 0, nil
		}
		pos += n
	}
}

// ReadKeyValue reads the key-value from input-data.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
// FIXME check/redo size-checks, especially inside the loops
// FIXME support non-terminated key-entry
// FIXME return a *KeyValue instead? (Would match better with SequenceValue and MapValue due to "by-reference" nature of those inner types)
func ReadKeyValue(data []byte, _hdr *Header) (uint, KeyValue) {
	if len(data) < 1 {
		return 0, KeyValue{}
	}
	var pos, n uint
	var h Header
	if n, h = readOrCopyHeader(data, _hdr); (_hdr == nil && n == 0) || h.Vtype != TYPE_KEYVALUE {
		return 0, KeyValue{}
	}
	pos += n
	if len(data[pos:]) < int(h.Size)+1 {
		return 0, KeyValue{}
	}
	key := string(data[pos : pos+uint(h.Size)])
	pos += uint(h.Size)
	var val Value
	if n, val = ReadValue(data[pos:]); n == 0 {
		return 0, KeyValue{}
	}
	return pos + n, KeyValue{K: key, V: val}
}

// ReadSequence reads a sequence-value from input-data.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
// FIXME check/redo size-checks, especially inside the loops
func ReadSequence(data []byte, _hdr *Header) (uint, SequenceValue) {
	if len(data) < 1 {
		return 0, nil
	}
	var pos, n uint
	var h Header
	if n, h = readOrCopyHeader(data, _hdr); (_hdr == nil && n == 0) || h.Vtype != TYPE_SEQUENCE {
		return 0, nil
	}
	pos += n
	var entries = make([]Value, 0, h.Size)
	for {
		for i := uint(0); i < uint(h.Size); i++ {
			n, entry := ReadValue(data[pos:])
			if n == 0 {
				return 0, nil
			}
			entries = append(entries, entry)
			pos += n
		}
		if h.Terminated {
			return pos, entries
		}
		if n, h = ReadHeader(data[pos:]); n == 0 || h.Vtype != TYPE_SEQUENCE {
			return 0, nil
		}
		pos += n
	}
}

// ReadMap reads a map-value from input-data.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
// TODO map assumes distinct keys, hence count is exact number of map entries.
func ReadMap(data []byte, _hdr *Header) (uint, MapValue) {
	var pos, n uint
	var h Header
	if n, h = readOrCopyHeader(data, _hdr); (_hdr == nil && n == 0) || h.Vtype != TYPE_MAP {
		return 0, nil
	}
	pos += n
	entries := make(map[string]Value, h.Size)
	var v KeyValue
	for {
		for i := uint(0); i < uint(h.Size); i++ {
			if n, v = ReadKeyValue(data[pos:], nil); n == 0 {
				return 0, nil
			}
			entries[v.K] = v.V
			pos += n
		}
		if h.Terminated {
			return pos, entries
		}
		if n, h = ReadHeader(data[pos:]); n == 0 || h.Vtype != TYPE_MAP {
			return 0, nil
		}
		pos += n
	}
}

// ReadValue reads a value of any type from input-data.
// - data: input-data
// TODO currently borrows data from input-array when constructing types, i.e. no cloning.
// TODO future: add support for custom mapping of type-to-readFunction mapping for custom types
func ReadValue(data []byte) (uint, Value) {
	var h Header
	var n uint
	if n, h = ReadHeader(data); n == 0 {
		return 0, nil
	}
	switch h.Vtype {
	case TYPE_BYTES:
		return ReadBytes(data[n:], &h)
	case TYPE_KEYVALUE:
		return ReadKeyValue(data[n:], &h)
	case TYPE_SEQUENCE:
		return ReadSequence(data[n:], &h)
	case TYPE_MAP:
		return ReadMap(data[n:], &h)
	default:
		panic("BUG: should not be reached")
	}
}
