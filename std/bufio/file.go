// SPDX-License-Identifier: LGPL-3.0-only

package bufio

import (
	"bufio"
	"io"
	"os"

	"github.com/cobratbq/goutils/std/errors"
	io_ "github.com/cobratbq/goutils/std/io"
)

// OpenFileProcessStringLinesFunc opens the file - specified by name - and reads lines denoted by
// `delim` in order to process their contents as strings passing them on to `process`. After
// processing has terminated (contents exhausted or completion signaled, see
// `ReadProcessStringLines`), the file is closed and results returned.
//
// Errors:
//   - returns no result and error, in case of error while opening the input file.
//   - returns partial results (successfully processed lines) and error, in case of error produced
//     inside the `process` func while processing a line. (See `ReadProcessStringLines`)
func OpenFileProcessStringLinesFunc[V any](filename string, delim byte, process func(string) (V, error)) ([]V, error) {
	reader, closer, inputErr := OpenFileReadOnly(filename)
	if inputErr != nil {
		return nil, inputErr
	}
	// NOTE: assuming for now that nothing significant can go wrong *if* a failure during closing
	// even happens. Logging for transparency but should be fine.
	defer io_.CloseLogged(closer, "Failed to gracefully close the input file")
	return readProcessTypedLinesFunc(ReadStringNoDelim, reader, delim, process)
}

// OpenFileProcessBytesLinesFunc opens the file - specified by name - and reads lines denoted by
// `delim` in order to process their contents as strings passing them on to `process`. After
// processing has terminated (contents exhausted or completion signaled, see
// `ReadProcessStringLines`), the file is closed and results returned.
//
// Errors:
//   - returns no result and error, in case of error while opening the input file.
//   - returns partial results (successfully processed lines) and error, in case of error produced
//     inside the `process` func while processing a line. (See `ReadProcessStringLines`)
func OpenFileProcessBytesLinesFunc[V any](filename string, delim byte, process func([]byte) (V, error)) ([]V, error) {
	reader, closer, inputErr := OpenFileReadOnly(filename)
	if inputErr != nil {
		return nil, inputErr
	}
	// NOTE: assuming for now that nothing significant can go wrong *if* a failure during closing
	// even happens. Logging for transparency but should be fine.
	defer io_.CloseLogged(closer, "Failed to gracefully close the input file")
	return readProcessTypedLinesFunc(ReadBytesNoDelim, reader, delim, process)
}

// - ErrProcessingIgnore to skip irrelevant value, ErrProcessingCompleted to signal to stop reading, ...
// FIXME document function (and reference other options for other use cases)
func ReadProcessStringLinesFunc[V any](reader *bufio.Reader, delim byte, process func(string) (V, error)) ([]V, error) {
	return readProcessTypedLinesFunc(ReadStringNoDelim, reader, delim, process)
}

// FIXME document function (and reference other options for other use cases)
func ReadProcessBytesLinesFunc[V any](reader *bufio.Reader, delim byte, process func([]byte) (V, error)) ([]V, error) {
	return readProcessTypedLinesFunc(ReadBytesNoDelim, reader, delim, process)
}

func readProcessTypedLinesFunc[T ~[]byte | ~string, V any](
	read func(*bufio.Reader, byte) (T, error), reader *bufio.Reader, delim byte,
	process func(T) (V, error)) ([]V, error) {

	var results []V
	for {
		line, readErr := read(reader, delim)
		if readErr != nil && !errors.Is(readErr, io.EOF) {
			// Error occurred while reading line, so abort. Return results that are available, let
			// user judge whether those are useful.
			return results, readErr
		}
		v, procErr := process(line)
		if procErr == ErrProcessingIgnore {
			// Error indicates resulting value should be ignored.
		} else if procErr == ErrProcessingCompleted {
			// Allow `process` function to signal early exit.
			results = append(results, v)
			break
		} else if procErr != nil {
			// Error occurred while processing line, so abort with processing failure.
			return results, errors.Context(ErrProcessingFailure, procErr.Error())
		} else {
			results = append(results, v)
		}
		if errors.Is(readErr, io.EOF) {
			break
		}
	}
	return results, nil
}

