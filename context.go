package jsonrpc

import (
	"context"
	"net/http"
)

type responseWriter struct{}

type headers struct{}

type cookies struct{}

type сookieGetter func(name string) (*http.Cookie, error)

type method struct{}

type requestId struct{}


func Method(c context.Context) string {
	return c.Value(method{}).(string)
}

func SetMethod(c context.Context, val string) context.Context {
	return context.WithValue(c, method{}, val)
}

func Headers(c context.Context) http.Header {
	return c.Value(headers{}).(http.Header)
}

func SetHeaders(c context.Context, h http.Header) context.Context {
	return context.WithValue(c, headers{}, h)
}

func ResponseWriter(c context.Context) http.ResponseWriter {
	return c.Value(responseWriter{}).(http.ResponseWriter)
}

func SetResponseWriter(c context.Context, writer http.ResponseWriter) context.Context {
	return context.WithValue(c, responseWriter{}, writer)
}

func Cookie(c context.Context, name string) (*http.Cookie, error) {
	return c.Value(cookies{}).(сookieGetter)(name)
}

func SetCookie(c context.Context, cookie сookieGetter) context.Context {
	return context.WithValue(c, cookies{}, cookie)
}

// RequestId takes request id from context.
func RequestId(c context.Context) interface{} {
	return c.Value(requestId{})
}

// WithRequestId adds request id to context.
func SetRequestId(c context.Context, id interface{}) context.Context {
	return context.WithValue(c, requestId{}, id)
}
