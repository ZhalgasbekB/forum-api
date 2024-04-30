package convert

import (
	"encoding/json"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
)

func LikePostConvertor(r *http.Request, userID int) (*model.LikePost, int, error) {
	createLike := &model.LikePostDTO{}
	if err := json.NewDecoder(r.Body).Decode(createLike); err != nil {
		return nil, -1, err
	}
	return &model.LikePost{
		UserId:     userID,
		PostId:     createLike.PostId,
		LikeStatus: createLike.LikeStatus,
	}, createLike.UserId, nil
}

func LikeCommentConvertor(r *http.Request, userID int) (*model.LikeComment, error) {
	createLike := &model.LikeCommentDTO{}
	if err := json.NewDecoder(r.Body).Decode(createLike); err != nil {
		return nil, err
	}
	return &model.LikeComment{
		UserId:     userID,
		CommentId:  createLike.CommentId,
		LikeStatus: createLike.LikeStatus,
	}, nil
}
