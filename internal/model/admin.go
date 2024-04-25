package model

// ????
type AdminPage struct { // Issues []IssueModerator `json:"issues"`
}


type RoleDTO struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}

type UserDeleteDTO struct {
	UserID int `json:"user_id"`
}

type UserUpdateDTO struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
