package note

import (
	"log"
	"note-hub/internal/handler"
	"note-hub/internal/model"
	"note-hub/internal/repository"
	"strconv"
)

type CreateParameters struct {
	Name    string `json:"name" field:"required max_length(64)"`
	Content string `json:"content" field:"required max_length(256)"`
}

func Create(request handler.Request) handler.Response {
	var params CreateParameters

	if err := request.ParseBody(&params); err != nil {
		return handler.BadRequest{"text": err.Error()}
	}

	note := repository.Note{
		Name:    params.Name,
		Content: params.Content,
		Author:  model.User(request.GetContextValue(handler.User).(repository.User)),
	}

	if _, err := note.Insert(); err != nil {
		return handler.BadRequest{"text": "internal server error"}
	}

	return handler.Success{
		"note": note,
	}
}

type GetParameters struct {
	Id string `field:"required max_length(10)"`
}

func Get(request handler.Request) handler.Response {
	var params GetParameters

	if err := request.ParseMuxVars(&params); err != nil {
		return handler.BadRequest{
			"text": err.Error(),
		}
	}

	id, err := strconv.ParseUint(params.Id, 10, 32)
	if err != nil {
		return handler.BadRequest{
			"text": "required uint32 'id' is missing",
		}
	}

	note := repository.Note{
		ID:     uint32(id),
		Author: model.User(request.GetContextValue(handler.User).(repository.User)),
	}

	if has, _ := note.Get(); !has {
		return handler.BadRequest{"text": "note not found"}
	}

	return handler.Success{
		"note": note,
	}
}

type ListParameters struct {
	Limit string `field:"required max_length(10)"`
	Start string `field:"required max_length(10)"`
}

func List(request handler.Request) handler.Response {
	var params ListParameters

	if err := request.ParseQueryVars(&params); err != nil {
		return handler.BadRequest{
			"text": err.Error(),
		}
	}

	limit, err := strconv.ParseUint(params.Limit, 10, 32)
	if err != nil {
		log.Println(err)
		return handler.BadRequest{
			"text": "required uint32 'limit' is missing",
		}
	}

	start, err := strconv.ParseUint(params.Start, 10, 32)
	if err != nil {
		return handler.BadRequest{
			"text": "required uint32 'start' is missing",
		}
	}

	user := request.GetContextValue(handler.User).(repository.User)

	var notes []repository.Note

	if err = user.WhereIn().Limit(int(limit), int(start)).Find(&notes); err != nil {
		return handler.BadRequest{"text": "note not found"}
	}

	return handler.Success{
		"notes": notes,
	}
}

type DeleteParameters struct {
	Id string `field:"required max_length(10)"`
}

func Delete(request handler.Request) handler.Response {
	var params DeleteParameters

	if err := request.ParseMuxVars(&params); err != nil {
		return handler.BadRequest{
			"text": err.Error(),
		}
	}

	id, err := strconv.ParseUint(params.Id, 10, 32)
	if err != nil {
		return handler.BadRequest{
			"text": "required uint32 'id' is missing",
		}
	}

	note := repository.Note{
		ID:     uint32(id),
		Author: model.User(request.GetContextValue(handler.User).(repository.User)),
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
