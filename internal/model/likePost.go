package model

type LikePost struct {
	UserId int `json:"user_id"`
	PostId int `json:"post_id"`
	Vote   int `json:"vote"`
}
