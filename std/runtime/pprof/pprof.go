// SPDX-License-Identifier: LGPL-3.0-or-later

// pprof provides utilities for `runtime/pprof`.
//
// References:
// - <https://pkg.go.dev/runtime/pprof>
// - <https://github.com/google/pprof/blob/main/doc/README.md>
package pprof

import (
	"io"
	"os"
	"runtime/pprof"

	"github.com/cobratbq/goutils/std/errors"
	io_ "github.com/cobratbq/goutils/std/io"
)

func StartCPUProfilingFileBacked(fileName string) (io.Closer, error) {
	dest, fileErr := os.Create(fileName)
	if fileErr != nil {
		return nil, errors.Context(fileErr, "failed to open file for file-backed profiling")
	}
	profilerCloser, profilerErr := StartCPUProfiling(dest)
	if profilerErr != nil {
		dest.Close()
		return nil, errors.Context(profilerErr, "failed to start CPU profiler (with file-backing)")
	}
	return io_.NewCloseSequence(profilerCloser, dest), nil
}

func StartCPUProfiling(output io.Writer) (io.Closer, error) {
	if err := pprof.StartCPUProfile(output); err != nil {
		return nil, err
	}
	return cpuProfilerCloser{}, nil
}

type cpuProfilerCloser struct{}

func (cpuProfilerCloser) Close() error {
	pprof.StopCPUProfile()
	return nil
}
