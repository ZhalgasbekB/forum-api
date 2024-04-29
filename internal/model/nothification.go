package model

import "time"

type Nothification struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	PostID        int       `json:"post_id"`
	Type          string    `json:"type"`
	CreatedUserID int       `json:"create_user_id"`
	Message       string    `json:"message"`
	IsRead        bool      `json:"is_read"`
	CreatedAt     time.Time `json:"created_at"`
}

type NothificationCreateDTO struct {
	UserId        int    `json:"user_id"`
	PostId        int    `json:"post_id"`
	Type          string `json:"type"`
	CreatedUserId int    `json:"created_user_id"`
}

type NothificationDTO2 struct {
	ID int `json:"id"`
}
