package error

import "errors"

var (
	ErrUniqueValid = errors.New("duplicate key value violates unique constraint")
	ErrNoRowsDB    = errors.New("doesn't not exist row in db")
	ErrEmailValid  = errors.New("not valid email for this field")
	ErrEmptyField  = errors.New("empty field need write something")
	ErrPassword    = errors.New("Not valid password must have a 8 characrters")
)
