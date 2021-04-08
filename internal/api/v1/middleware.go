package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/beryll1um/note-hub/internal/handler"
	"github.com/beryll1um/note-hub/internal/repository"
)

func MiddlewareAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var user repository.User
		if err := user.GetFromSession(request); err != nil {
			handler.WriteResponse(writer, handler.Unauthorized{
				"text": err.Error(),
			})
		} else {
			next.ServeHTTP(writer, request)
		}
	})
}

func MiddlewareRequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		method := "[" + request.Method + "]"
		log.Println(fmt.Sprintf("%s %s", method, request.RequestURI))
		next.ServeHTTP(writer, request)
	})
}
