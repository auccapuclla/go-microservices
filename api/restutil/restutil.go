package restutil

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrEmptyBody    = errors.New("body can't be empty")
	ErrUnauthorized = errors.New("Unauthorized")
)

type JError struct {
	Error string `json:"error"`
}

func WriteAsJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	e := JError{Error: err.Error()}
	if err != nil {
		e.Error = err.Error()
	}
	WriteAsJson(w, statusCode, e)

}
