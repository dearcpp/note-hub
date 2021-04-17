package repository

import (
	"github.com/beryll1um/note-hub/internal/database"
	"github.com/beryll1um/note-hub/internal/model"
)

type Note model.Note

func (note *Note) Insert() (int64, error) {
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

func (note *Note) Model() model.Note {
	return model.Note(*note)
}
