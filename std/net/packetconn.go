// SPDX-License-Identifier: AGPL-3.0-or-later

package net

import (
	"net"
	"time"
)

// NonClosingPacketConn is a wrapper to prevents wrapped net.PacketConn from being closed. This is
// useful when a function would otherwise prematurely close a connection. Every other function is
// directly wired to the original connection.
type NonClosingPacketConn struct {
	conn net.PacketConn
}

// NewNonClosingPacketConn creates a new NonClosingPacketConn instance.
func NewNonClosingPacketConn(conn net.PacketConn) NonClosingPacketConn {
	return NonClosingPacketConn{conn}
}

func (c *NonClosingPacketConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	return c.conn.ReadFrom(p)
}

func (c *NonClosingPacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	return c.conn.WriteTo(p, addr)
}

func (c *NonClosingPacketConn) Close() error {
	// Do not allow closing.
	return nil
}

func (c *NonClosingPacketConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *NonClosingPacketConn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *NonClosingPacketConn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *NonClosingPacketConn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}
