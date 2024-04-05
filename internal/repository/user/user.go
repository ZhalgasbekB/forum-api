package user

import (
	"database/sql"
	"fmt"

	"gitea.com/lzhuk/forum/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

const (
	createUserQuery   = `INSERT INTO users (name, email, password, is_admin, created_at) VALUES ($1,$2,$3,$4,$5)` 
	userByIDQuery     = `SELECT * FROM users WHERE id = $1`
	usersByEmailQuery = `SELECT * FROM users WHERE email = $1`
	usersQuery        = `SELECT * FROM users`
)

func (u *UserRepository) CreateUser(user *model.User) error {
	if _, err := u.db.Exec(createUserQuery, user.Name, user.Email, user.Password, user.IsAdmin, user.CreatedAt); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

 
func (u *UserRepository) UserByID(id int) (*model.User, error) {
	var user model.User
	if err := u.db.QueryRow(userByIDQuery, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsAdmin, &user.CreatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) UserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.QueryRow(usersByEmailQuery, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsAdmin, &user.CreatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Users() ([]model.User, error) {
	var users []model.User
	rows, err := u.db.Query(usersQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.IsAdmin, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
