package response

import (
	"io"
	"net/http"
)

type Streamable interface {
	Write(writer StreamWriter) error
}

type StreamWriter interface {
	io.Writer
	http.Flusher
}

type streamWriter struct {
	io.Writer
	http.Flusher
}

type dudFlusher struct{}

func (d *dudFlusher) Flush() {
	// stargaze
}

func NewStreamWriter(w io.Writer) StreamWriter {
	flusher, ok := w.(http.Flusher)
	if ok == false {
		flusher = &dudFlusher{}
	}

	return &streamWriter{w, flusher}
}
