package user

import (
	"net"

	"github.com/beryll1um/note-hub/internal/handler"
	"github.com/beryll1um/note-hub/internal/repository"
	"github.com/beryll1um/note-hub/internal/utilities"
)

func Register(request handler.Request) handler.Response {
	var email string
	var password string
	var err error

	if email, password, err = DecodeBodyCredentials(&request); err != nil {
		return handler.BadRequest{"text": err.Error()}
	}

	user := repository.User{
		Email:    email,
		Password: utilities.Md5EncryptToString(password),
	}

	if affected, _ := user.Insert(); affected == 0 {
		return handler.BadRequest{"text": "failed to register"}
	}

	return handler.Success{
		"text": "successfully registered",
	}
}

func Login(request handler.Request) handler.Response {
	var email string
	var password string
	var err error

	if email, password, err = DecodeBodyCredentials(&request); err != nil {
		return handler.BadRequest{"text": err.Error()}
	}

	user := repository.User{
		Email:    email,
		Password: utilities.Md5EncryptToString(password),
	}

	if has, _ := user.Get(); !has {
		return handler.BadRequest{"text": "invalid credentials"}
	}

	var remoteAddr string

	if remoteAddr, _, err = net.SplitHostPort(request.Data.RemoteAddr); err != nil {
		return handler.InternalServerError{"text": "internal server error"}
	}

	if err = user.CreateSession(remoteAddr, &request); err != nil {
		return handler.InternalServerError{"text": "internal server error"}
	}

	return handler.Success{
		"text": "successfully logged in",
	}
}

func Logout(request handler.Request) handler.Response {
	var user repository.User

	if err := user.GetFromSession(request.Data); err != nil {
		return handler.Unauthorized{"text": "session has not verified"}
	}

	if err := user.DestroySession(&request); err != nil {
		return handler.InternalServerError{"text": "internal server error"}
	}

	return handler.Success{
		"text": "successfully logged out",
	}
}
