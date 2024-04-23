package admin

import (
	"database/sql"
	"gitea.com/lzhuk/forum/internal/model"
)

type AdminRepository struct {
	DB *sql.DB
}

const (
	userQuery       = `SELECT * FROM users WHERE role != $1`
	updateUserQuery = `UPDATE users SET name = $1, email = $2, password = $3, role = $4 WHERE user_id = $5`
	deleteUserQuery = `DELETE FROM users WHERE user_id = $1`
)

func InitAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{
		DB: db,
	}
}

func (a *AdminRepository) Users() ([]model.User, error) {
	users := []model.User{}
	rows, err := a.DB.Query(userQuery, "ADMIN")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := model.User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (a *AdminRepository) UpdateUser(user model.User) error {
	if _, err := a.DB.Query(updateUserQuery, user.Name, user.Email, user.Password, user.Role, user.ID); err != nil {
		return nil
	}
	return nil
}

func (a *AdminRepository) DeleteUser(user_id int) error {
	if _, err := a.DB.Query(deleteUserQuery, user_id); err != nil {
		return err
	}
	return nil
}
