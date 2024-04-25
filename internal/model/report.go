package model

import "time"

type Report struct {
	ID            int       `json:"id"`
	PostID        int       `json:"post_id"`
	CommentID     int       `json:"comment_id"`
	UserID        int       `json:"user_id"`
	ModeratorID   int       `json:"moderator_id"`
	Status        string    `json:"status"`
	CategoryIssue string    `json:"category_issue"`
	Reason        string    `json:"reason"`
	AdminResponse string    `json:"admin_response"`
	CreateAt      time.Time `json:"created_at"`
	UpdateAt      time.Time `json:"update_at"`
}

type (
	ReportCreateDTO struct {
		PostID        int    `json:"post_id"`
		CommentID     int    `json:"comment_id"`
		UserID        int    `json:"user_id"`
		ModeratorID   int    `json:"moderator_id"`
		CategoryIssue string `json:"category_issue"`
		Reason        string `json:"reason"`
	}

	ReportResponseDTO struct {
		ReportID      int    `json:"report_id"`
		AdminResponse string `json:"response"`
		Status        string `json:"status"`
	}
)
