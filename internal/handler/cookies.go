package handler

import (
	"errors"
	"net"
	"note-hub/internal/repository"

	"github.com/gorilla/sessions"
)

type Session struct {
	*sessions.Session
	request *Request
}

var (
	cookies *sessions.CookieStore
)

func SetupCookieStore(key string) {
	cookies = sessions.NewCookieStore([]byte(key))
}

func Cookies(request *Request) (Session, error) {
	session, err := cookies.Get(request.Data, "note-session")
	return Session{
		session,
		request,
	}, err
}

func (session *Session) SetUser(user *repository.User) error {
	address, _, err := net.SplitHostPort(session.request.Data.RemoteAddr)
	if err != nil {
		return err
	}

	session.Values["email"] = user.Email
	session.Values["password"] = user.Password
	session.Values["address"] = address

	if err = session.Save(session.request.Data, session.request.Writer); err != nil {
		return err
	}

	return nil
}

func (session *Session) GetUser(user *repository.User) error {
	address, _, err := net.SplitHostPort(session.request.Data.RemoteAddr)
	if err != nil {
		return err
	}

	var ok bool

	user.Email, ok = session.Values["email"].(string)
	if !ok {
		return errors.New("string 'email' is not found")
	}

	user.Password, ok = session.Values["password"].(string)
	if !ok {
		return errors.New("string 'password' is not found")
	}

	var sessionAddress string

	sessionAddress, ok = session.Values["address"].(string)
	if !ok {
		return errors.New("string 'address' is not found")
	}

	if sessionAddress != address {
		return errors.New("session has not verified")
	}

	if has, _ := user.Get(); !has {
		return errors.New("session has not verified")
	}

	return nil
}

func (session *Session) Destroy() error {
	session.Options.MaxAge = -1
	return session.Save(session.request.Data, session.request.Writer)
}
