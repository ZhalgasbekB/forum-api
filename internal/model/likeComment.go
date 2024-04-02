package model

type LikeComment struct {
	UserId    int  `json:"user_id"`
	CommentId int  `json:"comment_id"`
	IsLike    bool `json:"is_like"`
	Like     int  `json:"like"`
	Dislike  int  `json:"dislike"`
}
