package handler

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Data   *http.Request
	Writer http.ResponseWriter
}

func (r *Request) DecodeBodyJson(result *map[string]interface{}) error {
	if err := json.NewDecoder(r.Data.Body).Decode(&result); err != nil {
		return err
	}
	return nil
}
