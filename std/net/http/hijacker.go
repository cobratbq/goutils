package http

import (
	"errors"
	"io"
	"net"
	"net/http"
)

// ErrNonHijackableWriter is returned when the connection cannot be hijacked.
var ErrNonHijackableWriter = errors.New("failed to acquire raw client connection: writer is not hijackable")

// HijackConnection acquires the underlying connection by inspecting the
// ResponseWriter provided. One may use the returned io.ReadWriter to perform
// communication operations. The net.Conn instance is provided for low-level
// connection maintenance.
// The user must close net.Conn!
func HijackConnection(resp http.ResponseWriter) (io.Reader, net.Conn, error) {
	hijacker, ok := resp.(http.Hijacker)
	if !ok {
		return nil, nil, ErrNonHijackableWriter
	}
	clientConn, bufrw, err := hijacker.Hijack()
	return bufrw, clientConn, err
}
