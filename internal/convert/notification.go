package convert

import (
	"encoding/json"
	"fmt"
	"gitea.com/lzhuk/forum/internal/model"
	"net/http"
)

func NothificationCreate(r *http.Request) (*model.Notification, error) {
	nothification := &model.NotificationCreateDTO{}
	if err := json.NewDecoder(r.Body).Decode(&nothification); err != nil {
		return nil, err
	}

	return &model.Notification{
		UserID:        nothification.UserId,
		PostID:        nothification.PostId,
		Type:          nothification.Type,
		CreatedUserID: nothification.CreatedUserId,
		Message:       fmt.Sprintf("Yoy get notification from user: %d, type of notification: %s, on your post: %d."),
	}, nil
}
