package repository

import (
	"note-hub/internal/database"
	"note-hub/internal/model"

	guuid "github.com/google/uuid"
	"xorm.io/xorm"
)

type User model.User

func (user *User) WhereIn() *xorm.Session {
	return database.Controller.Where("author = ?", user.ID)
}

func (user *User) Insert() (int64, error) {
	user.ID = guuid.New().ID()
	return database.Controller.Insert(user)
}

func (user *User) Get() (bool, error) {
	return database.Controller.Get(user)
}

func (user *User) Update(data *User) (int64, error) {
	return database.Controller.Update(data, user)
}

func (user *User) Delete() (int64, error) {
	return database.Controller.Delete(user)
}
