package model

type LikePost struct {
	UserId     int  `json:"user_id"`
	PostId     int  `json:"post_id"`
	LikeStatus bool `json:"status"` // false == 0 by defailt if 1 -1
}

type LikePostDTO struct {
	PostId     int  `json:"post_id"`
	LikeStatus bool `json:"status"` // false == 0 by defailt if 1 -1
}
