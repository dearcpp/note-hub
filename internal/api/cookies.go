package api

import (
	"errors"
	"net/http"

	"github.com/beryll1um/note-hub/internal/database"
	"github.com/beryll1um/note-hub/internal/models"
	"github.com/gorilla/sessions"
)

// @todo: getting the TEST_ENV_KEY and cookie name from the environment

var (
	cookiesStore *sessions.CookieStore = sessions.NewCookieStore([]byte("TEST_ENV_KEY"))
)

func GetCookieSession(request *http.Request) (*sessions.Session, error) {
	return cookiesStore.Get(request, "nh-session")
}

func GetUserFromSession(session *sessions.Session) (*models.User, error) {
	mail, okMail := session.Values["mail"].(string)
	password, okPassword := session.Values["password"].(string)

	if !okMail || !okPassword {
		return nil, errors.New("user is not authorized")
	}

	user := models.User{
		Mail:     mail,
		Password: password,
	}

	if has, _ := database.Controller.Get(&user); !has {
		return nil, errors.New("user is not found")
	}

	return &user, nil
}
