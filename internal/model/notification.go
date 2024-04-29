package model

import "time"

type Notification struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	PostID        int       `json:"post_id"`
	Type          string    `json:"type"`
	CreatedUserID int       `json:"create_user_id"`
	Message       string    `json:"message"`
	IsRead        bool      `json:"is_read"`
	CreatedAt     time.Time `json:"created_at"`
}

type NotificationCreateDTO struct {
	UserId        int    `json:"user_id"`
	PostId        int    `json:"post_id"`
	Type          string `json:"type"`
	CreatedUserId int    `json:"created_user_id"`
}

type NotificationDTO2 struct {
	ID int `json:"id"`
}
