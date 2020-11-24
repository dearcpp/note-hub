package database

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var (
	Controller *xorm.Engine
)

func Connect() error {
	var err error
	Controller, err = xorm.NewEngine("sqlite3", "database.sqlite")
	return err
}

func Close() error {
	return Controller.Close()
}
