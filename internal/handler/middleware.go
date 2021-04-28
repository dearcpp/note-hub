package handler

import "net/http"

type Middleware func(Request)

func (middlerware Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	middlerware(Request{
		Data:   request,
		Writer: writer,
	})
}
