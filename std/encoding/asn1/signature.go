package asn1

import (
	"bytes"
	"math/big"

	"github.com/cobratbq/goutils/assert"
)

// FromP1363 converts a P1363-formatted signature into DER-format (ASN.1-encoded).
// `signature` is expected to consist of two equal-length components `R` and `s`, and exact slice-length is
// needed.
func FromP1363(signature []byte) []byte {
	assert.Equal(0, len(signature)%2)
	var compLength = len(signature) / 2
	r := new(big.Int).SetBytes(signature[:compLength]).Bytes()
	if r[0] >= 0x80 {
		r = append([]byte{0}, r...)
	}
	s := new(big.Int).SetBytes(signature[compLength:]).Bytes()
	if s[0] >= 0x80 {
		s = append([]byte{0}, s...)
	}
	var b bytes.Buffer
	b.WriteByte(0x30)
	b.WriteByte(uint8(2 + len(r) + 2 + len(s)))
	b.WriteByte(0x02)
	b.WriteByte(uint8(len(r)))
	b.Write(r)
	b.WriteByte(0x02)
	b.WriteByte(uint8(len(s)))
	b.Write(s)
	return bytes.Clone(b.Bytes())
}
