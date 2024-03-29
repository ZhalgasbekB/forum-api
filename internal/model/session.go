package model

import "time"

type Sessinon struct {
	UUID     string    `json:"uuid"`
	UserID   int       `json:"user_id"`
	ExpireAt time.Time `json:"expire_at"`
}
