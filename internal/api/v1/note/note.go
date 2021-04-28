package note

import (
	"log"
	"note-hub/internal/handler"
	"note-hub/internal/model"
	"note-hub/internal/repository"
	"strconv"
)

type CreateParameters struct {
	Name    string `json:"name" clavis:"required max_length(64)"`
	Content string `json:"content" clavis:"required max_length(256)"`
}

func Create(request handler.Request) handler.Response {
	var params CreateParameters

	if err := request.ParseBody(&params); err != nil {
		return handler.BadRequest{"text": err.Error()}
	}

	note := repository.Note{
		Name:    params.Name,
		Content: params.Content,
		Author:  model.User(request.Get(handler.User).(repository.User)),
	}

	if _, err := note.Insert(); err != nil {
		return handler.BadRequest{"text": "internal server error"}
	}

	return handler.Success{
		"note": note,
	}
}

func Get(request handler.Request) handler.Response {
	id, err := strconv.ParseUint(request.GetMuxVar("id"), 10, 32)
	if err != nil {
		return handler.BadRequest{
			"text": "required uint32 'id' is missing",
		}
	}

	note := repository.Note{
		ID:     uint32(id),
		Author: model.User(request.Get(handler.User).(repository.User)),
	}

	if has, _ := note.Get(); !has {
		return handler.BadRequest{"text": "note not found"}
	}

	return handler.Success{
		"note": note,
	}
}

func List(request handler.Request) handler.Response {
	limitString, _ := request.GetQueryVar("limit")
	limit, err := strconv.ParseUint(limitString[0], 10, 32)
	if err != nil {
		log.Println(err)
		return handler.BadRequest{
			"text": "required uint32 'limit' is missing",
		}
	}

	startString, _ := request.GetQueryVar("start")
	start, err := strconv.ParseUint(startString[0], 10, 32)
	if err != nil {
		return handler.BadRequest{
			"text": "required uint32 'start' is missing",
		}
	}

	user := request.Get(handler.User).(repository.User)

	var notes []repository.Note

	if err = user.WhereIn().Limit(int(limit), int(start)).Find(&notes); err != nil {
		return handler.BadRequest{"text": "note not found"}
	}

	return handler.Success{
		"notes": notes,
	}
}

func Delete(request handler.Request) handler.Response {
	id, err := strconv.ParseUint(request.GetMuxVar("id"), 10, 32)
	if err != nil {
		return handler.BadRequest{
			"text": "required uint32 'id' is missing",
		}
	}

	note := repository.Note{
		ID:     uint32(id),
		Author: model.User(request.Get(handler.User).(repository.User)),
	}

	if has, _ := note.Get(); !has {
		return handler.BadRequest{"text": "note not found"}
	}

	if affected, _ := note.Delete(); affected == 0 {
		return handler.BadRequest{"text": "note not found"}
	}

	return handler.Success{
		"note": note,
	}
}
