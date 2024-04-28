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

func DeletePost(r *http.Request) (int, error) {
	post := &model.PostDeleteDTO{}
	if err := json.NewDecoder(r.Body).Decode(post); err != nil {
		return -1, err
	}
	return post.PostID, nil
}

func DeleteComment(r *http.Request) (int, error) {
	comment := &model.CommentDeleteAdminDTO{}
	if err := json.NewDecoder(r.Body).Decode(comment); err != nil {
		return -1, err
	}
	return comment.CommenID, nil
}

func AdminResponse(r *http.Request) (*model.AdminResponse, error) {
	resp := &model.AdminResponse{}
	if err := json.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, nil
	}
	return resp, nil
}
