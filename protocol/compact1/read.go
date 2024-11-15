package compact1

type CompositeType uint8

const (
	// value, singular = value, i.e. bytes
	TYPE_BYTES CompositeType = iota
	// key-value, singular = key-value-pair
	TYPE_KEYVALUE
	// value, multiple = sequence of values (ordered by virtue of position in data)
	TYPE_SEQUENCE
	// key-value, multiple = map of key-value-pairs (ordered by virtue of position in data)
	TYPE_MAP
)

// FIXME redo Read... implementations now that encoding/writing has matured

// FIXME proper error handling for n == 0, array-size-bounds
func ReadHeader(data []byte) (uint, CompositeType, uint16, bool) {
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
		return 1, vtype, size, term
	}
	size <<= 8
	size |= uint16(data[1])
	// Add `1` to size for 2-byte header, as we can already express 0, and adding 1 allows us to express
	// sizes/counts up to 2**12 == 4096.
	// TODO check if works as expected
	size += 1
	return 2, vtype, size, term
}

// FIXME support non-terminated key-entry
func ReadKeyValue(data []byte, keysize uint16) (uint, KeyValue) {
	// FIXME edge-case with key being non-continuous?
	// FIXME proper error handling for n == 0, array-size-bounds
	key := string(data[:keysize])
	n, val := ReadValue(data[keysize:])
	return uint(keysize) + n, KeyValue{K: key, V: val}
}

// FIXME proper error handling for n == 0, array-size-bounds
func ReadSequence(data []byte, count uint16) (uint, SequenceValue) {
	var entries = make([]Value, count)
	var pos, n uint
	for i := uint(0); i < uint(count); i++ {
		n, entries[i] = ReadValue(data[pos:])
		if n == 0 {
			// FIXME handle errors
			return 0, nil
		}
		pos += n
	}
	return pos, entries
}

// FIXME proper error handling for n == 0, array-size-bounds
func ReadMap(data []byte, count uint16) (uint, MapValue) {
	// TODO map assumes distinct keys, hence count is exact number of map entries.
	entries := make(map[string]Value, count)
	var pos, n uint
	var vtype CompositeType
	var size uint16
	for i := uint(0); i < uint(count); i++ {
		// FIXME support termination-flag for continued values
		n, vtype, size, _ = ReadHeader(data[pos:])
		if n == 0 || vtype > 0 {
			// require key-entry to be plain data (byte-array)
			return 0, nil
		}
		pos += n
		var key []byte
		var val Value
		// FIXME handle term==false for keys
		n, val := ReadKeyValue(data[pos:], size)
		if n == 0 {
			return 0, nil
		}
		entries[string(key)] = val
		pos += n
	}
	return pos, entries
}

func ReadValue(data []byte) (uint, Value) {
	// FIXME support termination-flag for continued values
	n, vtype, size, _ := ReadHeader(data)
	switch vtype {
	case TYPE_BYTES:
		return n + uint(size), Bytes(data[n : n+uint(size)])
	case TYPE_KEYVALUE:
		return ReadKeyValue(data[n:], size)
	case TYPE_SEQUENCE:
		return ReadSequence(data[n:], size)
	case TYPE_MAP:
		return ReadMap(data[n:], size)
	default:
		panic("BUG: should not be reached")
	}
}
