package models

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Err AppError
	Message string
	Code int
}

func (e *AppError) String() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
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
