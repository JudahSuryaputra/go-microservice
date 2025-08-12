package utils

import (
	"bytes"
	"net/http"
)

type BodyDumpResponseWriter struct {
	http.ResponseWriter
	Body *bytes.Buffer
}

func (w *BodyDumpResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b) // capture response body
	return w.ResponseWriter.Write(b)
}
