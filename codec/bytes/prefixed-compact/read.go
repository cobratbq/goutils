// SPDX-License-Identifier: LGPL-3.0-only

package prefixed

import "bytes"

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

func ReadBytes(data []byte, hdr *Header) (uint, Bytes) {
	if len(data) < 1 {
		return 0, nil
	}
	var n uint
	var h Header
	if hdr == nil {
		if n, h = ReadHeader(data); n == 0 {
			return 0, nil
		}
	} else {
		h = *hdr
	}
	// FIXME make expected Vtype a parameter and allow use of ReadBytes for both plain values and keys?
	if h.Vtype != TYPE_BYTES {
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

// FIXME support non-terminated key-entry
func ReadKeyValue(data []byte, keysize uint16) (uint, KeyValue) {
	if len(data) < int(keysize)+1 {
		return 0, KeyValue{}
	}
	key := string(data[:keysize])
	var n uint
	var val Value
	if n, val = ReadValue(data[keysize:]); n == 0 {
		return 0, KeyValue{}
	}
	return uint(keysize) + n, KeyValue{K: key, V: val}
}

// FIXME support non-terminated sequence
func ReadSequence(data []byte, count uint16) (uint, SequenceValue) {
	var entries = make([]Value, count)
	var pos, n uint
	for i := uint(0); i < uint(count); i++ {
		if n, entries[i] = ReadValue(data[pos:]); n == 0 {
			return 0, nil
		}
		pos += n
	}
	return pos, entries
}

// FIXME support non-terminated map
// TODO map assumes distinct keys, hence count is exact number of map entries.
func ReadMap(data []byte, count uint16) (uint, MapValue) {
	entries := make(map[string]Value, count)
	var pos, n uint
	var h Header
	for i := uint(0); i < uint(count); i++ {
		// FIXME support termination-flag for continued values
		if n, h = ReadHeader(data[pos:]); n == 0 || h.Vtype != TYPE_KEYVALUE {
			// require key-entry to be plain data (byte-array)
			return 0, nil
		}
		if len(data[pos+n:]) < int(h.Size)+1 {
			return 0, nil
		}
		pos += n
		var v KeyValue
		// FIXME handle term==false for keys
		if n, v = ReadKeyValue(data[pos:], h.Size); n == 0 {
			return 0, nil
		}
		entries[v.K] = v.V
		pos += n
	}
	return pos, entries
}

// TODO currently borrows data from input-array when constructing types, i.e. no cloning.
// TODO future: add support for custom mapping of type-to-readFunction mapping for custom types
func ReadValue(data []byte) (uint, Value) {
	// FIXME support termination-flag for continued values
	var h Header
	var n uint
	if n, h = ReadHeader(data); n == 0 {
		return 0, nil
	}
	switch h.Vtype {
	case TYPE_BYTES:
		return ReadBytes(data[n:], &h)
	case TYPE_KEYVALUE:
		return ReadKeyValue(data[n:], h.Size)
	case TYPE_SEQUENCE:
		return ReadSequence(data[n:], h.Size)
	case TYPE_MAP:
		return ReadMap(data[n:], h.Size)
	default:
		panic("BUG: should not be reached")
	}
}
