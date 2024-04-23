package user

import (
	"database/sql"

	"gitea.com/lzhuk/forum/internal/errors"
	"gitea.com/lzhuk/forum/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

const (
	createUserQuery   = `INSERT INTO users (name, email, password, role, created_at) VALUES ($1,$2,$3,$4,$5)`
	userByIDQuery     = `SELECT * FROM users WHERE id = $1`
	usersByEmailQuery = `SELECT * FROM users WHERE email = $1`
	usersQuery        = `SELECT * FROM users`
)

func (u *UserRepository) CreateUser(user *model.User) error {
	if _, err := u.db.Exec(createUserQuery, user.Name, user.Email, user.Password, user.Role, user.CreatedAt); err != nil {
		switch err.Error() {
		case "UNIQUE constraint failed: users.email":
			return errors.ErrHaveDuplicateEmail
		case "UNIQUE constraint failed: users.name":
			return errors.ErrHaveDuplicateName
		default:
			return err
		}
	}
	return nil
}

func (u *UserRepository) UserByID(id int) (*model.User, error) {
	user := &model.User{}
	if err := u.db.QueryRow(userByIDQuery, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if err == errors.ErrSQLNoRows {
			return nil, errors.ErrNotFoundData
		}
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) UserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := u.db.QueryRow(usersByEmailQuery, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
		if err == errors.ErrSQLNoRows {
			return nil, errors.ErrInvalidCredentials
		}
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Users() ([]model.User, error) {
	var users []model.User
	rows, err := u.db.Query(usersQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
