// SPDX-License-Identifier: LGPL-3.0-only

// ## Specification: "prefixed-compact"
//
// A very basic encoding that allows for encoding values in a limited structure. The protocol essentially
// offers "byte-carriers", either in singular form or (multiple) in sequence. Values are either plain (a
// number of bytes) or key-value-pair. Given its streamable nature, one may assume that order of occurrence is
// meaningful.
//
// Note: although at first glance a key-value-pair wouldn't make much sense, it does allow encoding predefined
// keys, such that multiple versions of a protocol can identify which specific "labeled" value they receive as
// opposed to relying solely on (assumed) positional values. Furthermore, unknown labels may be ignored or
// alternatively may indicate unsupported format or unknown elements.
//
// 4 Flags are provided to indicate:
//
//   - termination: whether the value terminates now, or a follow-up record is expected carrying the
//     remainder,
//   - value-type: a plain value or key-value-pair ("labeled" value),
//   - multiplicity: a single value or multiple (collection of) values,
//   - header-size: a 1-byte header, indicating sizes `[0, 15]`, or 2-byte header, indicating sizes
//     `[1, 4096]`, with sizes encoded as big-endian unsigned integer, with first 4 bits (of first byte)
//     (implicitly, due to restricted ranges for size) reserved for flags.
//
// The `termination`-bit is used to indicate whether this value is completed, meaning that if the bit is set,
// this entry concludes a value (possibly in a single entry), while an unset bit indicates that the following
// entry will continue the present value, effectively as concatenated bytes.
// Note: the non-terminated value _must_ be continued with the same type.
//
// This encoding does not provide any redundancy, error-correction or "value-type finished"-indicators.
// Consequently, corruption of the header-bytes will result in misinterpretation of subsequent data. (This
// must be solved outside of the encoding, if there is risk of corruption.)
//
// Formal notation:
//
//	value = data | (key , value) | sequence | map
//	data = header , {byte}
//	key = header , {byte}
//	sequence = header , {value}
//	map = header , {key , value}
//	header = termination-bit , valuetype-bit , multiplicity-bit , headersize-bit , 4-bit-size , [8-bit-size]
//
//	Additionally:
//	1.) header will include the second byte, as part of big-endian unsigned size (12 bits in total) if-and-
//	only-if the `headersize`-bit is set. The 4-bit size ranges `[0,15]`. The 12-bit size ranges `[1,4096]`,
//	thus requiring an offset correction.
//	2.) the minimum necessary header-size should be used, i.e. sizes/counts ranging `[0,15]` expressed with
//	single-byte header.
//	3.) if termination-bit is unset, next value must be of same `valuetype | multiplicity`.
//	4.) length of raw data (`{byte}`) in `bytes` and `key` must correspond to length indicated in header;
//	actual number of entries in `sequence` and `map` must correspond to count indicated in header.
//
// ## Interpretation of values (bytes) is left to be determined by the reader (application).
//
// This includes the exact meaning of the data-types. For example, duplicate keys in a map may indicate an
// error, or a replacement value, or an addition of a second value to the key, or an concatenation onto the
// first value, or ...
//
// There is inherent order through the position in the data-stream, consequently a sequence (which itself
// defines the size, i.e. number of elements) has an order, which may be ignored (collection), or could
// indicate order (list), or priority (queue), or a stack, or age i.e. history of prior values, ...
//
// The termination-bit which is used to indicate the end of a value, could be used to partition data in
// smaller chunks, either to benefit limited processing capabilities of embedded devices, or to transfer data
// as it is incoming at irregular intervals, or to signal an artificial boundary as a new batch refreshes/
// updates the data from the previous batch, or for streams of data for which the number of entries is not yet
// known. (One could signal the end of the value-type with a 0-size terminated entry, if needed.)
//
// In terms of interpreting values into data-types: this encoding does not provide any support towards that
// goal. 8-byte values _could_ indicate a big-endian/little-endian signed/unsigned 64-bit integer value. This
// would need to be aligned between communicating parties, either by convention or by agreement.
// "Labeled" values could be used to indicate a format. However, e.g. inter-process communication may not be
// concerned with these kinds of issues if it is all built from the same basic functions/libraries.
//
// There is no literal to represent `null`. In many cases, 0-length entry or not (having to) encode a
// particular key-value entry perfectly reflects intention. Given that this is an encoding of data, rather
// than a full-blown data-structure with matching data-types, this seems to suffice. Furthermore, it unburdens
// developers from having to include logic/restrictions for properly processing data in both the code-base
// and encoded format. (In virtually all cases, data-types, structures and classes already require such logic
// in order to preserve their invariants during use.)
// TODO consider adding PrefixedEncode and PrefixedDecode interfaces for encoding arbitrary types depending on user implementations. (Possibly with some verification during encoding.)
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
// unset: plain value, set: key-value-pair, i.e. a "labeled" value.
// Note that we don't actually support type-information for values. We merely specify "containers"
// ("byte-carriers") of a certain size.
//
// Note: the key-value pair is redundant in the sense that it could be expressed as 2 plain values in
// sequence, however that is based in convention. By allowing singular key-value-pair as type, we can express
// syntactically a key with its corresponding value. This, in turn, can function as version-independent
// indicator for whatever encoded format is represented. Thus giving any expressed format _some_ syntactically
// enforced handles.
const FLAG_KEYVALUE uint8 = 1 << 6

// FLAG_MULTIPLICITY indicates whether this concerns a single value or multiple/series/collection of values.
// unset: single entry, set: multiple entries.
// In case of multiple entries, the size-bits indicate the number of entries.
const FLAG_MULTIPLICITY uint8 = 1 << 5

// FLAG_HEADERSIZE indicates the size of the header.
// unset: 1-byte header, set: 2-byte header.
//
// Consequently, 4 bits to represent size [0, 15], or 12-bits to represent size [1, 4096]. If the flag is set,
// upon interpreting the 12-bit size, SIZE_2BYTE_OFFSET should be added to correct for difference in range
// start.
//
// The 2-byte header, i.e. this flag, should only be used for values outside of the 1-byte range. It is under
// consideration to use the lowest redundant value(s) to signal for a future version. Current implementations
// would do good to reject these values `[1,15]` in 2-byte header size-range as illegal, as such rejecting a
// possible unknown and unsupported encoding.
const FLAG_HEADERSIZE uint8 = 1 << 4

// MASK_SIZEBITS is the mask that drops all the special flag-bits from the first byte of the big-endian
// uint16-value. (The second byte fully represents its part of the size.)
const MASK_SIZEBITS uint8 = 0b00001111

// SIZE_1BYTE_MAX is the (inclusive) maximum for 1-byte headers.
// 4 bits available to indicate size, to indicate 0-15 bytes/count.
const SIZE_1BYTE_MAX uint = 15

// SIZE_2BYTE_MAX is the (inclusive) maximum for 2-byte headers.
// 12 bits available to indicate size, to indicate 1-4096 bytes/count.
//
// note: implementations must correct for `±1`, as we can express 0 in 1-byte variant already. This allows us
// to express exactly up to size/count 4096, i.e. [1, 4096] or 4KiB, in 12 bits.
const SIZE_2BYTE_MAX uint = 4096

// SIZE_2BYTE_OFFSET represents the correction performed while the 2-byte size-value is stored.
const SIZE_2BYTE_OFFSET uint16 = 1

// Value is the collective type for bytes, key-value-pairs, sequences and maps.
type Value interface {
	io.WriterTo
	// Equal tests whether provided value is of same type and content as this value.
	Equal(other Value) bool
	// Len returns the count of bytes, sequence-entries or map-entries present in Value.
	Len() int
}
