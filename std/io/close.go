// SPDX-License-Identifier: LGPL-3.0-only

package io

import (
	"io"

	"github.com/cobratbq/goutils/assert"
	"github.com/cobratbq/goutils/std/errors"
	"github.com/cobratbq/goutils/std/log"
)

// CloseIgnored closes the closer and discards any possible error.
func CloseIgnored(c io.Closer) {
	err := c.Close()
	log.TracelnDepth(1, "`CloseIgnored` ignores error:", err)
}

// ClosePanicked closes the closer and panics with specified message in case of any error, except
// for io.ErrClosedPipe.
func ClosePanicked(c io.Closer, message string) {
	err := c.Close()
	if errors.Is(err, io.ErrClosedPipe) {
		return
	}
	assert.Success(err, message)
}

// CloseLogged closes the closer and logs specified message in case of error. Any error except for
// io.ErrClosedPipe is logged. The error message is logged as a warning.
// TODO consider logging without formatter
func CloseLogged(c io.Closer, message string) {
	if err := c.Close(); err != nil && !errors.Is(err, io.ErrClosedPipe) {
		log.Warnf(message, err.Error())
	}
}

// NopCloser is a no-op close ReadWriteCloser implementation.
//
// One would typically want to respect the Close method on writers. This implementation provides
// no-op closing for cases where `Close()` is already handled at a different level of nesting.
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
	log.TracelnDepth(1, "`NopCloser` not closing wrapped ReadWriter.")
	return nil
}

type closerWrapper struct {
	closer io.Closer
}

func NewCloserWrapper(closer io.Closer) *closerWrapper {
	return &closerWrapper{closer}
}

func (c *closerWrapper) Close() error {
	return c.closer.Close()
}

type closeSequence struct {
	seq []io.Closer
}

// NewCloseSequence creates a new composite sequential closer that closes in the order provided.
//
// The sequence will not halt on error. If closing behavior is dependent on other closers, this
// should be part of the closer's logic. Instead, errors are collected and an aggregate error is
// returned that includes error messages from all the failures that occurred while closing the
// sequence.
//
// Returns `ErrSequenceFailure` with context in case at least one of the closers fails to close.
//
// Panics are not mitigated in any way.
func NewCloseSequence(seq ...io.Closer) io.Closer {
	return &closeSequence{seq: seq}
}

// Close closes all closers in sequence. It will continue with the next closer regardless of whether
// an error occurred.
func (c *closeSequence) Close() error {
	var errs []error
	for _, closer := range c.seq {
		if err := closer.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if errs == nil {
		return nil
	}
	return errors.Aggregate(ErrSequenceFailure, "one or more failures occurred", errs...)
}

var ErrSequenceFailure = errors.NewStringError("error while executing sequence of closers")
