package server

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/beryll1um/note-hub/internal/api"
	"github.com/beryll1um/note-hub/internal/database"
	"github.com/beryll1um/note-hub/internal/models"
)

func hUserRegister(request api.Request) api.Response {
	var body map[string]interface{}
	if err := json.NewDecoder(request.Data.Body).Decode(&body); err != nil {
		return api.BadRequest{"text": "failed to parse request body"}
	}

	if _, ok := body["mail"]; !ok {
		return api.BadRequest{"text": "mail not specified"}
	}

	if _, ok := body["password"]; !ok {
		return api.BadRequest{"text": "password not specified"}
	}

	if _, ok := body["mail"].(string); !ok {
		return api.BadRequest{"text": "mail is not a string"}
	}

	if _, ok := body["password"].(string); !ok {
		return api.BadRequest{"text": "password is not a string"}
	}

	passwordHash := md5.Sum([]byte(body["password"].(string)))

	user := models.User{
		Mail:     body["mail"].(string),
		Password: hex.EncodeToString(passwordHash[:]),
	}

	if affected, _ := database.Controller.Insert(user); affected == 0 {
		return api.BadRequest{"text": "failed to signed up"}
	}

	return api.Success{
		"text": "successfully registered",
	}
}

func hUserLogin(request api.Request) api.Response {
	var body map[string]interface{}
	if err := json.NewDecoder(request.Data.Body).Decode(&body); err != nil {
		return api.BadRequest{"text": "failed to parse request body"}
	}

	if _, ok := body["mail"]; !ok {
		return api.BadRequest{"text": "mail not specified"}
	}

	if _, ok := body["password"]; !ok {
		return api.BadRequest{"text": "password not specified"}
	}

	if _, ok := body["mail"].(string); !ok {
		return api.BadRequest{"text": "mail is not a string"}
	}

	if _, ok := body["password"].(string); !ok {
		return api.BadRequest{"text": "password is not a string"}
	}

	passwordHash := md5.Sum([]byte(body["password"].(string)))
	password := hex.EncodeToString(passwordHash[:])

	user := models.User{
		Mail:     body["mail"].(string),
		Password: password,
	}

	if has, _ := database.Controller.Get(&user); !has {
		return api.BadRequest{"text": "invalid credentials"}
	}

	session, err := api.GetCookieSession(request.Data)
	if err != nil {
		return api.Unauthorized{
			"text": "internal server error",
		}
	}

	session.Values["mail"] = body["mail"].(string)
	session.Values["password"] = password

	session.Save(request.Data, *request.Writer)

	return api.Success{
		"text": "successfully signed in",
	}
}

func hUserLogout(request api.Request) api.Response {
	session, err := api.GetCookieSession(request.Data)
	if err != nil {
		return api.Unauthorized{
			"text": "internal server error",
		}
	}

	session.Options.MaxAge = -1
	session.Save(request.Data, *request.Writer)

	return api.Success{
		"text": "successfully signed out",
	}
}
