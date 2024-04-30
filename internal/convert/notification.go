package convert

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
)

const commentS = "comment"

func NotificationCreateLikes(like *model.LikePost, user_id int) *model.Notification {
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

func NotificationCreateComment(user_id int, comment *model.Comment) *model.Notification {
	return &model.Notification{
		UserID:        user_id,
		PostID:        comment.Post,
		Type:          commentS,
		CreatedUserID: comment.User,
		Message:       fmt.Sprintf("Yoy get notification from user: %d, type of notification: %s, on your post: %d.", comment.User, commentS, comment.Post),
	}
}

func NotificationUpdateComment(r *http.Request) (int, error) {
	n := &model.NotificationUpdateDTO{}
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		return -1, err
	}
	return n.ID, nil
}
