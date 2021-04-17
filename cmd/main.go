package main

import (
	"log"
	"net/http"
	"time"

	v1 "github.com/beryll1um/note-hub/internal/api/v1"

	"github.com/beryll1um/note-hub/internal/database"
	"github.com/beryll1um/note-hub/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := database.Sync(); err != nil {
		log.Fatal(err)
	}

	if err := handler.SetupCookieStore("config/session.conf"); err != nil {
		log.Fatal(err)
	}

	var router *mux.Router = mux.NewRouter()
	var http *http.Server = &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	v1Router := router.PathPrefix("/api/v1/").Subrouter()
	v1.Setup(v1Router)

	log.Println("Service successfully started!")

	if err := http.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
