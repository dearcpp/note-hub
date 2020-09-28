package api

import "net/http"

type Response interface {
	Result() map[string]interface{}
	StatusCode() int
}

type (
	Success    map[string]interface{}
	BadRequest map[string]interface{}
)

func (r Success) Result() map[string]interface{} {
	return map[string]interface{}{
		"result": r,
	}
}

func (r Success) StatusCode() int {
	return http.StatusOK
}

func (r BadRequest) Result() map[string]interface{} {
	return map[string]interface{}{
		"error": r,
	}
}

func (r BadRequest) StatusCode() int {
	return http.StatusBadRequest
}
