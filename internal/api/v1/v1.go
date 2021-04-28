package v1

import (
	"note-hub/internal/api/v1/note"
	"note-hub/internal/api/v1/user"
	"note-hub/internal/handler"

	"github.com/gorilla/mux"
)

func SetupHandlers(router *mux.Router) {
	userRouter := router.PathPrefix("/user/").Subrouter()

	userRouter.Handle("/register", handler.Func(user.Register)).Methods("POST")
	userRouter.Handle("/login", handler.Func(user.Login)).Methods("POST")
	userRouter.Handle("/logout", handler.Func(user.Logout)).Methods("DELETE")

	noteRouter := router.PathPrefix("/note/").Subrouter()

	noteRouter.Handle("/create", handler.Func(note.Create)).Methods("POST")
	noteRouter.Handle("/{id:[0-9]+}/get", handler.Func(note.Get)).Methods("GET")
	noteRouter.Handle("/list", handler.Func(note.List)).Methods("GET")
	noteRouter.Handle("/{id:[0-9]+}/delete", handler.Func(note.Delete)).Methods("DELETE")

	noteRouter.Use(MiddlewareAuthentication)
	router.Use(MiddlewareRequestLogger)
}
