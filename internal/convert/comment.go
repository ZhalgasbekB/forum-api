package convert

import (
	"encoding/json"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
)

func CreateCommentConvert(r *http.Request, userID int) (*model.Comment, error) {
	var createComment model.Comment
	if err := json.NewDecoder(r.Body).Decode(&createComment); err != nil {
		return nil, err
	}
	return &model.Comment{
		User:        userID,
		Post:        createComment.Post,
		Description: createComment.Description,
	}, nil
}

func UpdateCommentConvert(r *http.Request, user_id int) (*model.Comment, error) {
	var updateComment model.CommentUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&updateComment); err != nil {
		return nil, err
	}
	return &model.Comment{
		ID:          updateComment.ID,
		User:        user_id,
		Post:        updateComment.Post,
		Description: updateComment.Description,
	}, nil
}

func DeleteCommentConvert(r *http.Request, user_id int) (*model.Comment, error) {
	var deleteComment model.CommentDeleteDTO
	if err := json.NewDecoder(r.Body).Decode(&deleteComment); err != nil {
		return nil, err
	}
	return &model.Comment{
		ID:   deleteComment.ID,
		User: user_id,
		Post: deleteComment.Post,
	}, nil
}
