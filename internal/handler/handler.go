package handler

import "net/http"

type Func func(Request) Response

func (handler Func) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	WriteResponse(writer, handler(Request{
		Data:   request,
		Writer: writer,
	}))
}
