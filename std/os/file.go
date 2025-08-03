package os

import (
	"os"

	"github.com/cobratbq/goutils/std/builtin"
)

// Exists checks whether a file-system object exists. (Symlinks will be followed to their target.)
func Exists(filepath string) bool {
	// TODO checking only for the error might be too superficial, possibly producing false positives or false negatives
	return builtin.Error(os.Stat(filepath)) == nil
}

// ExistsIsDirectory checks whether a file-system entry exists and is a directory.
func ExistsIsDirectory(filepath string) bool {
	if fi, err := os.Stat(filepath); err == nil {
		return fi.IsDir()
	} else {
		return false
	}
}

// ExistsFile checks whether a file-system object exists and it is a regular file. If the path references a
// symlink, it is followed to its target.
func ExistsFile(filepath string) bool {
	if fi, err := os.Stat(filepath); err == nil {
		return fi.Mode()&os.ModeType == 0
	} else {
		return false
	}
}

// ExistsIsSymlink checks whether a file-system object exists and it is itself a symbolic link. A symlink is
// not followed to its target.
func ExistsIsSymlink(filepath string) bool {
	// FIXME need to use Lstat here? (otherwise symlink is followed)
	if fi, err := os.Lstat(filepath); err == nil {
		return fi.Mode()&os.ModeSymlink != 0
	} else {
		return false
	}
}
