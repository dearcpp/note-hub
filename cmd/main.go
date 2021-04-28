package main

import (
	"log"
	"net/http"
	"os"
	"time"

	v1 "note-hub/internal/api/v1"

	"note-hub/internal/database"
	"note-hub/internal/handler"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := database.Sync(); err != nil {
		log.Fatal(err)
	}

	handler.SetupCookieStore(os.Getenv("COOKIES_ENCRYPT_KEY"))

	var router *mux.Router = mux.NewRouter()
	var http *http.Server = &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	v1Router := router.PathPrefix("/api/v1/").Subrouter()
	v1.SetupHandlers(v1Router)

	if err := http.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
