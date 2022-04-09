package response_test

import (
	"errors"
	"github.com/go-lean/response"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func Test_CustomWithError(t *testing.T) {
	resp := response.New(http.StatusTeapot).
		WithError(errors.New("baba"), "you")

	require.Equal(t, http.StatusTeapot, resp.StatusCode())
	require.Nil(t, resp.Payload())
	require.Error(t, resp.Error())
	require.Contains(t, resp.Error().Error(), "baba")
	require.NotContains(t, resp.Error().Error(), "you")
	require.Contains(t, resp.ErrorMessage(), "you")
	require.NotContains(t, resp.ErrorMessage(), "baba")
}

func Test_CustomWithoutContent(t *testing.T) {
	resp := response.New(http.StatusTeapot).
		WithoutContent()

	require.Equal(t, http.StatusTeapot, resp.StatusCode())
	require.Nil(t, resp.Payload())
	require.Nil(t, resp.Error())
	require.Equal(t, response.PayloadEmpty, resp.PayloadType())
}

func Test_OKPayload(t *testing.T) {
	type data struct {
		data string
	}
	payload := data{"baba"}

	resp := response.OK().WithJSON(payload)
	respPayload, ok := resp.Payload().(data)

	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.Equal(t, response.PayloadJSON, resp.PayloadType())
	require.True(t, ok)
	require.Equal(t, "baba", respPayload.data)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_OKText(t *testing.T) {
	resp := response.OK().WithText("baba")
	respPayload, ok := resp.Payload().(string)

	require.Equal(t, http.StatusOK, resp.StatusCode())
	require.Equal(t, response.PayloadText, resp.PayloadType())
	require.True(t, ok)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
	require.Equal(t, "baba", respPayload)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_Created(t *testing.T) {
	type data struct {
		data string
	}
	payload := data{"baba"}
	resp := response.Created().WithJSON(payload)
	respPayload, ok := resp.Payload().(data)

	require.Equal(t, http.StatusCreated, resp.StatusCode())
	require.True(t, ok)
	require.Equal(t, "baba", respPayload.data)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_Accepted(t *testing.T) {
	resp := response.Accepted().WithJSON("baba")
	respPayload, ok := resp.Payload().(string)

	require.Equal(t, http.StatusAccepted, resp.StatusCode())
	require.True(t, ok)
	require.Equal(t, "baba", respPayload)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_NoContent(t *testing.T) {
	resp := response.NoContent()

	require.Equal(t, http.StatusNoContent, resp.StatusCode())
	require.Nil(t, resp.Payload())
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
}

func Test_Errors(t *testing.T) {
	tests := []struct {
		c int
		r *response.HttpResponse
	}{
		{http.StatusBadRequest, response.BadRequest(errors.New("baba"), "you")},
		{http.StatusUnauthorized, response.Unauthorized(errors.New("baba"), "you")},
		{http.StatusPaymentRequired, response.PaymentRequired(errors.New("baba"), "you")},
		{http.StatusForbidden, response.Forbidden(errors.New("baba"), "you")},
		{http.StatusNotFound, response.NotFound(errors.New("baba"), "you")},
		{http.StatusNotAcceptable, response.NotAcceptable(errors.New("baba"), "you")},
		{http.StatusConflict, response.Conflict(errors.New("baba"), "you")},
		{http.StatusInternalServerError, response.InternalServerError(errors.New("baba"), "you")},
		{http.StatusNotImplemented, response.NotImplemented(errors.New("baba"), "you")},
		{http.StatusServiceUnavailable, response.ServiceUnavailable(errors.New("baba"), "you")},
		{http.StatusInsufficientStorage, response.InsufficientStorage(errors.New("baba"), "you")},
	}

	for _, tc := range tests {
		require.Equal(t, tc.c, tc.r.StatusCode())
		require.Error(t, tc.r.Error())
		require.NotEmpty(t, tc.r.ErrorMessage())
		require.Contains(t, tc.r.Error().Error(), "baba")
		require.NotContains(t, tc.r.Error().Error(), "you")
		require.Contains(t, tc.r.ErrorMessage(), "you")
		require.NotContains(t, tc.r.ErrorMessage(), "baba")
	}
}
