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
	errSQLNoRows          = sql.ErrNoRows
	errHaveDuplicateEmail = errors.New("Email already exist")
	errSessionExpired     = errors.New("Time session expired")
)

func NewError(status int, message string) *ErrorCustom {
	return &ErrorCustom{
		Status:  status,
		Message: message,
	}
}
////////////////////