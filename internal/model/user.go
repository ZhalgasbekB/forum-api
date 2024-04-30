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
	Role     string `json:"role"`
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

type UserActivityDTO struct {
	CreatedPost   []Post            `json:"u_posts"`
	LikdedPost    []Post            `json:"u_likes"`
	DislikePost   []Post            `json:"u_dislikes"`
	CommentedPost []PostCommentsDTO `json:"U-comments"`
}
