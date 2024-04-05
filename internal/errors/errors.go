package errors

import (
	"database/sql"
	"errors"
)

type ErrorCustom struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	ErrSQLNoRows          = sql.ErrNoRows
	ErrHaveDuplicateEmail = errors.New("Email already exist")
	ErrSessionExpired     = errors.New("Time session expired")
	ErrInvalidCredentials = errors.New("Invalid Credentials")
)

func NewError(status int, message string) *ErrorCustom {
	return &ErrorCustom{
		Status:  status,
		Message: message,
	}
}