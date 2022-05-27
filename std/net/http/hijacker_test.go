package http

import (
	"bufio"
	"net"
	"net/http"
	"testing"
)

func TestNonHijackableWriter(t *testing.T) {
	var writer nonHijackableWriter
	_, conn, err := HijackConnection(&writer)
	if err != ErrNonHijackableWriter {
		t.Error("Expected to receive non-hijackable writer error but got:", err.Error())
	}
	if conn != nil {
		t.Error("Expected to get no connection but got something anyway:", conn)
	}
}

func TestHijackableWriter(t *testing.T) {
	var writer hijackableWriter
	_, conn, err := HijackConnection(&writer)
	if err != nil {
		t.Error("Expected to get no error but got:", err.Error())
	}
	if conn == nil {
		t.Error("Expected to get a connection but got nothing.")
	}
	defer conn.Close()
}

type nonHijackableWriter struct{}

func (*nonHijackableWriter) Header() http.Header {
	return nil
}

func (*nonHijackableWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (*nonHijackableWriter) WriteHeader(int) {
}

type hijackableWriter struct {
	nonHijackableWriter
}

func (*hijackableWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	pipeR, pipeW := net.Pipe()
	return pipeR, bufio.NewReadWriter(bufio.NewReader(pipeR), bufio.NewWriter(pipeW)), nil
}
