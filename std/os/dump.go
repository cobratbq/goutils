// SPDX-License-Identifier: LGPL-3.0-only

package os

import (
	"io"
	"os"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/errors"
	io_ "github.com/cobratbq/goutils/std/io"
)

// DumpToFile writes provided content to file at `path`. This is a one-shot function that writes the content
// and closes the file. This is particularly useful during development and for debugging.
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

// MustDumpToFile executes DumpToFile and panics on failure.
func MustDumpToFile(path string, content []byte) {
	assert.Success(DumpToFile(path, content), "Failed to dump content to file.")
}

// DumpToFileThroughFunc1 dumps contents to file that are produced by a function that performs something like
// a serialization and accepts a writer as input.
func DumpToFileThroughFunc1(path string, writeIntoFunc func(out io.Writer) error) error {
	var out, err = os.Create(path)
	if err != nil {
		return errors.Context(err, "failed to create/truncate out-file for writing")
	}
	defer io_.CloseLogged(out, "failed to gracefully close out-file")
	return writeIntoFunc(out)
}

// DumpToFileThroughFunc2 dumps contents to file that are produced by a function that performs something like
// a serialization and accepts a writer as input.
func DumpToFileThroughFunc2(path string, writeIntoFunc func(out io.Writer) (int, error)) (int, error) {
	var out, err = os.Create(path)
	if err != nil {
		return 0, errors.Context(err, "failed to create/truncate out-file for writing")
	}
	defer io_.CloseLogged(out, "failed to gracefully close out-file")
	return writeIntoFunc(out)
}

// DumpReaderToFile dumps the reader to a file at `path`. This is a one-shot function that writes the content
// and closes the file. This is particularly useful during development and for debugging.
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

// MustDumpReaderToFile executes DumpReaderToFile and panics on failure.
func MustDumpReaderToFile(path string, in io.Reader) {
	assert.Success(DumpReaderToFile(path, in), "Failed to stream `in` to out-file.")
}
