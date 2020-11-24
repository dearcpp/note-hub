package main

import (
	"github.com/beryll1um/note-hub/cmd/server"
	"github.com/beryll1um/note-hub/internal/service"
)

func main() {

	server := service.Manager.Create(service.Manager{}, server.Service, server.Observer)
	server.Run()

}
