package api

import (
	"net/http"
)

type Request struct {
	Data   *http.Request
	Writer *http.ResponseWriter
}

type HandlerFunc func(Request) Response

func (handler HandlerFunc) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	WriteResponse(writer, handler(Request{
		Data:   request,
		Writer: &writer,
	}))
}
