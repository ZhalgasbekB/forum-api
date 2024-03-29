package convert

import (
	"encoding/json"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
)

func UserRegisterRequestBody(r *http.Request) (*model.User, error) {
	user := &model.UserRegisterDTO{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return nil, err
	}
	return &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  false,
	}, nil
}

func UserLoginRequestBody(r *http.Request) (*model.User, error) {
	user := &model.UserLoginDTO{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return nil, err
	}
	return &model.User{
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
