// SPDX-License-Identifier: LGPL-3.0-or-later

package bufio

import (
	"bufio"
	"io"
	"os"

	"github.com/cobratbq/goutils/std/errors"
	io_ "github.com/cobratbq/goutils/std/io"
)

func OpenFileReadOnly(path string) (*bufio.Reader, io.Closer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	return bufio.NewReader(file), io_.NewCloserWrapper(file), nil
}

// ReadBytesLinesNoDelimFunc
// Deprecated: use ReadBytesLinesFunc
func ReadBytesLinesNoDelimFunc(reader *bufio.Reader, delim byte, process func(line []byte) error) error {
	return ReadBytesLinesFunc(reader, delim, process)
}

// ReadBytesLinesFunc reads lines from the reader until `io.EOF` and calls `process` to
// process each line one-by-one, or until the first processing error. Anything except `io.EOF“ is
// treated as an error. In case `io.EOF` occurs, it is assumed that the acquired input is still
// valid/complete and is passed on for processing.
//
// Opinionated: in case of any error other than io.EOF, the last read input is lost and the error is
// returned. This is different from the basic `Read` functions, as these return whatever it still
// managed to read. Here we deviate because this util already sets assumptions on reading whole
// lines as part of its purpose.
//
// Returns nil if all lines are successfully processed and `io.EOF` was encountered. Returns
// `ErrProcessingFailure` (with context-information) if error was encountered during call to
// `process` closure. Returns IO-related errors for failures during reading.
//
// If `process` returns `ErrProcessingCompleted`, the read-loop will be interrupted to return early.
func ReadBytesLinesFunc(reader *bufio.Reader, delim byte, process func(line []byte) error) error {
	for {
		line, readErr := ReadBytesNoDelim(reader, delim)
		if readErr != nil && !errors.Is(readErr, io.EOF) {
			// Error occurred while reading line, so abort.
			return readErr
		}
		procErr := process(line)
		if procErr == ErrProcessingCompleted {
			// Allow `process` function to signal early exit.
			break
		}
		if procErr != nil {
			// Error occurred while processing line, so abort with processing failure.
			return errors.Context(ErrProcessingFailure, procErr.Error())
		}
		if errors.Is(readErr, io.EOF) {
			break
		}
	}
	return nil
}

// ReadStringLinesNoDelimFunc
// Deprecated: use ReadStringLinesFunc
func ReadStringLinesNoDelimFunc(reader *bufio.Reader, delim byte, process func(line string) error) error {
	return ReadStringLinesFunc(reader, delim, process)
}

// ReadStringLinesFunc reads lines from the reader until `io.EOF` and calls `process` to
// process each line one-by-one, or until the first processing error. Anything except `io.EOF“ is
// treated as an error. In case `io.EOF` occurs, it is assumed that the acquired input is still
// valid/complete and is passed on for processing.
//
// Opinionated: in case of any error other than io.EOF, the last read input is lost and the error is
// returned. This is different from the basic `Read` functions, as these return whatever it still
// managed to read. Here we deviate because this util already sets assumptions on reading whole
// lines as part of its purpose.
//
// Returns nil if all lines are successfully processed and `io.EOF` was encountered. Returns
// `ErrProcessingFailure` (with context-information) if error was encountered during call to
// `process` closure. Returns IO-related errors for failures during reading.
//
// If `process` returns `ErrProcessingCompleted`, the read-loop will be interrupted to return early.
func ReadStringLinesFunc(reader *bufio.Reader, delim byte, process func(line string) error) error {
	for {
		line, readErr := ReadStringNoDelim(reader, delim)
		if readErr != nil && !errors.Is(readErr, io.EOF) {
			// Error occurred while reading line, so abort.
			return readErr
		}
		procErr := process(line)
		if procErr == ErrProcessingCompleted {
			// Allow `process` function to signal early exit.
			break
		}
		if procErr != nil {
			// Error occurred while processing line, so abort with processing failure.
			return errors.Context(ErrProcessingFailure, procErr.Error())
		}
		if errors.Is(readErr, io.EOF) {
			break
		}
	}
	return nil
}

var ErrProcessingFailure = errors.NewStringError("failure encountered during processing")

var ErrProcessingCompleted = errors.NewStringError("processing completed")

func ReadBytesNoDelim(reader *bufio.Reader, delim byte) ([]byte, error) {
	buffer, err := reader.ReadBytes(delim)
	if err != nil {
		return buffer, err
	}
	return buffer[: len(buffer)-1 : len(buffer)-1], nil
}

func ReadStringNoDelim(reader *bufio.Reader, delim byte) (string, error) {
	buffer, err := reader.ReadString(delim)
	if err != nil {
		return buffer, err
	}
	return buffer[:len(buffer)-1], nil
}
