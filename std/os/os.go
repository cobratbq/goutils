package os

import (
	"os"

	"github.com/cobratbq/goutils/std/errors"
)

// WorkingDir gets the working directory and panics on failure.
func WorkingDir() string {
	wd, err := os.Getwd()
	errors.RequireSuccess(err, "cannot acquire working directory: %+v")
	return wd
}
