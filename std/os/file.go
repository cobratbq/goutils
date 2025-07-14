package os

import (
	"os"

	"github.com/cobratbq/goutils/std/builtin"
)

// Exists checks whether a file-system object exists.
func Exists(filepath string) bool {
	// TODO checking only for the error might be too superficial, possibly producing false positives or false negatives
	return builtin.Error(os.Stat(filepath)) == nil
}

// ExistsFile checks whether a file-system object exists and it is a regular file.
func ExistsFile(filepath string) bool {
	if fi, err := os.Stat(filepath); err == nil {
		return fi.Mode()&os.ModeType == 0
	} else {
		return false
	}
}

// ExistsSymlink checks whether a file-system object exists and it is a symbolic link.
func ExistsSymlink(filepath string) bool {
	if fi, err := os.Stat(filepath); err == nil {
		return fi.Mode()&os.ModeSymlink != 0
	} else {
		return false
	}
}
