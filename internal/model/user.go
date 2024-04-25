package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRegisterDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role  string `json:"role"` // role
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserAuthDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
