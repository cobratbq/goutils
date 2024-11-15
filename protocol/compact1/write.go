package compact1

import (
	"io"

	"github.com/cobratbq/goutils/codec/bytes/bigendian"
)

// FIXME do error handling
func writeData(out io.Writer, data []byte, typeflags byte) (int64, error) {
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
		countNext, _ = writeData(out, data[SIZE_2BYTE_MAX:], typeflags)
		return int64(count) + countNext, nil
	}
}

// Type 0, 0 (singular, plain value) for any length.
type Bytes []byte

func (v Bytes) _sealed() {}

func (v Bytes) WriteTo(out io.Writer) (int64, error) {
	return writeData(out, v, 0)
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

func (v KeyValue) _sealed() {}

// FIXME add proper error handling
func (v KeyValue) WriteTo(out io.Writer) (int64, error) {
	writeData(out, []byte(v.K), FLAG_KEYVALUE)
	return v.V.WriteTo(out)
}

// Type 1, 0 (multiple, plain value) for any length.
type SequenceValue []Value

func (v SequenceValue) _sealed() {}

// FIXME add proper error handling
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

func (v MapValue) _sealed() {}

func (v MapValue) WriteTo(out io.Writer) (int64, error) {
	var count int64
	if len(v) <= int(SIZE_1BYTE_MAX) {
		out.Write([]byte{byte(len(v)) | FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE})
		for key, value := range v {
			writeData(out, []byte(key), FLAG_KEYVALUE)
			value.WriteTo(out)
		}
		return count, nil
	} else if len(v) <= int(SIZE_2BYTE_MAX) {
		header := bigendian.FromUint16(uint16(len(v)) - 1)
		header[0] |= FLAG_TERMINATION | FLAG_MULTIPLICITY | FLAG_KEYVALUE | FLAG_HEADERSIZE
		out.Write(header)
		for key, value := range v {
			writeData(out, []byte(key), FLAG_KEYVALUE)
			value.WriteTo(out)
		}
		return count, nil
	} else {
		// FIXME implement exceedingly large maps
		panic("To be implemented: exceedingly large maps")
	}
}
