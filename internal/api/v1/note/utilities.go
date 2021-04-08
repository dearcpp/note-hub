package note

import (
	"errors"
	"strconv"

	"github.com/beryll1um/note-hub/internal/handler"
	"github.com/gorilla/mux"
)

func DecodeBodyNote(request *handler.Request) (string, string, error) {
	var body map[string]interface{}
	var ok bool

	if err := request.DecodeBodyJson(&body); err != nil {
		return "", "", errors.New("failed to parse request body")
	}

	var name string

	if name, ok = body["name"].(string); !ok {
		return "", "", errors.New("name string is not found")
	}

	var content string

	if content, ok = body["content"].(string); !ok {
		return "", "", errors.New("content string is not found")
	}

	return name, content, nil
}

func ParseGetVarInt64(request *handler.Request, name string) (int64, error) {
	var valueArrayString []string
	var ok bool

	if valueArrayString, ok = request.Data.URL.Query()[name]; !ok {
		return 0, errors.New("value is not found")
	}

	var valueInt int64
	var err error

	if valueInt, err = strconv.ParseInt(valueArrayString[0], 10, 64); err != nil {
		return 0, err
	}

	return valueInt, nil
}

func ParseMuxVarInt64(request *handler.Request, name string) (int64, error) {
	return strconv.ParseInt(mux.Vars(request.Data)[name], 10, 64)
}
