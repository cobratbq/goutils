// SPDX-License-Identifier: LGPL-3.0-only

package net

import (
	"context"
	"net"
	"syscall"

	"github.com/cobratbq/goutils/std/log"
)

// ListenWithOptions listens on the specified network address, using the options set through controlfunc.
// options have format: {option-level (syscall.SOL_*), option (syscall.SO_*, syscall.IP_*, syscall.TCP_*, ..)}
// -> value (integer value)
//
// controlfunc accepts a raw file descriptor for the raw established connection, that can then be configured
// according to low-level socket-options. See syscall package for various socket-options and levels. (See
// syscall.RawConn.Control for details on controlfunc.)
func ListenWithOptions(ctx context.Context, network, addr string, options map[Option]int) (net.Listener, error) {
	config := net.ListenConfig{Control: func(network, address string, conn syscall.RawConn) error {
		return conn.Control(func(fd uintptr) {
			for opt, value := range options {
				err := syscall.SetsockoptInt(int(fd), opt.Level, opt.Option, value)
				if log.Tracing() && err != nil {
					log.Traceln("Failed to set option", opt, "with value", value, ":", err.Error())
				}
			}
		})
	}}
	return config.Listen(ctx, network, addr)
}

// Option combines socket-option option-level and option-ID in a single entry.
//
// See `syscall.SOL_*` constants for available levels, and `syscall.SO_*`, `syscall.IP_*`, `syscall.TCP_*`
// for available options per level.
type Option struct {
	Level  int
	Option int
}
