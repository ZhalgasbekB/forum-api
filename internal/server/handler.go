package server

import (
	"fmt"
	"net/http"

	"gitea.com/lzhuk/forum/internal/model"
	"gitea.com/lzhuk/forum/internal/service"
)

type cookieKey string
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

const key = cookieKey("UserData")

func CookieFromContext(r *http.Request) (*model.User, error) {
	user, exist := r.Context().Value(key).(*model.User)
	if !exist {
		return nil, fmt.Errorf("NOT EXIST USER")
	}
	return user, nil
}
