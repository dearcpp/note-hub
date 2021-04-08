package database

import "github.com/beryll1um/note-hub/internal/model"

func Sync() error {
	var err error

	err = Controller.Sync2(new(model.Note))
	if err != nil {
		return err
	}

	err = Controller.Sync2(new(model.User))
	if err != nil {
		return err
	}

	return nil
}
