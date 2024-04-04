package model

import "time"

type Comment struct {
	ID          int       `json:"id"`
	User        int       `json:"user_id"`
	Post        int       `json:"post_id"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"created_at"`
	UpdatedDate time.Time `json:"updated_at"`

	Name    string `json:"name"`
	Like    int    `json:"likes"`
	Dislike int    `json:"dislikes"`
}

type CreateCommentDTO struct {
	Post        int    `json:"post_id"`
	Description string `json:"description"`
}

type CommentUpdateDTO struct {
	ID          int    `json:"id"`
	Post        int    `json:"post_id"`
	Description string `json:"description"`
}

type CommentDeleteDTO struct {
	ID   int `json:"id"`
	Post int `json:"post_id"`
}
