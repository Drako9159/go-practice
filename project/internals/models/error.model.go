package models

import (
	"fmt"
	"net/http"


)

type AppError struct {
	Err error
	Message string
	Code int
}

func (e *AppError) String() string {
	return fmt.Sprintf("error: %s, (code: %d)", e.Message, e.Code)
}

const (
	ErrBadRequest = http.StatusBadRequest
	ErrUnauthorized = http.StatusUnauthorized
	ErrNotFound = http.StatusNotFound
	ErrForbidden = http.StatusForbidden
	ErrInternalServerError = http.StatusInternalServerError
	ErrUnprocessableEntity = http.StatusUnprocessableEntity
	ErrTooManyRequests = http.StatusTooManyRequests
	ErrServiceUnavailable = http.StatusServiceUnavailable
	ErrMethodNotAllowed = http.StatusMethodNotAllowed
	ErrConflict = http.StatusConflict
)

