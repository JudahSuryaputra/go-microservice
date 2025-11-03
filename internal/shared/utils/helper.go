package utils

import (
	"io"
	"net/http"
)

type BodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *BodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
