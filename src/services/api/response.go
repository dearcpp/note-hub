package api

import "net/http"

type Response interface {
	Result() interface{}
	StatusCode() int
}

type (
	Success             map[string]interface{}
	BadRequest          map[string]interface{}
	InternalServerError string
	Unauthorized        map[string]interface{}
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

func (r InternalServerError) Result() interface{} {
	return r
}

func (r InternalServerError) StatusCode() int {
	return http.StatusInternalServerError
}

func (r Unauthorized) Result() interface{} {
	return map[string]interface{}{
		"error": r,
	}
}

func (r Unauthorized) StatusCode() int {
	return http.StatusUnauthorized
}
