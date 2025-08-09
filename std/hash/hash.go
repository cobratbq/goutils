// SPDX-License-Identifier: LGPL-3.0-only

package hash

import (
	"hash"
	"io"
	"os"

	"github.com/cobratbq/goutils/std/errors"
	io_ "github.com/cobratbq/goutils/std/io"
)

// HashFrom hashes the contents of the reader using the provided hash.
func HashFrom(hash hash.Hash, reader io.Reader) ([]byte, error) {
	if _, err := io.Copy(hash, reader); err != nil {
		return nil, errors.Context(err, "failure while hashing contents")
	}
	return hash.Sum(nil), nil
}

// HashFile hashes the contents of the file using the provided hash.
func HashFile(hash hash.Hash, filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errors.Context(err, "hash file '"+filepath+"'")
	}
	defer io_.CloseLogged(f, "Failed to gracefully close file")
	return HashFrom(hash, f)
}
