package server

import (
	"encoding/json"
	"strconv"

	"github.com/beryll1um/note-hub/internal/api"
	"github.com/beryll1um/note-hub/internal/database"
	"github.com/beryll1um/note-hub/internal/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func hNoteCreate(request api.Request) api.Response {
	var body map[string]interface{}

	if err := json.NewDecoder(request.Data.Body).Decode(&body); err != nil {
		return api.BadRequest{"text": "failed to parse request body"}
	}

	var ok bool
	if _, ok = body["name"]; !ok {
		return api.BadRequest{"text": "name not specified"}
	}

	if _, ok = body["content"]; !ok {
		return api.BadRequest{"text": "content not specified"}
	}

	var name string

	if name, ok = body["name"].(string); !ok {
		return api.BadRequest{"text": "name is not a string"}
	}

	var content string

	if content, ok = body["content"].(string); !ok {
		return api.BadRequest{"text": "content is not a string"}
	}

	var session *sessions.Session
	var err error

	if session, err = api.GetCookieSession(request.Data); err != nil {
		return api.Unauthorized{
			"text": "internal server error",
		}
	}

	var user *models.User

	if user, err = api.GetUserFromSession(session); err != nil {
		return api.Unauthorized{
			"text": err,
		}
	}

	note := models.Note{
		Name:    name,
		Content: content,
		Author:  *user,
	}

	if _, err := database.Controller.Insert(&note); err != nil {
		return api.BadRequest{"text": "bad parameters provided"}
	}

	return api.Success{
		"note": note,
	}
}

func hNoteGet(request api.Request) api.Response {
	var result int64
	var err error

	if result, err = strconv.ParseInt(mux.Vars(request.Data)["id"], 10, 64); err != nil {
		return api.BadRequest{"text": "bad parameters provided"}
	}

	var session *sessions.Session

	if session, err = api.GetCookieSession(request.Data); err != nil {
		return api.Unauthorized{
			"text": "internal server error",
		}
	}

	var user *models.User

	if user, err = api.GetUserFromSession(session); err != nil {
		return api.Unauthorized{
			"text": err,
		}
	}

	var has bool

	note := models.Note{
		ID:     result,
		Author: *user,
	}

	if has, err = database.Controller.Get(&note); !has {
		return api.BadRequest{"text": "note not found"}
	}

	return api.Success{
		"note": note,
	}
}

func hNoteList(request api.Request) api.Response {
	var limitArray []string
	var ok bool

	if limitArray, ok = request.Data.URL.Query()["limit"]; !ok {
		return api.BadRequest{"text": "limit not specified"}
	}

	var startArray []string
	if startArray, ok = request.Data.URL.Query()["start"]; !ok {
		return api.BadRequest{"text": "start not specified"}
	}

	var limit int
	var err error

	if limit, err = strconv.Atoi(limitArray[0]); err != nil {
		return api.BadRequest{"text": "limit is not a string"}
	}

	var start int

	if start, err = strconv.Atoi(startArray[0]); err != nil {
		return api.BadRequest{"text": "start is not a string"}
	}

	var session *sessions.Session

	if session, err = api.GetCookieSession(request.Data); err != nil {
		return api.Unauthorized{
			"text": "internal server error",
		}
	}

	var user *models.User

	if user, err = api.GetUserFromSession(session); err != nil {
		return api.Unauthorized{
			"text": err,
		}
	}

	var notes []models.Note

	if err := database.Controller.Where("author = ?", user.ID).Limit(limit, start).Find(&notes); err != nil {
		return api.BadRequest{"text": "bad parameters provided"}
	}

	if notes == nil {
		return api.Success{
			"text": "notes not found",
		}
	}

	return api.Success{
		"notes": notes,
	}
}

func hNoteDelete(request api.Request) api.Response {
	var result int64
	var err error

	if result, err = strconv.ParseInt(mux.Vars(request.Data)["id"], 10, 64); err != nil {
		return api.BadRequest{"text": "bad parameters provided"}
	}

	var session *sessions.Session

	if session, err = api.GetCookieSession(request.Data); err != nil {
		return api.Unauthorized{
			"text": "internal server error",
		}
	}

	var user *models.User

	if user, err = api.GetUserFromSession(session); err != nil {
		return api.Unauthorized{
			"text": err,
		}
	}

	note := models.Note{
		ID:     result,
		Author: *user,
	}

	var affected int64
	if affected, err = database.Controller.Delete(&note); affected == 0 {
		return api.BadRequest{"text": "note not found"}
	}

	return api.Success{
		"text": "successfully deleted",
	}
}
