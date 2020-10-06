package api

import (
	"encoding/json"
	"github.com/beryll1um/note-hub/src/database"
	"github.com/beryll1um/note-hub/src/models"
	"github.com/gorilla/mux"
	"strconv"
)

func hNoteCreate(request Request, user *models.User) Response {
	var body map[string]interface{}
	if err := json.NewDecoder(request.Data.Body).Decode(&body); err != nil {
		return BadRequest{"text": "failed to parse request body"}
	}

	var ok bool
	if _, ok = body["name"]; !ok {
		return BadRequest{"text": "name not specified"}
	}

	if _, ok = body["content"]; !ok {
		return BadRequest{"text": "content not specified"}
	}

	var name string
	if name, ok = body["name"].(string); !ok {
		return BadRequest{"text": "name is not a string"}
	}

	var content string
	if content, ok = body["content"].(string); !ok {
		return BadRequest{"text": "content is not a string"}
	}

	note := models.Note{
		Name:    name,
		Content: content,
		Author:  *user,
	}

	if _, err := database.Controller.Insert(&note); err != nil {
		return BadRequest{"text": "bad parameters provided"}
	}

	return Success{
		"note": note,
	}
}

func hNoteGet(request Request, user *models.User) Response {
	var result int64
	var err error

	if result, err = strconv.ParseInt(mux.Vars(request.Data)["id"], 10, 64); err != nil {
		return BadRequest{"text": "bad parameters provided"}
	}

	var has bool
	note := models.Note{
		ID:     result,
		Author: *user,
	}

	if has, err = database.Controller.Get(&note); !has {
		return BadRequest{"text": "note not found"}
	}

	return Success{
		"note": note,
	}
}

func hNoteList(request Request, user *models.User) Response {
	var ok bool
	var limitArray []string
	if limitArray, ok = request.Data.URL.Query()["limit"]; !ok {
		return BadRequest{"text": "limit not specified"}
	}

	var startArray []string
	if startArray, ok = request.Data.URL.Query()["start"]; !ok {
		return BadRequest{"text": "start not specified"}
	}

	var err error
	var limit int
	if limit, err = strconv.Atoi(limitArray[0]); err != nil {
		return BadRequest{"text": "limit is not a string"}
	}

	var start int
	if start, err = strconv.Atoi(startArray[0]); err != nil {
		return BadRequest{"text": "start is not a string"}
	}

	var notes []models.Note
	if err := database.Controller.Where("author = ?", user.ID).Limit(limit, start).Find(&notes); err != nil {
		return BadRequest{"text": "bad parameters provided"}
	}

	if notes == nil {
		return Success{
			"text": "notes not found",
		}
	}

	return Success{
		"notes": notes,
	}
}

func hNoteDelete(request Request, user *models.User) Response {
	var result int64
	var err error

	if result, err = strconv.ParseInt(mux.Vars(request.Data)["id"], 10, 64); err != nil {
		return BadRequest{"text": "bad parameters provided"}
	}

	note := models.Note{
		ID:     result,
		Author: *user,
	}

	var affected int64
	if affected, err = database.Controller.Delete(&note); affected == 0 {
		return BadRequest{"text": "note not found"}
	}

	return Success{
		"text": "successfully deleted",
	}
}
