package database

import "github.com/beryll1um/note-hub/internal/models"

func Sync() error {
	var err error

	err = Controller.Sync2(new(models.Note))
	if err != nil {
		return err
	}

	err = Controller.Sync2(new(models.User))
	if err != nil {
		return err
	}

	return nil
}