var ErrProcessingIgnore = errors.NewStringError("processing resulted in irrelevant result, ignore")

// ReadBytesLinesNoDelimFunc
// Deprecated: use ReadBytesLinesFunc
func ReadBytesLinesNoDelimFunc(reader *bufio.Reader, delim byte, process func([]byte) error) error {
	return ReadBytesLinesFunc(reader, delim, process)
}

// ReadBytesLinesFunc reads lines from the reader until `io.EOF` and calls `process` to
// process each line one-by-one, or until the first processing error. Anything except `io.EOF“ is
// treated as an error. In case `io.EOF` occurs, it is assumed that the acquired input is still
// valid/complete and is passed on for processing.
//
// Opinionated:
//   - any data read prior to `io.EOF` is considered a full line, therefore processed as normal.
//   - in case of any error other than `io.EOF`, the last read data is lost and the error is
//     returned immediately. This is different from the basic `Read` functions, as these return
//     whatever it still managed to read. Here we deviate because this util already sets assumptions
//     on reading whole lines as part of its purpose.
//
// Returns nil if all lines are successfully processed and `io.EOF` was reached or
// `ErrProcessingCompleted` was returned by `process` function. Returns `ErrProcessingFailure` (with
// context-information) if error was encountered during call to `process` closure. Returns
// IO-related errors for failures during reading.
func ReadBytesLinesFunc(reader *bufio.Reader, delim byte, process func([]byte) error) error {
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
func ReadStringLinesNoDelimFunc(reader *bufio.Reader, delim byte, process func(string) error) error {
	return ReadStringLinesFunc(reader, delim, process)
}

// ReadStringLinesFunc reads lines from the reader until `io.EOF` and calls `process` to
// process each line one-by-one, or until the first processing error. Anything except `io.EOF“ is
// treated as an error. In case `io.EOF` occurs, it is assumed that the acquired input is still
// valid/complete and is passed on for processing.
//
// Opinionated:
//   - any data read prior to `io.EOF` is considered a full line, therefore processed as normal.
//   - in case of any error other than `io.EOF`, the last read data is lost and the error is
//     returned immediately. This is different from the basic `Read` functions, as these return
//     whatever it still managed to read. Here we deviate because this util already sets assumptions
//     on reading whole lines as part of its purpose.
//
// Returns nil if all lines are successfully processed and `io.EOF` was reached or
// `ErrProcessingCompleted` was returned by `process` function. Returns `ErrProcessingFailure` (with
// context-information) if error was encountered during call to `process` closure. Returns
// IO-related errors for failures during reading.
func ReadStringLinesFunc(reader *bufio.Reader, delim byte, process func(string) error) error {
	for {
		line, readErr := ReadStringNoDelim(reader, delim)
		if readErr != nil && !errors.Is(readErr, io.EOF) {
			// Error occurred while reading line, so abort.
			return readErr
		}
		processErr := process(line)
		if processErr == ErrProcessingCompleted {
			// Allow `process` function to signal early exit.
			break
		}
		if processErr != nil {
			// Error occurred while processing line, so abort with processing failure.
			return errors.Context(ErrProcessingFailure, processErr.Error())
		}
		if errors.Is(readErr, io.EOF) {
			break
		}
	}
	return nil
}

// ErrProcessingFailure will be returned if the `process` func produces an error. This is
// interpreted as a failure that is critical enough to stop further processing. Instead,
// ErrProcessingCompleted will be returned with the processing error as context information.
var ErrProcessingFailure = errors.NewStringError("failure encountered during processing")

// ErrProcessingCompleted signals that all expected processing has been completed and we do not want
// to continue to process any possible lines to follow. Instead, when this is received, we shall
// break out of the processing loop.
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

func OpenFileReadOnly(path string) (*bufio.Reader, io.Closer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	return bufio.NewReader(file), io_.NewCloserWrapper(file), nil
}
