package jsonrpc

import (
	"context"
	"testing"

	"github.com/intel-go/fastjson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTakeMethod(t *testing.T) {

	mr := NewMethodRepository()

	r := &Request{}
	_, err := mr.TakeMethod(r)
	require.IsType(t, &Error{}, err)
	assert.Equal(t, ErrorCodeInvalidParams, err.Code)

	r.Method = "test"
	_, err = mr.TakeMethod(r)
	require.IsType(t, &Error{}, err)
	assert.Equal(t, ErrorCodeInvalidParams, err.Code)

	r.Version = "2.0"
	_, err = mr.TakeMethod(r)
	require.IsType(t, &Error{}, err)
	assert.Equal(t, ErrorCodeMethodNotFound, err.Code)

	require.NoError(t, mr.Handle("test", SampleHandler()))

	f, err := mr.TakeMethod(r)
	require.Nil(t, err)
	assert.NotEmpty(t, f)
}

func TestRegisterMethod(t *testing.T) {

	mr := NewMethodRepository()

	err := mr.Handle("", nil)
	require.Error(t, err)

	err = mr.Handle("test", nil)
	require.Error(t, err)

	err = mr.Handle("test", SampleHandler())
	require.NoError(t, err)
}

func TestMethods(t *testing.T) {

	mr := NewMethodRepository()

	err := mr.Handle("JsonRpc.Sample", SampleHandler())
	require.NoError(t, err)

	ml := mr.Methods()
	require.NotEmpty(t, ml)
	assert.NotEmpty(t, ml["JsonRpc.Sample"].Handler)
}

func SampleHandler() Handler {
	h := handler{}
	h.F = func(c context.Context, params *fastjson.RawMessage) (result interface{}, err *Error) {
		return nil, nil
	}
	return &h
}
