package model

type LikePost struct {
	UserId int  `json:"user_id"`
	PostId int  `json:"post_id"`
	IsLike bool `json:"vote"`
	Up     int  `json:"up"`
	Down   int  `json:"down"`
}


