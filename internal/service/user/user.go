package user

import (
	"time"

	"gitea.com/lzhuk/forum/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	CreateUser(*model.User) error
	UpdateUser(*model.User, int) error
	DeleteUserByID(int) error
	UserByID(int) (*model.User, error)
	UserByEmail(string) (*model.User, error)
	Users() ([]model.User, error)
}

type IUserService interface {
	CreateUserService(*model.User) error
	UpdateUserService(*model.User, int) error
	DeleteUserByIDService(int) error
	UserByIDService(int) (*model.User, error)
	UserByEmailService(string, string) (*model.User, error)
	UsersService() ([]model.User, error)
}

type UserService struct {
	iUserRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{iUserRepository: userRepository}
}

func (us *UserService) CreateUserService(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	return us.iUserRepository.CreateUser(user)
}

func (us *UserService) UpdateUserService(user *model.User, id int) error {
	return us.iUserRepository.UpdateUser(user, id)
}

func (us *UserService) DeleteUserByIDService(id int) error {
	return us.iUserRepository.DeleteUserByID(id)
}

func (us *UserService) UserByIDService(id int) (*model.User, error) {
	return us.iUserRepository.UserByID(id)
}

func (us *UserService) UsersService() ([]model.User, error) {
	return us.iUserRepository.Users()
}

func (us *UserService) UserByEmailService(email, password string) (*model.User, error) {
	user, err := us.iUserRepository.UserByEmail(email)
	if err != nil {
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, err
	}
	return user, nil
}
