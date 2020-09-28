package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type HandlerFunc func(request *http.Request) Response

func (handler HandlerFunc) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	response := handler(request)
	bytes, err := json.Marshal(response.Result())
	if err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte("internal error"))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(response.StatusCode())
	writer.Write(bytes)
}

func SetupRouter(router *mux.Router) {
	router.Handle("/note/create", HandlerFunc(hNoteCreate))
	router.Handle("/note/{id:[0-9]+}/get", HandlerFunc(hNoteGet))
}
