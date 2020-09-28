package main

import (
	"github.com/beryll1um/note-hub/src/database"
	"github.com/beryll1um/note-hub/src/models"
	"github.com/beryll1um/note-hub/src/services/api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	database.Connect()
	defer database.Close()

	if err := database.Controller.Sync(
		new(models.Note),
	); err != nil {
		log.Fatal("Failed to sync database")
	}

	log.Println("Successfully started")

	mainRouter := mux.NewRouter()
	api.SetupRouter(mainRouter)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mainRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
