package os

import (
	"os"

	"github.com/cobratbq/goutils/std/builtin"
)

// WorkingDir gets the working directory and panics on failure.
func WorkingDir() string {
	wd, err := os.Getwd()
	builtin.RequireSuccess(err, "cannot acquire working directory: %+v")
	return wd
}
