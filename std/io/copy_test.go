// SPDX-License-Identifier: LGPL-3.0-only

package io

import (
	"bytes"
	"errors"
	"sync"
	"testing"

	assert "github.com/cobratbq/goutils/std/testing"
)

func TestTransfer(t *testing.T) {
	var src = "Hello world, this is a bunch of data that is being transferred during a CONNECT session."
	srcBuf := bytes.NewBufferString(src)
	dstBuf := closeBuffer{}
	var wg sync.WaitGroup
	wg.Add(1)
	Transfer(&wg, &dstBuf, srcBuf)
	assert.Equal(t, dstBuf.String(), src)
}

func TestTransferError(t *testing.T) {
	var src = "Hello world, this is a bunch of data that is being transferred during a CONNECT session."
	srcBuf := bytes.NewBufferString(src)
	dstBuf := errorBuffer{}
	var wg sync.WaitGroup
	wg.Add(1)
	Transfer(&wg, &dstBuf, srcBuf)
	assert.Equal(t, dstBuf.String(), "Hello worl")
}

type closeBuffer struct {
	bytes.Buffer
}

func (*closeBuffer) Close() error {
	return nil
}

type errorBuffer struct {
	bytes.Buffer
}

func (e *errorBuffer) Write(p []byte) (int, error) {
	_, _ = e.Buffer.Write(p[:10])
	return 10, errors.New("bad stuff happened")
}

func (*errorBuffer) Close() error {
	return nil
}
