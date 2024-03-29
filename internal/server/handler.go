package server

import "gitea.com/lzhuk/forum/internal/service"

type Handler struct {
	Services service.Service
}

func NewHandler(services service.Service) Handler {
	return Handler{
		Services: services,
	}
}
