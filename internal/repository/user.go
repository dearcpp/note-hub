package repository

import (
	"errors"
	"net"
	"net/http"

	"github.com/beryll1um/note-hub/internal/database"
	"github.com/beryll1um/note-hub/internal/handler"
	"github.com/beryll1um/note-hub/internal/model"
	"xorm.io/xorm"
)

type User model.User

func (user *User) CreateSession(address string, request *handler.Request) error {
	session, err := handler.GetCookieSession(request.Data)

	if err != nil {
		return err
	}

	session.Values["email"] = user.Email
	session.Values["password"] = user.Password
	session.Values["address"] = address

	err = session.Save(request.Data, request.Writer)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) DestroySession(request *handler.Request) error {
	session, err := handler.GetCookieSession(request.Data)

	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	err = session.Save(request.Data, request.Writer)

	if err != nil {
		return err
	}

	return nil
}

func (user *User) GetFromSession(request *http.Request) error {
	session, err := handler.GetCookieSession(request)

	if err != nil {
		return err
	}

	var ok bool

	if user.Email, ok = session.Values["email"].(string); !ok {
		return errors.New("email string is not found")
	}

	if user.Password, ok = session.Values["password"].(string); !ok {
		return errors.New("password string is not found")
	}

	var address string

	if address, ok = session.Values["address"].(string); !ok {
		return errors.New("address string is not found")
	}

	var currentAddress string

	if currentAddress, _, err = net.SplitHostPort(request.RemoteAddr); err != nil {
		return err
	}

	if address != currentAddress {
		return errors.New("session has not verified")
	}

	if has, _ := user.Get(); !has {
		return errors.New("user deleted or does not exist")
	}

	return nil
}

func (user *User) WhereIn() *xorm.Session {
	return database.Controller.Where("author = ?", user.ID)
}

func (user *User) Insert() (int64, error) {
	return database.Controller.Insert(user)
}

func (user *User) Get() (bool, error) {
	return database.Controller.Get(user)
}

func (user *User) Update(data *Note) (int64, error) {
	return database.Controller.Update(data, user)
}

func (user *User) Delete() (int64, error) {
	return database.Controller.Delete(user)
}

func (user *User) Model() model.User {
	return model.User(*user)
}
