package database

import (
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var (
	Controller *xorm.Engine
)

func Connect() {
	postgresURI, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		data, err := ioutil.ReadFile("config/postgres.conf")
		if err != nil {
			log.Fatal(err.Error())
		}
		postgresURI = string(data)
	}

	var err error
	Controller, err = xorm.NewEngine("postgres", postgresURI)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Close() {
	if err := Controller.Close(); err != nil {
		log.Fatal(err.Error())
	}
}
