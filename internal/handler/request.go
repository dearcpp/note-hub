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

func parsePostConditions(dst interface{}) error {
	fields := reflect.ValueOf(dst).Elem()
	for i := 0; i < fields.NumField(); i++ {
		tags := fields.Type().Field(i).Tag.Get("field")
		for _, tag := range strings.Split(tags, " ") {
			if tag == "required" && fields.Field(i).IsZero() {
				fieldName := strings.ToLower(fields.Type().Field(i).Name)
				output := fmt.Sprintf("required field '%s' is missing", fieldName)
				return errors.New(output)
			} else if fields.Field(i).Type().String() == "string" {
				if matches := regexpAnnotationLength.FindStringSubmatch(tag); len(matches) != 0 {
					number, _ := strconv.Atoi(matches[1])
					if len(fields.Field(i).String()) > number {
						fieldName := strings.ToLower(fields.Type().Field(i).Name)
						output := fmt.Sprintf("field '%s' is too long", fieldName)
						return errors.New(output)
					}
				}
			}
		}
	}
	return nil
}

func (request *Request) ParseQueryVars(dst interface{}) error {
	fields := reflect.ValueOf(dst).Elem()

	for key, value := range request.Data.URL.Query() {
		field := fields.FieldByName(strings.Title(strings.ToLower(key)))

		if !field.IsValid() {
			continue
		}

		switch typeName := field.Type().Name(); typeName {
		case "string":
			field.SetString(value[0])
		default:
			fieldName := strings.ToLower(field.Type().Name())
			output := fmt.Sprintf("field '%s' has undefined type", fieldName)
			return errors.New(output)
		}
	}

	return parsePostConditions(dst)
}

func (request *Request) ParseMuxVars(dst interface{}) error {
	fields := reflect.ValueOf(dst).Elem()

	for key, value := range mux.Vars(request.Data) {
		field := fields.FieldByName(strings.Title(strings.ToLower(key)))

		if !field.IsValid() {
			continue
		}

		switch typeName := field.Type().Name(); typeName {
		case "string":
			field.SetString(value)
		default:
			fieldName := strings.ToLower(field.Type().Name())
			output := fmt.Sprintf("field '%s' has undefined type", fieldName)
			return errors.New(output)
		}
	}

	return parsePostConditions(dst)
}

func (request *Request) ParseBody(dst interface{}) error {
	if err := json.NewDecoder(request.Data.Body).Decode(dst); err != nil {
		return errors.New("invalid json structure provided")
	}

	return parsePostConditions(dst)
}

func (request *Request) GetQueryVar(key string) ([]string, bool) {
	array, ok := request.Data.URL.Query()[key]
	return array, ok
}

func (request *Request) GetMuxVar(key string) string {
	return mux.Vars(request.Data)[key]
}

func (request *Request) GetSession() (Session, error) {
	return Cookies(request)
}
