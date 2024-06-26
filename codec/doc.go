// SPDX-License-Identifier: LGPL-3.0-only

// encoding contains all encodings. An encoding is defined as a transformation from (typically) any
// collection of types into a target type, often a primitive type. The most prevalent target type
// is (raw) bytes. However, big integer numbers may be encoded into uint limbs for purpose of
// efficient computation. Any target type gets its own package, such that encodings are easily
// located. 'encoding', as a package, is separated from std as it captures an interplay between two
// types, with varying source-types, and a predefined target-type.
// TODO consider adding Kim encoding (Douglas Crockford) for text encoding, as it is mere text representation.
// TODO consider adding Bitcoin's CompactSize unsigned integer encoding.
package codec
