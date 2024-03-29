package model

type Vote struct {
	UserId int `json:"user_id"`
	PostId int `json:"post_id"`
	Vote   int `json:"vote"`
}
