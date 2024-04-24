package admin

import (
	"database/sql"
	"fmt"

	"gitea.com/lzhuk/forum/internal/model"
)

type AdminRepository struct {
	DB *sql.DB
}

const (
	userQuery       = `SELECT * FROM users WHERE role != $1 ORDER BY CASE WHEN role = $2 THEN 1 WHEN role = $3 THEN 2 ELSE 3 END;`
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

func (a *AdminRepository) UpdateUser(user model.User) error {
	fmt.Println(user.Role)
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

/// FULL UPDATE USER AND ADD TABLE WHICH TAKES A SOME DATA FROM MSSAGE FROM MODERATORS AND SERVE IT (???) need some think
/// 1. Update USER FULL
/// 2. MODERATOR ISSUES (???)
