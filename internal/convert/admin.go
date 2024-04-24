package convert

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(user.UserID)
		return -1, err
	}
	return user.UserID, nil
}
