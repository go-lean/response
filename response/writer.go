package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrUnsupportedPayloadType = errors.New("unsupported payload type")
	ErrNotStreamable          = errors.New("provided payload is not streamable")
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
	case PayloadStream:
		return writeStream(response, writer)
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

	streamable, ok := response.payload.(Streamable)
	if ok {
		writer.Header().Set("X-Content-Type-Options", "nosniff")
		writer.Header().Set("Connection", "Keep-Alive")
		writer.WriteHeader(response.statusCode)

		streamedWriter := NewStreamWriter(writer)
		err := streamable.Write(streamedWriter)
		streamedWriter.Flush()

		return err
	}

	writer.WriteHeader(response.statusCode)
	return json.NewEncoder(writer).Encode(response.payload)
}

func writeStream(response *HttpResponse, writer http.ResponseWriter) error {
	streamer, ok := response.payload.(Streamable)
	if ok == false {
		return ErrNotStreamable
	}

	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.Header().Set("Connection", "Keep-Alive")
	writer.Header().Set(contentTypeKey, response.contentType)
	writer.WriteHeader(response.statusCode)

	stream := NewStreamWriter(writer)
	if err := streamer.Write(stream); err != nil {
		return err
	}

	stream.Flush()
	return nil
}
