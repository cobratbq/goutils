package http

import (
	"fmt"
	"io"
	"net/http"
	"os"

	io_ "github.com/cobratbq/goutils/std/io"
)

// DownloadToFile downloads content from the given URL into the specified file
// name.
func DownloadToFile(fileName, URL string) error {
	dstFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer io_.ClosePanicked(dstFile, "failed to close destination file: %+v")
	return DownloadFromURL(dstFile, URL)
}

// DownloadFromURL downloads content from the specified URL and immediately
// writes it to the destination. In case an unexpected http status code is
// received, an error is returned containing that status code.
func DownloadFromURL(dst io.Writer, url string) error {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return err
	}
	defer io_.ClosePanicked(resp.Body, "failed to close response body: %+v")
	_, err = io.Copy(dst, resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return ErrStatusCode(resp.StatusCode)
	}
	return nil
}

// ErrStatusCode indicates that a status code other than the expected status
// code is received.
type ErrStatusCode int

func (code ErrStatusCode) Error() string {
	return fmt.Sprintf("http status code: %d", code)
}
