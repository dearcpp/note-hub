package handler

import "context"

type ContextValue int

const (
	User ContextValue = iota
)

func (request *Request) SetContextValue(key ContextValue, value interface{}) {
	ctx := context.WithValue(request.Data.Context(), key, value)
	request.Data = request.Data.WithContext(ctx)
}

func (request *Request) GetContextValue(key ContextValue) interface{} {
	return request.Data.Context().Value(key)
}
