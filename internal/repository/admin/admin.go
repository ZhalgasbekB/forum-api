package admin

import (
	"database/sql"
	"time"

	"gitea.com/lzhuk/forum/internal/model"
)

type AdminRepository struct {
	DB *sql.DB
}

const (
	usersQuery         = `SELECT * FROM users WHERE role != $1 ORDER BY CASE WHEN role = $2 THEN 1 WHEN role = $3 THEN 2 ELSE 3 END;`
	updateUserQuery    = `UPDATE users SET role = $1 WHERE id = $2`
	deleteUserQuery    = `DELETE FROM users WHERE id = $1`
	deletePostQuery    = `DELETE FROM posts WHERE id = $1`
	deleteCommentQuery = `DELETE FROM comments WHERE id = $1`

	categoryCreate = `INSERT INTO categories (category) VALUES ($1)`
	categoryDelete = `DELETE FROM categories WHERE category = $1 `
)

func InitAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{
		DB: db,
	}
}

func (a *AdminRepository) Users() ([]model.User, error) {
	users := []model.User{}
	rows, err := a.DB.Query(usersQuery, "ADMIN", "MODERATOR", "USER")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (a *AdminRepository) ChangeRole(user *model.User) error {
	if _, err := a.DB.Exec(updateUserQuery, user.Role, user.ID); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) DeleteUser(user_id int) error {
	if _, err := a.DB.Exec(deleteUserQuery, user_id); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) DeletePost(post_id int) error {
	if _, err := a.DB.Exec(deletePostQuery, post_id); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) DeleteComment(comment_id int) error {
	if _, err := a.DB.Exec(deleteCommentQuery, comment_id); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) CreateCategory(category string) error {
	if _, err := a.DB.Exec(categoryCreate, category); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) DeleteCategory(category string) error {
	if _, err := a.DB.Exec(categoryDelete, category); err != nil {
		return err
	}
	return nil
}

/// 2. MODERATOR ISSUES

const (
	reportCreateQuery     = `INSERT INTO reports (post_id, comment_id, user_id, moderator, category_issue, reason) VALUES ($1, $2, $3, $4, $5, $6)`
	reportUpdateQuery     = `UPDATE reports SET  status = $1, admin_response = $2, updated_at = $3 WHERE report_id = $4`
	reportDeleteQuery     = `DELETE FROM reports WHERE report_id = $1`
	reportsQuery          = `SELECT * FROM reports WHERE status = FALSE`
	reportsModeratorQuery = `SELECT * FROM reports WHERE moderator = $1`
)

func (a *AdminRepository) CreateReportModerator(report *model.ReportCreateDTO) error {
	if _, err := a.DB.Exec(reportCreateQuery, report.PostID, report.CommentID, report.UserID, report.ModeratorID, report.CategoryIssue, report.Reason); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) DeleteReport(id int) error {
	if _, err := a.DB.Exec(reportDeleteQuery, id); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) ResponseReportAdmin(update *model.ReportResponseDTO) error {
	updatedTime := time.Now()
	if _, err := a.DB.Exec(reportUpdateQuery, update.Status, update.AdminResponse, updatedTime, update.ReportID); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) ReportsByStatus() ([]model.Report, error) {
	reports := []model.Report{}
	rows, err := a.DB.Query(reportsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		report := &model.Report{}
		if err := rows.Scan(&report.ID, &report.PostID, &report.CommentID, &report.UserID, &report.ModeratorID, &report.Status, &report.CategoryIssue, &report.Reason, &report.AdminResponse, &report.CreateAt, &report.UpdateAt); err != nil {
			return nil, err
		}
		reports = append(reports, *report)
	}
	return reports, nil
}

// CHECK
func (a *AdminRepository) MonderatorReports(moderator_id int) ([]model.Report, error) {
	reports := []model.Report{}
	rows, err := a.DB.Query(reportsModeratorQuery, moderator_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		report := &model.Report{}
		if err := rows.Scan(&report.ID, &report.PostID, &report.CommentID, &report.UserID, &report.ModeratorID, &report.Status, &report.CategoryIssue, &report.Reason, &report.AdminResponse, &report.CreateAt, &report.UpdateAt); err != nil {
			return nil, err
		}
		reports = append(reports, *report)
	}
	return reports, nil
}

///  3 USER ISSUE

const (
	wantQuery       = `INSERT INTO wants (user_id, user_name) VALUES($1, $2)`
	wantsQuery      = `SELECT user_id, user_name, created_at FROM wants WHERE status = 0`
	updateWantQuery = `UPDATE wants SET status = $1 WHERE user_id = $2`
	wantsUser       = `SELECT status, created_at FROM wants WHERE user_id = $1`
)

func (a *AdminRepository) UserWant(w *model.WantsDTO) error {
	if _, err := a.DB.Exec(wantQuery, w.UserID, w.UserName); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) UsersWantRole() ([]model.Wants2DTO, error) {
	wants := []model.Wants2DTO{}
	rows, err := a.DB.Query(wantsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		want := &model.Wants2DTO{}
		if err := rows.Scan(&want.UserID, &want.UserName, &want.CreateAt); err != nil {
			return nil, err
		}
		wants = append(wants, *want)
	}
	return wants, err
}

func (a *AdminRepository) UpdateUserWantStatus(u *model.AdminResponse) error {
	if _, err := a.DB.Exec(updateWantQuery, u.Status, u.UserID); err != nil {
		return err
	}
	return nil
}

// CHECK

func (a *AdminRepository) UserWants(user_id int) ([]model.Wants1DTO, error) {
	wants := []model.Wants1DTO{}
	rows, err := a.DB.Query(wantsUser, user_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		want := &model.Wants1DTO{}
		if err := rows.Scan(&want.Status, &want.CreatedAt); err != nil {
			return nil, err
		}
		wants = append(wants, *want)
	}
	return wants, nil
}
