package user

import (
	"errors"

	"github.com/beryll1um/note-hub/internal/handler"
)

func DecodeBodyCredentials(request *handler.Request) (string, string, error) {
	var body map[string]interface{}
	var ok bool

	if err := request.DecodeBodyJson(&body); err != nil {
		return "", "", errors.New("failed to parse request body")
	}

	var email string

	if email, ok = body["email"].(string); !ok {
		return "", "", errors.New("email string is not found")
	}

	var password string

	if password, ok = body["password"].(string); !ok {
		return "", "", errors.New("password string is not found")
	}

	return email, password, nil
}
