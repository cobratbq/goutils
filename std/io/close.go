// SPDX-License-Identifier: LGPL-3.0-or-later
package io

import (
	"io"
	"log"

	"github.com/cobratbq/goutils/assert"
)

// CloseDiscarded closes the closer and discards any possible error.
func CloseIgnored(c io.Closer) {
	c.Close()
}

// ClosePanicked closes the closer and panics with specified message in case of
// any error, except for io.ErrClosedPipe.
func ClosePanicked(c io.Closer, message string) {
	err := c.Close()
	if err == io.ErrClosedPipe {
		return
	}
	assert.Success(err, message)
}

// CloseLogged closes the closer and logs specified message in case of error.
// Any error except for io.ErrClosedPipe is logged.
func CloseLogged(c io.Closer, message string) {
	if err := c.Close(); err != nil && err != io.ErrClosedPipe {
		log.Printf(message, err)
	}
}

// NopCloser is a no-op close ReadWriteCloser implementation.
//
// One would typically want to respect the Close method on writers. This
// implementation provides no-op closing for cases where `Close()` is already
// handled at a different level of nesting.
type NopCloser struct {
	Rw io.ReadWriter
}

func (n *NopCloser) Read(p []byte) (int, error) {
	return n.Rw.Read(p)
}

func (n *NopCloser) Write(p []byte) (int, error) {
	return n.Rw.Write(p)
}

// Close is a no-op.
func (n *NopCloser) Close() error {
	return nil
}
