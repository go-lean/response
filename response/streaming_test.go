package response_test

import (
	"errors"
	"github.com/go-lean/response"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type streamer struct {
	fun func(response.StreamWriter) error
}

type dudWriter struct{}

func (d *dudWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (d *dudWriter) Header() http.Header {
	return http.Header{}
}

func (d *dudWriter) WriteHeader(_ int) {
	// stargaze
}

func (s *streamer) Write(stream response.StreamWriter) error {
	return s.fun(stream)
}

func TestStream(t *testing.T) {
	w := httptest.NewRecorder()

	str := &streamer{fun: func(stream response.StreamWriter) error {
		tests := []struct {
			addition string
			expected string
		}{
			{"baba", "baba"},
			{" is", "baba is"},
			{" you", "baba is you"},
		}

		for _, tc := range tests {
			_, err := stream.Write([]byte(tc.addition))
			require.Nil(t, err)
			require.Equal(t, tc.expected, w.Body.String())
			stream.Flush()
		}

		return nil
	}}

	resp := response.OK().WithStreamable(str, "application/json")
	writer := response.NewWriter()

	err := writer.Write(resp, w)

	require.Nil(t, err)
	require.Equal(t, "baba is you", w.Body.String())
}

func TestStreamError(t *testing.T) {
	str := &streamer{fun: func(stream response.StreamWriter) error {
		stream.Flush()
		return errors.New("baba")
	}}

	resp := response.OK().WithStreamable(str, "application/json")
	w := httptest.NewRecorder()
	writer := response.NewWriter()

	err := writer.Write(resp, w)

	require.Error(t, err)
	require.Contains(t, err.Error(), "baba")
}

func TestStreamNilStreamer(t *testing.T) {
	resp := response.OK().WithStreamable(nil, "application/json")
	w := httptest.NewRecorder()
	writer := response.NewWriter()

	err := writer.Write(resp, w)

	require.Error(t, err)
	require.Contains(t, err.Error(), "streamable")
}

func TestStreamNonFlusherWriter(t *testing.T) {
	str := &streamer{fun: func(stream response.StreamWriter) error {
		stream.Flush()
		return nil
	}}

	resp := response.OK().WithStreamable(str, "application/json")
	w := &dudWriter{}
	writer := response.NewWriter()

	err := writer.Write(resp, w)

	require.Nil(t, err)
}
