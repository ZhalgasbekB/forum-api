package model

type AdminPage struct {
	Issues []IssueModerator `json:"issues"`
	//Reports []Report         `json:"reports"`
}

type IssueModerator struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

//type Report struct {
//	Status  bool   `json:"status"`
//	Message string `json:"message"`
//}

type Moderator struct {
	Report      string `json:"report"`
	ModeratorId int    `json:"moderator_id"`
}
type UserIssue struct {
}
