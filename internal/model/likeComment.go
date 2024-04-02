package model

type LikeComment struct {
	UserId     int  `json:"user_id"`
	CommentId  int  `json:"comment_id"`
	LikeStatus bool `json:"status"`
}

type LikeCommentDTO struct {
	CommentId  int  `json:"comment_id"`
	LikeStatus bool `json:"status"`
}
