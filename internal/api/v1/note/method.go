package note

import (
	"github.com/beryll1um/note-hub/internal/handler"
	"github.com/beryll1um/note-hub/internal/repository"
)

func Create(request handler.Request) handler.Response {
	var user repository.User

	if err := user.GetFromSession(request.Data); err != nil {
		return handler.InternalServerError{"text": "internal server error"}
	}

	var err error
	var content string
	var name string

	if name, content, err = DecodeBodyNote(&request); err != nil {
		return handler.BadRequest{"text": err.Error()}
	}

	note := repository.Note{
		Name:    name,
		Content: content,
		Author:  user.Model(),
	}

	if _, err := note.Insert(); err != nil {
		return handler.BadRequest{"text": "internal server error"}
	}

	return handler.Success{
		"note": note,
	}
}

func Get(request handler.Request) handler.Response {
	var id int64
	var err error

	if id, err = ParseMuxVarInt64(&request, "id"); err != nil {
		return handler.BadRequest{"text": "id number is not found"}
	}

	note := repository.Note{
		ID: id,
	}

	if has, _ := note.Get(); !has {
		if err != nil {
			return handler.InternalServerError{"text": "internal server error"}
		} else {
			return handler.BadRequest{"text": "note not found"}
		}
	}

	return handler.Success{
		"note": note,
	}
}

func List(request handler.Request) handler.Response {
	var user repository.User
	var err error

	if err = user.GetFromSession(request.Data); err != nil {
		return handler.InternalServerError{"text": "internal server error"}
	}

	var limit int64

	if limit, err = ParseGetVarInt64(&request, "limit"); err != nil {
		return handler.BadRequest{"text": "start number is not found"}
	}

	var start int64

	if start, err = ParseGetVarInt64(&request, "start"); err != nil {
		return handler.BadRequest{"text": "start number is not found"}
	}

	var notes []repository.Note

	if err = user.WhereIn().Limit(int(limit), int(start)).Find(&notes); err != nil {
		return handler.InternalServerError{"text": "internal server error"}
	}

	return handler.Success{
		"notes": notes,
	}
}

func Delete(request handler.Request) handler.Response {
	var id int64
	var err error

	if id, err = ParseMuxVarInt64(&request, "id"); err != nil {
		return handler.BadRequest{"text": "id number is not found"}
	}

	note := repository.Note{
		ID: id,
	}

	var has bool

	if has, err = note.Get(); !has {
		if err != nil {
			return handler.InternalServerError{"text": "internal server error"}
		} else {
			return handler.BadRequest{"text": "note not found"}
		}
	}

	var affected int64

	if affected, err = note.Delete(); affected == 0 {
		if err != nil {
			return handler.InternalServerError{"text": "internal server error"}
		} else {
			return handler.BadRequest{"text": "note not found"}
		}
	}

	return handler.Success{
		"note": note,
	}
}
