package validation

import (
	"fmt"
	"regexp"
	"strings"
)

type Form struct {
	Errors map[string][]string
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func New() *Form {
	return &Form{make(map[string][]string)}
}

func (f *Form) EmailValid(email string) bool {
	return emailRegex.Match([]byte(email))
}

func (f *Form) EmptyFieldValid(val string) bool {
	return strings.TrimSpace(val) != ""
}

func (f *Form) MinLengthValid(val string, length int) bool {
	return len(val) > length
}

func (f *Form) MaxLengthValid(val string, length int) bool {
	return len(val) < length
}

var (
	emailRegex    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func (f *Form) CheckField(ok bool, field, err string) error {
	if !ok {
		f.Errors[field] = append(f.Errors[field], err)
		return fmt.Errorf(err)
	}
	return nil
} // ???
