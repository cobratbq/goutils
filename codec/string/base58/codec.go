// SPDX-License-Identifier: LGPL-3.0-only

// base58 is a binary-to-text codec that encodes into a ASCII/ANSI-string (single bytes) and decodes back to
// the original data in binary.
package base58

import (
	"bytes"
	"crypto/sha256"
	"crypto/subtle"
	"math/big"
	"slices"

	"github.com/cobratbq/goutils/std/errors"
	"github.com/cobratbq/goutils/std/strconv"
)

var index [58]byte = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func factor() big.Int {
	var f big.Int
	f.SetUint64(58)
	return f
}

// Encode encodes the content into Base58.
func encode(data []byte) []byte {
	var x big.Int
	x.SetBytes(data)
	var zero big.Int
	factor := factor()
	var rem big.Int
	var result []byte
	for x.Cmp(&zero) != 0 {
		x.QuoRem(&x, &factor, &rem)
		result = append(result, index[rem.Uint64()])
	}
	for i := 0; i < len(data) && data[i] == 0; i++ {
		result = append(result, '1')
	}
	slices.Reverse(result)
	return result
}

// ChecksumEncode calculates the 4-byte checksum, concatenates the checksum, then encodes the content into
// Base58.
func ChecksumEncode(data []byte) []byte {
	check := sha256.Sum256(data)
	check = sha256.Sum256(check[:])
	content := make([]byte, 0, len(data)+4)
	content = append(content, data...)
	content = append(content, check[:4]...)
	return encode(content)
}

// CheckEncode checks the concatenated 4-byte checksum then encodes the data to Base58.
func CheckEncode(dataWithChecksum []byte) ([]byte, error) {
	check := sha256.Sum256(dataWithChecksum[:len(dataWithChecksum)-4])
	check = sha256.Sum256(check[:])
	if subtle.ConstantTimeCompare(check[:4], dataWithChecksum[len(dataWithChecksum)-4:]) != 1 {
		return nil, errors.Context(errors.ErrIllegal, "Base58 check-code does not match")
	}
	return encode(dataWithChecksum), nil
}

// Decode decodes the Base58-encoded content.
func Decode(encoded []byte) ([]byte, error) {
	factor := factor()
	var x big.Int
	var v big.Int
	var val int
	for i := 0; i < len(encoded); i++ {
		x.Mul(&x, &factor)
		val = bytes.IndexByte(index[:], encoded[i])
		if val == -1 {
			return nil, errors.Context(errors.ErrIllegal, "unexpected value in Base58-encoded content: "+strconv.FormatUint(encoded[i], 16))
		}
		v.SetInt64(int64(val))
		x.Add(&x, &v)
	}
	var result []byte
	for i := 0; i < len(encoded) && encoded[i] == '1'; i++ {
		result = append(result, 0)
	}
	result = append(result, x.Bytes()...)
	return result, nil
}

// Decodes the Base58-content then checks the checksum.
// Returns content without the checksum.
func CheckDecode(encoded []byte) ([]byte, error) {
	result, err := Decode(encoded)
	if err != nil {
		return nil, err
	}
	check := sha256.Sum256(result[:len(result)-4])
	check = sha256.Sum256(check[:])
	if subtle.ConstantTimeCompare(check[:4], result[len(result)-4:]) != 1 {
		return nil, errors.Context(errors.ErrIllegal, "Base58 check-code does not match")
	}
	return result[:len(result)-4], err
}
