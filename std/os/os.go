// SPDX-License-Identifier: LGPL-3.0-or-later
package os

import (
	"os"

	"github.com/cobratbq/goutils/assert"
)

// WorkingDir gets the working directory and panics on failure.
func WorkingDir() string {
	wd, err := os.Getwd()
	assert.RequireSuccess(err, "cannot acquire working directory: %+v")
	return wd
}

// CreateEmptyFile creates an empty file at the specified path and returns an
// error if problems occur while creating this path. Existing file will be
// truncated.
func CreateEmptyFile(path string) error {
	f, err := os.Create(path)
	if err == nil {
		err = f.Close()
	}
	return err
}
