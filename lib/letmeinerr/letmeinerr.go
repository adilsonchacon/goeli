package letmeinerr

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrNotFound            = errors.New("not found")
	ErrForbidden           = errors.New("forbidden")
	ErrBadRequest          = errors.New("bad request")
	ErrGeneral             = errors.New("general error")
)

type LetmeinError struct {
	MainError  error
	Body       []byte
	StatusCode int
}

func (e *LetmeinError) Error() string {
	return fmt.Sprintf("Letmein Error: %s", e.MainError)
}

func New(statusCode int, body []byte) *LetmeinError {
	var err error
	switch statusCode {
	case http.StatusUnprocessableEntity:
		err = ErrUnprocessableEntity
	case http.StatusNotFound:
		err = ErrNotFound
	case http.StatusForbidden:
		err = ErrForbidden
	case http.StatusBadRequest:
		err = ErrBadRequest
	default:
		err = ErrGeneral
	}

	return &LetmeinError{StatusCode: statusCode, Body: body, MainError: err}
}
