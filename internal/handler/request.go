package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var (
	regexpAnnotationLength = regexp.MustCompile(`max_length\(([[:digit:]]+)\)`)
)

type Request struct {
	Data   *http.Request
	Writer http.ResponseWriter
}

func (request *Request) ParseBody(dst interface{}) error {
	if err := json.NewDecoder(request.Data.Body).Decode(dst); err != nil {
		return errors.New("invalid json structure provided")
	}

	fields := reflect.ValueOf(dst).Elem()
	for i := 0; i < fields.NumField(); i++ {
		tags := fields.Type().Field(i).Tag.Get("clavis")
		for _, tag := range strings.Split(tags, " ") {
			if tag == "required" && fields.Field(i).IsZero() {
				field := strings.ToLower(fields.Type().Field(i).Name)
				output := fmt.Sprintf("required field '%s' is missing", field)
				return errors.New(output)
			} else if fields.Field(i).Type().String() == "string" {
				if matches := regexpAnnotationLength.FindStringSubmatch(tag); len(matches) != 0 {
					number, _ := strconv.Atoi(matches[1])
					if len(fields.Field(i).String()) > number {
						field := strings.ToLower(fields.Type().Field(i).Name)
						output := fmt.Sprintf("field '%s' is too long", field)
						return errors.New(output)
					}
				}
			}
		}
	}

	return nil
}

func (request *Request) GetQueryVar(key string) []string {
	array, _ := request.Data.URL.Query()[key]
	return array
}

func (request *Request) GetMuxVar(key string) string {
	return mux.Vars(request.Data)[key]
}

func (request *Request) GetSession() (Session, error) {
	return Cookies(request)
}
