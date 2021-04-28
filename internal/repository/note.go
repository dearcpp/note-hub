package repository

import (
	"note-hub/internal/database"
	"note-hub/internal/model"

	guuid "github.com/google/uuid"
)

type Note model.Note

func (note *Note) Insert() (int64, error) {
	note.ID = guuid.New().ID()
	return database.Controller.Insert(note)
}

func (note *Note) Get() (bool, error) {
	return database.Controller.Get(note)
}

func (note *Note) Update(data *Note) (int64, error) {
	return database.Controller.Update(data, note)
}

func (note *Note) Delete() (int64, error) {
	return database.Controller.Delete(note)
}
