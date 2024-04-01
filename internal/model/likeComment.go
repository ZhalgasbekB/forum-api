package model

type LikeComment struct {
	UserId    int  `json:"user_id"`
	CommentId int  `json:"post_id"`
	IsLike    bool `json:"vote"`
	Up        int  `json:"up"`
	Down      int  `json:"down"`
}
