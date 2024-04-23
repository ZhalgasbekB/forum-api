package model

type AdminPage struct {
	Issues []IssueModerator `json:"issues"`
	// Reports []Report         `json:"reports"`
}

type IssueModerator struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

type Moderator struct {
	Report      string `json:"report"`
	ModeratorId int    `json:"moderator_id"`
}

type RoleDTO struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}
