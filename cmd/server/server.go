package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/beryll1um/note-hub/internal/api"
	"github.com/beryll1um/note-hub/internal/database"
	"github.com/beryll1um/note-hub/internal/service"
	"github.com/gorilla/mux"
)

// @todo: make getting TEST_ENV_KEY from environment

var (
	EventChan service.EventChanType
)

func Observer(c service.EventChanType) {
	for {
		fmt.Println((<-c).Message)
	}
}

func Service(c service.EventChanType) {

	EventChan = c

	database.Connect()
	defer database.Close()

	if err := database.Sync(); err != nil {
		c <- service.Event.Create(service.Event{}, service.Error, "Failed to synchronize database!")
		c <- service.Event.Create(service.Event{}, service.Close, "Fatal error: Service terminating...")
		return
	}

	c <- service.Event.Create(service.Event{}, service.Ready, "Service successfully started!")

	var router *mux.Router = mux.NewRouter()
	var http *http.Server = &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	authRouter := router.PathPrefix("/auth/").Subrouter()
	authRouter.Handle("/register", api.HandlerFunc(hUserRegister)).Methods("POST")
	authRouter.Handle("/login", api.HandlerFunc(hUserLogin)).Methods("POST")
	authRouter.Handle("/logout", api.HandlerFunc(hUserLogout))

	userRouter := router.PathPrefix("/note/").Subrouter()
	userRouter.Handle("/create", api.HandlerFunc(hNoteCreate)).Methods("POST")
	userRouter.Handle("/list", api.HandlerFunc(hNoteList)).Methods("GET")
	userRouter.Handle("/{id:[0-9]+}/get", api.HandlerFunc(hNoteGet)).Methods("GET")
	userRouter.Handle("/{id:[0-9]+}/delete", api.HandlerFunc(hNoteDelete)).Methods("DELETE")
	userRouter.Use(MiddlewareUserAuthentication)

	router.Use(MiddlewareRequestLogger)

	if err := http.ListenAndServe(); err != nil {
		c <- service.Event.Create(service.Event{}, service.Error, err.Error())
		c <- service.Event.Create(service.Event{}, service.Close, "Fatal error: Service terminating...")
	}

}
