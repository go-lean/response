package response_test

import (
	"encoding/json"
	"errors"
	"github.com/go-lean/response"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CustomErrorWriter(t *testing.T) {
	type testData struct {
		Error      string
		StatusCode int
	}
	w := httptest.NewRecorder()
	writer := response.NewWriter()
	writer.ErrorWriterFunc = func(httpResponse *response.HttpResponse, writer http.ResponseWriter) error {
		data := testData{
			httpResponse.ErrorMessage(),
			httpResponse.StatusCode(),
		}

		return json.NewEncoder(writer).Encode(data)
	}
	resp := response.BadRequest(errors.New("baba"), "you")

	err := writer.Write(resp, w)

	require.Nil(t, err)
	var data testData
	require.Nil(t, json.NewDecoder(w.Body).Decode(&data))
	require.Contains(t, data.Error, "you", w.Body.String())
	require.Equal(t, http.StatusBadRequest, data.StatusCode)
}

func Test_WriteText(t *testing.T) {
	w := httptest.NewRecorder()
	writer := response.NewWriter()
	resp := response.Created().WithText("baba")

	err := writer.Write(resp, w)

	require.Nil(t, err)
	require.Contains(t, w.Body.String(), "baba", w.Body.String())
	require.Equal(t, http.StatusCreated, w.Code)
	require.Contains(t, w.Header().Get("Content-Type"), "text/plain")
}

func Test_WritePayload(t *testing.T) {
	type testData struct {
		Data string
		Code int
	}
	data := testData{
		"baba",
		201,
	}
	w := httptest.NewRecorder()
	writer := response.NewWriter()
	resp := response.Created().WithJSON(data)

	err := writer.Write(resp, w)

	require.Nil(t, err)
	var respPayload testData
	require.Nil(t, json.NewDecoder(w.Body).Decode(&respPayload))
	require.Equal(t, "baba", respPayload.Data)
	require.Equal(t, 201, respPayload.Code)
	require.Equal(t, http.StatusCreated, w.Code)
	require.Contains(t, w.Header().Get("Content-Type"), "json")
}

func Test_WriteError(t *testing.T) {
	w := httptest.NewRecorder()
	writer := response.NewWriter()
	resp := response.BadRequest(errors.New("baba"), "you")

	err := writer.Write(resp, w)

	require.Nil(t, err)
	require.Equal(t, http.StatusBadRequest, w.Code)
	require.Contains(t, w.Header().Get("Content-Type"), "text/plain")
	require.Contains(t, w.Body.String(), "you")
	require.NotContains(t, w.Body.String(), "baba")
}

func Test_WriteEmpty(t *testing.T) {
	w := httptest.NewRecorder()
	writer := response.NewWriter()
	resp := response.NoContent()

	err := writer.Write(resp, w)

	require.Nil(t, err)
	require.Equal(t, http.StatusNoContent, w.Code)
	require.Empty(t, w.Body.String())
}
