package convert

import (
	"encoding/json"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
)

func LikeConvertor(r *http.Request, userID int) (*model.LikePost, error) {
	createLike := &model.LikePostDTO{}
	if err := json.NewDecoder(r.Body).Decode(createLike); err != nil {
		return nil, err
	}
	return &model.LikePost{
		UserId:     userID,
		PostId:     createLike.PostId,
		LikeStatus: createLike.LikeStatus,
	}, nil
}
