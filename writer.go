package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrUnsupportedPayloadType = errors.New("unsupported payload type")
	ErrNotAStreamer           = errors.New("provided payload is not a streamer")
)

const (
	contentTypeKey        = "Content-Type"
	contentTypeOptionsKey = "X-Content-Type-Options"
)

type ErrorWriterFunc func(*HttpResponse, http.ResponseWriter) error

type Writer struct {
	TextContentType string
	JSONContentType string
	ErrorWriterFunc ErrorWriterFunc
}

func NewWriter() Writer {
	return Writer{
		"text/plain; charset=UTF-8",
		"application/json; charset=UTF-8",
		nil,
	}
}

func (w *Writer) Write(response *HttpResponse, writer http.ResponseWriter) error {
	if response.errMessage != "" || response.logError != nil {
		if w.ErrorWriterFunc != nil {
			return w.ErrorWriterFunc(response, writer)
		}

		return w.writeError(response, writer)
	}

	switch response.payloadType {
	case PayloadEmpty:
		writer.WriteHeader(response.statusCode)
		return nil
	case PayloadText:
		return w.writeText(response, writer)
	case PayloadJSON:
		return w.writeJSON(response, writer)
	case PayloadStreaming:
		return w.writeStream(response, writer)
	default:
		return ErrUnsupportedPayloadType
	}
}

func (w *Writer) writeError(response *HttpResponse, writer http.ResponseWriter) error {
	writer.Header().Set(contentTypeKey, w.TextContentType)
	writer.Header().Set(contentTypeOptionsKey, "nosniff")
	writer.WriteHeader(response.statusCode)

	_, err := writer.Write([]byte(response.errMessage))

	return err
}

func (w *Writer) writeText(response *HttpResponse, writer http.ResponseWriter) error {
	writer.Header().Set(contentTypeKey, w.TextContentType)
	writer.WriteHeader(response.statusCode)

	_, err := writer.Write([]byte(response.payload.(string)))

	return err
}

func (w *Writer) writeJSON(response *HttpResponse, writer http.ResponseWriter) error {
	writer.Header().Set(contentTypeKey, w.JSONContentType)
	writer.WriteHeader(response.statusCode)

	return json.NewEncoder(writer).Encode(response.payload)
}

func (w *Writer) writeStream(response *HttpResponse, writer http.ResponseWriter) error {
	stream := NewStream(writer)
	streamer, ok := response.payload.(Streamer)
	if ok == false {
		return ErrNotAStreamer
	}

	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.Header().Set("Connection", "Keep-Alive")
	writer.WriteHeader(response.statusCode)

	if err := streamer.Stream(stream); err != nil {
		return err
	}

	stream.Flush()
	return nil
}
