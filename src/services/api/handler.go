package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
)

type Request struct {
	Data    *http.Request
	Session *sessions.Session
}

type HandlerFunc func(request Request) Response

var store = sessions.NewCookieStore([]byte("TEST_ENV_KEY"))

func (handler HandlerFunc) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var response Response
	var err error

	defer func() {
		var bytes []byte
		bytes, err = json.Marshal(response.Result())
		if err != nil {
			response = InternalServerError("internal server error")
		} else {
			writer.Header().Set("Content-Type", "application/json")
		}
		writer.WriteHeader(response.StatusCode())
		writer.Write(bytes)
	}()

	var session *sessions.Session
	if session, err = store.Get(request, "nh-session"); err != nil {
		response = Unauthorized{
			"text": "internal server error",
		}
		return
	}

	response = handler(Request{
		Data:    request,
		Session: session,
	})

	sessions.Save(request, writer)
}

func SetupRouter(router *mux.Router) {
	router.Handle("/user/sign_out", HandlerFunc(hUserSignOut))
	router.Handle("/user/sign_in", HandlerFunc(hUserSignIn))
	router.Handle("/user/sign_up", HandlerFunc(hUserSignUp))
	router.Handle("/note/create", HandlerFunc(RequiresUser(hNoteCreate)))
	router.Handle("/note/{id:[0-9]+}/get", HandlerFunc(RequiresUser(hNoteGet)))
}
