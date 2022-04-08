package response_test

import (
	"errors"
	"github.com/go-lean/response"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_OKPayload(t *testing.T) {
	type data struct {
		data string
	}
	payload := data{"baba"}

	resp := response.OKPayload(payload)
	respPayload, ok := resp.Payload().(data)

	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.True(t, ok)
	require.Equal(t, "baba", respPayload.data)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_OKText(t *testing.T) {
	resp := response.OKText("baba")
	respPayload, ok := resp.Payload().(string)

	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.True(t, ok)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
	require.Equal(t, "baba", respPayload)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_Created(t *testing.T) {
	resp := response.Created("baba")
	respPayload, ok := resp.Payload().(string)

	require.Equal(t, http.StatusCreated, resp.StatusCode())
	require.True(t, ok)
	require.Equal(t, "baba", respPayload)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_Accepted(t *testing.T) {
	resp := response.Accepted("baba")
	respPayload, ok := resp.Payload().(string)

	require.Equal(t, http.StatusAccepted, resp.StatusCode())
	require.True(t, ok)
	require.Equal(t, "baba", respPayload)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_NoContent(t *testing.T) {
	resp := response.NoContent()

	require.Equal(t, http.StatusAccepted, resp.StatusCode())
	require.Nil(t, resp.Payload())
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_BadRequest(t *testing.T) {
	resp := response.BadRequest(errors.New("baba"), "dodo")

	require.Equal(t, http.StatusBadRequest, resp.StatusCode())
	require.Error(t, resp.Error())
	require.NotEmpty(t, resp.ErrorMessage())
	require.Contains(t, resp.Error().Error(), "baba")
	require.NotContains(t, resp.Error().Error(), "dodo")
	require.Contains(t, resp.ErrorMessage(), "dodo")
	require.NotContains(t, resp.ErrorMessage(), "baba")
}
