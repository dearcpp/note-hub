package database

import "note-hub/internal/model"

var models = []interface{}{
	new(model.Note),
	new(model.User),
}

func Sync() error {
	return Controller.Sync(models...)
}

func Drop() error {
	return Controller.DropTables(models...)
}
