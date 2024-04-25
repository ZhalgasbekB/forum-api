package model

import "time"

type Report struct {
	ID            int       `json:"id"`
	PostId        int       `json:"post_id"`
	CommentId     int       `json:"comment_id"`
	ReportedBy    int       `json:"moderator_id"`
	AssignedTo    int       `json:"admin_id"`
	Status        string    `json:"status"`
	Reason        string    `json:"reason"`
	AdminResponse string    `json:"admin_response"`
	CreateAt      time.Time `json:"created_at"`
	UpdateAt      time.Time `json:"update_at"`
}

type (
	ReportDTO  struct{}
	ReportDTO1 struct{}
	ReportDTO2 struct{}
	ReportDTO3 struct{}
	ReportDTO4 struct{}
)
