package admin

import (
	"database/sql"

	"gitea.com/lzhuk/forum/internal/model"
)

type AdminRepository struct {
	DB *sql.DB
}

const (
	userQuery       = `SELECT * FROM users WHERE role != $1 ORDER BY CASE WHEN role = $2 THEN 1 WHEN role = $3 THEN 2 ELSE 3 END;`
	updateAllQuery  = `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	updateUserQuery = `UPDATE users SET role = $1 WHERE id = $2`
	deleteUserQuery = `DELETE FROM users WHERE id = $1`
)

func InitAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{
		DB: db,
	}
}

func (a *AdminRepository) Users() ([]model.User, error) {
	users := []model.User{}
	rows, err := a.DB.Query(userQuery, "ADMIN", "MODERATOR", "USER")
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

func (a *AdminRepository) UpdateUser(user *model.User) error {
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

func (a *AdminRepository) UpdateUserNewDate(user *model.User) error {
	if _, err := a.DB.Exec(updateAllQuery, user.Name, user.Email, user.ID); err != nil {
		return err
	}
	return nil
}

/// 2. MODERATOR ISSUES (???)

const (
	reportCreateQuery = `INSERT INTO reports (post_id, comment_id, user_id, moderator, category_issue, reason) VALUES ($1, $2, $3, $4, $5, $6)`
	reportUpdateQuery = `UPDATE reports SET admin = $1`
	reportDeleteQuery = `DELETE FROM reports WHERE report_id = $1`
)

func (a *AdminRepository) CreateReportRepository(report *model.ReportCreateDTO) error {
	if _, err := a.DB.Exec(reportCreateQuery, report.PostID, report.CommentID, report.UserID, report.ModeratorID, report.CategoryIssue, report.Reason); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) UpdateReport() error {
	return nil
}

func (a *AdminRepository) DeleteReport(id int) error {
	if _, err := a.DB.Exec(reportDeleteQuery, id); err != nil {
		return err
	}
	return nil
}

func (a *AdminRepository) Report() error {
	return nil
}

// more code
