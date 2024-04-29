package convert

import (
	"encoding/json"
	"fmt"
	"gitea.com/lzhuk/forum/internal/model"
	"net/http"
)

func NotificationCreate(r *http.Request) (*model.Notification, error) {
	notification := &model.NotificationCreateDTO{}
	if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
		return nil, err
	}

	return &model.Notification{
		UserID:        notification.UserId,
		PostID:        notification.PostId,
		Type:          notification.Type,
		CreatedUserID: notification.CreatedUserId,
		Message:       fmt.Sprintf("Yoy get notification from user: %d, type of notification: %s, on your post: %d.", notification.UserId, notification.Type, notification.PostId),
	}, nil
}
