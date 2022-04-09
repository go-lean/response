package response

import (
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func Test_WriteEmpty(t *testing.T) {
	w := httptest.NewRecorder()
	writer := NewWriter()
	resp := NoContent()
	resp.payloadType = -1

	err := writer.Write(resp, w)

	require.Error(t, err)
	require.Contains(t, err.Error(), "unsupported")
}
