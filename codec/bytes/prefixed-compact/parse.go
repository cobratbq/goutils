// SPDX-License-Identifier: LGPL-3.0-only

// TODO at some point implement read-functions for reading from io.Reader
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

type Header struct {
	Vtype      CompositeType
	Size       uint16
	Terminated bool
}

// ParseHeader reads the 1-byte or 2-byte header from input-data
// - data: input-data
func ParseHeader(data []byte) (uint, Header) {
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

func parseOrCopyHeader(data []byte, _hdr *Header) (uint, Header) {
	if _hdr != nil {
		return 0, *_hdr
	}
	return ParseHeader(data)
}

// ParseBytes reads plain bytes.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
func ParseBytes(data []byte, _hdr *Header) (uint, Bytes) {
	if len(data) < 1 {
		return 0, nil
	}
	var n uint
	var h Header
	if n, h = parseOrCopyHeader(data, _hdr); (_hdr == nil && n == 0) || h.Vtype != TYPE_BYTES {
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
		if n, h = ParseHeader(data[pos:]); n == 0 || h.Vtype != TYPE_BYTES {
			return 0, nil
		}
		pos += n
	}
}

// ParseKeyValue reads the key-value from input-data.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
// FIXME check/redo size-checks, especially inside the loops
// FIXME support non-terminated key-entry
func ParseKeyValue(data []byte, _hdr *Header) (uint, *KeyValue) {
	if len(data) < 1 {
		return 0, nil
	}
	var pos, n uint
	var h Header
	if n, h = parseOrCopyHeader(data, _hdr); (_hdr == nil && n == 0) || h.Vtype != TYPE_KEYVALUE {
		return 0, nil
	}
	pos += n
	if len(data[pos:]) < int(h.Size)+1 {
		return 0, nil
	}
	key := string(data[pos : pos+uint(h.Size)])
	pos += uint(h.Size)
	var val Value
	if n, val = ParseValue(data[pos:]); n == 0 {
		return 0, nil
	}
	return pos + n, &KeyValue{K: key, V: val}
}

// ParseSequence reads a sequence-value from input-data.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
func ParseSequence(data []byte, _hdr *Header) (uint, SequenceValue) {
	if len(data) < 1 {
		return 0, nil
	}
	var pos, n uint
	var h Header
	if n, h = parseOrCopyHeader(data, _hdr); (_hdr == nil && n == 0) || h.Vtype != TYPE_SEQUENCE {
		return 0, nil
	}
	pos += n
	var entries = make([]Value, 0, h.Size)
	for {
		for i := uint16(0); i < h.Size; i++ {
			n, entry := ParseValue(data[pos:])
			if n == 0 {
				return 0, nil
			}
			entries = append(entries, entry)
			pos += n
		}
		if h.Terminated {
			return pos, entries
		}
		if n, h = ParseHeader(data[pos:]); n == 0 || h.Vtype != TYPE_SEQUENCE {
			return 0, nil
		}
		pos += n
	}
}

// ParseMap reads a map-value from input-data.
// - data: input-data
// - _hdr: the header is first read if it is not already provided.
// TODO map assumes distinct keys, hence count is exact number of map entries.
func ParseMap(data []byte, _hdr *Header) (uint, MapValue) {
	var pos, n uint
	var h Header
	if n, h = parseOrCopyHeader(data, _hdr); (_hdr == nil && n == 0) || h.Vtype != TYPE_MAP {
		return 0, nil
	}
	pos += n
	entries := make(map[string]Value, h.Size)
	var v *KeyValue
	for {
		for i := uint16(0); i < h.Size; i++ {
			if n, v = ParseKeyValue(data[pos:], nil); n == 0 {
				return 0, nil
			}
			entries[v.K] = v.V
			pos += n
		}
		if h.Terminated {
			return pos, entries
		}
		if n, h = ParseHeader(data[pos:]); n == 0 || h.Vtype != TYPE_MAP {
			return 0, nil
		}
		pos += n
	}
}

// ParseValue reads a value of any type from input-data.
// - data: input-data
// TODO currently borrows data from input-array when constructing types, i.e. no cloning.
// TODO future: add support for custom mapping of type-to-readFunction mapping for custom types
func ParseValue(data []byte) (uint, Value) {
	var n_h uint
	var h Header
	if n_h, h = ParseHeader(data); n_h == 0 {
		return 0, nil
	}
	var n_v uint
	var v Value
	switch h.Vtype {
	case TYPE_BYTES:
		n_v, v = ParseBytes(data[n_h:], &h)
	case TYPE_KEYVALUE:
		n_v, v = ParseKeyValue(data[n_h:], &h)
	case TYPE_SEQUENCE:
		n_v, v = ParseSequence(data[n_h:], &h)
	case TYPE_MAP:
		n_v, v = ParseMap(data[n_h:], &h)
	default:
		panic("BUG: should not be reached")
	}
	return n_h + n_v, v
}
