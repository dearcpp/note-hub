package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/beryll1um/note-hub/src/database"
	"github.com/beryll1um/note-hub/src/models"
)

func RequiresUser(callback func(Request, *models.User) Response) HandlerFunc {
	return func(request Request) Response {
		mail, okMail := request.Session.Values["mail"].(string)
		password, okPassword := request.Session.Values["password"].(string)

		if !okMail || !okPassword {
			return Unauthorized{
				"text": "you are not authorized",
			}
		}

		user := models.User{
			Mail:     mail,
			Password: password,
		}

		if has, _ := database.Controller.Get(&user); !has {
			return Unauthorized{
				"text": "session is out of date",
			}
		}

		return callback(request, &user)
	}
}

func hUserSignOut(request Request) Response {
	request.Session.Options.MaxAge = -1
	return Success{
		"text": "successfully signed out",
	}
}

func hUserSignIn(request Request) Response {
	var body map[string]interface{}
	if err := json.NewDecoder(request.Data.Body).Decode(&body); err != nil {
		return BadRequest{"text": "failed to parse request body"}
	}

	if _, ok := body["mail"]; !ok {
		return BadRequest{"text": "mail not specified"}
	}

	if _, ok := body["password"]; !ok {
		return BadRequest{"text": "password not specified"}
	}

	if _, ok := body["mail"].(string); !ok {
		return BadRequest{"text": "mail is not a string"}
	}

	if _, ok := body["password"].(string); !ok {
		return BadRequest{"text": "password is not a string"}
	}

	passwordHash := md5.Sum([]byte(body["password"].(string)))
	password := hex.EncodeToString(passwordHash[:])

	user := models.User{
		Mail:     body["mail"].(string),
		Password: password,
	}

	if has, _ := database.Controller.Get(&user); !has {
		return BadRequest{"text": "invalid credentials"}
	}

	request.Session.Values["mail"] = body["mail"].(string)
	request.Session.Values["password"] = password

	return Success{
		"text": "successfully signed in",
	}
}

func hUserSignUp(request Request) Response {
	var body map[string]interface{}
	if err := json.NewDecoder(request.Data.Body).Decode(&body); err != nil {
		return BadRequest{"text": "failed to parse request body"}
	}

	if _, ok := body["mail"]; !ok {
		return BadRequest{"text": "mail not specified"}
	}

	if _, ok := body["password"]; !ok {
		return BadRequest{"text": "password not specified"}
	}

	if _, ok := body["mail"].(string); !ok {
		return BadRequest{"text": "mail is not a string"}
	}

	if _, ok := body["password"].(string); !ok {
		return BadRequest{"text": "password is not a string"}
	}

	passwordHash := md5.Sum([]byte(body["password"].(string)))

	user := models.User{
		Mail:     body["mail"].(string),
		Password: hex.EncodeToString(passwordHash[:]),
	}

	if affected, _ := database.Controller.Insert(user); affected == 0 {
		return BadRequest{"text": "failed to signed up"}
	}

	return Success{
		"text": "successfully signed up",
	}
}
