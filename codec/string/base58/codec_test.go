package base58

import (
	"encoding/hex"
	"testing"

	"github.com/cobratbq/goutils/std/builtin"
	assert "github.com/cobratbq/goutils/std/testing"
)

func TestEncodeBase58Check(t *testing.T) {
	testdata := []struct {
		data    []byte
		encoded []byte
	}{
		{data: builtin.Expect(hex.DecodeString("00f54a5851e9372b87810a8e60cdd2e7cfd80b6e31c7f18fe8")), encoded: []byte("1PMycacnJaSqwwJqjawXBErnLsZ7RkXUAs")},
		{data: builtin.Expect(hex.DecodeString("800C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D507A5B8D")), encoded: []byte("5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")},
	}
	for _, entry := range testdata {
		result, err := CheckEncode(entry.data)
		assert.Nil(t, err)
		assert.EqualSlices(t, entry.encoded, result)
	}
}

func TestDecodeBase58Check(t *testing.T) {
	testdata := []struct {
		data    []byte
		decoded []byte
	}{
		{decoded: builtin.Expect(hex.DecodeString("00f54a5851e9372b87810a8e60cdd2e7cfd80b6e31c7f18fe8")), data: []byte("1PMycacnJaSqwwJqjawXBErnLsZ7RkXUAs")},
		{decoded: builtin.Expect(hex.DecodeString("800C28FCA386C7A227600B2FE50B7CAE11EC86D3BF1FBE471BE89827E19D72AA1D507A5B8D")), data: []byte("5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ")},
	}
	for _, entry := range testdata {
		result, err := CheckDecode(entry.data)
		assert.Nil(t, err)
		assert.EqualSlices(t, entry.decoded, result)
	}
}
