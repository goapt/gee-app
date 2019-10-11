package testutil

import (
	"bufio"
	"bytes"
	"net"
	"net/http"
)

type BufferWriter struct {
	header     http.Header
	buf        *bytes.Buffer
	statusCode int
}

func (w *BufferWriter) Header() http.Header {
	return w.header
}

func (w *BufferWriter) Write(data []byte) (int, error) {
	return w.buf.Write(data)
}

func (w *BufferWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (*BufferWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	panic("implement me")
}

func (*BufferWriter) Flush() {}

func (*BufferWriter) CloseNotify() <-chan bool {
	return make(chan bool)
}

func (w *BufferWriter) Status() int {
	return w.statusCode
}

func (w *BufferWriter) Size() int {
	return w.buf.Len()
}

func (w *BufferWriter) WriteString(s string) (int, error) {
	return w.buf.WriteString(s)
}

func (w *BufferWriter) Written() bool {
	return true
}

func (*BufferWriter) WriteHeaderNow() {}

func (*BufferWriter) Pusher() http.Pusher {
	return nil
}
