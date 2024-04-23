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
