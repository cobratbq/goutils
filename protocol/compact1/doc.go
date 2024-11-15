// Specification: "compact1" (temporary name)
//
// A very basic protocol and encoding that allows for encoding values in a limited structure. The protocol
// essentially offers "byte-carriers", either in singular form or (multiple) in sequence. Values are either
// plain (a number of bytes) or key-value-pair.
//
// Note: although at first glance a key-value pair wouldn't make much sense, it does allow specifying
// predefined keys such that multiple versions of a protocol can identify which specific "labeled" value they
// receive as opposed to relying solely on (assumed) positional values.
//
// 4 flags are provided to indicate: termination, value-type, and multiplicity, header-size.
//
// The `termination`-bit is used to indicate whether this value is completed, meaning that if the bit is set,
// this entry concludes a value (possibly in a single entry), while an unset bit indicates that the following
// entry will continue the present value, effectively as concatenated bytes. The continuation must be of the
// same type.
//
// Interpretation of the values is left to be determined by the reader (application).
// TODO In-development, changes are likely...
// TODO document, make explicit that size-bits are encoded in little-endian, such that MSB from first byte are available to use as flags.
package compact1

import "io"

// FLAG_TERMINATIONBIT indicates whether this completes the current value. Usually set, unless large
// values/lists/maps are represented. If unset, indicates that the next entry is a continuation of this entry,
// effectively a concatenation of the next value-bytes onto the current value-bytes. The continuation-value
// must have the same bits for value-type and multiplicity as the initial value, otherwise the continuation is
// illegal, i.e. protocol violated.
//
// Regardless of whether the termination-bit is set, the remaining bits, that indicate the size, should be
// respected.
// TODO there was originally an idea to make a zero-size without termination-bit a special case for a list/map of unspecified length. However, this requires some form of termination to indicate the end, and that is currently not yet decided on.
const FLAG_TERMINATION uint8 = 1 << 7

// FLAG_VALUETYPE indicates whether this concerns just a value or a key-value-pair. In case of the key-value-
// pair, `size` indicates the size of the key.
// unset: plain value, set: key-value-pair.
// Note that we don't actually support type-information for values. We merely specify "containers"
// ("byte-carriers") of a certain size.
// The key-value pair is redundant in the sense that it could be expressed as 2 plain values in sequence,
// however that is based on convention. By allowing singular key-value-pair as type, we can express
// syntactically a key with corresponding value. This, in turn, can function as version-independent indicator
// for whatever encoded format is represented. Thus giving any expressed format _some_ syntactically enforced
// handles for recognizing/interpreting an encoded payload.
// TODO determine if 1-byte header indicates 0-15 byte keys or 1-16 byte keys
// FIXME consider renaming to FLAG_KEYPAIR
const FLAG_KEYVALUE uint8 = 1 << 6

// FLAG_MULTIPLICITY indicates whether this concerns a single value or multiple/series/collection of values.
// unset: single entry, set: multiple entries.
// In case of multiple entries, the size-bits indicate the number of entries.
const FLAG_MULTIPLICITY uint8 = 1 << 5

// FLAG_HEADERSIZE indicates the size of the header.
// unset: 1-byte header, set: 2-byte header.
const FLAG_HEADERSIZE uint8 = 1 << 4

// MASK_SIZEBITS is the mask that drops all the special flag-bits from the uint16-value.
// TODO for what types (value-type, multiplicity) should size be increment by 1 because size of 0 makes no sense?
const MASK_SIZEBITS uint8 = 0b00001111

// SIZE_1BYTE_MAX is the (inclusive) maximum for 1-byte headers.
// 4 bits available to indicate size, to indicate 0-15 bytes/count.
// TODO note that implementations should prefer 1-byte header for values < 16.
const SIZE_1BYTE_MAX uint = 15

// SIZE_2BYTE_MAX is the (inclusive) maximum for 2-byte headers.
// 12 bits available to indicate size, to indicate 1-4096 bytes/count.
//
// note: implementations must add `1`, as we can express 0 in 1-byte variant already. This allows us to
// express exactly up to size/count 4096, i.e. [1, 4096] or 4KiB.
const SIZE_2BYTE_MAX uint = 4096

// Value is the collective type for bytes, key-value-pairs, sequences and maps.
//
// note: `_sealed()` prevents alternative implementations not part of this package.
type Value interface {
	_sealed()
	io.WriterTo
}
