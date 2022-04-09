package response

import (
	"io"
	"net/http"
)

type Stream struct {
	io.Writer
	flusher http.Flusher
}

type dudFlusher struct{}

func (d *dudFlusher) Flush() {
	// stargaze
}

func (s *Stream) Flush() {
	s.flusher.Flush()
}

func NewStream(w io.Writer) *Stream {
	flusher, ok := w.(http.Flusher)
	if ok == false {
		flusher = &dudFlusher{}
	}

	return &Stream{w, flusher}
}

type Streamer interface {
	Stream(stream *Stream) error
}
