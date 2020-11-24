package server

import (
	"fmt"
	"net/http"

	"github.com/beryll1um/note-hub/internal/api"
	"github.com/beryll1um/note-hub/internal/database"
	"github.com/beryll1um/note-hub/internal/models"
	"github.com/beryll1um/note-hub/internal/service"
	"github.com/fatih/color"
)

func MiddlewareRequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		methodColored := color.HiCyanString("[" + request.Method + "]")
		EventChan <- service.Event.Create(service.Event{}, service.Message, fmt.Sprintf("%s %s", methodColored, request.RequestURI))
		next.ServeHTTP(writer, request)
	})
}

func MiddlewareUserAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		session, err := api.GetCookieSession(request)
		if err != nil {
			api.WriteResponse(writer, api.Unauthorized{
				"text": "internal server error",
			})
			return
		}

		mail, hasMail := session.Values["mail"].(string)
		password, hasPassword := session.Values["password"].(string)

		if !hasMail || !hasPassword {
			api.WriteResponse(writer, api.Unauthorized{
				"text": "you are not authorized",
			})
			return
		}

		user := models.User{
			Mail:     mail,
			Password: password,
		}

		if has, _ := database.Controller.Get(&user); !has {
			api.WriteResponse(writer, api.Unauthorized{
				"text": "session is out of date",
			})
			return
		}

		next.ServeHTTP(writer, request)
	})
}
