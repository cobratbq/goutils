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
