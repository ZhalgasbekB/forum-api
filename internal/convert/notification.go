package convert

import (
	"fmt"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
)

func NothificationCreateLikes(r *http.Request, like *model.LikePost, user_id int) *model.Notification {
	typeN := "dislike"
	if like.LikeStatus {
		typeN = "like"
	}

	return &model.Notification{
		UserID:        user_id,
		PostID:        like.PostId,
		Type:          typeN,
		CreatedUserID: like.UserId,
		Message:       fmt.Sprintf("Yoy get notification from user: %d, type of notification: %s, on your post: %d.", like.UserId, typeN, like.PostId),
	}
}
