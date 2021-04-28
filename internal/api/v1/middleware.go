package v1

import (
	"fmt"
	"log"
	"net/http"
	"note-hub/internal/handler"
	"note-hub/internal/repository"
)

func MiddlewareAuthentication(next http.Handler) http.Handler {
	return handler.Middleware(func(request handler.Request) {
		session, err := request.GetSession()
		if err != nil {
			handler.WriteResponse(request.Writer, handler.InternalServerError{
				"text": "internal server error",
			})
			return
		}

		var user repository.User

		if err := session.GetUser(&user); err != nil {
			handler.WriteResponse(request.Writer, handler.Unauthorized{
				"text": "session has not verified",
			})
			return
		}

		request.Set(handler.User, user)

		next.ServeHTTP(request.Writer, request.Data)
	})
}

func MiddlewareRequestLogger(next http.Handler) http.Handler {
	return handler.Middleware(func(request handler.Request) {
		log.Println(fmt.Sprintf("[%s] %s", request.Data.Method, request.Data.RequestURI))
		next.ServeHTTP(request.Writer, request.Data)
	})
}
