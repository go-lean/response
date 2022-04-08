package response_test

import (
	"errors"
	"github.com/go-lean/response"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_OKPayload(t *testing.T) {
	type data struct {
		data string
	}
	payload := data{"baba"}

	resp := response.OKPayload(payload)
	respPayload, ok := resp.Payload().(data)

	require.True(t, ok)
	require.Equal(t, "baba", respPayload.data)
}

func Test_OKText(t *testing.T) {
	resp := response.OKText("baba")
	respPayload, ok := resp.Payload().(string)

	require.True(t, ok)
	require.Nil(t, resp.Error())
	require.Empty(t, resp.ErrorMessage())
	require.Equal(t, "baba", respPayload)
}

func Test_BadRequest(t *testing.T) {
	resp := response.BadRequest(errors.New("baba"), "dodo")

	require.Error(t, resp.Error())
	require.NotEmpty(t, resp.ErrorMessage())
	require.Contains(t, resp.Error().Error(), "baba")
	require.NotContains(t, resp.Error().Error(), "dodo")
	require.Contains(t, resp.ErrorMessage(), "dodo")
	require.NotContains(t, resp.ErrorMessage(), "baba")
}
