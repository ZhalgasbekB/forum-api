package model

import "time"

type UserDeleteDTO struct {
	UserID int `json:"user_id"`
}

type PostDeleteDTO struct {
	PostID int `json:"post_id"`
}

type CommentDeleteAdminDTO struct {
	CommenID int `json:"comment_id"`
}

type CategoryDTO struct {
	CategoryName string `json:"category"`
}

type RoleDTO struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}

type WantsDTO struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type Wants1DTO struct {
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
type Wants2DTO struct {
	UserID   int       `json:"user_id"`
	UserName string    `json:"user_name"`
	CreateAt time.Time `json:"create_at"`
}

type AdminResponse struct {
	UserID int `json:"user_id"`
	Status int `json:"number"`
}
