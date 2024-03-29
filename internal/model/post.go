package model

import "time"

type CreatePost struct {
	UserId       int       `json:"user_id"`
	CategoryName string    `json:"category_name"`
	Title        string    `json:"title"`
	Discription  string    `json:"discription"`
	CreateDate   time.Time `json:"create_at"`
}

type UpdatePost struct {
	Title       string    `json:"title"`
	Discription string    `json:"discription"`
	CreateDate  time.Time `json:"create_at"`
	PostId      int       `json:"id"`
	UserId      int       `json:"user_id"`
}

type DeletePost struct {
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
}

type Post struct {
	PostId       int       `json:"post_id"`
	UserId       int       `json:"user_id"`
	CategoryName string    `json:"category_name"`
	Title        string    `json:"title"`
	Discription  string    `json:"discription"`
	CreateDate   time.Time `json:"create_at"`
}
