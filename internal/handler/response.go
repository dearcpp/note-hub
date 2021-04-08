package handler

import (
	"encoding/json"
	"net/http"
)

type Response interface {
	Result() interface{}
	StatusCode() int
}

type (
	Success             map[string]interface{}
	BadRequest          map[string]interface{}
	Unauthorized        map[string]interface{}
	InternalServerError map[string]interface{}
)

func (r Success) Result() interface{} {
	return map[string]interface{}{
		"result": r,
	}
}

func (r Success) StatusCode() int {
	return http.StatusOK
}

func (r BadRequest) Result() interface{} {
	return map[string]interface{}{
		"error": r,
	}
}

func (r BadRequest) StatusCode() int {
	return http.StatusBadRequest
}

func (r Unauthorized) Result() interface{} {
	return map[string]interface{}{
		"error": r,
	}
}

func (r Unauthorized) StatusCode() int {
	return http.StatusUnauthorized
}

func (r InternalServerError) Result() interface{} {
	return map[string]interface{}{
		"error": r,
	}
}

func (r InternalServerError) StatusCode() int {
	return http.StatusInternalServerError
}

func WriteResponse(writer http.ResponseWriter, response Response) {
	writer.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(response.Result())

	if err != nil {
		bytes = []byte("{ \"error\": { \"text\": \"internal server error\" } }")
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(response.StatusCode())
	}

	writer.Write(bytes)
}
