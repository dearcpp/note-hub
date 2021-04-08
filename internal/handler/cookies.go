package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	cookiesStore *sessions.CookieStore
)

func SetupCookieStore(keyFile string) error {
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}
	cookiesStore = sessions.NewCookieStore(key)
	return nil
}

func GetCookieSession(request *http.Request) (*sessions.Session, error) {
	return cookiesStore.Get(request, "nh-session")
}
