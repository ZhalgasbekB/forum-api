package model

import "time"

type Nothification struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	TypeNoth          string    `json:"type"`
	PostID            int       `json:"post_id"`
	CreatedNothUserID int       `json:"noth_user_id"`
	Message           string    `json:"message"`
	ISRead            bool      `json:"is_read"`
	CreatedAt         time.Time `json:"create_at"`
}

type NothificationDTO struct {
	ID int `json:"id"`
}

type NothificationDTO1 struct {
	ID int `json:"id"`
}

type NothificationDTO2 struct {
	ID int `json:"id"`
}
