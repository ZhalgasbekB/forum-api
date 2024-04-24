package model

// ????
type AdminPage struct {
	Issues []IssueModerator `json:"issues"`
}

type IssueModerator struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

type Moderator struct {
	Report      string `json:"report"`
	ModeratorId int    `json:"moderator_id"`
}

/// USE
type RoleDTO struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}

type UserDeleteDTO struct {
	UserID int `json:"user_id"`
}

type UserUpdateDTO struct {
	UserID int `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
}
