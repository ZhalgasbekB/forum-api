package server

import (
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
	"gitea.com/lzhuk/forum/internal/service"
)

type KeyUser string
type (
	Handler struct {
		Services service.Service
	}
)

func NewHandler(services service.Service) Handler {
	return Handler{
		Services: services,
	}
}

const key = KeyUser("UserData")

func userFromContext(r *http.Request) *model.User {
	user, exist := r.Context().Value(key).(*model.User)
	if !exist {
		return nil
	}
	return user
}
