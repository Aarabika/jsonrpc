package jsonrpc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
)


func TestHeader(t *testing.T) {

	c := context.Background()
	headers := http.Header{
		"header": []string{
			"simple header 1",
			"simple header 2",
		},
	}
	c = SetHeaders(c, headers)
	var pick http.Header
	require.NotPanics(t, func() {
		pick = Headers(c)
	})
	require.Equal(t, headers, pick)
}

func TestResponseWriter(t *testing.T) {

	c := context.Background()
	w := httptest.NewRecorder()
	c = SetResponseWriter(c, w)
	var pick http.ResponseWriter
	require.NotPanics(t, func() {
		pick = ResponseWriter(c)
	})
	require.Equal(t, w, pick)
}