package model

type LikePost struct {
	UserId     int  `json:"user_id"`
	PostId     int  `json:"post_id"`
	LikeStatus bool `json:"status"`
}

type LikePostDTO struct {
	UserId     int  `json:"user_id"`
	PostId     int  `json:"post_id"`
	LikeStatus bool `json:"status"`
}
