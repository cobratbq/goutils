package bufio

import (
	"bufio"
	"io"
	"os"

	io_ "github.com/cobratbq/goutils/std/io"
)

func OpenFileReadOnly(path string) (*bufio.Reader, io.Closer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	return bufio.NewReader(file), io_.NewCloserWrapper(file), nil
}

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
