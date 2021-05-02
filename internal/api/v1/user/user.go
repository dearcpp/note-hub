package user

import (
	"note-hub/internal/handler"
	"note-hub/internal/repository"
	"note-hub/internal/utilities"
)

type RegisterParameters struct {
	Email    string `json:"email" field:"required max_length(320)"`
	Password string `json:"password" field:"required max_length(128)"`
}

func Register(request handler.Request) handler.Response {
	var params RegisterParameters

	if err := request.ParseBody(&params); err != nil {
		return handler.BadRequest{
			"text": err.Error(),
		}
	}

	user := repository.User{
		Email:    params.Email,
		Password: utilities.Md5EncryptToString(params.Password),
	}

	if affected, _ := user.Insert(); affected == 0 {
		return handler.BadRequest{"text": "failed to register"}
	}

	return handler.Success{
		"text": "successfully registered",
	}
}

type LoginParameters struct {
	Email    string `json:"email" field:"required max_length(320)"`
	Password string `json:"password" field:"required max_length(128)"`
}

func Login(request handler.Request) handler.Response {
	var params RegisterParameters

	if err := request.ParseBody(&params); err != nil {
		return handler.BadRequest{
			"text": err.Error(),
		}
	}

	session, err := request.GetSession()
	if err != nil {
		return handler.InternalServerError{
			"text": "internal server error",
		}
	}

	user := repository.User{
		Email:    params.Email,
		Password: utilities.Md5EncryptToString(params.Password),
	}

	if has, _ := user.Get(); !has {
		return handler.BadRequest{
			"text": "invalid credentials",
		}
	}

	if err := session.SetUser(&user); err != nil {
		return handler.InternalServerError{
			"text": "internal server error",
		}
	}

	return handler.Success{
		"text": "successfully logged in",
	}
}

func Logout(request handler.Request) handler.Response {
	session, err := request.GetSession()
	if err != nil {
		return handler.InternalServerError{
			"text": "internal server error",
		}
	}

	var user repository.User

	if err := session.GetUser(&user); err != nil {
		return handler.Unauthorized{
			"text": "session has not verified",
		}
	}

	if err := session.Destroy(); err != nil {
		return handler.InternalServerError{
			"text": "internal server error",
		}
	}

	return handler.Success{
		"text": "successfully logged out",
	}
}
