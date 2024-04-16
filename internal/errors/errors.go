package errors

import (
	"database/sql"
	"errors"
	"net/http"

	"gitea.com/lzhuk/forum/internal/helpers/json"
)

type ErrorCustom struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	ErrSQLNoRows    = sql.ErrNoRows
	ErrNotFoundData = errors.New("Not Found Any Data")

	ErrInvalidPassword    = errors.New("Invalid Password")
	ErrHaveDuplicateName  = errors.New("Name already exist")
	ErrHaveDuplicateEmail = errors.New("Email already exist")
	ErrInvalidCredentials = errors.New("Invalid Credentials")
)

func ErrorSend(w http.ResponseWriter, status int, message string) {
	custom := &ErrorCustom{
		Status:  status,
		Message: message,
	}
	json.WriteJSON(w, custom.Status, custom)
}
