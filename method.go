package jsonrpc

import (
	"errors"
	"sync"
)

type (
	// A MethodRepository has JSON-RPC method functions.
	MethodRepository struct {
		m sync.RWMutex
		r map[string]Metadata
	}
	// Metadata has method meta data.
	Metadata struct {
		Handler Handler
	}
)

// NewMethodRepository returns new MethodRepository.
func NewMethodRepository() *MethodRepository {
	return &MethodRepository{
		m: sync.RWMutex{},
		r: map[string]Metadata{},
	}
}

// TakeMethod takes jsonrpc.Func in MethodRepository.
func (mr *MethodRepository) TakeMethod(r *Request) (Handler, *Error) {
	if r.Method == "" || r.Version != Version {
		return nil, ErrInvalidParams()
	}

	mr.m.RLock()
	md, ok := mr.r[r.Method]
	mr.m.RUnlock()
	if !ok {
		return nil, ErrMethodNotFound()
	}

	return md.Handler, nil
}

// Handle registers jsonrpc.Func to MethodRepository.
func (mr *MethodRepository) Handle(method string, h Handler) error {
	if method == "" || h == nil {
		return errors.New("jsonrpc: method name and function should not be empty")
	}
	mr.m.Lock()
	mr.r[method] = Metadata{
		Handler: h,
	}
	mr.m.Unlock()
	return nil
}

// Methods returns registered methods.
func (mr *MethodRepository) Methods() map[string]Metadata {
	mr.m.RLock()
	ml := make(map[string]Metadata, len(mr.r))
	for k, md := range mr.r {
		ml[k] = md
	}
	mr.m.RUnlock()
	return ml
}
