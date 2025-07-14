package os

import (
	"os"

	"github.com/cobratbq/goutils/std/builtin"
)

// Exists checks whether an object exists at filepath.
func Exists(filepath string) bool {
	// TODO checking only for the error might be too superficial, possibly producing false positives or false negatives
	return builtin.Error(os.Stat(filepath)) == nil
}
