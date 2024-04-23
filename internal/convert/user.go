package convert

import (
	"encoding/json"
	"gitea.com/lzhuk/forum/internal/helpers/roles"
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
		Role:     roles.USER,
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

func AuthenticateUserDTO(r *http.Request) (*model.User, error) {
	user := &model.UserAuthDTO{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return nil, err
	}
	return &model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
