package convert

import (
	"encoding/json"
	"net/http"

	"gitea.com/lzhuk/forum/internal/helpers/roles"
	"gitea.com/lzhuk/forum/internal/model"
)

func UpdateRole(r *http.Request) (*model.User, error) {
	var role string
	uRole := &model.RoleDTO{}
	if err := json.NewDecoder(r.Body).Decode(uRole); err != nil {
		return nil, err
	}
	if uRole.Role == roles.MODERATOR {
		role = uRole.Role
	} else {
		role = roles.USER
	}
	return &model.User{
		ID:   uRole.UserID,
		Role: role,
	}, nil
}

func DeleteUser(r *http.Request) (int, error) {
	user := &model.UserDeleteDTO{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return -1, err
	}
	return user.UserID, nil
}

func UpdateUserAdmin(r *http.Request) (*model.User, error) {
	userUpdate := &model.UserUpdateDTO{}
	if err := json.NewDecoder(r.Body).Decode(userUpdate); err != nil {
		return nil, nil
	}
	return &model.User{
		ID:    userUpdate.UserID,
		Name:  userUpdate.Name,
		Email: userUpdate.Email,
	}, nil
}
