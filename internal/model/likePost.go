package model

type LikePost struct {
	UserId     int  `json:"user_id"`
	PostId     int  `json:"post_id"`
	LikeStatus bool `json:"is_like"`   // false == 0 by defailt if 1 -1
	LikeCode   int  `json:"like_code"` //  1, -1
}

type LikePostDTO struct {
	LikeStatus bool `json:"is_like"`   // false == 0 by defailt if 1 -1
	LikeCode   int  `json:"like_code"` //  1, -1
}
