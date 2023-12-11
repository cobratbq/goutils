// SPDX-License-Identifier: LGPL-3.0-only

package net

import (
	"context"
	"net"
	"syscall"
)

// ListenWithOptions listens on the specified network address, using the options set through controlfunc.
//
// controlfunc accepts a raw file descriptor for the raw established connection, that can then be configured
// according to low-level socket-options. See syscall package for various socket-options and levels. (See
// syscall.RawConn.Control for details on controlfunc.)
func ListenWithOptions(ctx context.Context, network, addr string, controlfunc func(fd uintptr)) (net.Listener, error) {
	config := net.ListenConfig{Control: func(network, address string, conn syscall.RawConn) error {
		return conn.Control(controlfunc)
	}}
	return config.Listen(ctx, network, addr)
}
