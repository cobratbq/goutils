package http

import (
	"io"
	"net/http"
	"os"

	io_ "github.com/cobratbq/goutils/std/io"
)

func DownloadToFile(fileName, URL string) error {
	dstFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer io_.ClosePanicked(dstFile, "failed to close destination file: %+v")
	return DownloadFromURL(dstFile, URL)
}

func DownloadFromURL(dst io.Writer, url string) error {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return err
	}
	defer io_.ClosePanicked(resp.Body, "failed to close response body: %+v")
	_, err = io.Copy(dst, resp.Body)
	return err
}
