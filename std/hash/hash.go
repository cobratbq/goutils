package hash

import (
	"hash"
	"io"
	"os"

	"github.com/cobratbq/goutils/std/errors"
	io_ "github.com/cobratbq/goutils/std/io"
)

func HashFile(h hash.Hash, filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Context(err, "hash file '"+filepath+"'")
	}
	defer io_.CloseLogged(f, "Failed to gracefully close file")
	if _, err := io.Copy(h, f); err != nil {
		return nil, errors.Context(err, "failure while hashing contents")
	}
	return h.Sum(nil), nil
}
