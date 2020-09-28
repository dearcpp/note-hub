package api

import (
	"encoding/json"
	"github.com/beryll1um/note-hub/src/database"
	"github.com/beryll1um/note-hub/src/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func hNoteCreate(request *http.Request) Response {
	var body map[string]interface{}
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		return BadRequest{"text": "failed to parse request body"}
	}

	if _, ok := body["name"]; !ok {
		return BadRequest{"text": "name not specified"}
	}

	if _, ok := body["content"]; !ok {
		return BadRequest{"text": "content not specified"}
	}

	if _, ok := body["name"].(string); !ok {
		return BadRequest{"text": "name is not a string"}
	}

	if _, ok := body["content"].(string); !ok {
		return BadRequest{"text": "content is not a string"}
	}

	note := models.Note{
		Name:    body["name"].(string),
		Content: body["content"].(string),
	}

	if _, err := database.Controller.Insert(&note); err != nil {
		return BadRequest{"text": "bad parameters provided"}
	}

	return Success{
		"note": note,
	}
}

func hNoteGet(request *http.Request) Response {
	var result int64
	var err error

	if result, err = strconv.ParseInt(mux.Vars(request)["id"], 10, 64); err != nil {
		return BadRequest{"text": "bad parameters provided"}
	}

	var has bool
	note := models.Note{ID: result}
	if has, err = database.Controller.Get(&note); !has || err != nil {
		return BadRequest{"text": "note not found"}
	}

	return Success{
		"note": note,
	}
}
