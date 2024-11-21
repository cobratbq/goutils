// Specification: "prefixed-compact" (temporary name)
//
// A very basic protocol and encoding that allows for encoding values in a limited structure. The protocol
// essentially offers "byte-carriers", either in singular form or (multiple) in sequence. Values are either
// plain (a number of bytes) or key-value-pair.
// Given its streamable nature, one may assume that order of occurrence is meaningful, e.g. earlier occurrence
// indicating higher priority or later occurrence superseding earlier occurrence, as needed.
//
// Note: although at first glance a key-value-pair wouldn't make much sense, it does allow specifying
// predefined keys, such that multiple versions of a protocol can identify which specific "labeled" value they
// receive as opposed to relying solely on (assumed) positional values. Furthermore, unknown labels may be
// ignored or alternative may indicate unsupported or unknown elements.
//
// 4 Flags are provided to indicate:
//
// - termination: whether the value terminates now, or a follow-up record is expected carrying the remainder,
// - value-type: a plain value or key-value-pair ("labeled" value),
// - multiplicity: a single value or multiple (sequence of) values,
// - header-size: a 1-byte header, indicating sizes `[0, 15]`, or 2-byte header, indicating sizes `[1, 4096]`,
//
// The `termination`-bit is used to indicate whether this value is completed, meaning that if the bit is set,
// this entry concludes a value (possibly in a single entry), while an unset bit indicates that the following
// entry will continue the present value, effectively as concatenated bytes.
// Note: the continuation must be of the same type.
//
// Interpretation of values (bytes) is left to be determined by the reader (application).
// TODO decide on a name once proven to be effective.
// TODO In-development, changes are likely...
// TODO document, make explicit that size-bits are encoded in big-endian, such that MSB from first byte are available to use as flags.
// TODO consider making shortest possible header mandatory, i.e. any size/count of <= 15, must use 1-byte header, that way the initial overlap of 2-byte header could be used to signal other characteristics in the future(?)
package prefixed

import "io"

// FLAG_TERMINATIONBIT indicates whether this completes the current value. Usually set, unless large
// values/keys/lists/maps are represented. If unset, indicates that the next entry is a continuation of this
// entry, effectively a concatenation of the next value-bytes onto the current value-bytes. The
// continuation-value must have the same bits for value-type and multiplicity as the initial value, otherwise
// the continuation is illegal.
//
// The termination-bit may be set or unset for any size/count. This makes it possible to reconstruct values
// from parts of varying lengths, possibly benefiting small devices or variable/unpredictable input-streams.
// TODO there was originally an idea to make a zero-size without termination-bit a special case for a list/map of unspecified length. However, this requires some form of termination to indicate the end, and that is currently not yet decided on.
// TODO document: if not set, must be followed by entry of same type-flags optionally FLAG_TERMINATION set, of any (allowed) size.
const FLAG_TERMINATION uint8 = 1 << 7

// FLAG_KEYVALUE indicates whether this concerns just a value or a key-value-pair. In case of the key-value-
// pair, `size` indicates the size of the key.
// unset: plain value, set: key-value-pair.
// Note that we don't actually support type-information for values. We merely specify "containers"
// ("byte-carriers") of a certain size.
// The key-value pair is redundant in the sense that it could be expressed as 2 plain values in sequence,
// however that is based on convention. By allowing singular key-value-pair as type, we can express
// syntactically a key with its corresponding value. This, in turn, can function as version-independent
// indicator for whatever encoded format is represented. Thus giving any expressed format _some_ syntactically
// enforced handles for recognizing/interpreting an encoded payload.
// TODO document: FLAG_KEYVALUE indicates the key, must be followed by a value of any type, i.e. the key cannot be the last entry.
const FLAG_KEYVALUE uint8 = 1 << 6

// FLAG_MULTIPLICITY indicates whether this concerns a single value or multiple/series/collection of values.
// unset: single entry, set: multiple entries.
// In case of multiple entries, the size-bits indicate the number of entries.
// TODO document: i.e. size-bits indicate count, instead of single length for value/key
const FLAG_MULTIPLICITY uint8 = 1 << 5

// FLAG_HEADERSIZE indicates the size of the header.
// unset: 1-byte header, set: 2-byte header.
// Consequently, 4 bits to represent size [0, 15], or 12-bits to represent size [1, 4096].
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
// note: implementations must correct for `±1`, as we can express 0 in 1-byte variant already. This allows us
// to express exactly up to size/count 4096, i.e. [1, 4096] or 4KiB, in 12 bits.
const SIZE_2BYTE_MAX uint = 4096

// SIZE_2BYTE_OFFSET represents the correction performed while the 2-byte size-value is stored.
// TODO consider if we want to shift values by +16, such that first 2-byte-header value is 16, and last is 4095+16 (Does not seem to touch on any significant benefits, unless we consider a few increments past 4096 to be an advantage)
const SIZE_2BYTE_OFFSET uint16 = 1

// Value is the collective type for bytes, key-value-pairs, sequences and maps.
//
// note: `_sealed()` prevents alternative implementations not part of this package.
type Value interface {
	// TODO does `_sealed` properly prevent third-party implementations?
	io.WriterTo
}
