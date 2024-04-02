package model

import "time"

type Post struct {
	PostId       int       `json:"post_id"`
	UserId       int       `json:"user_id"`
	CategoryName string    `json:"category_name"`
	Title        string    `json:"title"`
	Discription  string    `json:"discription"`
	CreateDate   time.Time `json:"create_at"`
}

type CreatePostDTO struct {
	CategoryName string //`json:"category_name"`
	Title        string `json:"title"`
	Discription  string `json:"discription"`
}

type UpdatePostDTO struct {
	Title       string `json:"title"`
	Discription string `json:"discription"`
}

type PostCommentsDTO struct {
	Post    *Post      `json:"post"`
	Comments []*Comment `json:"comments"`
}
