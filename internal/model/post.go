package model

import "time"

type Post struct {
	PostId       int       `json:"post_id"`
	UserId       int       `json:"user_id"`
	CategoryName string    `json:"category_name"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreateDate   time.Time `json:"create_at"`

	Author  string `json:"name"`
	Like    int    `json:"likes"`
	Dislike int    `json:"dislikes"`
}

type UploadPostImage struct {
	Path   string `json:"path"`
	PostId int    `json:"post_id"`
}

type PathImagePost struct {
	Path string `json:"path"`
}

type CreatePostDTO struct {
	CategoryName string `json:"category_name"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

type UpdatePostDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostCommentsDTO struct {
	Post     *Post     `json:"post"`
	Comments []Comment `json:"comments"`
}

type PostsNotification struct {
	Posts           []*Post `json:"posts"`
	NewNotification bool    `json:"new_notifications"`
}
