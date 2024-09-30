package os

import (
	"io"
	"os"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/errors"
	io_ "github.com/cobratbq/goutils/std/io"
)

func DumpToFile(path string, content []byte) error {
	var out, err = os.Create(path)
	if err != nil {
		return errors.Context(err, "failed to create/truncate file for writing")
	}
	defer io_.CloseLogged(out, "failed to gracefully close out-file")
	var n int
	n, err = out.Write(content)
	if err != nil {
		return errors.Context(err, "failed to write content to out-file")
	}
	if n < len(content) {
		return errors.Context(io.ErrShortWrite, "failed to write full content to file")
	}
	return nil
}

func MustDumpToFile(path string, content []byte) {
	assert.Success(DumpToFile(path, content), "Failed to dump content to file.")
}

func DumpReaderToFile(path string, in io.Reader) error {
	var out, err = os.Create(path)
	if err != nil {
		return errors.Context(err, "failed to create/truncate out-file for writing")
	}
	defer io_.CloseLogged(out, "failed to gracefully close out-file")
	if _, err = out.ReadFrom(in); err != nil {
		return errors.Context(err, "failed to stream `in` to out-file")
	}
	return nil
}

func MustReaderToFile(path string, in io.Reader) {
	assert.Success(DumpReaderToFile(path, in), "Failed to stream `in` to out-file.")
}
